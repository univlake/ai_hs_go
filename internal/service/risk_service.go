package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// RiskService 聚合风险管理 5 个模块的 CRUD：
//
//   - risk_distribution_maps      风险源管理（per-company singleton）
//   - monthly_investigations      隐患管理主表（含 reviews/views 子表整体替换）
//   - job_management              作业管理
//   - three_togethers             三同时管理
//   - assess_detects              安全评价与检测
//
// 主表 partial update，singleton 走 First/Create/Update。
// MonthlyInvestigation 走 transaction，子表先删后插，Laravel 原写法保留。
type RiskService struct {
	db *gorm.DB
}

func NewRiskService(db *gorm.DB) *RiskService {
	return &RiskService{db: db}
}

// ── 风险分布图（per-company singleton） ─────────────────────────────────

func (s *RiskService) FirstRiskDistributionMap(ctx context.Context, companyID int64) (*model.RiskDistributionMap, error) {
	var m model.RiskDistributionMap
	err := s.db.WithContext(ctx).Where("company_id = ?", companyID).Order("id DESC").First(&m).Error
	if err == nil {
		return &m, nil
	}
	if ae, ok := errs.As(mapNotFound(err)); ok && ae.Code == errs.CodeNotFound {
		return nil, nil
	}
	return nil, err
}

func (s *RiskService) CreateRiskDistributionMap(ctx context.Context, m *model.RiskDistributionMap) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *RiskService) UpdateRiskDistributionMap(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.RiskDistributionMap{}, id, updates)
}

// ── 隐患管理（主表 + 整改/复查附件子表） ─────────────────────────────────

type MonthlyInvestigationListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Type      string
}

func (s *RiskService) ListMonthlyInvestigation(ctx context.Context, p MonthlyInvestigationListParams) ([]model.MonthlyInvestigation, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.MonthlyInvestigation{}).Where("company_id = ?", p.CompanyID)
	if v := strings.TrimSpace(p.Type); v != "" {
		q = q.Where("type = ?", v)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.MonthlyInvestigation
	if err := q.Preload("Reviews").Preload("Views").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *RiskService) GetMonthlyInvestigation(ctx context.Context, id int64) (*model.MonthlyInvestigation, error) {
	var m model.MonthlyInvestigation
	if err := s.db.WithContext(ctx).Preload("Reviews").Preload("Views").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type MonthlyInvestigationInput struct {
	Investigation *model.MonthlyInvestigation
	Reviews       []model.MonthlyInvestigationReview
	Views         []model.MonthlyInvestigationView
}

func (s *RiskService) CreateMonthlyInvestigation(ctx context.Context, in MonthlyInvestigationInput) error {
	if in.Investigation.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Investigation).Error; err != nil {
			return wrapWrite(err)
		}
		uid := in.Investigation.UserID
		cid := in.Investigation.CompanyID
		for i := range in.Reviews {
			in.Reviews[i].MonthlyInvestigationID = in.Investigation.ID
			in.Reviews[i].CompanyID = cid
			in.Reviews[i].UserID = uid
			in.Reviews[i].LastUserID = uid
		}
		if len(in.Reviews) > 0 {
			if err := tx.Create(&in.Reviews).Error; err != nil {
				return wrapWrite(err)
			}
		}
		for i := range in.Views {
			in.Views[i].MonthlyInvestigationID = in.Investigation.ID
			in.Views[i].CompanyID = cid
			in.Views[i].UserID = uid
			in.Views[i].LastUserID = uid
		}
		if len(in.Views) > 0 {
			if err := tx.Create(&in.Views).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

// UpdateMonthlyInvestigation — 子表 nil 表示不动；非 nil（含空切片）表示整体替换。
func (s *RiskService) UpdateMonthlyInvestigation(
	ctx context.Context,
	id int64,
	updates map[string]any,
	reviews []model.MonthlyInvestigationReview,
	views []model.MonthlyInvestigationView,
	replaceReviews, replaceViews bool,
	uid int64,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.MonthlyInvestigation{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		var inv model.MonthlyInvestigation
		if err := tx.First(&inv, id).Error; err != nil {
			return err
		}
		cid := inv.CompanyID
		if replaceReviews {
			if err := tx.Where("monthly_investigation_id = ?", id).Delete(&model.MonthlyInvestigationReview{}).Error; err != nil {
				return err
			}
			if len(reviews) > 0 {
				for i := range reviews {
					reviews[i].MonthlyInvestigationID = id
					reviews[i].CompanyID = cid
					reviews[i].UserID = uid
					reviews[i].LastUserID = uid
				}
				if err := tx.Create(&reviews).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		if replaceViews {
			if err := tx.Where("monthly_investigation_id = ?", id).Delete(&model.MonthlyInvestigationView{}).Error; err != nil {
				return err
			}
			if len(views) > 0 {
				for i := range views {
					views[i].MonthlyInvestigationID = id
					views[i].CompanyID = cid
					views[i].UserID = uid
					views[i].LastUserID = uid
				}
				if err := tx.Create(&views).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *RiskService) DeleteMonthlyInvestigation(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("monthly_investigation_id = ?", id).Delete(&model.MonthlyInvestigationReview{}).Error; err != nil {
			return err
		}
		if err := tx.Where("monthly_investigation_id = ?", id).Delete(&model.MonthlyInvestigationView{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.MonthlyInvestigation{}, id).Error
	})
}

// ── 作业管理 ───────────────────────────────────────────────────────────

type JobManagementListParams struct {
	Page          int
	PerPage       int
	CompanyID     int64
	Type          string
	IsOutsourcing string
	JobTime       string
}

func (s *RiskService) ListJobManagement(ctx context.Context, p JobManagementListParams) ([]model.JobManagement, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.JobManagement{}).Where("company_id = ?", p.CompanyID)
	if v := strings.TrimSpace(p.Type); v != "" {
		q = q.Where("type = ?", v)
	}
	if v := strings.TrimSpace(p.IsOutsourcing); v != "" {
		q = q.Where("is_outsourcing = ?", v)
	}
	if v := strings.TrimSpace(p.JobTime); v != "" {
		q = q.Where("job_time = ?", v)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.JobManagement
	if err := q.Preload("EvaluationUnit").Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *RiskService) GetJobManagement(ctx context.Context, id int64) (*model.JobManagement, error) {
	var m model.JobManagement
	if err := s.db.WithContext(ctx).Preload("EvaluationUnit").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *RiskService) CreateJobManagement(ctx context.Context, m *model.JobManagement) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *RiskService) UpdateJobManagement(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.JobManagement{}, id, updates)
}

func (s *RiskService) DeleteJobManagement(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.JobManagement{}, id).Error
}

// ── 三同时管理 ─────────────────────────────────────────────────────────

type ThreeTogetherListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *RiskService) ListThreeTogether(ctx context.Context, p ThreeTogetherListParams) ([]model.ThreeTogether, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ThreeTogether{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ThreeTogether
	if err := q.Preload("DesignUnit").Preload("EvaluationUnit").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *RiskService) GetThreeTogether(ctx context.Context, id int64) (*model.ThreeTogether, error) {
	var m model.ThreeTogether
	if err := s.db.WithContext(ctx).Preload("DesignUnit").Preload("EvaluationUnit").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *RiskService) CreateThreeTogether(ctx context.Context, m *model.ThreeTogether) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *RiskService) UpdateThreeTogether(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.ThreeTogether{}, id, updates)
}

func (s *RiskService) DeleteThreeTogether(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.ThreeTogether{}, id).Error
}

// ── 安全评价与检测 ─────────────────────────────────────────────────────

type AssessDetectListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *RiskService) ListAssessDetect(ctx context.Context, p AssessDetectListParams) ([]model.AssessDetect, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.AssessDetect{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.AssessDetect
	if err := q.Preload("SkillUnit").Preload("SkillUnitKid").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *RiskService) GetAssessDetect(ctx context.Context, id int64) (*model.AssessDetect, error) {
	var m model.AssessDetect
	if err := s.db.WithContext(ctx).Preload("SkillUnit").Preload("SkillUnitKid").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *RiskService) CreateAssessDetect(ctx context.Context, m *model.AssessDetect) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *RiskService) UpdateAssessDetect(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.AssessDetect{}, id, updates)
}

func (s *RiskService) DeleteAssessDetect(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.AssessDetect{}, id).Error
}
