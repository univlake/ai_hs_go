package risk

import (
	"context"
	"database/sql"
	"math"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

// newMockGorm spins up a postgres-flavored gorm wrapped over sqlmock with
// regex query matching. PreferSimpleProtocol keeps placeholders as $N so
// the regex patterns below stay portable across pgx tweaks.
func newMockGorm(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	t.Helper()
	raw, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("sqlmock new: %v", err)
	}
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 raw,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open: %v", err)
	}
	return gdb, mock, raw
}

func TestCaseWhenUpdateEmpty(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()
	if err := caseWhenUpdate(gdb, "companies", "inherent_risk", nil); err != nil {
		t.Errorf("nil values returned err: %v", err)
	}
	if err := caseWhenUpdate(gdb, "companies", "inherent_risk", map[int64]any{}); err != nil {
		t.Errorf("empty values returned err: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("no expectations should fire: %v", err)
	}
}

// TestCaseWhenUpdateSingleRow: with a single map entry, iteration order is
// deterministic — we can pin the full SQL and the bound args.
func TestCaseWhenUpdateSingleRow(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	want := regexp.QuoteMeta(`UPDATE companies SET "inherent_risk" = CASE id WHEN $1 THEN $2 END WHERE id IN (42)`)
	mock.ExpectExec(want).
		WithArgs(int64(42), 8.42).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := caseWhenUpdate(gdb, "companies", "inherent_risk", map[int64]any{42: 8.42}); err != nil {
		t.Fatalf("caseWhenUpdate: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestCaseWhenUpdateMultiRow: map iteration order is non-deterministic, so
// we assert structure (right table/column, right number of WHEN clauses,
// IN-list members, arg count) without pinning order.
func TestCaseWhenUpdateMultiRow(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	// Three WHEN ? THEN ? pairs, IN list with the same 3 ids in any order.
	pattern := `UPDATE companies SET "float_risk_grade" = CASE id` +
		`( WHEN \$\d+ THEN \$\d+){3} END WHERE id IN \([0-9, ]+\)`
	mock.ExpectExec(pattern).
		WithArgs(
			sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(),
		).
		WillReturnResult(sqlmock.NewResult(0, 3))

	updates := map[int64]any{1: "红", 2: "橙", 3: "黄"}
	if err := caseWhenUpdate(gdb, "companies", "float_risk_grade", updates); err != nil {
		t.Fatalf("caseWhenUpdate: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestBatchUpdateInherentEmpty: nil/empty updates short-circuit before any tx
// is opened. ExpectationsWereMet would fail if a BEGIN slipped through.
func TestBatchUpdateInherentEmpty(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()
	c := &Calculator{db: gdb}
	if err := c.batchUpdateInherent(context.Background(), nil); err != nil {
		t.Errorf("nil updates: %v", err)
	}
	if err := c.batchUpdateInherent(context.Background(), map[int64]inherentUpdate{}); err != nil {
		t.Errorf("empty updates: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("no DB calls should fire: %v", err)
	}
}

func TestBatchUpdateFloatEmpty(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()
	c := &Calculator{db: gdb}
	if err := c.batchUpdateFloat(context.Background(), nil); err != nil {
		t.Errorf("nil updates: %v", err)
	}
	if err := c.batchUpdateFloat(context.Background(), map[int64]floatUpdate{}); err != nil {
		t.Errorf("empty updates: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("no DB calls should fire: %v", err)
	}
}

// TestBatchUpdateInherentTwoColumns: confirms the wrapper opens a transaction
// and issues exactly two UPDATE statements (risk + grade) inside it.
func TestBatchUpdateInherentTwoColumns(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()
	c := &Calculator{db: gdb}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE companies SET "inherent_risk" = CASE id WHEN \$1 THEN \$2 END WHERE id IN \(7\)`).
		WithArgs(int64(7), 4.5).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`UPDATE companies SET "inherent_risk_grade" = CASE id WHEN \$1 THEN \$2 END WHERE id IN \(7\)`).
		WithArgs(int64(7), "红").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	updates := map[int64]inherentUpdate{7: {risk: 4.5, grade: "红"}}
	if err := c.batchUpdateInherent(context.Background(), updates); err != nil {
		t.Fatalf("batchUpdateInherent: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestComputeInherentBranches: each subtest activates exactly one
// contribution branch and asserts that field's weighted value. Companions
// to TestComputeInherentMinimal (which locks the all-zeros defaults) and
// the *Ladder tests (which lock priority-ordered switches).
func TestComputeInherentBranches(t *testing.T) {
	now := time.Date(2026, 4, 28, 0, 0, 0, 0, time.UTC)
	mkCompany := func() model.Company {
		return model.Company{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}}
	}

	cases := []struct {
		name      string
		bulk      *bulkData
		self      map[string]float64
		global    map[string]float64
		wantField func(model.InherentRiskLog) (float64, float64)
	}{
		{
			name: "SafetyArchitectureValue when arch count > 0",
			bulk: &bulkData{safetyArch: map[int64]int64{1: 1}},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.SafetyArchitectureValue, inherentWeights["safety_architecture_value"]
			},
		},
		{
			name:   "SecurityOfficerScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"security_officer_scale": 0.04},
			global: map[string]float64{"security_officer_scale": 0.08},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.SecurityOfficerScale, 0.5 * inherentWeights["security_officer_scale"]
			},
		},
		{
			name:   "AlreadyObtainPeopleScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"already_obtain_people_scale": 0.5},
			global: map[string]float64{"already_obtain_people_scale": 1.0},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.AlreadyObtainPeopleScale, 0.5 * inherentWeights["already_obtain_people_scale"]
			},
		},
		{
			name:   "LastYearIncomeScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"last_year_income_scale": 2.0},
			global: map[string]float64{"last_year_income_scale": 8.0},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.LastYearIncomeScale, 0.25 * inherentWeights["last_year_income_scale"]
			},
		},
		{
			name: "FundingExecutionScale when funding+budget both > 0",
			bulk: &bulkData{
				fundingExec: map[int64]float64{1: 80},
				secInv:      map[int64]secInvAgg{1: {CompanyID: 1, BudgetSum: 200}},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.FundingExecutionScale, 0.4 * inherentWeights["funding_execution_scale"]
			},
		},
		{
			name: "HealthSystemScale when row present",
			bulk: &bulkData{
				healthSystem: map[int64]*model.HealthSystem{1: {}},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HealthSystemScale, inherentWeights["health_system_scale"]
			},
		},
		{
			name: "AssessDetectScale when count > 0",
			bulk: &bulkData{assessDetect: map[int64]int64{1: 3}},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.AssessDetectScale, inherentWeights["assess_detect_scale"]
			},
		},
		{
			name: "HighSchoolDiplomaScale ratio when both > 0",
			bulk: &bulkData{
				personnel: map[int64]*model.Personnel{1: {TotalPeople: 100, HighSchoolDiploma: 60}},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HighSchoolDiplomaScale, 0.6 * inherentWeights["high_school_diploma_scale"]
			},
		},
		{
			name:   "AcreageScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"acreage_scale": 0.1},
			global: map[string]float64{"acreage_scale": 0.5},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.AcreageScale, 0.2 * inherentWeights["acreage_scale"]
			},
		},
		{
			name: "StaffScale ratio when both > 0",
			bulk: &bulkData{
				personnel: map[int64]*model.Personnel{1: {TotalPeople: 100, Staff: 25}},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.StaffScale, 0.25 * inherentWeights["staff_scale"]
			},
		},
		{
			name: "SafetyEngineeringMajorScale ratio when both > 0",
			bulk: &bulkData{
				personnel: map[int64]*model.Personnel{1: {SecurityOfficer: 10, SafetyEngineeringMajor: 2}},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.SafetyEngineeringMajorScale, 0.2 * inherentWeights["safety_engineering_major_scale"]
			},
		},
		{
			name:   "HazardousProcessesScale via normalizedInverse",
			bulk:   &bulkData{},
			self:   map[string]float64{"hazardous_processes_scale": 2},
			global: map[string]float64{"hazardous_processes_scale": 5},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HazardousProcessesScale, 0.6 * inherentWeights["hazardous_processes_scale"]
			},
		},
		{
			name: "HazardousProcessScale: 1 - majHaz/(majHaz+haz)",
			bulk: &bulkData{
				majHazProcTotal: map[int64]int64{1: 2},
				hazProcTotal:    map[int64]int64{1: 8},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HazardousProcessScale, 0.8 * inherentWeights["hazardous_process_scale"]
			},
		},
		{
			name: "HazardousProcessMoreScale: 1 - over/total",
			bulk: &bulkData{
				majHazProcTotal: map[int64]int64{1: 2}, // procSum = 10
				hazProcTotal:    map[int64]int64{1: 8},
				majHazProcOver:  map[int64]int64{1: 1},
				hazProcOver:     map[int64]int64{1: 4}, // overSum = 5
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HazardousProcessMoreScale, 0.5 * inherentWeights["hazardous_process_more_scale"]
			},
		},
		{
			name: "ProcessIsAutomaticScale: auto/procSum",
			bulk: &bulkData{
				majHazProcTotal: map[int64]int64{1: 4}, // procSum = 10
				hazProcTotal:    map[int64]int64{1: 6},
				hazProcAuto:     map[int64]int64{1: 3},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.ProcessIsAutomaticScale, 0.3 * inherentWeights["process_is_automatic_scale"]
			},
		},
		{
			name: "SpecialEquipmentNumberScale: 1 - sp/(sp+dang)",
			bulk: &bulkData{
				spEqTotal:   map[int64]int64{1: 3}, // eqSum = 10
				dangEqTotal: map[int64]int64{1: 7},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.SpecialEquipmentNumberScale, 0.7 * inherentWeights["special_equipment_number_scale"]
			},
		},
		{
			name: "EquipmentMoreNumberScale: 1 - oldSum/eqSum",
			bulk: &bulkData{
				spEqTotal:   map[int64]int64{1: 5}, // eqSum = 10
				dangEqTotal: map[int64]int64{1: 5},
				spEqOld:     map[int64]int64{1: 1},
				dangEqOld:   map[int64]int64{1: 3}, // oldSum = 4
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.EquipmentMoreNumberScale, 0.6 * inherentWeights["equipment_more_number_scale"]
			},
		},
		{
			name: "HazardousChemicalScale: 1 - danger/total",
			bulk: &bulkData{
				hazChemStats: map[int64]hazChemStats{
					1: {CompanyID: 1, Total: 10, ExplosionCnt: 1, ToxinGasCnt: 1, FireACnt: 0, FireBCnt: 1},
				},
			},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HazardousChemicalScale, 0.7 * inherentWeights["hazardous_chemical_scale"]
			},
		},
		{
			name:   "HazardousChemicalsNumberScale via normalizedInverse",
			bulk:   &bulkData{},
			self:   map[string]float64{"hazardous_chemicals_scale": 1},
			global: map[string]float64{"hazardous_chemicals_scale": 4},
			wantField: func(r model.InherentRiskLog) (float64, float64) {
				return r.HazardousChemicalsNumberScale, 0.75 * inherentWeights["hazardous_chemicals_number_scale"]
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			self := c.self
			if self == nil {
				self = map[string]float64{}
			}
			global := c.global
			if global == nil {
				global = map[string]float64{}
			}
			state := &runState{
				bulk:         c.bulk,
				riskMaxByCid: map[int64]map[string]float64{1: self},
				globalMax:    global,
			}
			logs, _, _ := computeInherent(state, []model.Company{mkCompany()}, now, "t", time.UTC)
			got, want := c.wantField(logs[0])
			if math.Abs(got-want) > 1e-9 {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// TestComputeFloatBranches: same shape as TestComputeInherentBranches,
// activates one Phase-3 contribution per subtest.
func TestComputeFloatBranches(t *testing.T) {
	now := time.Date(2026, 4, 28, 0, 0, 0, 0, time.UTC)
	mkCompany := func() model.Company {
		return model.Company{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}}
	}

	cases := []struct {
		name      string
		bulk      *bulkData
		self      map[string]float64
		global    map[string]float64
		wantField func(model.FloatRiskLog) (float64, float64)
	}{
		{
			name:   "InternalMeetingScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"internal_meeting_scale": 3},
			global: map[string]float64{"internal_meeting_scale": 6},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.InternalMeetingScale, 0.5 * floatWeights["internal_meeting_scale"]
			},
		},
		{
			name:   "InternalTrainingScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"internal_training_scale": 1},
			global: map[string]float64{"internal_training_scale": 4},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.InternalTrainingScale, 0.25 * floatWeights["internal_training_scale"]
			},
		},
		{
			name:   "AccidentBulletinScale via normalizedInverse",
			bulk:   &bulkData{},
			self:   map[string]float64{"accident_bulletin_scale": 2},
			global: map[string]float64{"accident_bulletin_scale": 5},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.AccidentBulletinScale, 0.6 * floatWeights["accident_bulletin_scale"]
			},
		},
		{
			name: "AccidentReportURLScale when reported_count > 0",
			bulk: &bulkData{
				accidentStats: map[int64]accidentStats{1: {CompanyID: 1, ReportedCnt: 1, OtherCnt: 1}},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.AccidentReportURLScale, floatWeights["accident_report_url_scale"]
			},
		},
		{
			name:   "EmergencyTeamScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"emergency_team_scale": 0.4},
			global: map[string]float64{"emergency_team_scale": 0.8},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.EmergencyTeamScale, 0.5 * floatWeights["emergency_team_scale"]
			},
		},
		{
			name:   "EmergencyDrillScale via normalized",
			bulk:   &bulkData{},
			self:   map[string]float64{"emergency_drill_scale": 2},
			global: map[string]float64{"emergency_drill_scale": 8},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.EmergencyDrillScale, 0.25 * floatWeights["emergency_drill_scale"]
			},
		},
		{
			name:   "JobManagementScale via normalizedInverse",
			bulk:   &bulkData{},
			self:   map[string]float64{"job_management_scale": 1.5},
			global: map[string]float64{"job_management_scale": 5},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.JobManagementScale, 0.7 * floatWeights["job_management_scale"]
			},
		},
		{
			name: "MonthlyInvestigationScale: done/total",
			bulk: &bulkData{
				monthlyInvTotal: map[int64]int64{1: 10},
				monthlyInvDone:  map[int64]int64{1: 4},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.MonthlyInvestigationScale, 0.4 * floatWeights["monthly_investigation_scale"]
			},
		},
		{
			name: "ThreeTogetherOverScale: over/total",
			bulk: &bulkData{
				threeTotal: map[int64]int64{1: 10},
				threeOver:  map[int64]int64{1: 3},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.ThreeTogetherOverScale, 0.3 * floatWeights["three_together_over_scale"]
			},
		},
		{
			name: "UnitSumNumberScale: sum-qual / sum-total across 3 unit kinds",
			bulk: &bulkData{
				skillUnitTotal:  map[int64]int64{1: 4}, // total = 12
				evalUnitTotal:   map[int64]int64{1: 6},
				designUnitTotal: map[int64]int64{1: 2},
				skillUnitQual:   map[int64]int64{1: 1}, // qual = 3
				evalUnitQual:    map[int64]int64{1: 1},
				designUnitQual:  map[int64]int64{1: 1},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.UnitSumNumberScale, 0.25 * floatWeights["unit_sum_number_scale"]
			},
		},
		{
			name: "EvaluationUnitScale: agree/total",
			bulk: &bulkData{
				evalUnitTotal: map[int64]int64{1: 8},
				evalUnitAgree: map[int64]int64{1: 2},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.EvaluationUnitScale, 0.25 * floatWeights["evaluation_unit_scale"]
			},
		},
		{
			name: "EvaluationUnitSchoolScale: trained/total",
			bulk: &bulkData{
				evalUnitTotal:   map[int64]int64{1: 5},
				evalUnitTrained: map[int64]int64{1: 1},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.EvaluationUnitSchoolScale, 0.2 * floatWeights["evaluation_unit_school_scale"]
			},
		},
		{
			name: "HealthCheckupPeopleScale: checkupSum/totalPeople",
			bulk: &bulkData{
				personnel:     map[int64]*model.Personnel{1: {TotalPeople: 100}},
				healthCheckup: map[int64]float64{1: 30},
			},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.HealthCheckupPeopleScale, 0.3 * floatWeights["health_checkup_people_scale"]
			},
		},
		{
			name:   "HealthCheckupScale via normalizedInverse",
			bulk:   &bulkData{},
			self:   map[string]float64{"health_checkup_scale": 0.2},
			global: map[string]float64{"health_checkup_scale": 0.8},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.HealthCheckupScale, 0.75 * floatWeights["health_checkup_scale"]
			},
		},
		{
			name: "HealthReportScale: any reports filed → 0",
			bulk: &bulkData{healthReport: map[int64]int64{1: 1}},
			wantField: func(r model.FloatRiskLog) (float64, float64) {
				return r.HealthReportScale, 0
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			self := c.self
			if self == nil {
				self = map[string]float64{}
			}
			global := c.global
			if global == nil {
				global = map[string]float64{}
			}
			state := &runState{
				bulk:         c.bulk,
				riskMaxByCid: map[int64]map[string]float64{1: self},
				globalMax:    global,
			}
			logs, _, _ := computeFloat(state, []model.Company{mkCompany()}, now, "t", time.UTC)
			got, want := c.wantField(logs[0])
			if math.Abs(got-want) > 1e-9 {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestBatchUpdateFloatTwoColumns(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()
	c := &Calculator{db: gdb}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE companies SET "float_risk" = CASE id WHEN \$1 THEN \$2 END WHERE id IN \(9\)`).
		WithArgs(int64(9), 12.34).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`UPDATE companies SET "float_risk_grade" = CASE id WHEN \$1 THEN \$2 END WHERE id IN \(9\)`).
		WithArgs(int64(9), "黄").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	updates := map[int64]floatUpdate{9: {risk: 12.34, grade: "黄"}}
	if err := c.batchUpdateFloat(context.Background(), updates); err != nil {
		t.Fatalf("batchUpdateFloat: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
