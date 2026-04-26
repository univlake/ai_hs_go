package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// HealthFireService 聚合职业健康（5 模块 + 1 小模块）+ 消防安全（3 模块）的 CRUD：
//
//   - occupational_hazards    职业危害因素清单
//   - health_systems          职业健康体系（per-company singleton）
//   - health_files            健康档案
//   - health_reports          职业病报告
//   - health_checkups         职业健康体检
//   - hygiene_evaluations     职业卫生评价（关联 skill_units / skill_unit_kids）
//   - fire_safeties           消防安全（关联 skill_units / skill_unit_kids）
//   - fire_equipment          消防器材（含 copy 批量复制）
//   - fire_facilities         消防设施（含 copy 批量复制）
//
// 主表 partial update，singleton 走 First/Create/Update。
type HealthFireService struct {
	db *gorm.DB
}

func NewHealthFireService(db *gorm.DB) *HealthFireService {
	return &HealthFireService{db: db}
}

// ── 职业危害清单 ───────────────────────────────────────────────────────

type OccupationalHazardListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
	IsDeclare string
}

func (s *HealthFireService) ListOccupationalHazard(ctx context.Context, p OccupationalHazardListParams) ([]model.OccupationalHazard, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.OccupationalHazard{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	if v := strings.TrimSpace(p.IsDeclare); v != "" {
		q = q.Where("is_declare = ?", v)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.OccupationalHazard
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetOccupationalHazard(ctx context.Context, id int64) (*model.OccupationalHazard, error) {
	var m model.OccupationalHazard
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateOccupationalHazard(ctx context.Context, m *model.OccupationalHazard) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateOccupationalHazard(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.OccupationalHazard{}, id, updates)
}

func (s *HealthFireService) DeleteOccupationalHazard(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.OccupationalHazard{}, id).Error
}

// ── 职业健康体系（singleton per company） ──────────────────────────────

func (s *HealthFireService) FirstHealthSystem(ctx context.Context, companyID int64) (*model.HealthSystem, error) {
	var m model.HealthSystem
	err := s.db.WithContext(ctx).Where("company_id = ?", companyID).Order("id DESC").First(&m).Error
	if err == nil {
		return &m, nil
	}
	if ae, ok := errs.As(mapNotFound(err)); ok && ae.Code == errs.CodeNotFound {
		return nil, nil
	}
	return nil, err
}

func (s *HealthFireService) CreateHealthSystem(ctx context.Context, m *model.HealthSystem) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateHealthSystem(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.HealthSystem{}, id, updates)
}

// ── 健康档案 ───────────────────────────────────────────────────────────

type HealthFileListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *HealthFireService) ListHealthFile(ctx context.Context, p HealthFileListParams) ([]model.HealthFile, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.HealthFile{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HealthFile
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetHealthFile(ctx context.Context, id int64) (*model.HealthFile, error) {
	var m model.HealthFile
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateHealthFile(ctx context.Context, m *model.HealthFile) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateHealthFile(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.HealthFile{}, id, updates)
}

func (s *HealthFireService) DeleteHealthFile(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.HealthFile{}, id).Error
}

// ── 职业病报告 ─────────────────────────────────────────────────────────

type HealthReportListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *HealthFireService) ListHealthReport(ctx context.Context, p HealthReportListParams) ([]model.HealthReport, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.HealthReport{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HealthReport
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetHealthReport(ctx context.Context, id int64) (*model.HealthReport, error) {
	var m model.HealthReport
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateHealthReport(ctx context.Context, m *model.HealthReport) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateHealthReport(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.HealthReport{}, id, updates)
}

func (s *HealthFireService) DeleteHealthReport(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.HealthReport{}, id).Error
}

// ── 职业健康体检 ───────────────────────────────────────────────────────

type HealthCheckupListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Type      string
}

func (s *HealthFireService) ListHealthCheckup(ctx context.Context, p HealthCheckupListParams) ([]model.HealthCheckup, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.HealthCheckup{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Type); kw != "" {
		q = q.Where("type ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HealthCheckup
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetHealthCheckup(ctx context.Context, id int64) (*model.HealthCheckup, error) {
	var m model.HealthCheckup
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateHealthCheckup(ctx context.Context, m *model.HealthCheckup) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateHealthCheckup(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.HealthCheckup{}, id, updates)
}

func (s *HealthFireService) DeleteHealthCheckup(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.HealthCheckup{}, id).Error
}

// ── 职业卫生评价（with skill_unit / skill_unit_kid） ───────────────────

type HygieneEvaluationListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *HealthFireService) ListHygieneEvaluation(ctx context.Context, p HygieneEvaluationListParams) ([]model.HygieneEvaluation, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.HygieneEvaluation{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.HygieneEvaluation
	if err := q.Preload("SkillUnit").Preload("SkillUnitKid").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetHygieneEvaluation(ctx context.Context, id int64) (*model.HygieneEvaluation, error) {
	var m model.HygieneEvaluation
	if err := s.db.WithContext(ctx).Preload("SkillUnit").Preload("SkillUnitKid").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateHygieneEvaluation(ctx context.Context, m *model.HygieneEvaluation) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateHygieneEvaluation(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.HygieneEvaluation{}, id, updates)
}

func (s *HealthFireService) DeleteHygieneEvaluation(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.HygieneEvaluation{}, id).Error
}

// ── 消防安全（with skill_unit / skill_unit_kid） ───────────────────────

type FireSafetyListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *HealthFireService) ListFireSafety(ctx context.Context, p FireSafetyListParams) ([]model.FireSafety, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.FireSafety{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.FireSafety
	if err := q.Preload("SkillUnit").Preload("SkillUnitKid").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetFireSafety(ctx context.Context, id int64) (*model.FireSafety, error) {
	var m model.FireSafety
	if err := s.db.WithContext(ctx).Preload("SkillUnit").Preload("SkillUnitKid").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateFireSafety(ctx context.Context, m *model.FireSafety) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateFireSafety(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.FireSafety{}, id, updates)
}

func (s *HealthFireService) DeleteFireSafety(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.FireSafety{}, id).Error
}

// ── 消防器材 ───────────────────────────────────────────────────────────

type FireEquipmentListParams struct {
	Page          int
	PerPage       int
	CompanyID     int64
	EquipmentName string
	SpecialType   string
}

func (s *HealthFireService) ListFireEquipment(ctx context.Context, p FireEquipmentListParams) ([]model.FireEquipment, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.FireEquipment{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.EquipmentName); kw != "" {
		q = q.Where("equipment_name ILIKE ?", "%"+kw+"%")
	}
	if kw := strings.TrimSpace(p.SpecialType); kw != "" {
		q = q.Where("special_type ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.FireEquipment
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetFireEquipment(ctx context.Context, id int64) (*model.FireEquipment, error) {
	var m model.FireEquipment
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateFireEquipment(ctx context.Context, m *model.FireEquipment) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateFireEquipment(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.FireEquipment{}, id, updates)
}

func (s *HealthFireService) DeleteFireEquipment(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.FireEquipment{}, id).Error
}

// CopyFireEquipment 复制选中行（去掉 id / 时间戳，保留业务字段 + company_id）。
func (s *HealthFireService) CopyFireEquipment(ctx context.Context, ids []int64, uid int64) ([]int64, error) {
	if len(ids) == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 ids")
	}
	var originals []model.FireEquipment
	if err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&originals).Error; err != nil {
		return nil, err
	}
	if len(originals) == 0 {
		return nil, errs.New(errs.CodeNotFound, "未找到指定记录")
	}
	newIDs := make([]int64, 0, len(originals))
	for _, o := range originals {
		clone := model.FireEquipment{
			EquipmentName:    o.EquipmentName,
			UtilizedLocation: o.UtilizedLocation,
			SpecialType:      o.SpecialType,
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

// ── 消防设施 ───────────────────────────────────────────────────────────

type FireFacilityListParams struct {
	Page          int
	PerPage       int
	CompanyID     int64
	EquipmentName string
	SpecialType   string
}

func (s *HealthFireService) ListFireFacility(ctx context.Context, p FireFacilityListParams) ([]model.FireFacility, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.FireFacility{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.EquipmentName); kw != "" {
		q = q.Where("equipment_name ILIKE ?", "%"+kw+"%")
	}
	if kw := strings.TrimSpace(p.SpecialType); kw != "" {
		q = q.Where("special_type ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.FireFacility
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *HealthFireService) GetFireFacility(ctx context.Context, id int64) (*model.FireFacility, error) {
	var m model.FireFacility
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *HealthFireService) CreateFireFacility(ctx context.Context, m *model.FireFacility) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *HealthFireService) UpdateFireFacility(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.FireFacility{}, id, updates)
}

func (s *HealthFireService) DeleteFireFacility(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.FireFacility{}, id).Error
}

func (s *HealthFireService) CopyFireFacility(ctx context.Context, ids []int64, uid int64) ([]int64, error) {
	if len(ids) == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 ids")
	}
	var originals []model.FireFacility
	if err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&originals).Error; err != nil {
		return nil, err
	}
	if len(originals) == 0 {
		return nil, errs.New(errs.CodeNotFound, "未找到指定记录")
	}
	newIDs := make([]int64, 0, len(originals))
	for _, o := range originals {
		clone := model.FireFacility{
			EquipmentName:    o.EquipmentName,
			UtilizedLocation: o.UtilizedLocation,
			SpecialType:      o.SpecialType,
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
