package service

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// ── Certificate statuses ────────────────────────────────────────

type CertificateStatusService struct{ db *gorm.DB }

func NewCertificateStatusService(db *gorm.DB) *CertificateStatusService {
	return &CertificateStatusService{db: db}
}

type CertificateStatusListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Type      *int16 // nil = no filter; 1-3 三类人员; 4-6 消防
}

func (s *CertificateStatusService) List(ctx context.Context, p CertificateStatusListParams) ([]model.CertificateStatus, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.CertificateStatus{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if p.Type != nil {
		q = q.Where("type = ?", *p.Type)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.CertificateStatus
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type CertificateStatusInput struct {
	Name                *string
	Type                int16
	ShouldObtainPeople  *int
	AlreadyObtainPeople *int
	NoObtainPeople      *string
	TrainPeople         *int
	CompanyID           int64
	UserID              int64
}

func (s *CertificateStatusService) Create(ctx context.Context, in CertificateStatusInput) (*model.CertificateStatus, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	if in.Type == 0 {
		in.Type = 1
	}
	row := &model.CertificateStatus{
		Name:                in.Name,
		Type:                in.Type,
		ShouldObtainPeople:  in.ShouldObtainPeople,
		AlreadyObtainPeople: in.AlreadyObtainPeople,
		NoObtainPeople:      in.NoObtainPeople,
		TrainPeople:         in.TrainPeople,
		CompanyID:           in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *CertificateStatusService) Update(ctx context.Context, id int64, in CertificateStatusInput) (*model.CertificateStatus, error) {
	var row model.CertificateStatus
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Name != nil {
		updates["name"] = in.Name
	}
	if in.Type != 0 {
		updates["type"] = in.Type
	}
	if in.ShouldObtainPeople != nil {
		updates["should_obtain_people"] = in.ShouldObtainPeople
	}
	if in.AlreadyObtainPeople != nil {
		updates["already_obtain_people"] = in.AlreadyObtainPeople
	}
	if in.NoObtainPeople != nil {
		updates["no_obtain_people"] = in.NoObtainPeople
	}
	if in.TrainPeople != nil {
		updates["train_people"] = in.TrainPeople
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

// ── Production standardization (singleton per company) ─────────

type ProductionStandardizationService struct{ db *gorm.DB }

func NewProductionStandardizationService(db *gorm.DB) *ProductionStandardizationService {
	return &ProductionStandardizationService{db: db}
}

func (s *ProductionStandardizationService) First(ctx context.Context, companyID int64) (*model.ProductionStandardization, error) {
	var row model.ProductionStandardization
	err := s.db.WithContext(ctx).Where("company_id = ?", companyID).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &row, err
}

type ProductionStandardizationInput struct {
	Grade       *string
	GetTime     *string
	ValidPeriod *string
	AnnexURL    *string
	CompanyID   int64
	UserID      int64
}

func (s *ProductionStandardizationService) Create(ctx context.Context, in ProductionStandardizationInput) (*model.ProductionStandardization, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.ProductionStandardization{
		Grade:       in.Grade,
		GetTime:     parseDateOnly(deref(in.GetTime)),
		ValidPeriod: parseDateOnly(deref(in.ValidPeriod)),
		AnnexURL:    in.AnnexURL,
		CompanyID:   in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *ProductionStandardizationService) Update(ctx context.Context, id int64, in ProductionStandardizationInput) (*model.ProductionStandardization, error) {
	var row model.ProductionStandardization
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Grade != nil {
		updates["grade"] = in.Grade
	}
	if in.GetTime != nil {
		if t := parseDateOnly(*in.GetTime); t != nil {
			updates["get_time"] = t
		} else if *in.GetTime == "" {
			updates["get_time"] = nil
		}
	}
	if in.ValidPeriod != nil {
		if t := parseDateOnly(*in.ValidPeriod); t != nil {
			updates["valid_period"] = t
		} else if *in.ValidPeriod == "" {
			updates["valid_period"] = nil
		}
	}
	if in.AnnexURL != nil {
		updates["annex_url"] = in.AnnexURL
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

// ── Occupational health safety (singleton per company) ─────────

type OccupationalHealthSafetyService struct{ db *gorm.DB }

func NewOccupationalHealthSafetyService(db *gorm.DB) *OccupationalHealthSafetyService {
	return &OccupationalHealthSafetyService{db: db}
}

func (s *OccupationalHealthSafetyService) First(ctx context.Context, companyID int64) (*model.OccupationalHealthSafety, error) {
	var row model.OccupationalHealthSafety
	err := s.db.WithContext(ctx).Where("company_id = ?", companyID).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &row, err
}

type OccupationalHealthSafetyInput struct {
	CertificationBody *string
	StandardEdition   *string
	GetTime           *string
	ValidPeriod       *string
	AnnexURL          *string
	CompanyID         int64
	UserID            int64
}

func (s *OccupationalHealthSafetyService) Create(ctx context.Context, in OccupationalHealthSafetyInput) (*model.OccupationalHealthSafety, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.OccupationalHealthSafety{
		CertificationBody: in.CertificationBody,
		StandardEdition:   in.StandardEdition,
		GetTime:           parseDateOnly(deref(in.GetTime)),
		ValidPeriod:       parseDateOnly(deref(in.ValidPeriod)),
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

func (s *OccupationalHealthSafetyService) Update(ctx context.Context, id int64, in OccupationalHealthSafetyInput) (*model.OccupationalHealthSafety, error) {
	var row model.OccupationalHealthSafety
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.CertificationBody != nil {
		updates["certification_body"] = in.CertificationBody
	}
	if in.StandardEdition != nil {
		updates["standard_edition"] = in.StandardEdition
	}
	if in.GetTime != nil {
		if t := parseDateOnly(*in.GetTime); t != nil {
			updates["get_time"] = t
		} else if *in.GetTime == "" {
			updates["get_time"] = nil
		}
	}
	if in.ValidPeriod != nil {
		if t := parseDateOnly(*in.ValidPeriod); t != nil {
			updates["valid_period"] = t
		} else if *in.ValidPeriod == "" {
			updates["valid_period"] = nil
		}
	}
	if in.AnnexURL != nil {
		updates["annex_url"] = in.AnnexURL
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

// ── Funding executions (with computed proportion) ──────────────

type FundingExecutionService struct{ db *gorm.DB }

func NewFundingExecutionService(db *gorm.DB) *FundingExecutionService {
	return &FundingExecutionService{db: db}
}

type FundingExecutionListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Year      string
}

// List eager-loads SecurityInvestment + computes `proportion` per row by
// dividing this row's budget against the total budget of all
// security_investments belonging to the same company AND year.
func (s *FundingExecutionService) List(ctx context.Context, p FundingExecutionListParams) ([]model.FundingExecution, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.FundingExecution{}).Preload("SecurityInvestment")
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if p.Year != "" {
		q = q.Where("year = ?", p.Year)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.FundingExecution
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}

	// Sum security_investments.budget grouped by (company_id, year) for
	// every distinct year present in the result page. Single roundtrip
	// avoids N+1.
	years := map[string]struct{}{}
	for _, r := range rows {
		if r.Year != nil && *r.Year != "" {
			years[*r.Year] = struct{}{}
		}
	}
	totalsByYear := make(map[string]float64, len(years))
	if len(years) > 0 && p.CompanyID > 0 {
		yearList := make([]string, 0, len(years))
		for y := range years {
			yearList = append(yearList, y)
		}
		type sumRow struct {
			Year  string
			Total float64
		}
		var sums []sumRow
		if err := s.db.WithContext(ctx).
			Table("security_investments").
			Select("year, COALESCE(SUM(budget), 0) AS total").
			Where("company_id = ? AND year IN ? AND deleted_at IS NULL", p.CompanyID, yearList).
			Group("year").
			Scan(&sums).Error; err != nil {
			return nil, 0, err
		}
		for _, s := range sums {
			totalsByYear[s.Year] = s.Total
		}
	}
	for i := range rows {
		rows[i].Proportion = computeProportion(rows[i].Budget, rows[i].Year, totalsByYear)
	}
	return rows, total, nil
}

func computeProportion(budget *float64, year *string, totals map[string]float64) string {
	if budget == nil || year == nil {
		return ""
	}
	annual, ok := totals[*year]
	if !ok || annual == 0 {
		return ""
	}
	pct := (*budget / annual) * 100
	return fmt.Sprintf("%.2f%%", pct)
}

type FundingExecutionInput struct {
	Year                 *string
	SecurityInvestmentID *int64
	SimpleContent        *string
	Content              *string
	Budget               *float64
	CompanyID            int64
	UserID               int64
}

func (s *FundingExecutionService) Create(ctx context.Context, in FundingExecutionInput) (*model.FundingExecution, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.FundingExecution{
		Year:                 in.Year,
		SecurityInvestmentID: in.SecurityInvestmentID,
		SimpleContent:        in.SimpleContent,
		Content:              in.Content,
		Budget:               in.Budget,
		CompanyID:            in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *FundingExecutionService) Update(ctx context.Context, id int64, in FundingExecutionInput) (*model.FundingExecution, error) {
	var row model.FundingExecution
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
	if in.SecurityInvestmentID != nil {
		updates["security_investment_id"] = in.SecurityInvestmentID
	}
	if in.SimpleContent != nil {
		updates["simple_content"] = in.SimpleContent
	}
	if in.Content != nil {
		updates["content"] = in.Content
	}
	if in.Budget != nil {
		updates["budget"] = *in.Budget
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *FundingExecutionService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.FundingExecution{}, id).Error
}

// ── Safety management staff ────────────────────────────────────

type SafetyManagementStaffService struct{ db *gorm.DB }

func NewSafetyManagementStaffService(db *gorm.DB) *SafetyManagementStaffService {
	return &SafetyManagementStaffService{db: db}
}

type SafetyManagementStaffListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	TypeID    *int64
}

func (s *SafetyManagementStaffService) List(ctx context.Context, p SafetyManagementStaffListParams) ([]model.SafetyManagementStaff, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	// Preload the type dictionary so the table + detail drawer can show
	// the type label without an extra round-trip.
	q := s.db.WithContext(ctx).Model(&model.SafetyManagementStaff{}).
		Preload("SafetyManagementType")
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if p.TypeID != nil {
		q = q.Where("safety_management_type_id = ?", *p.TypeID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SafetyManagementStaff
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type SafetyManagementStaffInput struct {
	Name                   *string
	SafetyManagementTypeID int64
	Tel                    *string
	CertTitle              *string
	CertNumber             *string
	ValidPeriod            *string
	CompanyID              int64
	UserID                 int64
}

func (s *SafetyManagementStaffService) Create(ctx context.Context, in SafetyManagementStaffInput) (*model.SafetyManagementStaff, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	if in.SafetyManagementTypeID == 0 {
		in.SafetyManagementTypeID = 1
	}
	row := &model.SafetyManagementStaff{
		Name:                   in.Name,
		SafetyManagementTypeID: in.SafetyManagementTypeID,
		Tel:                    in.Tel,
		CertTitle:              in.CertTitle,
		CertNumber:             in.CertNumber,
		ValidPeriod:            parseDateOnly(deref(in.ValidPeriod)),
		CompanyID:              in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *SafetyManagementStaffService) Update(ctx context.Context, id int64, in SafetyManagementStaffInput) (*model.SafetyManagementStaff, error) {
	var row model.SafetyManagementStaff
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Name != nil {
		updates["name"] = in.Name
	}
	if in.SafetyManagementTypeID != 0 {
		updates["safety_management_type_id"] = in.SafetyManagementTypeID
	}
	if in.Tel != nil {
		updates["tel"] = in.Tel
	}
	if in.CertTitle != nil {
		updates["cert_title"] = in.CertTitle
	}
	if in.CertNumber != nil {
		updates["cert_number"] = in.CertNumber
	}
	if in.ValidPeriod != nil {
		if t := parseDateOnly(*in.ValidPeriod); t != nil {
			updates["valid_period"] = t
		} else if *in.ValidPeriod == "" {
			updates["valid_period"] = nil
		}
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *SafetyManagementStaffService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.SafetyManagementStaff{}, id).Error
}

// ── Special equipment staff (model already in safety_ledger.go) ────

type SpecialEquipmentStaffService struct{ db *gorm.DB }

func NewSpecialEquipmentStaffService(db *gorm.DB) *SpecialEquipmentStaffService {
	return &SpecialEquipmentStaffService{db: db}
}

type SpecialEquipmentStaffListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
}

func (s *SpecialEquipmentStaffService) List(ctx context.Context, p SpecialEquipmentStaffListParams) ([]model.SpecialEquipmentStaff, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.SpecialEquipmentStaff{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.SpecialEquipmentStaff
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type SpecialEquipmentStaffInput struct {
	Name        *string
	CertTitle   *string
	CertNumber  *string
	ValidPeriod *string
	CompanyID   int64
	UserID      int64
}

func (s *SpecialEquipmentStaffService) Create(ctx context.Context, in SpecialEquipmentStaffInput) (*model.SpecialEquipmentStaff, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.SpecialEquipmentStaff{
		Name:        in.Name,
		CertTitle:   in.CertTitle,
		CertNumber:  in.CertNumber,
		ValidPeriod: parseDateOnly(deref(in.ValidPeriod)),
		CompanyID:   in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *SpecialEquipmentStaffService) Update(ctx context.Context, id int64, in SpecialEquipmentStaffInput) (*model.SpecialEquipmentStaff, error) {
	var row model.SpecialEquipmentStaff
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Name != nil {
		updates["name"] = in.Name
	}
	if in.CertTitle != nil {
		updates["cert_title"] = in.CertTitle
	}
	if in.CertNumber != nil {
		updates["cert_number"] = in.CertNumber
	}
	if in.ValidPeriod != nil {
		if t := parseDateOnly(*in.ValidPeriod); t != nil {
			updates["valid_period"] = t
		} else if *in.ValidPeriod == "" {
			updates["valid_period"] = nil
		}
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *SpecialEquipmentStaffService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.SpecialEquipmentStaff{}, id).Error
}

// ── Elimination control staff ──────────────────────────────────

type EliminationControlStaffService struct{ db *gorm.DB }

func NewEliminationControlStaffService(db *gorm.DB) *EliminationControlStaffService {
	return &EliminationControlStaffService{db: db}
}

type EliminationControlStaffListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Type      *int16
}

func (s *EliminationControlStaffService) List(ctx context.Context, p EliminationControlStaffListParams) ([]model.EliminationControlStaff, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.EliminationControlStaff{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if p.Type != nil {
		q = q.Where("type = ?", *p.Type)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.EliminationControlStaff
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type EliminationControlStaffInput struct {
	Name        *string
	Type        int16
	CertTitle   *string
	CertNumber  *string
	ValidPeriod *string
	CompanyID   int64
	UserID      int64
}

func (s *EliminationControlStaffService) Create(ctx context.Context, in EliminationControlStaffInput) (*model.EliminationControlStaff, error) {
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	if in.Type == 0 {
		in.Type = 1
	}
	row := &model.EliminationControlStaff{
		Name:        in.Name,
		Type:        in.Type,
		CertTitle:   in.CertTitle,
		CertNumber:  in.CertNumber,
		ValidPeriod: parseDateOnly(deref(in.ValidPeriod)),
		CompanyID:   in.CompanyID,
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *EliminationControlStaffService) Update(ctx context.Context, id int64, in EliminationControlStaffInput) (*model.EliminationControlStaff, error) {
	var row model.EliminationControlStaff
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if in.Name != nil {
		updates["name"] = in.Name
	}
	if in.Type != 0 {
		updates["type"] = in.Type
	}
	if in.CertTitle != nil {
		updates["cert_title"] = in.CertTitle
	}
	if in.CertNumber != nil {
		updates["cert_number"] = in.CertNumber
	}
	if in.ValidPeriod != nil {
		if t := parseDateOnly(*in.ValidPeriod); t != nil {
			updates["valid_period"] = t
		} else if *in.ValidPeriod == "" {
			updates["valid_period"] = nil
		}
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *EliminationControlStaffService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.EliminationControlStaff{}, id).Error
}
