package risk

import (
	"context"
	"slices"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

// inherentUpdate is the per-company write produced by Phase 2.
type inherentUpdate struct {
	risk  float64
	grade string
}

// floatUpdate is the per-company write produced by Phase 3.
type floatUpdate struct {
	risk  float64
	grade string
}

// computeInherent walks every company, computing 21 inherent-risk metrics.
// Mirrors the trait's computeInherentRisk: each contribution is conditional
// (skip when source data is missing) and accumulated into a *log row + a
// risk_log delta. Final score is round(sum * industryWeight * 100, 2).
//
// Pure compute — returns the per-company log rows, type=1 risk_logs deltas,
// and the company UPDATE payload, but does not persist anything. The caller
// (Run) flushes via batchUpdateInherent + CreateInBatches. Split this way so
// tests can lock down the math without a live DB.
func computeInherent(
	state *runState,
	companies []model.Company,
	now time.Time,
	content string,
	tz *time.Location,
) ([]model.InherentRiskLog, []model.RiskLog, map[int64]inherentUpdate) {
	logs := make([]model.InherentRiskLog, 0, len(companies))
	riskLogs := make([]model.RiskLog, 0, len(companies))
	updates := make(map[int64]inherentUpdate, len(companies))
	dateInt := dateInt(now, tz)

	for i := range companies {
		company := &companies[i]
		cid := company.ID
		personnel := state.bulk.personnel[cid]
		self := state.riskMaxByCid[cid]
		w := industryWeight(company.EnterpriseTypeID)

		var row model.InherentRiskLog
		row.CompanyID = cid
		row.CreatedAt = now
		row.UpdatedAt = now

		if state.bulk.safetyArch[cid] > 0 {
			row.SafetyArchitectureValue = inherentWeights["safety_architecture_value"]
		}

		row.SecurityOfficerScale = state.normalized(self, "security_officer_scale") * inherentWeights["security_officer_scale"]
		row.AlreadyObtainPeopleScale = state.normalized(self, "already_obtain_people_scale") * inherentWeights["already_obtain_people_scale"]
		row.LastYearIncomeScale = state.normalized(self, "last_year_income_scale") * inherentWeights["last_year_income_scale"]

		fundingSum := state.bulk.fundingExec[cid]
		var secInvBudget float64
		if sec, ok := state.bulk.secInv[cid]; ok {
			secInvBudget = sec.BudgetSum
		}
		if fundingSum > 0 && secInvBudget > 0 {
			row.FundingExecutionScale = fundingSum / secInvBudget * inherentWeights["funding_execution_scale"]
		}

		if prodStd := state.bulk.productionStd[cid]; prodStd != nil && prodStd.Grade != nil {
			switch *prodStd.Grade {
			case standardizationLevel1:
				row.ProductionStandardizationScale = 0.9 * inherentWeights["production_standardization_max"]
			case standardizationLevel2:
				row.ProductionStandardizationScale = 0.75 * inherentWeights["production_standardization_max"]
			case standardizationLevel3:
				row.ProductionStandardizationScale = 0.6 * inherentWeights["production_standardization_max"]
			}
		}

		if _, ok := state.bulk.healthSystem[cid]; ok {
			row.HealthSystemScale = inherentWeights["health_system_scale"]
		}

		if state.bulk.assessDetect[cid] > 0 {
			row.AssessDetectScale = inherentWeights["assess_detect_scale"]
		}

		if personnel != nil && personnel.HighSchoolDiploma > 0 && personnel.TotalPeople > 0 {
			row.HighSchoolDiplomaScale = float64(personnel.HighSchoolDiploma) / float64(personnel.TotalPeople) *
				inherentWeights["high_school_diploma_scale"]
		}

		row.AcreageScale = state.normalized(self, "acreage_scale") * inherentWeights["acreage_scale"]

		if personnel != nil && personnel.Staff > 0 && personnel.TotalPeople > 0 {
			row.StaffScale = float64(personnel.Staff) / float64(personnel.TotalPeople) * inherentWeights["staff_scale"]
		}

		if personnel != nil && personnel.SafetyEngineeringMajor > 0 && personnel.SecurityOfficer > 0 {
			row.SafetyEngineeringMajorScale = float64(personnel.SafetyEngineeringMajor) /
				float64(personnel.SecurityOfficer) * inherentWeights["safety_engineering_major_scale"]
		}

		row.HazardousProcessesScale = state.normalizedInverse(self, "hazardous_processes_scale") *
			inherentWeights["hazardous_processes_scale"]

		majHazCount := state.bulk.majHazProcTotal[cid]
		hazProcCount := state.bulk.hazProcTotal[cid]
		procSum := majHazCount + hazProcCount
		if majHazCount > 0 && procSum > 0 {
			row.HazardousProcessScale = (1 - float64(majHazCount)/float64(procSum)) *
				inherentWeights["hazardous_process_scale"]
		}

		majHazOver := state.bulk.majHazProcOver[cid]
		hazProcOver := state.bulk.hazProcOver[cid]
		overSum := majHazOver + hazProcOver
		if overSum > 0 && procSum > 0 {
			row.HazardousProcessMoreScale = (1 - float64(overSum)/float64(procSum)) *
				inherentWeights["hazardous_process_more_scale"]
		}

		hazProcAuto := state.bulk.hazProcAuto[cid]
		if hazProcAuto > 0 && procSum > 0 {
			row.ProcessIsAutomaticScale = float64(hazProcAuto) / float64(procSum) *
				inherentWeights["process_is_automatic_scale"]
		}

		spEqCount := state.bulk.spEqTotal[cid]
		dangEqCount := state.bulk.dangEqTotal[cid]
		eqSum := spEqCount + dangEqCount
		if spEqCount > 0 && eqSum > 0 {
			row.SpecialEquipmentNumberScale = (1 - float64(spEqCount)/float64(eqSum)) *
				inherentWeights["special_equipment_number_scale"]
		}

		spEqOldCount := state.bulk.spEqOld[cid]
		dangEqOldCount := state.bulk.dangEqOld[cid]
		oldSum := spEqOldCount + dangEqOldCount
		if oldSum > 0 && eqSum > 0 {
			row.EquipmentMoreNumberScale = (1 - float64(oldSum)/float64(eqSum)) *
				inherentWeights["equipment_more_number_scale"]
		}

		var hazChemCount, hazChemDangerCount int64
		if hcs, ok := state.bulk.hazChemStats[cid]; ok {
			hazChemCount = hcs.Total
			hazChemDangerCount = hcs.ExplosionCnt + hcs.ToxinGasCnt + hcs.FireACnt + hcs.FireBCnt
		}
		if hazChemDangerCount > 0 && hazChemCount > 0 {
			row.HazardousChemicalScale = (1 - float64(hazChemDangerCount)/float64(hazChemCount)) *
				inherentWeights["hazardous_chemical_scale"]
		}

		row.HazardousChemicalsNumberScale = state.normalizedInverse(self, "hazardous_chemicals_scale") *
			inherentWeights["hazardous_chemicals_number_scale"]

		levels := state.bulk.majorHazard[cid]
		switch {
		case slices.Contains(levels, majorHazardLevel1):
			row.MajorHazardScale = 0.6 * inherentWeights["major_hazard_max"]
		case slices.Contains(levels, majorHazardLevel2):
			row.MajorHazardScale = 0.7 * inherentWeights["major_hazard_max"]
		case slices.Contains(levels, majorHazardLevel3):
			row.MajorHazardScale = 0.8 * inherentWeights["major_hazard_max"]
		case slices.Contains(levels, majorHazardLevel4):
			row.MajorHazardScale = 0.9 * inherentWeights["major_hazard_max"]
		default:
			row.MajorHazardScale = 1.0 * inherentWeights["major_hazard_max"]
		}

		sum := row.SafetyArchitectureValue + row.SecurityOfficerScale + row.AlreadyObtainPeopleScale +
			row.LastYearIncomeScale + row.FundingExecutionScale + row.ProductionStandardizationScale +
			row.HealthSystemScale + row.AssessDetectScale + row.HighSchoolDiplomaScale + row.AcreageScale +
			row.StaffScale + row.SafetyEngineeringMajorScale + row.HazardousProcessesScale +
			row.HazardousProcessScale + row.HazardousProcessMoreScale + row.ProcessIsAutomaticScale +
			row.SpecialEquipmentNumberScale + row.EquipmentMoreNumberScale + row.HazardousChemicalScale +
			row.HazardousChemicalsNumberScale + row.MajorHazardScale
		score := round2(sum * w * 100)
		grade := riskGrade(score)

		logs = append(logs, row)
		riskLogs = append(riskLogs, buildRiskLog(company, 1, company.InherentRisk, score,
			company.InherentRiskGrade, grade, content, dateInt, now))

		updates[cid] = inherentUpdate{risk: score, grade: grade}
	}
	return logs, riskLogs, updates
}

// computeFloat mirrors computeInherent for the 17 dynamic-risk metrics.
// Final score is round(sum * 100, 2) — no industry multiplier on Phase 3.
//
// Pure compute, same split rationale as computeInherent.
func computeFloat(
	state *runState,
	companies []model.Company,
	now time.Time,
	content string,
	tz *time.Location,
) ([]model.FloatRiskLog, []model.RiskLog, map[int64]floatUpdate) {
	logs := make([]model.FloatRiskLog, 0, len(companies))
	riskLogs := make([]model.RiskLog, 0, len(companies))
	updates := make(map[int64]floatUpdate, len(companies))
	dateInt := dateInt(now, tz)

	for i := range companies {
		company := &companies[i]
		cid := company.ID
		personnel := state.bulk.personnel[cid]
		self := state.riskMaxByCid[cid]

		var row model.FloatRiskLog
		row.CompanyID = cid
		row.CreatedAt = now
		row.UpdatedAt = now

		row.InternalMeetingScale = state.normalized(self, "internal_meeting_scale") * floatWeights["internal_meeting_scale"]
		row.InternalTrainingScale = state.normalized(self, "internal_training_scale") * floatWeights["internal_training_scale"]
		row.AccidentBulletinScale = state.normalizedInverse(self, "accident_bulletin_scale") * floatWeights["accident_bulletin_scale"]

		var lightCount, seriousCount, otherCount, reportedCount int64
		if acc, ok := state.bulk.accidentStats[cid]; ok {
			lightCount = acc.LightCnt
			seriousCount = acc.SeriousCnt
			otherCount = acc.OtherCnt
			reportedCount = acc.ReportedCnt
		}
		switch {
		case otherCount > 0:
			row.AccidentBulletinSevereScale = 0
		case seriousCount > 0:
			row.AccidentBulletinSevereScale = 0.3 * floatWeights["accident_bulletin_severe_max"]
		case lightCount > 0:
			row.AccidentBulletinSevereScale = 0.6 * floatWeights["accident_bulletin_severe_max"]
		default:
			row.AccidentBulletinSevereScale = 1.0 * floatWeights["accident_bulletin_severe_max"]
		}
		if reportedCount > 0 {
			row.AccidentReportURLScale = floatWeights["accident_report_url_scale"]
		}

		row.EmergencyTeamScale = state.normalized(self, "emergency_team_scale") * floatWeights["emergency_team_scale"]
		row.EmergencyDrillScale = state.normalized(self, "emergency_drill_scale") * floatWeights["emergency_drill_scale"]
		row.JobManagementScale = state.normalizedInverse(self, "job_management_scale") * floatWeights["job_management_scale"]
		// risk_source_scale is hard-zeroed in the trait — keep that explicit.
		row.RiskSourceScale = 0

		invTotal := state.bulk.monthlyInvTotal[cid]
		invDone := state.bulk.monthlyInvDone[cid]
		if invTotal > 0 && invDone > 0 {
			row.MonthlyInvestigationScale = float64(invDone) / float64(invTotal) * floatWeights["monthly_investigation_scale"]
		}

		threeTotal := state.bulk.threeTotal[cid]
		threeOver := state.bulk.threeOver[cid]
		if threeOver > 0 && threeTotal > 0 {
			row.ThreeTogetherOverScale = float64(threeOver) / float64(threeTotal) * floatWeights["three_together_over_scale"]
		}

		skillQual := state.bulk.skillUnitQual[cid]
		skillTotal := state.bulk.skillUnitTotal[cid]
		evalQual := state.bulk.evalUnitQual[cid]
		evalTotal := state.bulk.evalUnitTotal[cid]
		designQual := state.bulk.designUnitQual[cid]
		designTotal := state.bulk.designUnitTotal[cid]
		unitSumQual := skillQual + evalQual + designQual
		unitSumTotal := skillTotal + evalTotal + designTotal
		if unitSumQual > 0 && unitSumTotal > 0 {
			row.UnitSumNumberScale = float64(unitSumQual) / float64(unitSumTotal) * floatWeights["unit_sum_number_scale"]
		}

		evalAgree := state.bulk.evalUnitAgree[cid]
		if evalAgree > 0 && evalTotal > 0 {
			row.EvaluationUnitScale = float64(evalAgree) / float64(evalTotal) * floatWeights["evaluation_unit_scale"]
		}

		evalTrained := state.bulk.evalUnitTrained[cid]
		if evalTrained > 0 && evalTotal > 0 {
			row.EvaluationUnitSchoolScale = float64(evalTrained) / float64(evalTotal) * floatWeights["evaluation_unit_school_scale"]
		}

		healthCheckupPeople := state.bulk.healthCheckup[cid]
		if healthCheckupPeople > 0 && personnel != nil && personnel.TotalPeople > 0 {
			row.HealthCheckupPeopleScale = healthCheckupPeople / float64(personnel.TotalPeople) *
				floatWeights["health_checkup_people_scale"]
		}

		row.HealthCheckupScale = state.normalizedInverse(self, "health_checkup_scale") * floatWeights["health_checkup_scale"]

		// "No occupational-health report filed" awards full weight (legacy semantics).
		if state.bulk.healthReport[cid] == 0 {
			row.HealthReportScale = floatWeights["health_report_scale"]
		}

		sum := row.InternalMeetingScale + row.InternalTrainingScale + row.AccidentBulletinScale +
			row.AccidentBulletinSevereScale + row.AccidentReportURLScale + row.EmergencyTeamScale +
			row.EmergencyDrillScale + row.JobManagementScale + row.RiskSourceScale +
			row.MonthlyInvestigationScale + row.ThreeTogetherOverScale + row.UnitSumNumberScale +
			row.EvaluationUnitScale + row.EvaluationUnitSchoolScale + row.HealthCheckupPeopleScale +
			row.HealthCheckupScale + row.HealthReportScale
		score := round2(sum * 100)
		grade := riskGrade(score)

		logs = append(logs, row)
		riskLogs = append(riskLogs, buildRiskLog(company, 2, company.FloatRisk, score,
			company.FloatRiskGrade, grade, content, dateInt, now))

		updates[cid] = floatUpdate{risk: score, grade: grade}
	}
	return logs, riskLogs, updates
}

// buildRiskLog constructs the per-company risk_logs delta row matching the
// trait's buildRiskLogRow shape. type=1 inherent / type=2 float.
func buildRiskLog(
	company *model.Company,
	logType int16,
	oldRisk *float64,
	newRisk float64,
	oldGrade *string,
	newGrade string,
	content string,
	date int,
	now time.Time,
) model.RiskLog {
	contentPtr := content
	companyName := company.Title
	return model.RiskLog{
		CompanyID:    company.ID,
		Type:         logType,
		OldRiskNum:   oldRisk,
		RiskNum:      newRisk,
		Date:         date,
		Content:      &contentPtr,
		CompanyName:  &companyName,
		OldRiskGrade: oldGrade,
		RiskGrade:    &newGrade,
		Audit: model.Audit{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

// dateInt formats the day in YYYYMMDD as an int — matches Laravel's `date('Ymd')`.
func dateInt(now time.Time, tz *time.Location) int {
	d := now.In(tz)
	return d.Year()*10000 + int(d.Month())*100 + d.Day()
}

// batchUpdateInherent applies the per-company inherent_risk + grade updates
// inside a single transaction. Uses a CASE WHEN per attribute to keep the
// number of round-trips at 2 (one per column), matching the trait.
func (c *Calculator) batchUpdateInherent(ctx context.Context, updates map[int64]inherentUpdate) error {
	if len(updates) == 0 {
		return nil
	}
	risks := make(map[int64]any, len(updates))
	grades := make(map[int64]any, len(updates))
	for id, u := range updates {
		risks[id] = u.risk
		grades[id] = u.grade
	}
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := caseWhenUpdate(tx, "companies", "inherent_risk", risks); err != nil {
			return err
		}
		return caseWhenUpdate(tx, "companies", "inherent_risk_grade", grades)
	})
}

func (c *Calculator) batchUpdateFloat(ctx context.Context, updates map[int64]floatUpdate) error {
	if len(updates) == 0 {
		return nil
	}
	risks := make(map[int64]any, len(updates))
	grades := make(map[int64]any, len(updates))
	for id, u := range updates {
		risks[id] = u.risk
		grades[id] = u.grade
	}
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := caseWhenUpdate(tx, "companies", "float_risk", risks); err != nil {
			return err
		}
		return caseWhenUpdate(tx, "companies", "float_risk_grade", grades)
	})
}

// caseWhenUpdate executes
//
//	UPDATE <table> SET <col> = CASE id WHEN ? THEN ? ... END WHERE id IN (...)
//
// — one round-trip per column for any size of `values`. Mirrors the trait's
// batchUpdateCompanies. The ID list is interpolated as integers so binding
// count stays at 2 per company + zero for the IN list.
func caseWhenUpdate(tx *gorm.DB, table, column string, values map[int64]any) error {
	if len(values) == 0 {
		return nil
	}
	ids := make([]int64, 0, len(values))
	args := make([]any, 0, 2*len(values))
	caseSQL := ""
	for id, v := range values {
		caseSQL += " WHEN ? THEN ?"
		args = append(args, id, v)
		ids = append(ids, id)
	}
	idsList := intsCSV(ids)
	sql := "UPDATE " + table + ` SET "` + column + `" = CASE id` + caseSQL + " END WHERE id IN (" + idsList + ")"
	return tx.Exec(sql, args...).Error
}

func intsCSV(ids []int64) string {
	parts := make([]string, len(ids))
	for i, id := range ids {
		parts[i] = strconv.FormatInt(id, 10)
	}
	return strings.Join(parts, ",")
}
