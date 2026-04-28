package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/repository"
)

// RiskMonitorService — 风险监测栏目（集团视图）共用服务，覆盖 5 个模块的所有
// 聚合 / 关注企业 / 在线巡检告警 / 特种设备汇总接口。
//
// 模块映射：
//   - 企业分布图           复用 CompanyService（不在此 service）
//   - 重大危险源分布图     复用 MajorHazardService（不在此 service）
//   - 风险监控             InherentRiskNum / FloatRiskNum / *RiskStatistics / RiskLog / HiddenDangerSort
//                          + FocusCompany CRUD（重点关注企业 Tab）
//   - 在线巡检             InspectionAlerts
//   - 特种设备数据统计     SpecialEquipmentStatistics
type RiskMonitorService struct {
	db          *gorm.DB
	companyRepo *repository.CompanyRepo
}

func NewRiskMonitorService(db *gorm.DB, companyRepo *repository.CompanyRepo) *RiskMonitorService {
	return &RiskMonitorService{db: db, companyRepo: companyRepo}
}

// ─── helpers ───────────────────────────────────────────────────────────

// gradeFromScore 跟 Laravel BlocRisksController 的分桶完全一致：
// <5 红 / <10 橙 / <15 黄 / 其它 蓝。
func gradeFromScore(score float64) string {
	switch {
	case score < 5:
		return "红"
	case score < 10:
		return "橙"
	case score < 15:
		return "黄"
	default:
		return "蓝"
	}
}

// round2 精度 2 位的对齐 Laravel round($x, 2)。
func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

// gradeBucket 表示按 grade 聚合的一行结果。num 是该 grade 的企业数，
// sum 是该 grade 的风险值合计（仅在计算综合指数时使用）。
type gradeBucket struct {
	Grade *string `gorm:"column:risk_grade" json:"-"`
	Num   int64   `gorm:"column:num" json:"num"`
	Sum   float64 `gorm:"column:sum_value" json:"-"`
}

// computeIndex 把按 grade 聚合的桶折回综合指数（与 Laravel 算法 1:1 对齐）：
//
//	risk = Σ (sum_i / num_i * num_i / total_count) = Σ sum_i / total_count
//
// 即所有有等级企业的算术平均（按企业权重，不按桶权重）。
func computeIndex(buckets []gradeBucket) float64 {
	var totalCount int64
	var totalSum float64
	for _, b := range buckets {
		if b.Num == 0 {
			continue
		}
		totalCount += b.Num
		totalSum += b.Sum
	}
	if totalCount == 0 {
		return 0
	}
	return round2(totalSum / float64(totalCount))
}

// produceCompanyIDs 返回 rootID 子树中 status=1 且对应 grade 字段非空的所有企业 id。
// gradeColumn ∈ {"inherent_risk_grade", "float_risk_grade"}。
func (s *RiskMonitorService) produceCompanyIDs(ctx context.Context, rootID int64, gradeColumn string) ([]int64, error) {
	descendants, err := s.companyRepo.DescendantIDs(ctx, rootID)
	if err != nil {
		return nil, err
	}
	if len(descendants) == 0 {
		return nil, nil
	}
	var ids []int64
	if err := s.db.WithContext(ctx).
		Model(&model.Company{}).
		Where("id IN ? AND status = ? AND "+gradeColumn+" IS NOT NULL", descendants, 1).
		Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

// ─── 风险监控 / 综合指数 ──────────────────────────────────────────────

// RiskNumResult 综合指数 + 等级。复用同一个结构体，handler 决定 JSON 字段名
// （inherent_* / float_*），跟 Laravel 接口字段对齐。
type RiskNumResult struct {
	RiskNum   float64
	RiskGrade string
}

func (s *RiskMonitorService) inherentOrFloatNum(ctx context.Context, rootID int64, gradeColumn, valueColumn string) (RiskNumResult, error) {
	ids, err := s.produceCompanyIDs(ctx, rootID, gradeColumn)
	if err != nil {
		return RiskNumResult{}, err
	}
	if len(ids) == 0 {
		return RiskNumResult{RiskNum: 0, RiskGrade: gradeFromScore(0)}, nil
	}
	var buckets []gradeBucket
	if err := s.db.WithContext(ctx).
		Model(&model.Company{}).
		Select(gradeColumn+" AS risk_grade, COUNT(*) AS num, COALESCE(SUM("+valueColumn+"), 0) AS sum_value").
		Where("id IN ?", ids).
		Group(gradeColumn).
		Scan(&buckets).Error; err != nil {
		return RiskNumResult{}, err
	}
	score := computeIndex(buckets)
	return RiskNumResult{RiskNum: score, RiskGrade: gradeFromScore(score)}, nil
}

func (s *RiskMonitorService) InherentRiskNum(ctx context.Context, rootID int64) (RiskNumResult, error) {
	return s.inherentOrFloatNum(ctx, rootID, "inherent_risk_grade", "inherent_risk")
}

func (s *RiskMonitorService) FloatRiskNum(ctx context.Context, rootID int64) (RiskNumResult, error) {
	return s.inherentOrFloatNum(ctx, rootID, "float_risk_grade", "float_risk")
}

// GradeStatRow 集团端饼图行（仅 grade + num，前端按 grade 取色）。
// JSON 字段名跟 Laravel BlocRisksController.{inherent,float}RiskStatistics 对齐：
// inherent_risk_grade / float_risk_grade + num。handler 在序列化时切换 key。
type GradeStatRow struct {
	Grade string `json:"-"`
	Num   int64  `json:"num"`
}

func (s *RiskMonitorService) riskStatistics(ctx context.Context, rootID int64, enterpriseTypeID int64, gradeColumn string) ([]GradeStatRow, error) {
	ids, err := s.produceCompanyIDs(ctx, rootID, gradeColumn)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []GradeStatRow{}, nil
	}
	q := s.db.WithContext(ctx).
		Model(&model.Company{}).
		Select(gradeColumn+" AS grade, COUNT(*) AS num").
		Where("id IN ?", ids).
		Group(gradeColumn)
	if enterpriseTypeID > 0 {
		q = q.Where("enterprise_type_id = ?", enterpriseTypeID)
	}
	var rows []GradeStatRow
	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *RiskMonitorService) InherentRiskStatistics(ctx context.Context, rootID int64, enterpriseTypeID int64) ([]GradeStatRow, error) {
	return s.riskStatistics(ctx, rootID, enterpriseTypeID, "inherent_risk_grade")
}

func (s *RiskMonitorService) FloatRiskStatistics(ctx context.Context, rootID int64, enterpriseTypeID int64) ([]GradeStatRow, error) {
	return s.riskStatistics(ctx, rootID, enterpriseTypeID, "float_risk_grade")
}

// ─── 7 日风险趋势 ──────────────────────────────────────────────────────

// DayRiskRow 单日双类型风险快照（与前端 DayLog 字段一致）。
type DayRiskRow struct {
	Title             string  `json:"title"`
	InherentRiskNum   float64 `json:"inherent_risk_num"`
	InherentRiskGrade string  `json:"inherent_risk_grade"`
	FloatRiskNum      float64 `json:"float_risk_num"`
	FloatRiskGrade    string  `json:"float_risk_grade"`
}

// dayLogBucket 日志聚合行：date+type+grade 三键分组，num+sum 聚合。
type dayLogBucket struct {
	Date  int     `gorm:"column:date"`
	Type  int16   `gorm:"column:type"`
	Grade *string `gorm:"column:risk_grade"`
	Num   int64   `gorm:"column:num"`
	Sum   float64 `gorm:"column:sum_value"`
}

// RiskLog7Days 返回最近 7 天（含今日，按时间正序）的集团双类型风险曲线。
//
// total_count 对齐 Laravel：每天分母都是 rootID 子树中 status=1 的企业总数（与是否
// 有 grade 无关），跟 BlocRisksController.getRiskLog 行为一致。
func (s *RiskMonitorService) RiskLog7Days(ctx context.Context, rootID int64) ([]DayRiskRow, error) {
	descendants, err := s.companyRepo.DescendantIDs(ctx, rootID)
	if err != nil {
		return nil, err
	}
	var totalCount int64
	if len(descendants) > 0 {
		if err := s.db.WithContext(ctx).
			Model(&model.Company{}).
			Where("id IN ? AND status = ?", descendants, 1).
			Count(&totalCount).Error; err != nil {
			return nil, err
		}
	}

	dates := make([]int, 7)
	now := time.Now()
	for i := 6; i >= 0; i-- {
		d := now.AddDate(0, 0, -i)
		dates[6-i] = d.Year()*10000 + int(d.Month())*100 + d.Day()
	}

	rows := make([]DayRiskRow, len(dates))
	for i, d := range dates {
		rows[i] = DayRiskRow{
			Title:             fmt.Sprintf("%08d", d),
			InherentRiskGrade: gradeFromScore(0),
			FloatRiskGrade:    gradeFromScore(0),
		}
	}
	if totalCount == 0 || len(descendants) == 0 {
		return rows, nil
	}

	var buckets []dayLogBucket
	if err := s.db.WithContext(ctx).
		Model(&model.RiskLog{}).
		Select(`"date" AS date, "type" AS type, risk_grade AS risk_grade,
                COUNT(*) AS num, COALESCE(SUM(risk_num), 0) AS sum_value`).
		Where("company_id IN ? AND date IN ?", descendants, dates).
		Group(`"date", "type", risk_grade`).
		Scan(&buckets).Error; err != nil {
		return nil, err
	}

	// bucketed[date][type] = []bucket
	type key struct {
		Date int
		Type int16
	}
	byKey := make(map[key][]gradeBucket)
	for _, b := range buckets {
		k := key{Date: b.Date, Type: b.Type}
		byKey[k] = append(byKey[k], gradeBucket{Grade: b.Grade, Num: b.Num, Sum: b.Sum})
	}

	for i, d := range dates {
		inherentBuckets := byKey[key{Date: d, Type: 1}]
		floatBuckets := byKey[key{Date: d, Type: 2}]
		// Laravel 的算法分母是 totalCount（全部生产端企业），不是有日志的企业数。
		var inherentSum, floatSum float64
		for _, b := range inherentBuckets {
			inherentSum += b.Sum
		}
		for _, b := range floatBuckets {
			floatSum += b.Sum
		}
		inh := round2(inherentSum / float64(totalCount))
		flt := round2(floatSum / float64(totalCount))
		rows[i].InherentRiskNum = inh
		rows[i].InherentRiskGrade = gradeFromScore(inh)
		rows[i].FloatRiskNum = flt
		rows[i].FloatRiskGrade = gradeFromScore(flt)
	}
	return rows, nil
}

// ─── 隐患整改排序 ──────────────────────────────────────────────────────

// HiddenDangerSortRow 单个下属企业的隐患整改进度（与前端 HazardItem 字段对齐）。
type HiddenDangerSortRow struct {
	CompanyID                  int64   `json:"company_id"`
	CompanyName                string  `json:"company_name"`
	MonthlyInvestigationScale  float64 `json:"monthly_investigation_scale"`
	JQR                        float64 `json:"jqr"`
}

// hiddenDangerAggRow 单个企业的 4 字段聚合：total / total_old / resolved / resolved_old。
type hiddenDangerAggRow struct {
	CompanyID    int64 `gorm:"column:company_id"`
	Total        int64 `gorm:"column:total"`
	TotalOld     int64 `gorm:"column:total_old"`
	Resolved     int64 `gorm:"column:resolved"`
	ResolvedOld  int64 `gorm:"column:resolved_old"`
}

const correctResultRectified = "已整改"

// HiddenDangerSort 返回下属企业（不含根集团本身）的隐患整改完成率排序，
// 按 monthly_investigation_scale 降序。
func (s *RiskMonitorService) HiddenDangerSort(ctx context.Context, rootID int64) ([]HiddenDangerSortRow, error) {
	descendants, err := s.companyRepo.DescendantIDs(ctx, rootID)
	if err != nil {
		return nil, err
	}
	companyIDs := make([]int64, 0, len(descendants))
	for _, id := range descendants {
		if id != rootID {
			companyIDs = append(companyIDs, id)
		}
	}
	if len(companyIDs) == 0 {
		return []HiddenDangerSortRow{}, nil
	}

	today := time.Now().Format("2006-01-02")
	var aggs []hiddenDangerAggRow
	if err := s.db.WithContext(ctx).
		Model(&model.MonthlyInvestigation{}).
		Select(`company_id,
                COUNT(*) AS total,
                COUNT(*) FILTER (WHERE updated_at::date <> ?) AS total_old,
                COUNT(*) FILTER (WHERE correct_result = ?) AS resolved,
                COUNT(*) FILTER (WHERE correct_result = ? AND updated_at::date <> ?) AS resolved_old`,
			today, correctResultRectified, correctResultRectified, today).
		Where("company_id IN ?", companyIDs).
		Group("company_id").
		Scan(&aggs).Error; err != nil {
		return nil, err
	}
	aggByCompany := make(map[int64]hiddenDangerAggRow, len(aggs))
	for _, a := range aggs {
		aggByCompany[a.CompanyID] = a
	}

	var titles []struct {
		ID    int64  `gorm:"column:id"`
		Title string `gorm:"column:title"`
	}
	if err := s.db.WithContext(ctx).
		Model(&model.Company{}).
		Select("id, title").
		Where("id IN ?", companyIDs).
		Scan(&titles).Error; err != nil {
		return nil, err
	}
	titleByID := make(map[int64]string, len(titles))
	for _, t := range titles {
		titleByID[t.ID] = t.Title
	}

	rows := make([]HiddenDangerSortRow, 0, len(companyIDs))
	for _, cid := range companyIDs {
		agg := aggByCompany[cid]
		var scale float64
		if agg.Total > 0 && agg.Resolved > 0 {
			// 与 Laravel 一致：round(resolved/total, 2) * 100，先两位再 *100。
			scale = round2(float64(agg.Resolved)/float64(agg.Total)) * 100
		}
		var jqr float64
		if agg.TotalOld > 0 && agg.ResolvedOld > 0 {
			jqr = scale - round2(float64(agg.ResolvedOld)/float64(agg.TotalOld))*100
		}
		rows = append(rows, HiddenDangerSortRow{
			CompanyID:                 cid,
			CompanyName:               titleByID[cid],
			MonthlyInvestigationScale: scale,
			JQR:                       jqr,
		})
	}

	// 稳定降序：按 scale desc。
	for i := 0; i < len(rows); i++ {
		for j := i + 1; j < len(rows); j++ {
			if rows[j].MonthlyInvestigationScale > rows[i].MonthlyInvestigationScale {
				rows[i], rows[j] = rows[j], rows[i]
			}
		}
	}
	return rows, nil
}

// ─── 重点关注企业 ──────────────────────────────────────────────────────

type FocusCompanyListParams struct {
	Page      int
	PerPage   int
	CompanyID int64 // 集团 id
	FocusID   int64 // 可选：只查某一关注企业关系
}

func (s *RiskMonitorService) ListFocusCompanies(ctx context.Context, p FocusCompanyListParams) ([]model.FocusCompany, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.FocusCompany{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if p.FocusID > 0 {
		q = q.Where("focus_id = ?", p.FocusID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.FocusCompany
	if err := q.
		Preload("Company", func(db *gorm.DB) *gorm.DB {
			// 仅取面板需要的字段。
			return db.Select("id", "title", "inherent_risk", "inherent_risk_grade", "float_risk", "float_risk_grade")
		}).
		Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *RiskMonitorService) CreateFocusCompany(ctx context.Context, fc *model.FocusCompany) error {
	if fc.CompanyID == 0 || fc.FocusID == nil || *fc.FocusID == 0 {
		return errs.New(errs.CodeUnprocessable, "company_id 与 focus_id 不能为空")
	}
	if err := s.db.WithContext(ctx).Create(fc).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "写入失败", err)
	}
	return nil
}

func (s *RiskMonitorService) DeleteFocusCompany(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.FocusCompany{}, id)
	if res.Error != nil {
		return errs.Wrap(errs.CodeInternal, "删除失败", res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "关注企业不存在")
	}
	return nil
}

// ─── 在线巡检告警 ──────────────────────────────────────────────────────

type InspectionAlerts struct {
	ExpiredStaffCerts int64 `json:"expired_staff_certs"`
	ExpiredEquipment  int64 `json:"expired_equipment"`
	UnrevisedHazards  int64 `json:"unrevised_hazards"`
}

// InspectionAlerts 计数：rootID 子树（含自身）中
//   - 过期人员证件     special_equipment_staff.valid_period < today
//   - 过期特种设备     special_equipment.valid_period < today
//   - 未整改隐患       monthly_investigations.correct_result <> '已整改'
func (s *RiskMonitorService) InspectionAlerts(ctx context.Context, rootID int64) (InspectionAlerts, error) {
	ids, err := s.companyRepo.DescendantIDs(ctx, rootID)
	if err != nil {
		return InspectionAlerts{}, err
	}
	if len(ids) == 0 {
		return InspectionAlerts{}, nil
	}
	today := time.Now().Format("2006-01-02")

	var out InspectionAlerts
	if err := s.db.WithContext(ctx).
		Model(&model.SpecialEquipmentStaff{}).
		Where("company_id IN ? AND valid_period < ?", ids, today).
		Count(&out.ExpiredStaffCerts).Error; err != nil {
		return InspectionAlerts{}, err
	}
	if err := s.db.WithContext(ctx).
		Model(&model.SpecialEquipment{}).
		Where("company_id IN ? AND valid_period < ?", ids, today).
		Count(&out.ExpiredEquipment).Error; err != nil {
		return InspectionAlerts{}, err
	}
	if err := s.db.WithContext(ctx).
		Model(&model.MonthlyInvestigation{}).
		Where("company_id IN ? AND (correct_result IS NULL OR correct_result <> ?)", ids, correctResultRectified).
		Count(&out.UnrevisedHazards).Error; err != nil {
		return InspectionAlerts{}, err
	}
	return out, nil
}

// ─── 特种设备数据统计 ──────────────────────────────────────────────────

// SpecialEquipmentStatsType 单一特种设备类型，含 id + title。
type SpecialEquipmentStatsType struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// SpecialEquipmentStatsRow 单个企业的设备数量分布（counts 按 string(type_id) 索引）。
type SpecialEquipmentStatsRow struct {
	CompanyID    int64            `json:"company_id"`
	CompanyTitle string           `json:"company_title"`
	Counts       map[string]int64 `json:"counts"`
	Total        int64            `json:"total"`
}

type SpecialEquipmentStatsResult struct {
	Types []SpecialEquipmentStatsType `json:"types"`
	Rows  []SpecialEquipmentStatsRow  `json:"rows"`
}

// SpecialEquipmentStatistics 按 rootID 的直接下级企业 × 全部 SpecialType 汇总
// 设备数量。无下级企业时退化为对自身企业聚合（与 Laravel 行为一致）。
func (s *RiskMonitorService) SpecialEquipmentStatistics(ctx context.Context, rootID int64) (SpecialEquipmentStatsResult, error) {
	type companyRow struct {
		ID    int64  `gorm:"column:id"`
		Title string `gorm:"column:title"`
	}
	var companies []companyRow
	if err := s.db.WithContext(ctx).
		Model(&model.Company{}).
		Select("id, title").
		Where("parent_id = ? AND status = ?", rootID, 1).
		Order("id").
		Scan(&companies).Error; err != nil {
		return SpecialEquipmentStatsResult{}, err
	}
	if len(companies) == 0 {
		if err := s.db.WithContext(ctx).
			Model(&model.Company{}).
			Select("id, title").
			Where("id = ?", rootID).
			Scan(&companies).Error; err != nil {
			return SpecialEquipmentStatsResult{}, err
		}
	}

	var types []SpecialEquipmentStatsType
	if err := s.db.WithContext(ctx).
		Table("special_types").
		Select("id, title").
		Where("deleted_at IS NULL").
		Order("id").
		Scan(&types).Error; err != nil {
		return SpecialEquipmentStatsResult{}, err
	}

	type aggRow struct {
		CompanyID     int64 `gorm:"column:company_id"`
		SpecialTypeID int64 `gorm:"column:special_type_id"`
		Count         int64 `gorm:"column:count"`
	}
	companyIDs := make([]int64, len(companies))
	for i, c := range companies {
		companyIDs[i] = c.ID
	}
	var aggs []aggRow
	if len(companyIDs) > 0 {
		if err := s.db.WithContext(ctx).
			Model(&model.SpecialEquipment{}).
			Select("company_id, special_type_id, COUNT(*) AS count").
			Where("company_id IN ?", companyIDs).
			Group("company_id, special_type_id").
			Scan(&aggs).Error; err != nil {
			return SpecialEquipmentStatsResult{}, err
		}
	}
	byCompanyType := make(map[int64]map[int64]int64, len(companyIDs))
	for _, a := range aggs {
		m := byCompanyType[a.CompanyID]
		if m == nil {
			m = make(map[int64]int64)
			byCompanyType[a.CompanyID] = m
		}
		m[a.SpecialTypeID] = a.Count
	}

	rows := make([]SpecialEquipmentStatsRow, 0, len(companies))
	for _, c := range companies {
		counts := make(map[string]int64, len(types))
		var total int64
		for _, t := range types {
			n := byCompanyType[c.ID][t.ID]
			counts[fmt.Sprintf("%d", t.ID)] = n
			total += n
		}
		rows = append(rows, SpecialEquipmentStatsRow{
			CompanyID:    c.ID,
			CompanyTitle: c.Title,
			Counts:       counts,
			Total:        total,
		})
	}
	return SpecialEquipmentStatsResult{Types: types, Rows: rows}, nil
}
