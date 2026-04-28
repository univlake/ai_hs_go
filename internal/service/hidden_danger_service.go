// Package service — 隐患信息业务层。
//
// HiddenDangerService 实现 Laravel HiddenDangerInfoController 的全部接口：
//   CRUD + rectify (整改提交) + confirmRectification (整改确认) + archive (归档)
//
// 三类附件 (hiddenExamineAnnex / hiddenReviewAnnex / hiddenRectifyAnnex) 通过 is_review
// 字段区分，按 hidden_danger_info_id 整体替换：传 nil 不动，传非 nil（含空切片）= 全替换。
package service

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

type HiddenDangerService struct {
	db *gorm.DB
}

func NewHiddenDangerService(db *gorm.DB) *HiddenDangerService {
	return &HiddenDangerService{db: db}
}

// ── 列表 ─────────────────────────────────────────────────────────────

type HiddenDangerListParams struct {
	Page               int
	PerPage            int
	ExpertInspectionID int64
	ExpertCompanyID    int64
	CheckCompanyID     int64
	CompanyID          int64
	ExpertUserID       int64
	Area               string
	HiddenLevel        string
	HiddenType         string
	CorrectStatus      *int
}

func (s *HiddenDangerService) List(ctx context.Context, p HiddenDangerListParams) ([]model.HiddenDangerInfo, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.HiddenDangerInfo{})
	if p.ExpertInspectionID > 0 {
		q = q.Where("expert_inspection_id = ?", p.ExpertInspectionID)
	}
	if p.ExpertCompanyID > 0 {
		q = q.Where("expert_company_id = ?", p.ExpertCompanyID)
	}
	if p.CheckCompanyID > 0 {
		q = q.Where("check_company_id = ?", p.CheckCompanyID)
	}
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if p.ExpertUserID > 0 {
		// 跟 Laravel：当传了 expert_user_id，把范围限制到该专家被指派的所有 expert_companies。
		q = q.Where("expert_company_id IN (?)",
			s.db.Model(&model.ExpertInspectionUser{}).
				Select("expert_company_id").
				Where("expert_user_id = ?", p.ExpertUserID),
		)
	}
	if v := strings.TrimSpace(p.Area); v != "" {
		q = q.Where("area ILIKE ?", "%"+v+"%")
	}
	if v := strings.TrimSpace(p.HiddenLevel); v != "" {
		q = q.Where("hidden_level = ?", v)
	}
	if v := strings.TrimSpace(p.HiddenType); v != "" {
		q = q.Where("hidden_type = ?", v)
	}
	if p.CorrectStatus != nil {
		q = q.Where("correct_status = ?", *p.CorrectStatus)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HiddenDangerInfo
	if err := q.
		Preload("CheckCompany", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order("created_at DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	if err := s.attachAnnexes(ctx, rows); err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// Get — show 接口；含三类附件 + checkCompany。
func (s *HiddenDangerService) Get(ctx context.Context, id int64) (*model.HiddenDangerInfo, error) {
	var m model.HiddenDangerInfo
	if err := s.db.WithContext(ctx).
		Preload("CheckCompany", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	rows := []model.HiddenDangerInfo{m}
	if err := s.attachAnnexes(ctx, rows); err != nil {
		return nil, err
	}
	return &rows[0], nil
}

// AnnexInput — 创建/更新隐患时一并提交三类附件。nil = 不动；非 nil（含空切片）= 整体替换。
type AnnexInput struct {
	Geolocation string
	AnnexURL    string
}

type HiddenDangerInput struct {
	Record         *model.HiddenDangerInfo
	ExamineAnnexes *[]AnnexInput
	ReviewAnnexes  *[]AnnexInput
	RectifyAnnexes *[]AnnexInput
}

func (s *HiddenDangerService) Create(ctx context.Context, in HiddenDangerInput) (*model.HiddenDangerInfo, error) {
	if in.Record.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Record).Error; err != nil {
			return wrapWrite(err)
		}
		if err := s.replaceAnnexes(tx, in.Record.ID, 1, in.ExamineAnnexes, in.Record.UserID, false); err != nil {
			return err
		}
		if err := s.replaceAnnexes(tx, in.Record.ID, 2, in.ReviewAnnexes, in.Record.UserID, false); err != nil {
			return err
		}
		return s.replaceAnnexes(tx, in.Record.ID, 3, in.RectifyAnnexes, in.Record.UserID, false)
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, in.Record.ID)
}

func (s *HiddenDangerService) Update(
	ctx context.Context,
	id int64,
	updates map[string]any,
	examineAnnexes, reviewAnnexes, rectifyAnnexes *[]AnnexInput,
	uid int64,
) (*model.HiddenDangerInfo, error) {
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.HiddenDangerInfo{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "隐患信息不存在")
		}
		if err := s.replaceAnnexes(tx, id, 1, examineAnnexes, uid, true); err != nil {
			return err
		}
		if err := s.replaceAnnexes(tx, id, 2, reviewAnnexes, uid, true); err != nil {
			return err
		}
		return s.replaceAnnexes(tx, id, 3, rectifyAnnexes, uid, true)
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

func (s *HiddenDangerService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("hidden_danger_info_id = ?", id).Delete(&model.HiddenUploadAnnex{}).Error; err != nil {
			return wrapWrite(err)
		}
		res := tx.Delete(&model.HiddenDangerInfo{}, id)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "隐患信息不存在")
		}
		return nil
	})
}

// Rectify — 受检企业或专家提交整改：写 zg_user_id / correct_date / review_description / correct_status=2，
// 同时整体替换 hiddenRectifyAnnex 附件。
func (s *HiddenDangerService) Rectify(ctx context.Context, id int64, uid int64, reviewDescription string, correctDate *model.Date, rectifyAnnexes []AnnexInput) (*model.HiddenDangerInfo, error) {
	if strings.TrimSpace(reviewDescription) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "整改描述不能为空")
	}
	cd := correctDate
	if cd == nil {
		d := model.Date(time.Now())
		cd = &d
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.HiddenDangerInfo{}).Where("id = ?", id).Updates(map[string]any{
			"zg_user_id":         uid,
			"correct_date":       cd,
			"review_description": reviewDescription,
			"correct_status":     2,
			"last_user_id":       uid,
		})
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "隐患信息不存在")
		}
		if len(rectifyAnnexes) == 0 {
			return nil
		}
		ax := rectifyAnnexes
		return s.replaceAnnexes(tx, id, 3, &ax, uid, true)
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

// ConfirmRectification — 专家或集团确认整改合格：写 review_user_id / review_time / correct_status=3，
// 同时整体替换 hiddenReviewAnnex。
func (s *HiddenDangerService) ConfirmRectification(ctx context.Context, id int64, uid int64, reviewAnnexes []AnnexInput) (*model.HiddenDangerInfo, error) {
	now := time.Now()
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.HiddenDangerInfo{}).Where("id = ?", id).Updates(map[string]any{
			"review_user_id": uid,
			"review_time":    now,
			"correct_status": 3,
			"last_user_id":   uid,
		})
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "隐患信息不存在")
		}
		if len(reviewAnnexes) == 0 {
			return nil
		}
		ax := reviewAnnexes
		return s.replaceAnnexes(tx, id, 2, &ax, uid, true)
	})
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, id)
}

// Archive — 仅集团可操作：写 archive_user_id / invest_status=2。
func (s *HiddenDangerService) Archive(ctx context.Context, id int64, uid int64) (*model.HiddenDangerInfo, error) {
	res := s.db.WithContext(ctx).Model(&model.HiddenDangerInfo{}).Where("id = ?", id).Updates(map[string]any{
		"archive_user_id": uid,
		"invest_status":   2,
		"last_user_id":    uid,
	})
	if res.Error != nil {
		return nil, wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, errs.New(errs.CodeNotFound, "隐患信息不存在")
	}
	return s.Get(ctx, id)
}

// ── 内部 helpers ─────────────────────────────────────────────────────

func (s *HiddenDangerService) replaceAnnexes(tx *gorm.DB, dangerID int64, isReview int, annexes *[]AnnexInput, uid int64, wipe bool) error {
	if annexes == nil {
		return nil
	}
	if wipe {
		if err := tx.Where("hidden_danger_info_id = ? AND is_review = ?", dangerID, isReview).
			Delete(&model.HiddenUploadAnnex{}).Error; err != nil {
			return wrapWrite(err)
		}
	}
	if len(*annexes) == 0 {
		return nil
	}
	rows := make([]model.HiddenUploadAnnex, 0, len(*annexes))
	for _, a := range *annexes {
		if a.AnnexURL == "" {
			continue
		}
		rows = append(rows, model.HiddenUploadAnnex{
			HiddenDangerInfoID: dangerID,
			Geolocation:        a.Geolocation,
			IsReview:           isReview,
			AnnexURL:           a.AnnexURL,
			UserID:             uid,
		})
	}
	if len(rows) == 0 {
		return nil
	}
	return wrapWrite(tx.Create(&rows).Error)
}

// attachAnnexes — Laravel preload 的三类附件靠 is_review 区分；GORM 的 Preload 不支持
// 同一关联按 where 拆三份，因此这里手写一次性拉取再 fan-out。
func (s *HiddenDangerService) attachAnnexes(ctx context.Context, rows []model.HiddenDangerInfo) error {
	if len(rows) == 0 {
		return nil
	}
	ids := make([]int64, 0, len(rows))
	idx := make(map[int64]int, len(rows))
	for i, r := range rows {
		ids = append(ids, r.ID)
		idx[r.ID] = i
	}
	var annexes []model.HiddenUploadAnnex
	if err := s.db.WithContext(ctx).
		Where("hidden_danger_info_id IN ?", ids).
		Order("id ASC").
		Find(&annexes).Error; err != nil {
		return err
	}
	for _, a := range annexes {
		i, ok := idx[a.HiddenDangerInfoID]
		if !ok {
			continue
		}
		switch a.IsReview {
		case 1:
			rows[i].HiddenExamineAnnex = append(rows[i].HiddenExamineAnnex, a)
		case 2:
			rows[i].HiddenReviewAnnex = append(rows[i].HiddenReviewAnnex, a)
		case 3:
			rows[i].HiddenRectifyAnnex = append(rows[i].HiddenRectifyAnnex, a)
		}
	}
	return nil
}
