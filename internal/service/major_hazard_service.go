package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// MajorHazardService — 重大危险源（major_hazards + major_constitutes）。
//
// 与 SafetyLedgerService.HazardousChemical 的写入语义保持一致：写主表 + 整体替换
// 子表（前端每次提交都把全量 constitutes 数组发上来），事务内保证两者一致。
type MajorHazardService struct {
	db *gorm.DB
}

func NewMajorHazardService(db *gorm.DB) *MajorHazardService {
	return &MajorHazardService{db: db}
}

// MajorHazardListParams 复用 ListParams 的语义但用 Name（前端搜索字段）。
type MajorHazardListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *MajorHazardService) List(ctx context.Context, p MajorHazardListParams) ([]model.MajorHazard, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.MajorHazard{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.MajorHazard
	if err := q.Preload("MajorConstitutes").
		Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *MajorHazardService) Get(ctx context.Context, id int64) (*model.MajorHazard, error) {
	var m model.MajorHazard
	if err := s.db.WithContext(ctx).Preload("MajorConstitutes").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

// CreateMajorHazardInput — handler 装配后传入。Constitutes 的 MajorHazardID/CompanyID
// 在 service 内统一赋值，handler 不需要操心。
type CreateMajorHazardInput struct {
	Hazard      *model.MajorHazard
	Constitutes []model.MajorConstitute
}

func (s *MajorHazardService) Create(ctx context.Context, in CreateMajorHazardInput) error {
	if in.Hazard.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Hazard).Error; err != nil {
			return wrapWrite(err)
		}
		for i := range in.Constitutes {
			in.Constitutes[i].MajorHazardID = in.Hazard.ID
			in.Constitutes[i].CompanyID = in.Hazard.CompanyID
		}
		if len(in.Constitutes) > 0 {
			if err := tx.Create(&in.Constitutes).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

// Update — 主表 partial update + 子表整体替换（仅当 constitutes != nil 时替换）。
func (s *MajorHazardService) Update(ctx context.Context, id int64, updates map[string]any, constitutes []model.MajorConstitute) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.MajorHazard{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		if constitutes != nil {
			if err := tx.Where("major_hazard_id = ?", id).Delete(&model.MajorConstitute{}).Error; err != nil {
				return err
			}
			if len(constitutes) > 0 {
				var hazard model.MajorHazard
				if err := tx.First(&hazard, id).Error; err != nil {
					return err
				}
				for i := range constitutes {
					constitutes[i].MajorHazardID = id
					constitutes[i].CompanyID = hazard.CompanyID
				}
				if err := tx.Create(&constitutes).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *MajorHazardService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.MajorHazard{}, id).Error
}
