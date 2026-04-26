package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// ProcessService 聚合工艺信息三个模块的 CRUD：
//   - major_hazardous_processes  重点监管危险工艺
//   - hazardous_processes        较大危险工艺
//   - space_works                有限空间辨识
//
// 三个模块各自独立、无子表，写入语义和 SafetyLedgerService.SpecialEquipment
// 一致：主表 partial update + 软删除。共享一个 service 是为了避免在 router 里
// 拉三个零行为的构造函数。
type ProcessService struct {
	db *gorm.DB
}

func NewProcessService(db *gorm.DB) *ProcessService {
	return &ProcessService{db: db}
}

// ── 重点监管危险工艺 ────────────────────────────────────────────────────

// MajorHazardousProcessListParams — 前端搜索字段是 name（key-hazardous/Index.tsx
// onSearchChange 写的是 name）。Keyword 兼容通用调用。
type MajorHazardousProcessListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *ProcessService) ListMajorHazardousProcess(ctx context.Context, p MajorHazardousProcessListParams) ([]model.MajorHazardousProcess, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.MajorHazardousProcess{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.MajorHazardousProcess
	if err := q.Preload("DangerousCraftType").
		Preload("HazardousChemical").
		Preload("DangerousEquipment").
		Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ProcessService) GetMajorHazardousProcess(ctx context.Context, id int64) (*model.MajorHazardousProcess, error) {
	var m model.MajorHazardousProcess
	if err := s.db.WithContext(ctx).
		Preload("DangerousCraftType").
		Preload("HazardousChemical").
		Preload("DangerousEquipment").
		First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ProcessService) CreateMajorHazardousProcess(ctx context.Context, m *model.MajorHazardousProcess) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ProcessService) UpdateMajorHazardousProcess(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.MajorHazardousProcess{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ProcessService) DeleteMajorHazardousProcess(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.MajorHazardousProcess{}, id).Error
}

// ── 较大危险工艺 ───────────────────────────────────────────────────────

// HazardousProcessListParams — 前端搜索字段是 position
// （hazardous/Index.tsx onSearchChange 写的是 position）。
type HazardousProcessListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Position  string
}

func (s *ProcessService) ListHazardousProcess(ctx context.Context, p HazardousProcessListParams) ([]model.HazardousProcess, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.HazardousProcess{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Position); kw != "" {
		q = q.Where("position ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HazardousProcess
	if err := q.Preload("ProcessType").
		Preload("HazardousChemical").
		Preload("DangerousEquipment").
		Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ProcessService) GetHazardousProcess(ctx context.Context, id int64) (*model.HazardousProcess, error) {
	var m model.HazardousProcess
	if err := s.db.WithContext(ctx).
		Preload("ProcessType").
		Preload("HazardousChemical").
		Preload("DangerousEquipment").
		First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ProcessService) CreateHazardousProcess(ctx context.Context, m *model.HazardousProcess) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ProcessService) UpdateHazardousProcess(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.HazardousProcess{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ProcessService) DeleteHazardousProcess(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.HazardousProcess{}, id).Error
}

// ── 有限空间辨识 ───────────────────────────────────────────────────────

// SpaceWorkListParams — 前端没显式搜索框，但保留 Name 字段以备后续接入。
type SpaceWorkListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *ProcessService) ListSpaceWork(ctx context.Context, p SpaceWorkListParams) ([]model.SpaceWork, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.SpaceWork{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SpaceWork
	if err := q.Preload("SpaceWorkType").
		Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ProcessService) GetSpaceWork(ctx context.Context, id int64) (*model.SpaceWork, error) {
	var m model.SpaceWork
	if err := s.db.WithContext(ctx).Preload("SpaceWorkType").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ProcessService) CreateSpaceWork(ctx context.Context, m *model.SpaceWork) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ProcessService) UpdateSpaceWork(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.SpaceWork{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ProcessService) DeleteSpaceWork(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.SpaceWork{}, id).Error
}
