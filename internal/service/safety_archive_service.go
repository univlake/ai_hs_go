package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// ── Safety architecture ─────────────────────────────────────────

// SafetyArchitectureService — single record per company. Frontend reads
// via index() with a company_id query, then either creates a new row or
// updates the existing one.
type SafetyArchitectureService struct {
	db *gorm.DB
}

func NewSafetyArchitectureService(db *gorm.DB) *SafetyArchitectureService {
	return &SafetyArchitectureService{db: db}
}

// First returns the company's single safety_architecture row (if any).
// nil + nil means no record yet — caller should hand back null to the
// frontend so it can decide between POST and PUT.
func (s *SafetyArchitectureService) First(ctx context.Context, companyID int64) (*model.SafetyArchitecture, error) {
	var row model.SafetyArchitecture
	err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &row, nil
}

type SafetyArchitectureInput struct {
	ThumbnailURL *string
	AnnexURL     *string
	CompanyID    int64
	UserID       int64
}

func (s *SafetyArchitectureService) Create(ctx context.Context, in SafetyArchitectureInput) (*model.SafetyArchitecture, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.SafetyArchitecture{
		ThumbnailURL: in.ThumbnailURL,
		AnnexURL:     in.AnnexURL,
		CompanyID:    &in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *SafetyArchitectureService) Update(ctx context.Context, id int64, in SafetyArchitectureInput) (*model.SafetyArchitecture, error) {
	var row model.SafetyArchitecture
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.ThumbnailURL != nil {
		updates["thumbnail_url"] = in.ThumbnailURL
	}
	if in.AnnexURL != nil {
		updates["annex_url"] = in.AnnexURL
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

// ── Filing permissions ──────────────────────────────────────────

type FilingPermissionService struct {
	db *gorm.DB
}

func NewFilingPermissionService(db *gorm.DB) *FilingPermissionService {
	return &FilingPermissionService{db: db}
}

type FilingPermissionListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
}

func (s *FilingPermissionService) List(ctx context.Context, p FilingPermissionListParams) ([]model.FilingPermission, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.FilingPermission{}).Preload("RecordType")
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.FilingPermission
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type FilingPermissionInput struct {
	RecordTypeID      *int64
	IssuingDepartment *string
	FilingRange       *string
	FilingTime        *string // ISO date; service parses
	ValidPeriod       *string
	Comment           *string
	AnnexURL          *string
	CompanyID         int64
	UserID            int64
}

func (s *FilingPermissionService) Create(ctx context.Context, in FilingPermissionInput) (*model.FilingPermission, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.FilingPermission{
		RecordTypeID:      in.RecordTypeID,
		IssuingDepartment: in.IssuingDepartment,
		FilingRange:       in.FilingRange,
		FilingTime:        parseDateOnly(deref(in.FilingTime)),
		ValidPeriod:       parseDateOnly(deref(in.ValidPeriod)),
		Comment:           in.Comment,
		AnnexURL:          in.AnnexURL,
		CompanyID:         in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *FilingPermissionService) Update(ctx context.Context, id int64, in FilingPermissionInput) (*model.FilingPermission, error) {
	var row model.FilingPermission
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.RecordTypeID != nil {
		updates["record_type_id"] = in.RecordTypeID
	}
	if in.IssuingDepartment != nil {
		updates["issuing_department"] = in.IssuingDepartment
	}
	if in.FilingRange != nil {
		updates["filing_range"] = in.FilingRange
	}
	if in.FilingTime != nil {
		if t := parseDateOnly(*in.FilingTime); t != nil {
			updates["filing_time"] = t
		} else if *in.FilingTime == "" {
			updates["filing_time"] = nil
		}
	}
	if in.ValidPeriod != nil {
		if t := parseDateOnly(*in.ValidPeriod); t != nil {
			updates["valid_period"] = t
		} else if *in.ValidPeriod == "" {
			updates["valid_period"] = nil
		}
	}
	if in.Comment != nil {
		updates["comment"] = in.Comment
	}
	if in.AnnexURL != nil {
		updates["annex_url"] = in.AnnexURL
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *FilingPermissionService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.FilingPermission{}, id).Error
}

// ── Buildings ───────────────────────────────────────────────────

type BuildingService struct {
	db *gorm.DB
}

func NewBuildingService(db *gorm.DB) *BuildingService {
	return &BuildingService{db: db}
}

type BuildingListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Keyword   string
}

func (s *BuildingService) List(ctx context.Context, p BuildingListParams) ([]model.Building, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.Building{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if kw := strings.TrimSpace(p.Keyword); kw != "" {
		q = q.Where("title ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.Building
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type BuildingInput struct {
	Title     *string
	Structure *string
	Area      *string
	FireClass *string
	CompanyID int64
	UserID    int64
}

func (s *BuildingService) Create(ctx context.Context, in BuildingInput) (*model.Building, error) {
	row := &model.Building{
		Title:     in.Title,
		Structure: in.Structure,
		Area:      in.Area,
		FireClass: in.FireClass,
		CompanyID: &in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *BuildingService) Update(ctx context.Context, id int64, in BuildingInput) (*model.Building, error) {
	var row model.Building
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Title != nil {
		updates["title"] = in.Title
	}
	if in.Structure != nil {
		updates["structure"] = in.Structure
	}
	if in.Area != nil {
		updates["area"] = in.Area
	}
	if in.FireClass != nil {
		updates["fire_class"] = in.FireClass
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *BuildingService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.Building{}, id).Error
}

// ── Security investments ───────────────────────────────────────

type SecurityInvestmentService struct {
	db *gorm.DB
}

func NewSecurityInvestmentService(db *gorm.DB) *SecurityInvestmentService {
	return &SecurityInvestmentService{db: db}
}

type SecurityInvestmentListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Year      string // optional year filter
}

func (s *SecurityInvestmentService) List(ctx context.Context, p SecurityInvestmentListParams) ([]model.SecurityInvestment, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.SecurityInvestment{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if y := strings.TrimSpace(p.Year); y != "" {
		q = q.Where("year = ?", y)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SecurityInvestment
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	for i := range rows {
		rows[i].Scale = computeScale(rows[i].Budget, rows[i].LastYearIncome)
	}
	return rows, total, nil
}

// computeScale returns "X.XX%" for budget / last_year_income, or "" when
// either side is missing/zero. Display-only — never persisted.
func computeScale(budget, lastYearIncome *float64) string {
	if budget == nil || lastYearIncome == nil || *lastYearIncome == 0 {
		return ""
	}
	pct := (*budget / *lastYearIncome) * 100
	return fmt.Sprintf("%.2f%%", pct)
}

type SecurityInvestmentInput struct {
	Year           *string
	Type           *string
	Title          *string
	Content        *string
	Budget         *float64
	LastYearIncome *float64
	CompanyID      int64
	UserID         int64
}

func (s *SecurityInvestmentService) Create(ctx context.Context, in SecurityInvestmentInput) (*model.SecurityInvestment, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.SecurityInvestment{
		Year:           in.Year,
		Type:           in.Type,
		Title:          in.Title,
		Content:        in.Content,
		Budget:         in.Budget,
		LastYearIncome: in.LastYearIncome,
		CompanyID:      in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *SecurityInvestmentService) Update(ctx context.Context, id int64, in SecurityInvestmentInput) (*model.SecurityInvestment, error) {
	var row model.SecurityInvestment
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Year != nil {
		updates["year"] = in.Year
	}
	if in.Type != nil {
		updates["type"] = in.Type
	}
	if in.Title != nil {
		updates["title"] = in.Title
	}
	if in.Content != nil {
		updates["content"] = in.Content
	}
	if in.Budget != nil {
		updates["budget"] = *in.Budget
	}
	if in.LastYearIncome != nil {
		updates["last_year_income"] = *in.LastYearIncome
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *SecurityInvestmentService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.SecurityInvestment{}, id).Error
}

// ── Personnels (model already exists in 0004) ──────────────────

type PersonnelService struct {
	db *gorm.DB
}

func NewPersonnelService(db *gorm.DB) *PersonnelService {
	return &PersonnelService{db: db}
}

type PersonnelListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
}

func (s *PersonnelService) List(ctx context.Context, p PersonnelListParams) ([]model.Personnel, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.Personnel{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.Personnel
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *PersonnelService) Get(ctx context.Context, id int64) (*model.Personnel, error) {
	var row model.Personnel
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	return &row, nil
}

type PersonnelInput struct {
	CompanyID              int64
	TotalPeople            int
	Staff                  int
	Service                int
	Manager                int
	HighSchoolDiploma      int
	SecurityOfficer        int
	SafetyEngineer         int
	SafetyEngineeringMajor int
	FireEngineer           int
	Other                  int
	UserID                 int64
}

func (s *PersonnelService) Create(ctx context.Context, in PersonnelInput) (*model.Personnel, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.Personnel{
		TotalPeople:            in.TotalPeople,
		Staff:                  in.Staff,
		Service:                in.Service,
		Manager:                in.Manager,
		HighSchoolDiploma:      in.HighSchoolDiploma,
		SecurityOfficer:        in.SecurityOfficer,
		SafetyEngineer:         in.SafetyEngineer,
		SafetyEngineeringMajor: in.SafetyEngineeringMajor,
		FireEngineer:           in.FireEngineer,
		Other:                  in.Other,
		CompanyID:              in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *PersonnelService) Update(ctx context.Context, id int64, in PersonnelInput) (*model.Personnel, error) {
	var row model.Personnel
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{
		"total_people":             in.TotalPeople,
		"staff":                    in.Staff,
		"service":                  in.Service,
		"manager":                  in.Manager,
		"high_school_diploma":      in.HighSchoolDiploma,
		"security_officer":         in.SecurityOfficer,
		"safety_engineer":          in.SafetyEngineer,
		"safety_engineering_major": in.SafetyEngineeringMajor,
		"fire_engineer":            in.FireEngineer,
		"other":                    in.Other,
		"last_user_id":             in.UserID,
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

// deref unwraps a *string into a string ("" when nil) — small helper used
// when piping optional string fields into parseDateOnly.
func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
