package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// ThirdPartyService 聚合第三方管理三个模块的 CRUD：
//   - design_units      设计单位
//   - evaluation_units  施工/评价单位
//   - skill_units (+ skill_unit_kids) 技术服务单位（带子表）
//
// 与 ProcessService / MajorHazardService 一致：主表 partial update，
// SkillUnit.kids 整体替换（前端每次提交都把全量 kids 数组发上来）。
type ThirdPartyService struct {
	db *gorm.DB
}

func NewThirdPartyService(db *gorm.DB) *ThirdPartyService {
	return &ThirdPartyService{db: db}
}

// ── 设计单位 ────────────────────────────────────────────────────────────

type DesignUnitListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *ThirdPartyService) ListDesignUnit(ctx context.Context, p DesignUnitListParams) ([]model.DesignUnit, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.DesignUnit{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.DesignUnit
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ThirdPartyService) GetDesignUnit(ctx context.Context, id int64) (*model.DesignUnit, error) {
	var m model.DesignUnit
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ThirdPartyService) CreateDesignUnit(ctx context.Context, m *model.DesignUnit) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ThirdPartyService) UpdateDesignUnit(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.DesignUnit{}, id, updates)
}

func (s *ThirdPartyService) DeleteDesignUnit(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.DesignUnit{}, id).Error
}

// ── 评价单位 ────────────────────────────────────────────────────────────

type EvaluationUnitListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *ThirdPartyService) ListEvaluationUnit(ctx context.Context, p EvaluationUnitListParams) ([]model.EvaluationUnit, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.EvaluationUnit{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.EvaluationUnit
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ThirdPartyService) GetEvaluationUnit(ctx context.Context, id int64) (*model.EvaluationUnit, error) {
	var m model.EvaluationUnit
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ThirdPartyService) CreateEvaluationUnit(ctx context.Context, m *model.EvaluationUnit) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ThirdPartyService) UpdateEvaluationUnit(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.EvaluationUnit{}, id, updates)
}

func (s *ThirdPartyService) DeleteEvaluationUnit(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.EvaluationUnit{}, id).Error
}

// ── 技术服务单位（含子表） ─────────────────────────────────────────────

type SkillUnitListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *ThirdPartyService) ListSkillUnit(ctx context.Context, p SkillUnitListParams) ([]model.SkillUnit, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.SkillUnit{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SkillUnit
	if err := q.Preload("Kids").Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ThirdPartyService) GetSkillUnit(ctx context.Context, id int64) (*model.SkillUnit, error) {
	var m model.SkillUnit
	if err := s.db.WithContext(ctx).Preload("Kids").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

// CreateSkillUnitInput — handler 装配后传入。Kids 的 SkillUnitID/CompanyID
// 在 service 内统一赋值。
type CreateSkillUnitInput struct {
	Unit *model.SkillUnit
	Kids []model.SkillUnitKid
}

func (s *ThirdPartyService) CreateSkillUnit(ctx context.Context, in CreateSkillUnitInput) error {
	if in.Unit.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Unit).Error; err != nil {
			return wrapWrite(err)
		}
		for i := range in.Kids {
			in.Kids[i].SkillUnitID = in.Unit.ID
			in.Kids[i].CompanyID = in.Unit.CompanyID
		}
		if len(in.Kids) > 0 {
			if err := tx.Create(&in.Kids).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

// UpdateSkillUnit — 主表 partial update + 子表整体替换（仅当 kids != nil 时替换）。
func (s *ThirdPartyService) UpdateSkillUnit(ctx context.Context, id int64, updates map[string]any, kids []model.SkillUnitKid) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.SkillUnit{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		if kids != nil {
			if err := tx.Where("skill_unit_id = ?", id).Delete(&model.SkillUnitKid{}).Error; err != nil {
				return err
			}
			if len(kids) > 0 {
				var unit model.SkillUnit
				if err := tx.First(&unit, id).Error; err != nil {
					return err
				}
				for i := range kids {
					kids[i].SkillUnitID = id
					kids[i].CompanyID = unit.CompanyID
				}
				if err := tx.Create(&kids).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *ThirdPartyService) DeleteSkillUnit(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("skill_unit_id = ?", id).Delete(&model.SkillUnitKid{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.SkillUnit{}, id).Error
	})
}

// ── helpers ─────────────────────────────────────────────────────────────

func normalizePage(page, per int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if per <= 0 || per > 100 {
		per = 20
	}
	return page, per
}

// updateRow — 单表 partial update + 行存在性检查，复用给所有不带子表的模块。
func updateRow(db *gorm.DB, modelPtr any, id int64, updates map[string]any) error {
	res := db.Model(modelPtr).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}
