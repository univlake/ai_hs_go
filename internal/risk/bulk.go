package risk

import (
	"context"
	"strconv"

	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

// bulkData holds every per-company aggregation Phase 1/2/3 needs, loaded
// in one batch up front. Mirrors the $this->bulk array in the Laravel trait.
type bulkData struct {
	personnel  map[int64]*model.Personnel
	cert       map[int64]*model.CertificateStatus
	productionStd  map[int64]*model.ProductionStandardization
	healthSystem   map[int64]*model.HealthSystem
	majorHazard    map[int64][]string

	secInv         map[int64]secInvAgg
	fundingExec    map[int64]float64
	hazChemStats   map[int64]hazChemStats
	accidentStats  map[int64]accidentStats

	majHazProcTotal  map[int64]int64
	majHazProcOver   map[int64]int64
	hazProcTotal     map[int64]int64
	hazProcOver      map[int64]int64
	hazProcAuto      map[int64]int64
	internalMeeting  map[int64]int64
	internalTraining map[int64]int64
	emergencyTeam    map[int64]float64
	emergencyDrill   map[int64]int64
	jobManagement    map[int64]int64
	occHaz           map[int64]float64
	safetyArch       map[int64]int64
	assessDetect     map[int64]int64
	spEqTotal        map[int64]int64
	spEqOld          map[int64]int64
	dangEqTotal      map[int64]int64
	dangEqOld        map[int64]int64
	monthlyInvTotal  map[int64]int64
	monthlyInvDone   map[int64]int64
	threeTotal       map[int64]int64
	threeOver        map[int64]int64
	skillUnitTotal   map[int64]int64
	skillUnitQual    map[int64]int64
	evalUnitTotal    map[int64]int64
	evalUnitQual     map[int64]int64
	evalUnitAgree    map[int64]int64
	evalUnitTrained  map[int64]int64
	designUnitTotal  map[int64]int64
	designUnitQual   map[int64]int64
	healthCheckup    map[int64]float64
	healthReport     map[int64]int64
}

type secInvAgg struct {
	CompanyID int64   `gorm:"column:company_id"`
	BudgetSum float64 `gorm:"column:budget_sum"`
	IncomeSum float64 `gorm:"column:income_sum"`
	Cnt       int64   `gorm:"column:cnt"`
}

type hazChemStats struct {
	CompanyID    int64 `gorm:"column:company_id"`
	Total        int64 `gorm:"column:total"`
	ExplosionCnt int64 `gorm:"column:explosion_cnt"`
	ToxinGasCnt  int64 `gorm:"column:toxin_gas_cnt"`
	FireACnt     int64 `gorm:"column:fire_a_cnt"`
	FireBCnt     int64 `gorm:"column:fire_b_cnt"`
}

type accidentStats struct {
	CompanyID   int64 `gorm:"column:company_id"`
	Total       int64 `gorm:"column:total"`
	LightCnt    int64 `gorm:"column:light_cnt"`
	SeriousCnt  int64 `gorm:"column:serious_cnt"`
	OtherCnt    int64 `gorm:"column:other_cnt"`
	ReportedCnt int64 `gorm:"column:reported_cnt"`
}

// countRow / sumRow are scratch row types for GROUP BY company_id queries.
type countRow struct {
	CompanyID int64 `gorm:"column:company_id"`
	Cnt       int64 `gorm:"column:cnt"`
}

type sumRow struct {
	CompanyID int64   `gorm:"column:company_id"`
	Total     float64 `gorm:"column:total"`
}

// loadBulk runs every aggregation needed downstream, keyed by company_id.
//
// `currentYear` filters yearly buckets (security_investments, accidents, etc.)
// — passed in rather than re-derived so all callers see the same year value.
// `tenYearsAgo` is a YYYY-MM-DD cutoff for old-equipment counts. Both are
// already adjusted for the calculator's timezone by the caller.
//
// Several columns (operating_people_num, team_number, three_togethers.status)
// are stored as VARCHAR but the legacy PHP queries compared them numerically.
// Postgres won't auto-coerce, so the SQL below casts explicitly — flagged in
// docs/risk-calc-migration-plan.md Phase 0.
func (c *Calculator) loadBulk(ctx context.Context, companyIDs []int64, currentYear int, tenYearsAgo string) (*bulkData, error) {
	db := c.db.WithContext(ctx)
	yearStr := strconv.Itoa(currentYear)
	yearStart, yearEnd := yearRange(currentYear, c.tz)

	b := &bulkData{
		personnel:        map[int64]*model.Personnel{},
		cert:             map[int64]*model.CertificateStatus{},
		productionStd:    map[int64]*model.ProductionStandardization{},
		healthSystem:     map[int64]*model.HealthSystem{},
		majorHazard:      map[int64][]string{},
		secInv:           map[int64]secInvAgg{},
		fundingExec:      map[int64]float64{},
		hazChemStats:     map[int64]hazChemStats{},
		accidentStats:    map[int64]accidentStats{},
		majHazProcTotal:  map[int64]int64{},
		majHazProcOver:   map[int64]int64{},
		hazProcTotal:     map[int64]int64{},
		hazProcOver:      map[int64]int64{},
		hazProcAuto:      map[int64]int64{},
		internalMeeting:  map[int64]int64{},
		internalTraining: map[int64]int64{},
		emergencyTeam:    map[int64]float64{},
		emergencyDrill:   map[int64]int64{},
		jobManagement:    map[int64]int64{},
		occHaz:           map[int64]float64{},
		safetyArch:       map[int64]int64{},
		assessDetect:     map[int64]int64{},
		spEqTotal:        map[int64]int64{},
		spEqOld:          map[int64]int64{},
		dangEqTotal:      map[int64]int64{},
		dangEqOld:        map[int64]int64{},
		monthlyInvTotal:  map[int64]int64{},
		monthlyInvDone:   map[int64]int64{},
		threeTotal:       map[int64]int64{},
		threeOver:        map[int64]int64{},
		skillUnitTotal:   map[int64]int64{},
		skillUnitQual:    map[int64]int64{},
		evalUnitTotal:    map[int64]int64{},
		evalUnitQual:     map[int64]int64{},
		evalUnitAgree:    map[int64]int64{},
		evalUnitTrained:  map[int64]int64{},
		designUnitTotal:  map[int64]int64{},
		designUnitQual:   map[int64]int64{},
		healthCheckup:    map[int64]float64{},
		healthReport:     map[int64]int64{},
	}

	// Singletons-per-company: full row keyed by company_id.
	var personnel []model.Personnel
	if err := db.Where("company_id IN ?", companyIDs).Find(&personnel).Error; err != nil {
		return nil, err
	}
	for i := range personnel {
		b.personnel[personnel[i].CompanyID] = &personnel[i]
	}

	var certs []model.CertificateStatus
	if err := db.Where("company_id IN ? AND type = ? AND name = ?", companyIDs, 1, "安全管理员").
		Find(&certs).Error; err != nil {
		return nil, err
	}
	for i := range certs {
		b.cert[certs[i].CompanyID] = &certs[i]
	}

	var prodStd []model.ProductionStandardization
	if err := db.Where("company_id IN ?", companyIDs).Find(&prodStd).Error; err != nil {
		return nil, err
	}
	for i := range prodStd {
		b.productionStd[prodStd[i].CompanyID] = &prodStd[i]
	}

	var healthSys []model.HealthSystem
	if err := db.Where("company_id IN ?", companyIDs).Find(&healthSys).Error; err != nil {
		return nil, err
	}
	for i := range healthSys {
		b.healthSystem[healthSys[i].CompanyID] = &healthSys[i]
	}

	// MajorHazard: list of `level` strings per company (nil levels skipped).
	var majorHazards []model.MajorHazard
	if err := db.Where("company_id IN ?", companyIDs).Find(&majorHazards).Error; err != nil {
		return nil, err
	}
	for i := range majorHazards {
		if majorHazards[i].Level != nil {
			b.majorHazard[majorHazards[i].CompanyID] = append(b.majorHazard[majorHazards[i].CompanyID], *majorHazards[i].Level)
		}
	}

	// security_investments: sum budget + sum last_year_income + count, per company, this year.
	// year column is VARCHAR — compare with string literal.
	var secRows []secInvAgg
	if err := db.Table("security_investments").
		Select("company_id, COALESCE(SUM(budget),0) AS budget_sum, COALESCE(SUM(last_year_income),0) AS income_sum, COUNT(*) AS cnt").
		Where("company_id IN ? AND year = ? AND deleted_at IS NULL", companyIDs, yearStr).
		Group("company_id").Scan(&secRows).Error; err != nil {
		return nil, err
	}
	for _, r := range secRows {
		b.secInv[r.CompanyID] = r
	}

	if err := scanSumMap(db.Table("funding_executions").
		Select("company_id, COALESCE(SUM(budget),0) AS total").
		Where("company_id IN ? AND year = ? AND deleted_at IS NULL", companyIDs, yearStr).
		Group("company_id"), b.fundingExec); err != nil {
		return nil, err
	}

	// hazardous_processes: total / over-3-people / is_automatic; same shape for major_hazardous_processes.
	// operating_people_num is VARCHAR, cast for the >3 comparison.
	if err := scanCountMap(db.Table("major_hazardous_processes").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.majHazProcTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("major_hazardous_processes").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND CAST(NULLIF(operating_people_num, '') AS INTEGER) > ?", companyIDs, 3).
		Group("company_id"), b.majHazProcOver); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("hazardous_processes").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.hazProcTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("hazardous_processes").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND CAST(NULLIF(operating_people_num, '') AS INTEGER) > ?", companyIDs, 3).
		Group("company_id"), b.hazProcOver); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("hazardous_processes").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND is_automatic = ?", companyIDs, 1).
		Group("company_id"), b.hazProcAuto); err != nil {
		return nil, err
	}

	// hazardous_chemicals: per-company total + 4 conditional counts (mirrors the trait's CASE WHEN payload).
	var hazChemRows []hazChemStats
	if err := db.Table("hazardous_chemicals").
		Select(`company_id,
		        COUNT(*) AS total,
		        SUM(CASE WHEN is_explosion = 1 THEN 1 ELSE 0 END) AS explosion_cnt,
		        SUM(CASE WHEN is_toxin = 1 AND matter_state = ? THEN 1 ELSE 0 END) AS toxin_gas_cnt,
		        SUM(CASE WHEN fire_class = ? AND matter_state = ? THEN 1 ELSE 0 END) AS fire_a_cnt,
		        SUM(CASE WHEN fire_class = ? AND matter_state = ? THEN 1 ELSE 0 END) AS fire_b_cnt`,
			matterStateGas, fireClassA, matterStateGas, fireClassB, matterStateGas).
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id").Scan(&hazChemRows).Error; err != nil {
		return nil, err
	}
	for _, r := range hazChemRows {
		b.hazChemStats[r.CompanyID] = r
	}

	// year-scoped meeting / training counts; range filter so the timestamp index is usable.
	if err := scanCountMap(db.Table("internal_meetings").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND created_at >= ? AND created_at < ?", companyIDs, yearStart, yearEnd).
		Group("company_id"), b.internalMeeting); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("internal_trainings").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND created_at >= ? AND created_at < ?", companyIDs, yearStart, yearEnd).
		Group("company_id"), b.internalTraining); err != nil {
		return nil, err
	}

	// accident_bulletins: per-company year-scoped totals + 3 level buckets + reported count.
	var accRows []accidentStats
	if err := db.Table("accident_bulletins").
		Select(`company_id,
		        COUNT(*) AS total,
		        SUM(CASE WHEN accident_level = ? THEN 1 ELSE 0 END) AS light_cnt,
		        SUM(CASE WHEN accident_level = ? THEN 1 ELSE 0 END) AS serious_cnt,
		        SUM(CASE WHEN accident_level NOT IN (?, ?) THEN 1 ELSE 0 END) AS other_cnt,
		        SUM(CASE WHEN accident_report_url IS NOT NULL THEN 1 ELSE 0 END) AS reported_cnt`,
			accidentLevelLight, accidentLevelSerious, accidentLevelLight, accidentLevelSerious).
		Where("company_id IN ? AND deleted_at IS NULL AND created_at >= ? AND created_at < ?", companyIDs, yearStart, yearEnd).
		Group("company_id").Scan(&accRows).Error; err != nil {
		return nil, err
	}
	for _, r := range accRows {
		b.accidentStats[r.CompanyID] = r
	}

	// emergency_teams.team_number is VARCHAR — cast/strip empties before SUM.
	if err := scanSumMap(db.Table("emergency_teams").
		Select("company_id, COALESCE(SUM(COALESCE(NULLIF(team_number, ''), '0')::numeric), 0) AS total").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.emergencyTeam); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("emergency_drills").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.emergencyDrill); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("job_management").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND type = ? AND created_at >= ? AND created_at < ?",
			companyIDs, jobTypeSpecialEight, yearStart, yearEnd).
		Group("company_id"), b.jobManagement); err != nil {
		return nil, err
	}
	if err := scanSumMap(db.Table("occupational_hazards").
		Select("company_id, COALESCE(SUM(number),0) AS total").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.occHaz); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("safety_architectures").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.safetyArch); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("assess_detects").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.assessDetect); err != nil {
		return nil, err
	}

	// special / dangerous equipment: total and "use_time older than 10 years ago".
	if err := scanCountMap(db.Table("special_equipment").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.spEqTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("special_equipment").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND use_time < ?", companyIDs, tenYearsAgo).
		Group("company_id"), b.spEqOld); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("dangerous_equipment").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.dangEqTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("dangerous_equipment").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND use_time < ?", companyIDs, tenYearsAgo).
		Group("company_id"), b.dangEqOld); err != nil {
		return nil, err
	}

	// monthly investigations: total + closed (correct_result = '已整改').
	if err := scanCountMap(db.Table("monthly_investigations").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.monthlyInvTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("monthly_investigations").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND correct_result = ?", companyIDs, correctionRectified).
		Group("company_id"), b.monthlyInvDone); err != nil {
		return nil, err
	}

	// three_togethers.status is VARCHAR(255); Laravel's where('status', 4) becomes string compare.
	if err := scanCountMap(db.Table("three_togethers").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.threeTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("three_togethers").
		Select("company_id, COUNT(*) AS cnt").
		Where(`company_id IN ? AND deleted_at IS NULL AND "status" = ?`, companyIDs, "4").
		Group("company_id"), b.threeOver); err != nil {
		return nil, err
	}

	// SkillUnit + has('kids'): only count units that have at least one (non-deleted) kid row.
	if err := scanCountMap(db.Table("skill_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.skillUnitTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("skill_units AS su").
		Select("su.company_id AS company_id, COUNT(*) AS cnt").
		Where("su.company_id IN ? AND su.deleted_at IS NULL", companyIDs).
		Where("EXISTS (SELECT 1 FROM skill_unit_kids k WHERE k.skill_unit_id = su.id AND k.deleted_at IS NULL)").
		Group("su.company_id"), b.skillUnitQual); err != nil {
		return nil, err
	}

	if err := scanCountMap(db.Table("evaluation_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.evalUnitTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("evaluation_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND is_qualification = ?", companyIDs, 1).
		Group("company_id"), b.evalUnitQual); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("evaluation_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND agreement_annex_url IS NOT NULL", companyIDs).
		Group("company_id"), b.evalUnitAgree); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("evaluation_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND is_agreement = ?", companyIDs, 1).
		Group("company_id"), b.evalUnitTrained); err != nil {
		return nil, err
	}

	if err := scanCountMap(db.Table("design_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.designUnitTotal); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("design_units").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL AND qualification_name IS NOT NULL", companyIDs).
		Group("company_id"), b.designUnitQual); err != nil {
		return nil, err
	}

	if err := scanSumMap(db.Table("health_checkups").
		Select("company_id, COALESCE(SUM(number),0) AS total").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.healthCheckup); err != nil {
		return nil, err
	}
	if err := scanCountMap(db.Table("health_reports").
		Select("company_id, COUNT(*) AS cnt").
		Where("company_id IN ? AND deleted_at IS NULL", companyIDs).
		Group("company_id"), b.healthReport); err != nil {
		return nil, err
	}

	return b, nil
}

func scanCountMap(tx *gorm.DB, dst map[int64]int64) error {
	var rows []countRow
	if err := tx.Scan(&rows).Error; err != nil {
		return err
	}
	for _, r := range rows {
		dst[r.CompanyID] = r.Cnt
	}
	return nil
}

func scanSumMap(tx *gorm.DB, dst map[int64]float64) error {
	var rows []sumRow
	if err := tx.Scan(&rows).Error; err != nil {
		return err
	}
	for _, r := range rows {
		dst[r.CompanyID] = r.Total
	}
	return nil
}
