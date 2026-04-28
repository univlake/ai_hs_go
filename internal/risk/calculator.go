// Package risk hosts the daily risk-recalculation worker. It mirrors the
// Laravel CalculateRiskValue trait at app/Traits/CalculateRiskValue.php in
// the aifix_hs codebase — that file is the canonical spec; any divergence is
// a porting bug.
package risk

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

// Industry weight multipliers, keyed by Company.enterprise_type_id.
// Any unmatched id falls back to 1.0.
var industryWeights = map[int64]float64{
	1: 0.8,
	2: 0.6,
	3: 0.9,
}

// Phase-2 (inherent / structural) weight constants, keyed by metric.
// Magic numbers are inherited from the legacy implementation — extracted as
// named keys here so the math is auditable without changing it.
var inherentWeights = map[string]float64{
	"safety_architecture_value":         0.02496,
	"security_officer_scale":            0.007345,
	"already_obtain_people_scale":       0.007345,
	"last_year_income_scale":            0.007345,
	"funding_execution_scale":           0.007345,
	"production_standardization_max":    0.002665,
	"health_system_scale":               0.002665,
	"assess_detect_scale":               0.00533,
	"high_school_diploma_scale":         0.014235,
	"acreage_scale":                     0.007865,
	"staff_scale":                       0.007865,
	"safety_engineering_major_scale":    0.035035,
	"hazardous_processes_scale":         0.14775,
	"hazardous_process_scale":           0.04925,
	"hazardous_process_more_scale":      0.14775,
	"process_is_automatic_scale":        0.04925,
	"special_equipment_number_scale":    0.119,
	"equipment_more_number_scale":       0.119,
	"hazardous_chemical_scale":          0.158746,
	"hazardous_chemicals_number_scale":  0.0594405,
	"major_hazard_max":                  0.0198135,
}

// Phase-3 (float / dynamic) weight constants, keyed by metric.
var floatWeights = map[string]float64{
	"internal_meeting_scale":         0.031635,
	"internal_training_scale":        0.063365,
	"accident_bulletin_scale":        0.10075035,
	"accident_bulletin_severe_max":   0.058859415,
	"accident_report_url_scale":      0.017145235,
	"emergency_team_scale":           0.0441225,
	"emergency_drill_scale":          0.0441225,
	"job_management_scale":           0.228114,
	"monthly_investigation_scale":    0.083752,
	"three_together_over_scale":      0.048488,
	"unit_sum_number_scale":          0.012509904,
	"evaluation_unit_scale":          0.00509124,
	"evaluation_unit_school_scale":   0.030886856,
	"health_checkup_people_scale":    0.009345,
	"health_checkup_scale":           0.022962,
	"health_report_scale":            0.056693,
}

// Enum string values copied from the Laravel App\Enums namespace. Database
// rows store these literals, so the Go side compares against the same.
const (
	standardizationLevel1 = "一级"
	standardizationLevel2 = "二级"
	standardizationLevel3 = "三级"

	majorHazardLevel1 = "一级"
	majorHazardLevel2 = "二级"
	majorHazardLevel3 = "三级"
	majorHazardLevel4 = "四级"

	correctionRectified = "已整改"

	matterStateGas = "气态"
	fireClassA     = "甲类"
	fireClassB     = "乙类"

	jobTypeSpecialEight = "八大特殊作业"

	accidentLevelLight   = "轻伤事故"
	accidentLevelSerious = "重伤事故"
)

type Calculator struct {
	db     *gorm.DB
	logger *zap.Logger
	tz     *time.Location
}

func NewCalculator(db *gorm.DB, logger *zap.Logger, tz *time.Location) *Calculator {
	return &Calculator{db: db, logger: logger, tz: tz}
}

// runState holds per-run scratch space (bulk-loaded data, normalization tables).
// Kept off Calculator so concurrent Run invocations don't share state.
type runState struct {
	bulk         *bulkData
	riskMaxByCid map[int64]map[string]float64
	globalMax    map[string]float64
}

// Run executes a full daily risk recalculation pass. content is the
// human-readable trigger label persisted on each risk_logs row.
//
// Steps mirror app/Traits/CalculateRiskValue.php::countRisk:
//  1. Load all producing + active companies; bail if none.
//  2. Wipe today's existing risk_logs / inherent_risk_logs / float_risk_logs
//     so a re-run produces a clean set (failure is logged but not fatal).
//  3. Bulk-load every per-company aggregation in one shot.
//  4. Compute Phase-1 normalization tables.
//  5. Compute inherent (Phase 2) → write companies + log rows + risk_logs.
//  6. Compute float (Phase 3) → write companies + log rows + risk_logs.
//  7. Bulk-insert all log rows and the merged risk_logs.
//  8. Aggregate bloc averages from the just-written child values.
func (c *Calculator) Run(ctx context.Context, content string) error {
	now := time.Now()
	local := now.In(c.tz)
	currentYear := local.Year()

	startOfYear := time.Date(currentYear, 1, 1, 0, 0, 0, 0, c.tz)
	daysThisYear := int(local.Sub(startOfYear).Hours() / 24)
	if daysThisYear < 1 {
		daysThisYear = 1
	}
	tenYearsAgo := local.AddDate(-10, 0, 0).Format("2006-01-02")

	var companies []model.Company
	if err := c.db.WithContext(ctx).
		Where("is_produce = ? AND status = ?", 1, 1).
		Find(&companies).Error; err != nil {
		return err
	}
	if len(companies) == 0 {
		c.logger.Info("risk calc skipped (no producing companies)", zap.String("content", content))
		return nil
	}

	companyIDs := make([]int64, len(companies))
	for i := range companies {
		companyIDs[i] = companies[i].ID
	}

	c.resetTodayLogs(ctx, now)

	bulk, err := c.loadBulk(ctx, companyIDs, currentYear, tenYearsAgo)
	if err != nil {
		return err
	}

	state := &runState{bulk: bulk}
	state.computePhase1(companies, daysThisYear, currentYear, c.tz)

	inherentLogs, inherentRiskLogs, inherentUpdates := computeInherent(state, companies, now, content, c.tz)
	if err := c.batchUpdateInherent(ctx, inherentUpdates); err != nil {
		return err
	}
	floatLogs, floatRiskLogs, floatUpdates := computeFloat(state, companies, now, content, c.tz)
	if err := c.batchUpdateFloat(ctx, floatUpdates); err != nil {
		return err
	}

	if len(inherentLogs) > 0 {
		if err := c.db.WithContext(ctx).CreateInBatches(&inherentLogs, 500).Error; err != nil {
			return err
		}
	}
	if len(floatLogs) > 0 {
		if err := c.db.WithContext(ctx).CreateInBatches(&floatLogs, 500).Error; err != nil {
			return err
		}
	}
	riskLogs := append(inherentRiskLogs, floatRiskLogs...)
	if len(riskLogs) > 0 {
		if err := c.db.WithContext(ctx).CreateInBatches(&riskLogs, 500).Error; err != nil {
			return err
		}
	}

	if err := c.computeBlocs(ctx, content, now); err != nil {
		return err
	}

	c.logger.Info("risk calc completed",
		zap.String("content", content),
		zap.Int("companies", len(companies)),
		zap.Int("inherent_log_rows", len(inherentLogs)),
		zap.Int("float_log_rows", len(floatLogs)),
		zap.Int("risk_log_rows", len(riskLogs)),
	)
	return nil
}

// resetTodayLogs soft-deletes today's pending log rows so a re-run produces a
// single fresh set. Failures are logged but not fatal — matches the trait,
// since the inserts below will fail loudly if the tables are actually unwritable.
func (c *Calculator) resetTodayLogs(ctx context.Context, now time.Time) {
	start, end := dayRange(now, c.tz)
	db := c.db.WithContext(ctx)

	if err := db.Where("created_at >= ? AND created_at < ?", start, end).
		Delete(&model.RiskLog{}).Error; err != nil {
		c.logger.Warn("risk calc: reset of today's risk_logs failed", zap.Error(err))
	}
	if err := db.Where("created_at >= ? AND created_at < ?", start, end).
		Delete(&model.FloatRiskLog{}).Error; err != nil {
		c.logger.Warn("risk calc: reset of today's float_risk_logs failed", zap.Error(err))
	}
	if err := db.Where("created_at >= ? AND created_at < ?", start, end).
		Delete(&model.InherentRiskLog{}).Error; err != nil {
		c.logger.Warn("risk calc: reset of today's inherent_risk_logs failed", zap.Error(err))
	}
}

// computePhase1 builds per-company raw metric maps + a global per-metric
// maximum used as the normalization denominator. Pure in-memory pass over
// the bulk data — no DB I/O. Mirrors the trait's computePhase1Metrics.
//
// daysThisYear is the elapsed days since Jan 1 of `currentYear`, floor()'d
// and clamped to ≥1 (matches the trait); used to scale yearly counts to a
// 30-day window for meeting/training/job and to a 365-day window for accidents.
func (s *runState) computePhase1(companies []model.Company, daysThisYear int, currentYear int, tz *time.Location) {
	byCid := make(map[int64]map[string]float64, len(companies))
	maxKeys := []string{
		"security_officer_scale", "already_obtain_people_scale", "last_year_income_scale",
		"acreage_scale", "hazardous_processes_scale", "hazardous_chemicals_scale",
		"internal_meeting_scale", "internal_training_scale", "accident_bulletin_scale",
		"emergency_team_scale", "emergency_drill_scale", "job_management_scale",
		"health_checkup_scale",
	}
	globalMax := make(map[string]float64, len(maxKeys))
	daysThisYearF := float64(daysThisYear)

	for i := range companies {
		company := &companies[i]
		cid := company.ID
		row := map[string]float64{}
		personnel := s.bulk.personnel[cid]

		if personnel != nil && personnel.TotalPeople > 0 && personnel.SecurityOfficer > 0 {
			row["security_officer_scale"] = float64(personnel.SecurityOfficer) / float64(personnel.TotalPeople)
		}

		if cert := s.bulk.cert[cid]; cert != nil && personnel != nil &&
			cert.AlreadyObtainPeople != nil && *cert.AlreadyObtainPeople > 0 && personnel.SecurityOfficer > 0 {
			row["already_obtain_people_scale"] = float64(*cert.AlreadyObtainPeople) / float64(personnel.SecurityOfficer)
		}

		if sec, ok := s.bulk.secInv[cid]; ok && sec.BudgetSum > 0 && sec.IncomeSum > 0 {
			row["last_year_income_scale"] = sec.BudgetSum / sec.IncomeSum * float64(sec.Cnt)
		}

		if personnel != nil && company.Acreage != nil && *company.Acreage > 0 && personnel.TotalPeople > 0 {
			row["acreage_scale"] = float64(personnel.TotalPeople) / *company.Acreage
		}

		totalHazProc := s.bulk.majHazProcTotal[cid] + s.bulk.hazProcTotal[cid]
		if personnel != nil && totalHazProc > 0 && personnel.TotalPeople > 0 {
			row["hazardous_processes_scale"] = float64(totalHazProc) / float64(personnel.TotalPeople)
		}

		var hazChemCount int64
		if hcs, ok := s.bulk.hazChemStats[cid]; ok {
			hazChemCount = hcs.Total
		}
		if personnel != nil && hazChemCount > 0 && personnel.TotalPeople > 0 {
			row["hazardous_chemicals_scale"] = float64(hazChemCount) / float64(personnel.TotalPeople)
		}

		if cnt := s.bulk.internalMeeting[cid]; cnt > 0 {
			row["internal_meeting_scale"] = float64(cnt) / daysThisYearF * 30
		}
		if cnt := s.bulk.internalTraining[cid]; cnt > 0 {
			row["internal_training_scale"] = float64(cnt) / daysThisYearF * 30
		}

		var accidentCount int64
		if acc, ok := s.bulk.accidentStats[cid]; ok {
			accidentCount = acc.Total
		}
		if accidentCount > 0 {
			row["accident_bulletin_scale"] = float64(accidentCount) / daysThisYearF * 365
		}

		if teamSum := s.bulk.emergencyTeam[cid]; personnel != nil && teamSum > 0 && personnel.TotalPeople > 0 {
			row["emergency_team_scale"] = teamSum / float64(personnel.TotalPeople)
		}

		drillCount := s.bulk.emergencyDrill[cid]
		years := currentYear - company.CreatedAt.In(tz).Year()
		if years < 1 {
			years = 1
		}
		if drillCount > 0 {
			row["emergency_drill_scale"] = float64(drillCount) / float64(years)
		}

		if cnt := s.bulk.jobManagement[cid]; cnt > 0 {
			row["job_management_scale"] = float64(cnt) / daysThisYearF * 30
		}

		if occ := s.bulk.occHaz[cid]; personnel != nil && occ > 0 && personnel.TotalPeople > 0 {
			row["health_checkup_scale"] = occ / float64(personnel.TotalPeople)
		}

		byCid[cid] = row
		for _, k := range maxKeys {
			if v := row[k]; v > globalMax[k] {
				globalMax[k] = v
			}
		}
	}

	s.riskMaxByCid = byCid
	s.globalMax = globalMax
}

// normalized returns self[key] / globalMax[key], or 0 if either side is 0.
func (s *runState) normalized(self map[string]float64, key string) float64 {
	val := self[key]
	max := s.globalMax[key]
	if val <= 0 || max <= 0 {
		return 0
	}
	return val / max
}

// normalizedInverse returns 1 - (self[key] / globalMax[key]), or 0 if either
// side is 0. Used for "more is worse" metrics where a higher raw value should
// reduce the contribution.
func (s *runState) normalizedInverse(self map[string]float64, key string) float64 {
	val := self[key]
	max := s.globalMax[key]
	if val <= 0 || max <= 0 {
		return 0
	}
	return 1 - val/max
}

// industryWeight returns the industry multiplier for a company's
// enterprise_type_id, falling back to 1.0 for nil or unknown ids.
func industryWeight(enterpriseTypeID *int64) float64 {
	if enterpriseTypeID == nil {
		return 1.0
	}
	if w, ok := industryWeights[*enterpriseTypeID]; ok {
		return w
	}
	return 1.0
}

// riskGrade buckets a risk score into one of four labels. Boundaries are
// strict-less-than: <5 红, <10 橙, <15 黄, ≥15 蓝.
func riskGrade(value float64) string {
	switch {
	case value < 5:
		return "红"
	case value < 10:
		return "橙"
	case value < 15:
		return "黄"
	default:
		return "蓝"
	}
}

// yearRange returns [start, end) timestamps for the given calendar year in
// the calculator's timezone. Used in place of `whereYear(...)` so the
// timestamp index stays usable.
func yearRange(year int, tz *time.Location) (time.Time, time.Time) {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, tz)
	end := time.Date(year+1, 1, 1, 0, 0, 0, 0, tz)
	return start, end
}

// dayRange returns [start, end) timestamps for the given day in the
// calculator's timezone, used for the today's-logs reset.
func dayRange(day time.Time, tz *time.Location) (time.Time, time.Time) {
	d := day.In(tz)
	start := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, tz)
	end := start.Add(24 * time.Hour)
	return start, end
}

// round2 rounds to 2 decimal places. Mirrors PHP `round($x, 2)` half-away-from-zero.
func round2(v float64) float64 {
	if v >= 0 {
		return float64(int64(v*100+0.5)) / 100
	}
	return float64(int64(v*100-0.5)) / 100
}
