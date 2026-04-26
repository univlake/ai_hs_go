package service

import (
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

type SafetyLedgerService struct {
	db *gorm.DB
}

func NewSafetyLedgerService(db *gorm.DB) *SafetyLedgerService {
	return &SafetyLedgerService{db: db}
}

type ListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Keyword   string
}

func paginate(p *ListParams) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}
}

// ----- SpecialEquipment -----

func (s *SafetyLedgerService) ListSpecialEquipment(ctx context.Context, p ListParams) ([]model.SpecialEquipment, int64, error) {
	paginate(&p)
	q := s.db.WithContext(ctx).Model(&model.SpecialEquipment{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Keyword); kw != "" {
		like := "%" + kw + "%"
		q = q.Where("equipment_name ILIKE ? OR serial_number ILIKE ?", like, like)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SpecialEquipment
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *SafetyLedgerService) CreateSpecialEquipment(ctx context.Context, m *model.SpecialEquipment) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *SafetyLedgerService) UpdateSpecialEquipment(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.SpecialEquipment{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *SafetyLedgerService) DeleteSpecialEquipment(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.SpecialEquipment{}, id).Error
}

func (s *SafetyLedgerService) GetSpecialEquipment(ctx context.Context, id int64) (*model.SpecialEquipment, error) {
	var m model.SpecialEquipment
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

// CopySpecialEquipment 批量复制选中行（去掉 id / 时间戳，保留业务字段 + company_id）。
func (s *SafetyLedgerService) CopySpecialEquipment(ctx context.Context, ids []int64, uid int64) ([]int64, error) {
	if len(ids) == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 ids")
	}
	var originals []model.SpecialEquipment
	if err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&originals).Error; err != nil {
		return nil, err
	}
	if len(originals) == 0 {
		return nil, errs.New(errs.CodeNotFound, "未找到指定记录")
	}
	newIDs := make([]int64, 0, len(originals))
	for _, o := range originals {
		clone := model.SpecialEquipment{
			EquipmentName:    o.EquipmentName,
			SpecialTypeID:    o.SpecialTypeID,
			UtilizedLocation: o.UtilizedLocation,
			NormModel1:       o.NormModel1,
			NormModel2:       o.NormModel2,
			NormModel3:       o.NormModel3,
			NormModel4:       o.NormModel4,
			SerialNumber:     o.SerialNumber,
			UseTime:          o.UseTime,
			ValidPeriod:      o.ValidPeriod,
			InspectUnit:      o.InspectUnit,
			MaintainInfo:     o.MaintainInfo,
			Comment:          o.Comment,
			CompanyID:        o.CompanyID,
		}
		clone.UserID = uid
		clone.LastUserID = uid
		if err := s.db.WithContext(ctx).Create(&clone).Error; err != nil {
			return nil, wrapWrite(err)
		}
		newIDs = append(newIDs, clone.ID)
	}
	return newIDs, nil
}

// ----- DangerousEquipment -----

func (s *SafetyLedgerService) ListDangerousEquipment(ctx context.Context, p ListParams) ([]model.DangerousEquipment, int64, error) {
	paginate(&p)
	q := s.db.WithContext(ctx).Model(&model.DangerousEquipment{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Keyword); kw != "" {
		like := "%" + kw + "%"
		q = q.Where("equipment_name ILIKE ? OR serial_number ILIKE ?", like, like)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.DangerousEquipment
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *SafetyLedgerService) CreateDangerousEquipment(ctx context.Context, m *model.DangerousEquipment) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *SafetyLedgerService) UpdateDangerousEquipment(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.DangerousEquipment{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *SafetyLedgerService) DeleteDangerousEquipment(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.DangerousEquipment{}, id).Error
}

func (s *SafetyLedgerService) GetDangerousEquipment(ctx context.Context, id int64) (*model.DangerousEquipment, error) {
	var m model.DangerousEquipment
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

// CopyDangerousEquipment 批量复制选中行。
func (s *SafetyLedgerService) CopyDangerousEquipment(ctx context.Context, ids []int64, uid int64) ([]int64, error) {
	if len(ids) == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 ids")
	}
	var originals []model.DangerousEquipment
	if err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&originals).Error; err != nil {
		return nil, err
	}
	if len(originals) == 0 {
		return nil, errs.New(errs.CodeNotFound, "未找到指定记录")
	}
	newIDs := make([]int64, 0, len(originals))
	for _, o := range originals {
		clone := model.DangerousEquipment{
			EquipmentName:    o.EquipmentName,
			UtilizedLocation: o.UtilizedLocation,
			NormModel:        o.NormModel,
			SerialNumber:     o.SerialNumber,
			UseTime:          o.UseTime,
			AutomationType:   o.AutomationType,
			MaintainInfo:     o.MaintainInfo,
			Comment:          o.Comment,
			CompanyID:        o.CompanyID,
		}
		clone.UserID = uid
		clone.LastUserID = uid
		if err := s.db.WithContext(ctx).Create(&clone).Error; err != nil {
			return nil, wrapWrite(err)
		}
		newIDs = append(newIDs, clone.ID)
	}
	return newIDs, nil
}

// ----- HazardousChemical -----

func (s *SafetyLedgerService) ListHazardousChemicals(ctx context.Context, p ListParams) ([]model.HazardousChemical, int64, error) {
	paginate(&p)
	q := s.db.WithContext(ctx).Model(&model.HazardousChemical{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Keyword); kw != "" {
		like := "%" + kw + "%"
		q = q.Where("name ILIKE ? OR alias ILIKE ? OR cas ILIKE ?", like, like, like)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HazardousChemical
	if err := q.Preload("Ingredients").Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// CreateHazardousChemicalInput is split out so the handler can pass nested
// ingredients in a single payload.
type CreateHazardousChemicalInput struct {
	Chemical    *model.HazardousChemical
	Ingredients []model.ChemicalIngredient
}

func (s *SafetyLedgerService) CreateHazardousChemical(ctx context.Context, in CreateHazardousChemicalInput) error {
	if in.Chemical.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Chemical).Error; err != nil {
			return wrapWrite(err)
		}
		for i := range in.Ingredients {
			in.Ingredients[i].HazardousChemicalID = in.Chemical.ID
			in.Ingredients[i].CompanyID = in.Chemical.CompanyID
		}
		if len(in.Ingredients) > 0 {
			if err := tx.Create(&in.Ingredients).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

func (s *SafetyLedgerService) UpdateHazardousChemical(ctx context.Context, id int64, updates map[string]any, ingredients []model.ChemicalIngredient) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.HazardousChemical{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		if ingredients != nil {
			if err := tx.Where("hazardous_chemical_id = ?", id).Delete(&model.ChemicalIngredient{}).Error; err != nil {
				return err
			}
			if len(ingredients) > 0 {
				var chem model.HazardousChemical
				if err := tx.First(&chem, id).Error; err != nil {
					return err
				}
				for i := range ingredients {
					ingredients[i].HazardousChemicalID = id
					ingredients[i].CompanyID = chem.CompanyID
				}
				if err := tx.Create(&ingredients).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *SafetyLedgerService) DeleteHazardousChemical(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.HazardousChemical{}, id).Error
}

func (s *SafetyLedgerService) GetHazardousChemical(ctx context.Context, id int64) (*model.HazardousChemical, error) {
	var m model.HazardousChemical
	if err := s.db.WithContext(ctx).Preload("Ingredients").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func wrapWrite(err error) error {
	if err == nil {
		return nil
	}
	return errs.Wrap(errs.CodeInternal, "写入失败", err)
}

func mapNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return err
}
