// Package service — 考核评分模块业务层。
//
// 一个 service 覆盖 Laravel inspections.php 的 10 个控制器：
//
//	DailyCheckTemplates / DailyChecks            日常目标管理工作（filing 流：草稿→已提交→已评分）
//	AccidentCheckTemplates / AccidentChecks      安全生产事故考核（一次性评分）
//	AddPointsTemplates / AddPoints               加分项（一次性评分）
//	SystemChecks                                 活动参与及材料上报（compute 自动算 + 手动 CRUD）
//	NormChecks                                   安全生产标准化（每企业每年单条）
//	AssessmentTemplates                          跨三表年度并集 + 一键拷贝
//	AssessmentSummary                            6 维度汇总 + Word 导出
//
// 三张模板表（daily/accident/add_points）共享 (title,level,sort,year,pid,score_range)
// 形状，泛型 helper 驱动 List/Show/CRUD；填报数据三张表共享 (template_id, path, number,
// comment, year, company_id) 形状但 status 仅在 daily 表存在。
package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

type AssessmentService struct {
	db *gorm.DB
}

func NewAssessmentService(db *gorm.DB) *AssessmentService {
	return &AssessmentService{db: db}
}

// ── 通用模板树 helpers ────────────────────────────────────────────

// daily check templates ────────────────────────────────────────────

func (s *AssessmentService) ListDailyCheckTemplates(ctx context.Context, year int) ([]model.DailyCheckTemplate, error) {
	var tops []model.DailyCheckTemplate
	q := s.db.WithContext(ctx).Where("level = ?", 1)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	if err := q.Order("sort ASC").Find(&tops).Error; err != nil {
		return nil, err
	}
	if len(tops) == 0 {
		return tops, nil
	}
	pids := make([]int64, 0, len(tops))
	for _, t := range tops {
		pids = append(pids, t.ID)
	}
	var children []model.DailyCheckTemplate
	cq := s.db.WithContext(ctx).Where("pid IN ?", pids)
	if year > 0 {
		cq = cq.Where("year = ?", year)
	}
	if err := cq.Order("sort ASC").Find(&children).Error; err != nil {
		return nil, err
	}
	idx := make(map[int64]int, len(tops))
	for i, t := range tops {
		idx[t.ID] = i
	}
	for _, c := range children {
		if i, ok := idx[c.Pid]; ok {
			tops[i].Child = append(tops[i].Child, c)
		}
	}
	return tops, nil
}

func (s *AssessmentService) GetDailyCheckTemplate(ctx context.Context, id int64) (*model.DailyCheckTemplate, error) {
	var t model.DailyCheckTemplate
	if err := s.db.WithContext(ctx).First(&t, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	var children []model.DailyCheckTemplate
	if err := s.db.WithContext(ctx).Where("pid = ?", t.ID).Order("sort ASC").Find(&children).Error; err != nil {
		return nil, err
	}
	t.Child = children
	return &t, nil
}

func (s *AssessmentService) CreateDailyCheckTemplate(ctx context.Context, in *model.DailyCheckTemplate) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateDailyCheckTemplate(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.DailyCheckTemplate{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "日常检查模板不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteDailyCheckTemplate(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.DailyCheckTemplate{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "日常检查模板不存在")
	}
	return nil
}

// accident check templates ─────────────────────────────────────────

func (s *AssessmentService) ListAccidentCheckTemplates(ctx context.Context, year int) ([]model.AccidentCheckTemplate, error) {
	var tops []model.AccidentCheckTemplate
	q := s.db.WithContext(ctx).Where("level = ?", 1)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	if err := q.Order("sort ASC").Find(&tops).Error; err != nil {
		return nil, err
	}
	if len(tops) == 0 {
		return tops, nil
	}
	pids := make([]int64, 0, len(tops))
	for _, t := range tops {
		pids = append(pids, t.ID)
	}
	var children []model.AccidentCheckTemplate
	cq := s.db.WithContext(ctx).Where("pid IN ?", pids)
	if year > 0 {
		cq = cq.Where("year = ?", year)
	}
	if err := cq.Order("sort ASC").Find(&children).Error; err != nil {
		return nil, err
	}
	idx := make(map[int64]int, len(tops))
	for i, t := range tops {
		idx[t.ID] = i
	}
	for _, c := range children {
		if i, ok := idx[c.Pid]; ok {
			tops[i].Child = append(tops[i].Child, c)
		}
	}
	return tops, nil
}

func (s *AssessmentService) GetAccidentCheckTemplate(ctx context.Context, id int64) (*model.AccidentCheckTemplate, error) {
	var t model.AccidentCheckTemplate
	if err := s.db.WithContext(ctx).First(&t, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	var children []model.AccidentCheckTemplate
	if err := s.db.WithContext(ctx).Where("pid = ?", t.ID).Order("sort ASC").Find(&children).Error; err != nil {
		return nil, err
	}
	t.Child = children
	return &t, nil
}

func (s *AssessmentService) CreateAccidentCheckTemplate(ctx context.Context, in *model.AccidentCheckTemplate) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateAccidentCheckTemplate(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.AccidentCheckTemplate{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "事故检查模板不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteAccidentCheckTemplate(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.AccidentCheckTemplate{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "事故检查模板不存在")
	}
	return nil
}

// add points templates ─────────────────────────────────────────────

func (s *AssessmentService) ListAddPointsTemplates(ctx context.Context, year int) ([]model.AddPointsTemplate, error) {
	var tops []model.AddPointsTemplate
	q := s.db.WithContext(ctx).Where("level = ?", 1)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	if err := q.Order("sort ASC").Find(&tops).Error; err != nil {
		return nil, err
	}
	if len(tops) == 0 {
		return tops, nil
	}
	pids := make([]int64, 0, len(tops))
	for _, t := range tops {
		pids = append(pids, t.ID)
	}
	var children []model.AddPointsTemplate
	cq := s.db.WithContext(ctx).Where("pid IN ?", pids)
	if year > 0 {
		cq = cq.Where("year = ?", year)
	}
	if err := cq.Order("sort ASC").Find(&children).Error; err != nil {
		return nil, err
	}
	idx := make(map[int64]int, len(tops))
	for i, t := range tops {
		idx[t.ID] = i
	}
	for _, c := range children {
		if i, ok := idx[c.Pid]; ok {
			tops[i].Child = append(tops[i].Child, c)
		}
	}
	return tops, nil
}

func (s *AssessmentService) GetAddPointsTemplate(ctx context.Context, id int64) (*model.AddPointsTemplate, error) {
	var t model.AddPointsTemplate
	if err := s.db.WithContext(ctx).First(&t, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	var children []model.AddPointsTemplate
	if err := s.db.WithContext(ctx).Where("pid = ?", t.ID).Order("sort ASC").Find(&children).Error; err != nil {
		return nil, err
	}
	t.Child = children
	return &t, nil
}

func (s *AssessmentService) CreateAddPointsTemplate(ctx context.Context, in *model.AddPointsTemplate) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateAddPointsTemplate(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.AddPointsTemplate{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "加分项模板不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteAddPointsTemplate(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.AddPointsTemplate{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "加分项模板不存在")
	}
	return nil
}

// ── 日常检查（filing 流） ─────────────────────────────────────────

const (
	DailyCheckStatusDraft     = 0
	DailyCheckStatusCompleted = 1
	DailyCheckStatusScored    = 2
)

// DailySubmissionRow — submissions 列表行；company_id+year 聚合。
type DailySubmissionRow struct {
	CompanyID   int64      `json:"company_id"`
	CompanyName string     `json:"company_name"`
	Year        int        `json:"year"`
	Status      int        `json:"status"`
	TotalNumber float64    `json:"total_number"`
	RowCount    int        `json:"row_count"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
}

// ListDailySubmissions — 上级（pid）或本企业（company_id）视角的填报记录聚合。
func (s *AssessmentService) ListDailySubmissions(ctx context.Context, pid, companyID int64, status *int, year int) ([]DailySubmissionRow, error) {
	q := s.db.WithContext(ctx).Model(&model.DailyCheck{})

	if pid > 0 {
		var subordinateIDs []int64
		if err := s.db.WithContext(ctx).Model(&model.Company{}).
			Where("parent_id = ?", pid).Pluck("id", &subordinateIDs).Error; err != nil {
			return nil, err
		}
		if len(subordinateIDs) == 0 {
			return []DailySubmissionRow{}, nil
		}
		q = q.Where("company_id IN ?", subordinateIDs)
		if status != nil {
			q = q.Where("status = ?", *status)
		} else {
			// 上级视角默认隐藏草稿。
			q = q.Where("status IN ?", []int{DailyCheckStatusCompleted, DailyCheckStatusScored})
		}
	} else {
		q = q.Where("company_id = ?", companyID)
		if status != nil {
			q = q.Where("status = ?", *status)
		}
	}

	if year > 0 {
		q = q.Where("year = ?", year)
	}

	type aggRow struct {
		CompanyID   int64
		Year        int
		Status      int
		TotalNumber float64
		RowCount    int
		UpdatedAt   *time.Time
	}
	var aggs []aggRow
	if err := q.
		Select(`company_id, "year", MIN(status) AS status, COALESCE(SUM(number), 0) AS total_number, COUNT(*) AS row_count, MAX(updated_at) AS updated_at`).
		Group(`company_id, "year"`).
		Order(`"year" DESC`).
		Find(&aggs).Error; err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(aggs))
	for _, a := range aggs {
		ids = append(ids, a.CompanyID)
	}
	titles := s.companyTitles(ctx, ids)
	out := make([]DailySubmissionRow, 0, len(aggs))
	for _, a := range aggs {
		out = append(out, DailySubmissionRow{
			CompanyID:   a.CompanyID,
			CompanyName: titles[a.CompanyID],
			Year:        a.Year,
			Status:      a.Status,
			TotalNumber: roundFloat(a.TotalNumber, 1),
			RowCount:    a.RowCount,
			UpdatedAt:   a.UpdatedAt,
			Title:       "日常目标管理工作考核",
		})
	}
	return out, nil
}

// StartDailyForm — 拉模板树 + 本企业已有填报。is_new=true 时检查是否重复。
func (s *AssessmentService) StartDailyForm(ctx context.Context, companyID int64, year int, isNew bool) ([]model.DailyCheckTemplate, error) {
	if isNew {
		var cnt int64
		if err := s.db.WithContext(ctx).Model(&model.DailyCheck{}).
			Where("company_id = ? AND year = ?", companyID, year).Count(&cnt).Error; err != nil {
			return nil, err
		}
		if cnt > 0 {
			return nil, errs.New(errs.CodeUnprocessable, "该年份已有填报记录，如需修改请从列表中继续填报")
		}
	}
	tops, err := s.ListDailyCheckTemplates(ctx, year)
	if err != nil {
		return nil, err
	}
	if len(tops) == 0 {
		return tops, nil
	}
	// 拿所有二级模板的 ID，一次拉本企业现有评分，回填。
	childIDs := make([]int64, 0)
	for _, top := range tops {
		for _, c := range top.Child {
			childIDs = append(childIDs, c.ID)
		}
	}
	if len(childIDs) == 0 {
		return tops, nil
	}
	var checks []model.DailyCheck
	if err := s.db.WithContext(ctx).
		Where("daily_check_template_id IN ? AND company_id = ? AND year = ?", childIDs, companyID, year).
		Find(&checks).Error; err != nil {
		return nil, err
	}
	checkByTemplate := make(map[int64]*model.DailyCheck, len(checks))
	for i := range checks {
		checkByTemplate[checks[i].DailyCheckTemplateID] = &checks[i]
	}
	for ti := range tops {
		for ci := range tops[ti].Child {
			cid := tops[ti].Child[ci].ID
			if c, ok := checkByTemplate[cid]; ok {
				tops[ti].Child[ci].Check = c
			}
		}
	}
	return tops, nil
}

// DailyDraftItem — saveDraft 提交的单行。
type DailyDraftItem struct {
	DailyCheckTemplateID int64
	Path                 *string
}

func (s *AssessmentService) SaveDailyDraft(ctx context.Context, companyID int64, year int, userID int64, items []DailyDraftItem) error {
	// 锁检查：年份内若有非草稿状态的记录，禁止再修改。
	var lockedCnt int64
	if err := s.db.WithContext(ctx).Model(&model.DailyCheck{}).
		Where("company_id = ? AND year = ? AND status != ?", companyID, year, DailyCheckStatusDraft).
		Count(&lockedCnt).Error; err != nil {
		return err
	}
	if lockedCnt > 0 {
		return errs.New(errs.CodeForbidden, "该年份填报已提交，无法再修改")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, it := range items {
			if it.DailyCheckTemplateID == 0 {
				continue
			}
			var row model.DailyCheck
			err := tx.Where("daily_check_template_id = ? AND company_id = ? AND year = ?",
				it.DailyCheckTemplateID, companyID, year).First(&row).Error
			isNew := errors.Is(err, gorm.ErrRecordNotFound)
			if err != nil && !isNew {
				return wrapWrite(err)
			}
			row.DailyCheckTemplateID = it.DailyCheckTemplateID
			row.CompanyID = companyID
			row.Year = year
			row.Path = it.Path
			row.Status = DailyCheckStatusDraft
			if isNew {
				row.UserID = userID
			}
			row.LastUserID = userID
			if isNew {
				if err := tx.Create(&row).Error; err != nil {
					return wrapWrite(err)
				}
			} else {
				if err := tx.Save(&row).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *AssessmentService) CompleteDaily(ctx context.Context, companyID int64, year int, userID int64) error {
	res := s.db.WithContext(ctx).Model(&model.DailyCheck{}).
		Where("company_id = ? AND year = ? AND status = ?", companyID, year, DailyCheckStatusDraft).
		Updates(map[string]any{
			"status":       DailyCheckStatusCompleted,
			"last_user_id": userID,
		})
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeUnprocessable, "未找到可提交的草稿")
	}
	return nil
}

// DailyScoreItem — score 接口的单行评分。
type DailyScoreItem struct {
	DailyCheckTemplateID int64
	Number               *float64
	Comment              *string
}

func (s *AssessmentService) ScoreDaily(ctx context.Context, companyID int64, year int, userID int64, items []DailyScoreItem) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, it := range items {
			number := 0.0
			if it.Number != nil {
				number = roundFloat(*it.Number, 1)
			}
			res := tx.Model(&model.DailyCheck{}).
				Where("company_id = ? AND year = ? AND daily_check_template_id = ?",
					companyID, year, it.DailyCheckTemplateID).
				Updates(map[string]any{
					"number":       number,
					"comment":      it.Comment,
					"status":       DailyCheckStatusScored,
					"last_user_id": userID,
				})
			if res.Error != nil {
				return wrapWrite(res.Error)
			}
		}
		return nil
	})
}

// DailySubmissionDetail — showSubmission 返回 templates + meta。
type DailySubmissionDetail struct {
	CompanyID   int64                      `json:"company_id"`
	Year        int                        `json:"year"`
	Status      int                        `json:"status"`
	TotalNumber float64                    `json:"total_number"`
	Templates   []model.DailyCheckTemplate `json:"templates"`
}

func (s *AssessmentService) ShowDailySubmission(ctx context.Context, companyID int64, year int) (*DailySubmissionDetail, error) {
	tops, err := s.StartDailyForm(ctx, companyID, year, false)
	if err != nil {
		return nil, err
	}
	type metaRow struct {
		Status      int
		TotalNumber float64
	}
	var m metaRow
	if err := s.db.WithContext(ctx).Model(&model.DailyCheck{}).
		Where("company_id = ? AND year = ?", companyID, year).
		Select("COALESCE(MIN(status), 0) AS status, COALESCE(SUM(number), 0) AS total_number").
		Scan(&m).Error; err != nil {
		return nil, err
	}
	return &DailySubmissionDetail{
		CompanyID:   companyID,
		Year:        year,
		Status:      m.Status,
		TotalNumber: roundFloat(m.TotalNumber, 1),
		Templates:   tops,
	}, nil
}

func (s *AssessmentService) DestroyDailyForm(ctx context.Context, companyID int64, year int) error {
	var lockedCnt int64
	if err := s.db.WithContext(ctx).Model(&model.DailyCheck{}).
		Where("company_id = ? AND year = ? AND status != ?", companyID, year, DailyCheckStatusDraft).
		Count(&lockedCnt).Error; err != nil {
		return err
	}
	if lockedCnt > 0 {
		return errs.New(errs.CodeForbidden, "该年份填报已提交，无法删除")
	}
	if err := s.db.WithContext(ctx).
		Where("company_id = ? AND year = ? AND status = ?", companyID, year, DailyCheckStatusDraft).
		Delete(&model.DailyCheck{}).Error; err != nil {
		return wrapWrite(err)
	}
	return nil
}

// ListDailyChecks — legacy paginated index。
func (s *AssessmentService) ListDailyChecks(ctx context.Context, page, perPage int, companyID int64, year int) ([]model.DailyCheck, int64, error) {
	page, per := normalizePage(page, perPage)
	q := s.db.WithContext(ctx).Model(&model.DailyCheck{})
	if companyID > 0 {
		q = q.Where("company_id = ?", companyID)
	}
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.DailyCheck
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AssessmentService) GetDailyCheck(ctx context.Context, id int64) (*model.DailyCheck, error) {
	var m model.DailyCheck
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

// ── 事故检查（一次性评分） ────────────────────────────────────────

// AccidentSubmissionRow — submissions 行；事故扣分汇总后转最终得分（max(0, 20 - 累计扣分)）。
type AccidentSubmissionRow struct {
	CompanyID   int64      `json:"company_id"`
	CompanyName string     `json:"company_name"`
	Year        int        `json:"year"`
	RawNumber   float64    `json:"raw_number"`
	FinalNumber float64    `json:"final_number"`
	RowCount    int        `json:"row_count"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
}

func (s *AssessmentService) ListAccidentSubmissions(ctx context.Context, pid int64, year int) ([]AccidentSubmissionRow, error) {
	var subordinateIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("parent_id = ?", pid).Pluck("id", &subordinateIDs).Error; err != nil {
		return nil, err
	}
	if len(subordinateIDs) == 0 {
		return []AccidentSubmissionRow{}, nil
	}
	q := s.db.WithContext(ctx).Model(&model.AccidentCheck{}).Where("company_id IN ?", subordinateIDs)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	type aggRow struct {
		CompanyID int64
		Year      int
		RawNumber float64
		RowCount  int
		UpdatedAt *time.Time
	}
	var aggs []aggRow
	if err := q.
		Select(`company_id, "year", COALESCE(SUM(number), 0) AS raw_number, COUNT(*) AS row_count, MAX(updated_at) AS updated_at`).
		Group(`company_id, "year"`).
		Order(`"year" DESC`).
		Find(&aggs).Error; err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(aggs))
	for _, a := range aggs {
		ids = append(ids, a.CompanyID)
	}
	titles := s.companyTitles(ctx, ids)
	out := make([]AccidentSubmissionRow, 0, len(aggs))
	for _, a := range aggs {
		raw := roundFloat(a.RawNumber, 1)
		final := math.Max(0, 20-raw)
		out = append(out, AccidentSubmissionRow{
			CompanyID:   a.CompanyID,
			CompanyName: titles[a.CompanyID],
			Year:        a.Year,
			RawNumber:   raw,
			FinalNumber: roundFloat(final, 1),
			RowCount:    a.RowCount,
			UpdatedAt:   a.UpdatedAt,
			Title:       "安全生产安全生产事故考核指标",
		})
	}
	return out, nil
}

// StartAccidentForm — 拉模板树 + 本企业已有评分。
func (s *AssessmentService) StartAccidentForm(ctx context.Context, companyID int64, year int) ([]model.AccidentCheckTemplate, error) {
	tops, err := s.ListAccidentCheckTemplates(ctx, year)
	if err != nil {
		return nil, err
	}
	if len(tops) == 0 {
		return tops, nil
	}
	childIDs := make([]int64, 0)
	for _, top := range tops {
		for _, c := range top.Child {
			childIDs = append(childIDs, c.ID)
		}
	}
	if len(childIDs) == 0 {
		return tops, nil
	}
	var checks []model.AccidentCheck
	if err := s.db.WithContext(ctx).
		Where("accident_check_template_id IN ? AND company_id = ? AND year = ?", childIDs, companyID, year).
		Find(&checks).Error; err != nil {
		return nil, err
	}
	checkByTemplate := make(map[int64]*model.AccidentCheck, len(checks))
	for i := range checks {
		checkByTemplate[checks[i].AccidentCheckTemplateID] = &checks[i]
	}
	for ti := range tops {
		for ci := range tops[ti].Child {
			cid := tops[ti].Child[ci].ID
			if c, ok := checkByTemplate[cid]; ok {
				tops[ti].Child[ci].Check = c
			}
		}
	}
	return tops, nil
}

// AccidentScoreItem — saveScore 单行评分。
type AccidentScoreItem struct {
	AccidentCheckTemplateID int64
	Number                  *float64
	Comment                 *string
}

func (s *AssessmentService) SaveAccidentScore(ctx context.Context, companyID int64, year int, userID int64, items []AccidentScoreItem) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, it := range items {
			if it.AccidentCheckTemplateID == 0 {
				continue
			}
			var row model.AccidentCheck
			err := tx.Where("accident_check_template_id = ? AND company_id = ? AND year = ?",
				it.AccidentCheckTemplateID, companyID, year).First(&row).Error
			isNew := errors.Is(err, gorm.ErrRecordNotFound)
			if err != nil && !isNew {
				return wrapWrite(err)
			}
			row.AccidentCheckTemplateID = it.AccidentCheckTemplateID
			row.CompanyID = companyID
			row.Year = year
			if it.Number != nil {
				row.Number = roundFloat(*it.Number, 1)
			} else {
				row.Number = 0
			}
			row.Comment = it.Comment
			if isNew {
				row.UserID = userID
			}
			row.LastUserID = userID
			if isNew {
				if err := tx.Create(&row).Error; err != nil {
					return wrapWrite(err)
				}
			} else {
				if err := tx.Save(&row).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *AssessmentService) ListAccidentChecks(ctx context.Context, page, perPage int, companyID int64, year int) ([]model.AccidentCheck, int64, error) {
	page, per := normalizePage(page, perPage)
	q := s.db.WithContext(ctx).Model(&model.AccidentCheck{})
	if companyID > 0 {
		q = q.Where("company_id = ?", companyID)
	}
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.AccidentCheck
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AssessmentService) GetAccidentCheck(ctx context.Context, id int64) (*model.AccidentCheck, error) {
	var m model.AccidentCheck
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *AssessmentService) CreateAccidentCheck(ctx context.Context, in *model.AccidentCheck) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateAccidentCheck(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.AccidentCheck{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "事故检查记录不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteAccidentCheck(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.AccidentCheck{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "事故检查记录不存在")
	}
	return nil
}

// ── 加分项（一次性评分，跟事故检查同构） ──────────────────────────

// AddPointsSubmissionRow — 下级企业加分汇总。
type AddPointsSubmissionRow struct {
	CompanyID   int64      `json:"company_id"`
	CompanyName string     `json:"company_name"`
	Year        int        `json:"year"`
	TotalNumber float64    `json:"total_number"`
	RowCount    int        `json:"row_count"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
}

func (s *AssessmentService) ListAddPointsSubmissions(ctx context.Context, pid int64, year int) ([]AddPointsSubmissionRow, error) {
	var subordinateIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("parent_id = ?", pid).Pluck("id", &subordinateIDs).Error; err != nil {
		return nil, err
	}
	if len(subordinateIDs) == 0 {
		return []AddPointsSubmissionRow{}, nil
	}
	q := s.db.WithContext(ctx).Model(&model.AddPoint{}).Where("company_id IN ?", subordinateIDs)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	type aggRow struct {
		CompanyID   int64
		Year        int
		TotalNumber float64
		RowCount    int
		UpdatedAt   *time.Time
	}
	var aggs []aggRow
	if err := q.
		Select(`company_id, "year", COALESCE(SUM(number), 0) AS total_number, COUNT(*) AS row_count, MAX(updated_at) AS updated_at`).
		Group(`company_id, "year"`).
		Order(`"year" DESC`).
		Find(&aggs).Error; err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(aggs))
	for _, a := range aggs {
		ids = append(ids, a.CompanyID)
	}
	titles := s.companyTitles(ctx, ids)
	out := make([]AddPointsSubmissionRow, 0, len(aggs))
	for _, a := range aggs {
		out = append(out, AddPointsSubmissionRow{
			CompanyID:   a.CompanyID,
			CompanyName: titles[a.CompanyID],
			Year:        a.Year,
			TotalNumber: roundFloat(a.TotalNumber, 1),
			RowCount:    a.RowCount,
			UpdatedAt:   a.UpdatedAt,
			Title:       "加分项考核",
		})
	}
	return out, nil
}

// StartAddPointsForm — 跟 daily/accident 不同：Laravel 仅取 level=1 一级模板，不嵌套二级。
// 每个一级模板带 hasOne check（按 company_id + year）。
func (s *AssessmentService) StartAddPointsForm(ctx context.Context, companyID int64, year int) ([]model.AddPointsTemplate, error) {
	var tops []model.AddPointsTemplate
	if err := s.db.WithContext(ctx).
		Where("level = ? AND year = ?", 1, year).
		Order("sort ASC").
		Find(&tops).Error; err != nil {
		return nil, err
	}
	if len(tops) == 0 {
		return tops, nil
	}
	tplIDs := make([]int64, 0, len(tops))
	for _, t := range tops {
		tplIDs = append(tplIDs, t.ID)
	}
	var checks []model.AddPoint
	if err := s.db.WithContext(ctx).
		Where("add_points_template_id IN ? AND company_id = ? AND year = ?", tplIDs, companyID, year).
		Find(&checks).Error; err != nil {
		return nil, err
	}
	checkByTemplate := make(map[int64]*model.AddPoint, len(checks))
	for i := range checks {
		checkByTemplate[checks[i].AddPointsTemplateID] = &checks[i]
	}
	for i := range tops {
		if c, ok := checkByTemplate[tops[i].ID]; ok {
			tops[i].Check = c
		}
	}
	return tops, nil
}

// AddPointsScoreItem — saveScore 单行。
type AddPointsScoreItem struct {
	AddPointsTemplateID int64
	Number              *float64
	Comment             *string
}

func (s *AssessmentService) SaveAddPointsScore(ctx context.Context, companyID int64, year int, userID int64, items []AddPointsScoreItem) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, it := range items {
			if it.AddPointsTemplateID == 0 {
				continue
			}
			var row model.AddPoint
			err := tx.Where("add_points_template_id = ? AND company_id = ? AND year = ?",
				it.AddPointsTemplateID, companyID, year).First(&row).Error
			isNew := errors.Is(err, gorm.ErrRecordNotFound)
			if err != nil && !isNew {
				return wrapWrite(err)
			}
			row.AddPointsTemplateID = it.AddPointsTemplateID
			row.CompanyID = companyID
			row.Year = year
			if it.Number != nil {
				row.Number = roundFloat(*it.Number, 1)
			} else {
				row.Number = 0
			}
			row.Comment = it.Comment
			if isNew {
				row.UserID = userID
			}
			row.LastUserID = userID
			if isNew {
				if err := tx.Create(&row).Error; err != nil {
					return wrapWrite(err)
				}
			} else {
				if err := tx.Save(&row).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *AssessmentService) ListAddPoints(ctx context.Context, page, perPage int, companyID int64, year int) ([]model.AddPoint, int64, error) {
	page, per := normalizePage(page, perPage)
	q := s.db.WithContext(ctx).Model(&model.AddPoint{})
	if companyID > 0 {
		q = q.Where("company_id = ?", companyID)
	}
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.AddPoint
	if err := q.Preload("AddPointsTemplate").Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AssessmentService) GetAddPoint(ctx context.Context, id int64) (*model.AddPoint, error) {
	var m model.AddPoint
	if err := s.db.WithContext(ctx).Preload("AddPointsTemplate").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *AssessmentService) CreateAddPoint(ctx context.Context, in *model.AddPoint) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateAddPoint(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.AddPoint{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "加分项记录不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteAddPoint(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.AddPoint{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "加分项记录不存在")
	}
	return nil
}

// ── 系统检查（活动参与及材料上报） ────────────────────────────────

func (s *AssessmentService) ListSystemChecks(ctx context.Context, page, perPage int, companyID int64, year int) ([]model.SystemCheck, int64, error) {
	page, per := normalizePage(page, perPage)
	q := s.db.WithContext(ctx).Model(&model.SystemCheck{})
	if companyID > 0 {
		q = q.Where("company_id = ?", companyID)
	}
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SystemCheck
	if err := q.
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AssessmentService) GetSystemCheck(ctx context.Context, id int64) (*model.SystemCheck, error) {
	var m model.SystemCheck
	if err := s.db.WithContext(ctx).
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *AssessmentService) CreateSystemCheck(ctx context.Context, in *model.SystemCheck) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateSystemCheck(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.SystemCheck{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "系统考核记录不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteSystemCheck(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.SystemCheck{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "系统考核记录不存在")
	}
	return nil
}

// SystemCheckRow — submissions 行；展平 v1..v5 + total。
type SystemCheckRow struct {
	ID          int64      `json:"id"`
	CompanyID   int64      `json:"company_id"`
	CompanyName string     `json:"company_name"`
	Year        int        `json:"year"`
	Number      float64    `json:"number"`
	NumberV1    float64    `json:"number_v1"`
	NumberV2    float64    `json:"number_v2"`
	NumberV3    float64    `json:"number_v3"`
	NumberV4    float64    `json:"number_v4"`
	NumberV5    float64    `json:"number_v5"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
}

func (s *AssessmentService) ListSystemSubmissions(ctx context.Context, pid int64, year int) ([]SystemCheckRow, error) {
	var subordinateIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("parent_id = ?", pid).Pluck("id", &subordinateIDs).Error; err != nil {
		return nil, err
	}
	if len(subordinateIDs) == 0 {
		return []SystemCheckRow{}, nil
	}
	q := s.db.WithContext(ctx).Model(&model.SystemCheck{}).Where("company_id IN ?", subordinateIDs)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	var rows []model.SystemCheck
	if err := q.
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order(`"year" DESC`).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]SystemCheckRow, 0, len(rows))
	for _, r := range rows {
		ua := r.UpdatedAt
		title := ""
		if r.Company != nil {
			title = r.Company.Title
		}
		out = append(out, SystemCheckRow{
			ID:          r.ID,
			CompanyID:   r.CompanyID,
			CompanyName: title,
			Year:        r.Year,
			Number:      float64(r.Number),
			NumberV1:    float64(r.NumberV1),
			NumberV2:    float64(r.NumberV2),
			NumberV3:    float64(r.NumberV3),
			NumberV4:    float64(r.NumberV4),
			NumberV5:    float64(r.NumberV5),
			UpdatedAt:   &ua,
			Title:       "活动参与及材料上报情况",
		})
	}
	return out, nil
}

// ComputeSystemCheck — 自动算 v1..v5 并写入 system_checks。
//
//   v1 月报覆盖率 = 所辖企业（含本企业）当年覆盖月份数 / 12 * 4
//                  覆盖月份 = month_statistics 表中本企业 + 直接下级当年提交过月报的
//                  distinct 月份数；满覆盖（12 个月）= 4 分。
//   v2 年报、v3 通知反馈：业务停用，固定 0。
//   v4 任务反馈率 = 已反馈 OfficeNotifyScope / 总数 * 4
//   v5 信息稿条数 = min(条数/8, 1) * 2
//
//   number = round(v1+v2+v3+v4+v5)
func (s *AssessmentService) ComputeSystemCheck(ctx context.Context, companyID int64, year int, userID int64) (*model.SystemCheck, error) {
	v1 := 0.0
	{
		// 本企业 + 直接下级；按 created_at 月份去重计数。
		var subs []int64
		if err := s.db.WithContext(ctx).Model(&model.Company{}).
			Where("parent_id = ?", companyID).Pluck("id", &subs).Error; err != nil {
			return nil, err
		}
		ids := append([]int64{companyID}, subs...)
		var distinctMonths int64
		if err := s.db.WithContext(ctx).Model(&model.MonthStatistic{}).
			Where("company_id IN ?", ids).
			Where("EXTRACT(YEAR FROM created_at) = ?", year).
			Select("COUNT(DISTINCT EXTRACT(MONTH FROM created_at))").
			Scan(&distinctMonths).Error; err != nil {
			return nil, err
		}
		if distinctMonths > 12 {
			distinctMonths = 12
		}
		v1 = roundFloat(float64(distinctMonths)/12.0*4.0, 2)
	}
	v2 := 0.0
	v3 := 0.0

	v4 := 0.0
	{
		var total int64
		if err := s.db.WithContext(ctx).Model(&model.OfficeNotifyScope{}).
			Where("company_id = ? AND EXTRACT(YEAR FROM created_at) = ?", companyID, year).
			Count(&total).Error; err != nil {
			return nil, err
		}
		if total > 0 {
			var done int64
			if err := s.db.WithContext(ctx).Model(&model.OfficeNotifyScope{}).
				Where("company_id = ? AND EXTRACT(YEAR FROM created_at) = ? AND feedback_status = ?",
					companyID, year, "已反馈").
				Count(&done).Error; err != nil {
				return nil, err
			}
			v4 = roundFloat(float64(done)/float64(total)*4, 2)
		}
	}

	v5 := 0.0
	{
		var drafts int64
		if err := s.db.WithContext(ctx).Model(&model.InformationDraft{}).
			Where("company_id = ? AND EXTRACT(YEAR FROM created_at) = ?", companyID, year).
			Count(&drafts).Error; err != nil {
			return nil, err
		}
		if drafts >= 8 {
			v5 = 2.0
		} else {
			v5 = roundFloat(float64(drafts)/8*2, 2)
		}
	}

	final := int(math.Round(v1 + v2 + v3 + v4 + v5))

	var row model.SystemCheck
	err := s.db.WithContext(ctx).
		Where("company_id = ? AND year = ?", companyID, year).First(&row).Error
	isNew := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !isNew {
		return nil, err
	}
	row.CompanyID = companyID
	row.Year = year
	row.NumberV1 = float32(v1)
	row.NumberV2 = float32(v2)
	row.NumberV3 = float32(v3)
	row.NumberV4 = float32(v4)
	row.NumberV5 = float32(v5)
	row.Number = final
	if isNew {
		row.UserID = userID
	}
	row.LastUserID = userID
	if isNew {
		if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
			return nil, wrapWrite(err)
		}
	} else {
		if err := s.db.WithContext(ctx).Save(&row).Error; err != nil {
			return nil, wrapWrite(err)
		}
	}
	return &row, nil
}

// ── 标准化考核 ────────────────────────────────────────────────────

func (s *AssessmentService) ListNormChecks(ctx context.Context, page, perPage int, companyID int64, year int, kind string) ([]model.NormCheck, int64, error) {
	page, per := normalizePage(page, perPage)
	q := s.db.WithContext(ctx).Model(&model.NormCheck{})
	if companyID > 0 {
		q = q.Where("company_id = ?", companyID)
	}
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	if kind != "" {
		q = q.Where("type = ?", kind)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.NormCheck
	if err := q.
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AssessmentService) GetNormCheck(ctx context.Context, id int64) (*model.NormCheck, error) {
	var m model.NormCheck
	if err := s.db.WithContext(ctx).
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *AssessmentService) CreateNormCheck(ctx context.Context, in *model.NormCheck) error {
	return wrapWrite(s.db.WithContext(ctx).Create(in).Error)
}

func (s *AssessmentService) UpdateNormCheck(ctx context.Context, id int64, fields map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.NormCheck{}).Where("id = ?", id).Updates(fields)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "标准化检查记录不存在")
	}
	return nil
}

func (s *AssessmentService) DeleteNormCheck(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.NormCheck{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "标准化检查记录不存在")
	}
	return nil
}

// NormCheckRow — submissions 行；按年份折算 raw → final。
type NormCheckRow struct {
	ID          int64      `json:"id"`
	CompanyID   int64      `json:"company_id"`
	CompanyName string     `json:"company_name"`
	Year        int        `json:"year"`
	Type        string     `json:"type"`
	RawNumber   int        `json:"raw_number"`
	FinalNumber float64    `json:"final_number"`
	Path        *string    `json:"path"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
}

func (s *AssessmentService) ListNormSubmissions(ctx context.Context, pid int64, year int) ([]NormCheckRow, error) {
	var subordinateIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("parent_id = ?", pid).Pluck("id", &subordinateIDs).Error; err != nil {
		return nil, err
	}
	if len(subordinateIDs) == 0 {
		return []NormCheckRow{}, nil
	}
	q := s.db.WithContext(ctx).Model(&model.NormCheck{}).Where("company_id IN ?", subordinateIDs)
	if year > 0 {
		q = q.Where("year = ?", year)
	}
	var rows []model.NormCheck
	if err := q.
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order(`"year" DESC`).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]NormCheckRow, 0, len(rows))
	for _, r := range rows {
		scale := 20
		if r.Year > 2025 {
			scale = 30
		}
		final := roundFloat(float64(r.Number)/100*float64(scale), 2)
		ua := r.UpdatedAt
		title := ""
		if r.Company != nil {
			title = r.Company.Title
		}
		out = append(out, NormCheckRow{
			ID:          r.ID,
			CompanyID:   r.CompanyID,
			CompanyName: title,
			Year:        r.Year,
			Type:        r.Type,
			RawNumber:   r.Number,
			FinalNumber: final,
			Path:        r.Path,
			UpdatedAt:   &ua,
			Title:       "大中型企业安全生产标准化管理体系构建",
		})
	}
	return out, nil
}

// SaveNormScore — company+year 唯一，存在则覆盖。
func (s *AssessmentService) SaveNormScore(ctx context.Context, companyID int64, year int, userID int64, kind string, number int, path *string) (*model.NormCheck, error) {
	var row model.NormCheck
	err := s.db.WithContext(ctx).
		Where("company_id = ? AND year = ?", companyID, year).First(&row).Error
	isNew := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !isNew {
		return nil, err
	}
	row.CompanyID = companyID
	row.Year = year
	row.Type = kind
	row.Number = number
	row.Path = path
	if isNew {
		row.UserID = userID
	}
	row.LastUserID = userID
	if isNew {
		if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
			return nil, wrapWrite(err)
		}
	} else {
		if err := s.db.WithContext(ctx).Save(&row).Error; err != nil {
			return nil, wrapWrite(err)
		}
	}
	return &row, nil
}

// NormTemplateLink — 内置打分表下载地址（按 type 路由到 public/demo_check/ 下的 xls/xlsx）。
func NormTemplateLink(kind string) (url string, t string) {
	m := map[string]string{
		"机械": "/demo_check/浙江省机械制造行业安全生产标准化企业考核评定表.xls",
		"造纸": "/demo_check/造纸企业安全生产标准化评定标准.xlsx",
		"危化": "/demo_check/危险化学品从业单位安全生产标准化评审标准.xlsx",
	}
	if u, ok := m[kind]; ok {
		return u, kind
	}
	return "/demo_check/冶金等工贸企业安全生产标准化基本规范评分细则.xlsx", kind
}

// ── 考核模板（年度并集 + 一键拷贝） ───────────────────────────────

// AssessmentYears — 三张模板表已存在的年度并集，从大到小排序。
func (s *AssessmentService) AssessmentYears(ctx context.Context) ([]int, error) {
	years := make(map[int]struct{})
	for _, table := range []string{"daily_check_templates", "accident_check_templates", "add_points_templates"} {
		var ys []int
		if err := s.db.WithContext(ctx).Raw(
			fmt.Sprintf(`SELECT DISTINCT "year" FROM %s WHERE deleted_at IS NULL`, table),
		).Scan(&ys).Error; err != nil {
			return nil, err
		}
		for _, y := range ys {
			if y > 0 {
				years[y] = struct{}{}
			}
		}
	}
	out := make([]int, 0, len(years))
	for y := range years {
		out = append(out, y)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	return out, nil
}

// CopyKindResult — 单类型拷贝结果。
type CopyKindResult struct {
	Copied int `json:"copied"`
}

type CopyResult struct {
	FromYear int                       `json:"from_year"`
	ToYear   int                       `json:"to_year"`
	Summary  map[string]CopyKindResult `json:"summary"`
}

// CopyAssessmentTemplates — 把 from_year 的模板树拷贝到 to_year，保留 pid。
//
//   kinds：可选 daily / accident / add_points 子集，默认全部。
//   overwrite：true 时先清空 to_year，否则若已有任何记录就报错。
func (s *AssessmentService) CopyAssessmentTemplates(ctx context.Context, fromYear, toYear int, kinds []string, overwrite bool) (*CopyResult, error) {
	if fromYear == toYear {
		return nil, errs.New(errs.CodeUnprocessable, "from_year 与 to_year 不能相同")
	}
	if len(kinds) == 0 {
		kinds = []string{"daily", "accident", "add_points"}
	}
	summary := make(map[string]CopyKindResult, len(kinds))

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, kind := range kinds {
			var copied int
			var err error
			switch kind {
			case "daily":
				copied, err = copyDailyCheckTemplates(tx, fromYear, toYear, overwrite)
			case "accident":
				copied, err = copyAccidentCheckTemplates(tx, fromYear, toYear, overwrite)
			case "add_points":
				copied, err = copyAddPointsTemplates(tx, fromYear, toYear, overwrite)
			default:
				return errs.New(errs.CodeUnprocessable, "未知 kind: "+kind)
			}
			if err != nil {
				return err
			}
			summary[kind] = CopyKindResult{Copied: copied}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &CopyResult{FromYear: fromYear, ToYear: toYear, Summary: summary}, nil
}

// 三张模板表的拷贝逻辑虽然形状相同，但 Go 不支持给外部包类型加方法，无法用泛型一次写完。
// 这里直接展开三遍 — 每个 ~25 行，比硬上 reflection / generic interface 清楚。

func copyDailyCheckTemplates(tx *gorm.DB, fromYear, toYear int, overwrite bool) (int, error) {
	var targetCount int64
	if err := tx.Model(&model.DailyCheckTemplate{}).Where("year = ?", toYear).Count(&targetCount).Error; err != nil {
		return 0, err
	}
	if targetCount > 0 && !overwrite {
		return 0, errs.New(errs.CodeUnprocessable,
			fmt.Sprintf("日常检查 目标年度 %d 已存在 %d 条记录，如需覆盖请启用 overwrite", toYear, targetCount))
	}
	if overwrite && targetCount > 0 {
		if err := tx.Unscoped().Where("year = ?", toYear).Delete(&model.DailyCheckTemplate{}).Error; err != nil {
			return 0, wrapWrite(err)
		}
	}
	var sources []model.DailyCheckTemplate
	if err := tx.Where("year = ?", fromYear).Order("level ASC, pid ASC, sort ASC").Find(&sources).Error; err != nil {
		return 0, err
	}
	idMap := map[int64]int64{0: 0}
	for _, src := range sources {
		newPid := idMap[src.Pid]
		dup := src
		dup.ID = 0
		dup.Pid = newPid
		dup.Year = toYear
		if err := tx.Create(&dup).Error; err != nil {
			return 0, wrapWrite(err)
		}
		idMap[src.ID] = dup.ID
	}
	return len(sources), nil
}

func copyAccidentCheckTemplates(tx *gorm.DB, fromYear, toYear int, overwrite bool) (int, error) {
	var targetCount int64
	if err := tx.Model(&model.AccidentCheckTemplate{}).Where("year = ?", toYear).Count(&targetCount).Error; err != nil {
		return 0, err
	}
	if targetCount > 0 && !overwrite {
		return 0, errs.New(errs.CodeUnprocessable,
			fmt.Sprintf("事故检查 目标年度 %d 已存在 %d 条记录，如需覆盖请启用 overwrite", toYear, targetCount))
	}
	if overwrite && targetCount > 0 {
		if err := tx.Unscoped().Where("year = ?", toYear).Delete(&model.AccidentCheckTemplate{}).Error; err != nil {
			return 0, wrapWrite(err)
		}
	}
	var sources []model.AccidentCheckTemplate
	if err := tx.Where("year = ?", fromYear).Order("level ASC, pid ASC, sort ASC").Find(&sources).Error; err != nil {
		return 0, err
	}
	idMap := map[int64]int64{0: 0}
	for _, src := range sources {
		newPid := idMap[src.Pid]
		dup := src
		dup.ID = 0
		dup.Pid = newPid
		dup.Year = toYear
		if err := tx.Create(&dup).Error; err != nil {
			return 0, wrapWrite(err)
		}
		idMap[src.ID] = dup.ID
	}
	return len(sources), nil
}

func copyAddPointsTemplates(tx *gorm.DB, fromYear, toYear int, overwrite bool) (int, error) {
	var targetCount int64
	if err := tx.Model(&model.AddPointsTemplate{}).Where("year = ?", toYear).Count(&targetCount).Error; err != nil {
		return 0, err
	}
	if targetCount > 0 && !overwrite {
		return 0, errs.New(errs.CodeUnprocessable,
			fmt.Sprintf("加分项 目标年度 %d 已存在 %d 条记录，如需覆盖请启用 overwrite", toYear, targetCount))
	}
	if overwrite && targetCount > 0 {
		if err := tx.Unscoped().Where("year = ?", toYear).Delete(&model.AddPointsTemplate{}).Error; err != nil {
			return 0, wrapWrite(err)
		}
	}
	var sources []model.AddPointsTemplate
	if err := tx.Where("year = ?", fromYear).Order("level ASC, pid ASC, sort ASC").Find(&sources).Error; err != nil {
		return 0, err
	}
	idMap := map[int64]int64{0: 0}
	for _, src := range sources {
		newPid := idMap[src.Pid]
		dup := src
		dup.ID = 0
		dup.Pid = newPid
		dup.Year = toYear
		if err := tx.Create(&dup).Error; err != nil {
			return 0, wrapWrite(err)
		}
		idMap[src.ID] = dup.ID
	}
	return len(sources), nil
}

// ── 考核汇总（6 维度） ────────────────────────────────────────────

// AssessmentSummaryRow — summary 接口的输出形状。
type AssessmentSummaryRow struct {
	CompanyID   int64   `json:"company_id"`
	CompanyName string  `json:"company_name"`
	Year        int     `json:"year"`
	Type        string  `json:"type"`
	Number1     float64 `json:"number1"`
	Number2     float64 `json:"number2"`
	Number3     float64 `json:"number3"`
	Number4     float64 `json:"number4"`
	Number5     float64 `json:"number5"`
	Number6     float64 `json:"number6"`
	Number      float64 `json:"number"`
}

// computeCompanyScore — 算单个企业 6 维度总分。double_check 表暂未迁移，year ≤ 2025
// 的 number3 固定 0；2026+ 业务上也已废除。
func (s *AssessmentService) computeCompanyScore(ctx context.Context, companyID int64, year int) (AssessmentSummaryRow, error) {
	row := AssessmentSummaryRow{
		CompanyID: companyID,
		Year:      year,
		Type:      "暂无",
	}

	// number1 — 日常目标管理工作（直接求和）
	var n1 float64
	if err := s.db.WithContext(ctx).Model(&model.DailyCheck{}).
		Where("company_id = ? AND year = ?", companyID, year).
		Select("COALESCE(SUM(number), 0)").Scan(&n1).Error; err != nil {
		return row, err
	}
	row.Number1 = roundFloat(n1, 2)

	// number2 — 安全生产标准化（year>2025 系数 30，否则 20）
	var norm model.NormCheck
	err := s.db.WithContext(ctx).
		Where("company_id = ? AND year = ?", companyID, year).First(&norm).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return row, err
	}
	if err == nil {
		row.Type = norm.Type
		scale := 20.0
		if year > 2025 {
			scale = 30.0
		}
		row.Number2 = roundFloat(float64(norm.Number)/100*scale, 2)
	}

	// number3 — 双重预防 (year>2025 取消)，且 double_checks 表尚未迁；保持 0。
	row.Number3 = 0.0

	// number4 — 活动参与及材料上报（system_checks 总分）
	var n4 float64
	if err := s.db.WithContext(ctx).Model(&model.SystemCheck{}).
		Where("company_id = ? AND year = ?", companyID, year).
		Select("COALESCE(SUM(number), 0)").Scan(&n4).Error; err != nil {
		return row, err
	}
	row.Number4 = roundFloat(n4, 2)

	// number5 — 事故扣分；20 - SUM(number)，下限 0。
	var accCount int64
	if err := s.db.WithContext(ctx).Model(&model.AccidentCheck{}).
		Where("company_id = ? AND year = ?", companyID, year).Count(&accCount).Error; err != nil {
		return row, err
	}
	if accCount > 0 {
		var raw float64
		if err := s.db.WithContext(ctx).Model(&model.AccidentCheck{}).
			Where("company_id = ? AND year = ?", companyID, year).
			Select("COALESCE(SUM(number), 0)").Scan(&raw).Error; err != nil {
			return row, err
		}
		if raw >= 20 {
			row.Number5 = 0
		} else {
			row.Number5 = roundFloat(20-raw, 2)
		}
	}

	// number6 — 加分项
	var n6 float64
	if err := s.db.WithContext(ctx).Model(&model.AddPoint{}).
		Where("company_id = ? AND year = ?", companyID, year).
		Select("COALESCE(SUM(number), 0)").Scan(&n6).Error; err != nil {
		return row, err
	}
	row.Number6 = roundFloat(n6, 2)

	row.Number = roundFloat(row.Number1+row.Number2+row.Number3+row.Number4+row.Number5+row.Number6, 2)

	// 取 company.title。
	var title string
	_ = s.db.WithContext(ctx).Model(&model.Company{}).
		Where("id = ?", companyID).Select("title").Scan(&title).Error
	row.CompanyName = title

	return row, nil
}

// SummaryBloc — 集团视图：返回 pid 下所有 status=1 子企业的 6 维度评分。
func (s *AssessmentService) SummaryBloc(ctx context.Context, pid int64, year int) ([]AssessmentSummaryRow, error) {
	var subIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("parent_id = ? AND status = ?", pid, 1).Pluck("id", &subIDs).Error; err != nil {
		return nil, err
	}
	out := make([]AssessmentSummaryRow, 0, len(subIDs))
	for _, id := range subIDs {
		r, err := s.computeCompanyScore(ctx, id, year)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, nil
}

// SummaryCompany — 单企业视图。
func (s *AssessmentService) SummaryCompany(ctx context.Context, companyID int64, year int) (*AssessmentSummaryRow, error) {
	r, err := s.computeCompanyScore(ctx, companyID, year)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// CompanyTitle — 取单企业 title（导出文件名用）。
func (s *AssessmentService) CompanyTitle(ctx context.Context, id int64) string {
	var title string
	_ = s.db.WithContext(ctx).Model(&model.Company{}).
		Where("id = ?", id).Select("title").Scan(&title).Error
	return title
}

// ── 内部 helpers ──────────────────────────────────────────────────

// companyTitles — 给一组 ID 批量取 title。
func (s *AssessmentService) companyTitles(ctx context.Context, ids []int64) map[int64]string {
	if len(ids) == 0 {
		return map[int64]string{}
	}
	type row struct {
		ID    int64
		Title string
	}
	var titles []row
	_ = s.db.WithContext(ctx).Model(&model.Company{}).
		Where("id IN ?", ids).Select("id, title").Scan(&titles).Error
	m := make(map[int64]string, len(titles))
	for _, t := range titles {
		m[t.ID] = t.Title
	}
	return m
}

func roundFloat(v float64, digits int) float64 {
	shift := math.Pow(10, float64(digits))
	return math.Round(v*shift) / shift
}
