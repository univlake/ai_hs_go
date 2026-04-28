package service

import (
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// ExtraModulesService 聚合 5 个剩余迁移模块的 CRUD：
//   - change_manages         变更管理
//   - regulation_libraries   法律法规库（preload regulation_type）
//   - training_files         培训资料·文件
//   - training_videos        培训资料·视频
//   - 风险画像 / 风险日志    /risks/portrait 与 /risks/log
//
// 5 个模块互相独立但都是 per-company partial-update 软删除模式，集中放一个 service
// 是为了避免在 router 拉一长串构造函数；处理流程跟 ProcessService 完全一致。
type ExtraModulesService struct {
	db *gorm.DB
}

func NewExtraModulesService(db *gorm.DB) *ExtraModulesService {
	return &ExtraModulesService{db: db}
}

// ── 变更管理 ────────────────────────────────────────────────────────────

type ChangeManageListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
}

func (s *ExtraModulesService) ListChangeManage(ctx context.Context, p ChangeManageListParams) ([]model.ChangeManage, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.ChangeManage{}).Where("company_id = ?", p.CompanyID)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ChangeManage
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExtraModulesService) GetChangeManage(ctx context.Context, id int64) (*model.ChangeManage, error) {
	var m model.ChangeManage
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ExtraModulesService) CreateChangeManage(ctx context.Context, m *model.ChangeManage) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ExtraModulesService) UpdateChangeManage(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.ChangeManage{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ExtraModulesService) DeleteChangeManage(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.ChangeManage{}, id).Error
}

// ── 法律法规库 ─────────────────────────────────────────────────────────

type RegulationLibraryListParams struct {
	Page                    int
	PerPage                 int
	Title                   string
	RegulationLibraryTypeID *int64 // 显式指定的类型 id；nil 表示未指定
	TypeParentID            *int64 // 仅当上面 id 为 nil 时生效，按父级展开成 IN 列表
}

func (s *ExtraModulesService) ListRegulationLibrary(ctx context.Context, p RegulationLibraryListParams) ([]model.RegulationLibrary, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.RegulationLibrary{})
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ?", "%"+t+"%")
	}
	switch {
	case p.RegulationLibraryTypeID != nil:
		q = q.Where("regulation_library_type_id = ?", *p.RegulationLibraryTypeID)
	case p.TypeParentID != nil:
		// 跟 Laravel RegulationLibraryController 一致：按 type 父级展开成 child id 列表，
		// 没有子级则退化为按父级 id 自身过滤。
		var childIDs []int64
		if err := s.db.WithContext(ctx).
			Model(&model.RegulationLibraryType{}).
			Where("pid = ?", *p.TypeParentID).
			Pluck("id", &childIDs).Error; err != nil {
			return nil, 0, err
		}
		if len(childIDs) == 0 {
			q = q.Where("regulation_library_type_id = ?", *p.TypeParentID)
		} else {
			q = q.Where("regulation_library_type_id IN ?", childIDs)
		}
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.RegulationLibrary
	if err := q.Preload("RegulationType").
		Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExtraModulesService) GetRegulationLibrary(ctx context.Context, id int64) (*model.RegulationLibrary, error) {
	var m model.RegulationLibrary
	if err := s.db.WithContext(ctx).Preload("RegulationType").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ExtraModulesService) CreateRegulationLibrary(ctx context.Context, m *model.RegulationLibrary) error {
	if strings.TrimSpace(m.Title) == "" {
		return errs.New(errs.CodeUnprocessable, "标题不能为空")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ExtraModulesService) UpdateRegulationLibrary(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.RegulationLibrary{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ExtraModulesService) DeleteRegulationLibrary(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.RegulationLibrary{}, id).Error
}

// ── 培训资料·文件 ─────────────────────────────────────────────────────

type TrainingFileListParams struct {
	Page            int
	PerPage         int
	CompanyID       int64
	ParentCompanyID int64 // 集团端共享的资料：>0 时 OR 进 IN 列表
	Title           string
}

func (s *ExtraModulesService) ListTrainingFile(ctx context.Context, p TrainingFileListParams) ([]model.TrainingFile, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.TrainingFile{})
	if p.CompanyID != 0 {
		ids := []int64{p.CompanyID}
		if p.ParentCompanyID > 0 {
			ids = append(ids, p.ParentCompanyID)
		}
		q = q.Where("company_id IN ?", ids)
	}
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ?", "%"+t+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.TrainingFile
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExtraModulesService) GetTrainingFile(ctx context.Context, id int64) (*model.TrainingFile, error) {
	var m model.TrainingFile
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

// IncrementTrainingFileDownload 自增计数 + 写入下载流水。Laravel 在 show 接口里
// 同步做这两件事，handler 取到 file_url 后会调 OssService::signUrl 包装后返回。
func (s *ExtraModulesService) IncrementTrainingFileDownload(ctx context.Context, id, userID, companyID int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.TrainingFile{}).
			Where("id = ?", id).
			UpdateColumn("times_download", gorm.Expr("times_download + 1")).Error; err != nil {
			return err
		}
		log := &model.TrainingFileLog{
			TrainingFileID: id,
			CompanyID:      companyID,
		}
		log.UserID = userID
		log.LastUserID = userID
		return tx.Create(log).Error
	})
}

func (s *ExtraModulesService) CreateTrainingFile(ctx context.Context, m *model.TrainingFile) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ExtraModulesService) UpdateTrainingFile(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.TrainingFile{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ExtraModulesService) DeleteTrainingFile(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.TrainingFile{}, id).Error
}

// ── 培训资料·视频 ─────────────────────────────────────────────────────

type TrainingVideoListParams struct {
	Page            int
	PerPage         int
	CompanyID       int64
	ParentCompanyID int64
	Title           string
}

func (s *ExtraModulesService) ListTrainingVideo(ctx context.Context, p TrainingVideoListParams) ([]model.TrainingVideo, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.TrainingVideo{})
	if p.CompanyID != 0 {
		ids := []int64{p.CompanyID}
		if p.ParentCompanyID > 0 {
			ids = append(ids, p.ParentCompanyID)
		}
		q = q.Where("company_id IN ?", ids)
	}
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ?", "%"+t+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.TrainingVideo
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExtraModulesService) GetTrainingVideo(ctx context.Context, id int64) (*model.TrainingVideo, error) {
	var m model.TrainingVideo
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ExtraModulesService) IncrementTrainingVideoWatch(ctx context.Context, id, userID, companyID int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.TrainingVideo{}).
			Where("id = ?", id).
			UpdateColumn("times_watched", gorm.Expr("times_watched + 1")).Error; err != nil {
			return err
		}
		log := &model.TrainingVideoLog{
			TrainingVideoID: id,
			CompanyID:       companyID,
		}
		log.UserID = userID
		log.LastUserID = userID
		return tx.Create(log).Error
	})
}

func (s *ExtraModulesService) CreateTrainingVideo(ctx context.Context, m *model.TrainingVideo) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ExtraModulesService) UpdateTrainingVideo(ctx context.Context, id int64, updates map[string]any) error {
	res := s.db.WithContext(ctx).Model(&model.TrainingVideo{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

func (s *ExtraModulesService) DeleteTrainingVideo(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.TrainingVideo{}, id).Error
}

// ── 风险画像 / 风险日志 ─────────────────────────────────────────────────

// PortraitData 跟 Laravel RisksController.portrait 的返回格式对齐：
// 三个字段任一为空就返回空对象（Laravel 直接 return success([])，前端对应 if 分支）。
type PortraitData struct {
	InherentRiskLog *model.InherentRiskLog `json:"inherent_risk_log,omitempty"`
	FloatRiskLog    *model.FloatRiskLog    `json:"float_risk_log,omitempty"`
	RiskLog         *model.RiskLog         `json:"risk_log,omitempty"`
}

// CompanyPortrait 取最新的固有 / 浮动 / 风险日志各一行；
// 任一固有或浮动缺失时返回 nil，让 handler 决定是否返回空数组（沿 Laravel 行为）。
func (s *ExtraModulesService) CompanyPortrait(ctx context.Context, companyID int64) (*PortraitData, error) {
	if companyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	var inherent model.InherentRiskLog
	if err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Order("created_at DESC, id DESC").
		First(&inherent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var floatLog model.FloatRiskLog
	if err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Order("created_at DESC, id DESC").
		First(&floatLog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var risk model.RiskLog
	var riskPtr *model.RiskLog
	if err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Order("created_at DESC, id DESC").
		First(&risk).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err == nil {
		riskPtr = &risk
	}
	return &PortraitData{
		InherentRiskLog: &inherent,
		FloatRiskLog:    &floatLog,
		RiskLog:         riskPtr,
	}, nil
}

type RiskLogListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
}

// ListCompanyRiskLogs 是 /risks/log 的实现：单企业风险日志分页。
// 跟 RiskMonitorService.RiskLog 不同：那个聚合所有下属企业，这个只返回单 company_id。
func (s *ExtraModulesService) ListCompanyRiskLogs(ctx context.Context, p RiskLogListParams) ([]model.RiskLog, int64, error) {
	if p.CompanyID == 0 {
		return nil, 0, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.RiskLog{}).Where("company_id = ?", p.CompanyID)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.RiskLog
	if err := q.Order("created_at DESC, id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

