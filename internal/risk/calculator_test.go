package risk

import (
	"math"
	"testing"
	"time"

	"ai_hs_go/internal/model"
)

func TestRiskGrade(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{0, "红"},
		{4.99, "红"},
		{5, "橙"},
		{9.99, "橙"},
		{10, "黄"},
		{14.99, "黄"},
		{15, "蓝"},
		{100, "蓝"},
	}
	for _, c := range cases {
		if got := riskGrade(c.in); got != c.want {
			t.Errorf("riskGrade(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNormalized(t *testing.T) {
	s := &runState{globalMax: map[string]float64{"k": 10, "z": 0}}
	cases := []struct {
		name string
		self map[string]float64
		key  string
		want float64
	}{
		{"happy path", map[string]float64{"k": 5}, "k", 0.5},
		{"val zero", map[string]float64{"k": 0}, "k", 0},
		{"val negative", map[string]float64{"k": -1}, "k", 0},
		{"max zero", map[string]float64{"z": 5}, "z", 0},
		{"key missing in self", map[string]float64{}, "k", 0},
		{"key missing everywhere", map[string]float64{}, "missing", 0},
	}
	for _, c := range cases {
		if got := s.normalized(c.self, c.key); got != c.want {
			t.Errorf("%s: normalized = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestNormalizedInverse(t *testing.T) {
	s := &runState{globalMax: map[string]float64{"k": 10}}
	if got := s.normalizedInverse(map[string]float64{"k": 4}, "k"); math.Abs(got-0.6) > 1e-9 {
		t.Errorf("normalizedInverse(4/10) = %v, want 0.6", got)
	}
	// zero-guard mirrors normalized
	if got := s.normalizedInverse(map[string]float64{"k": 0}, "k"); got != 0 {
		t.Errorf("normalizedInverse(0/10) = %v, want 0", got)
	}
	if got := s.normalizedInverse(map[string]float64{"k": 5}, "missing"); got != 0 {
		t.Errorf("normalizedInverse missing max = %v, want 0", got)
	}
}

func TestIndustryWeight(t *testing.T) {
	id1 := int64(1)
	id2 := int64(2)
	id3 := int64(3)
	idMissing := int64(99)
	cases := []struct {
		name string
		in   *int64
		want float64
	}{
		{"nil → 1.0", nil, 1.0},
		{"id=1 → 0.8", &id1, 0.8},
		{"id=2 → 0.6", &id2, 0.6},
		{"id=3 → 0.9", &id3, 0.9},
		{"unknown id → 1.0", &idMissing, 1.0},
	}
	for _, c := range cases {
		if got := industryWeight(c.in); got != c.want {
			t.Errorf("%s: industryWeight = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestSubtreeDepth(t *testing.T) {
	t.Run("leaf returns 0", func(t *testing.T) {
		if d := subtreeDepth(1, companyTree{}); d != 0 {
			t.Errorf("leaf depth = %d, want 0", d)
		}
	})
	t.Run("single child returns 1", func(t *testing.T) {
		tree := companyTree{1: {2}}
		if d := subtreeDepth(1, tree); d != 1 {
			t.Errorf("single-child depth = %d, want 1", d)
		}
	})
	t.Run("two-level chain returns 2", func(t *testing.T) {
		tree := companyTree{1: {2}, 2: {3}}
		if d := subtreeDepth(1, tree); d != 2 {
			t.Errorf("chain depth = %d, want 2", d)
		}
	})
	t.Run("diamond — descendant shared via two paths", func(t *testing.T) {
		// 1 → 2,3; 2 → 4; 3 → 4 (shared leaf)
		// Both branches resolve to depth 2 from the root.
		tree := companyTree{1: {2, 3}, 2: {4}, 3: {4}}
		if d := subtreeDepth(1, tree); d != 2 {
			t.Errorf("diamond depth = %d, want 2", d)
		}
	})
	t.Run("uneven branches takes max", func(t *testing.T) {
		// 1 → 2 (leaf), 3 → 4 → 5 (depth 2 from 3, depth 3 from 1)
		tree := companyTree{1: {2, 3}, 3: {4}, 4: {5}}
		if d := subtreeDepth(1, tree); d != 3 {
			t.Errorf("uneven depth = %d, want 3", d)
		}
	})
}

func TestRound2(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{0, 0},
		{1.234, 1.23},
		{1.235, 1.24},
		{1.245, 1.25},
		{-1.235, -1.24},
		{99.999, 100},
		{0.005, 0.01},
		{0.004, 0},
	}
	for _, c := range cases {
		got := round2(c.in)
		if math.Abs(got-c.want) > 1e-9 {
			t.Errorf("round2(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestDateInt(t *testing.T) {
	tz := time.FixedZone("Asia/Shanghai", 8*3600)
	// 2026-04-28 00:30 in UTC = 2026-04-28 08:30 in CST → date 20260428.
	utc := time.Date(2026, 4, 28, 0, 30, 0, 0, time.UTC)
	if got := dateInt(utc, tz); got != 20260428 {
		t.Errorf("dateInt(UTC midnight-ish, CST tz) = %d, want 20260428", got)
	}
	// 2026-04-28 23:30 UTC = 2026-04-29 07:30 CST → next day.
	utc2 := time.Date(2026, 4, 28, 23, 30, 0, 0, time.UTC)
	if got := dateInt(utc2, tz); got != 20260429 {
		t.Errorf("dateInt crossing midnight in tz = %d, want 20260429", got)
	}
}

func TestYearRange(t *testing.T) {
	tz := time.FixedZone("Asia/Shanghai", 8*3600)
	start, end := yearRange(2026, tz)
	wantStart := time.Date(2026, 1, 1, 0, 0, 0, 0, tz)
	wantEnd := time.Date(2027, 1, 1, 0, 0, 0, 0, tz)
	if !start.Equal(wantStart) {
		t.Errorf("yearRange start = %v, want %v", start, wantStart)
	}
	if !end.Equal(wantEnd) {
		t.Errorf("yearRange end = %v, want %v", end, wantEnd)
	}
}

func TestDayRange(t *testing.T) {
	tz := time.FixedZone("Asia/Shanghai", 8*3600)
	// 2026-04-28 23:30 UTC = 2026-04-29 07:30 CST → range covers 2026-04-29 in CST.
	utc := time.Date(2026, 4, 28, 23, 30, 0, 0, time.UTC)
	start, end := dayRange(utc, tz)
	wantStart := time.Date(2026, 4, 29, 0, 0, 0, 0, tz)
	wantEnd := time.Date(2026, 4, 30, 0, 0, 0, 0, tz)
	if !start.Equal(wantStart) {
		t.Errorf("dayRange start = %v, want %v", start, wantStart)
	}
	if !end.Equal(wantEnd) {
		t.Errorf("dayRange end = %v, want %v", end, wantEnd)
	}
}

func TestBuildRiskLog(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	oldRisk := 12.5
	oldGrade := "黄"
	company := &model.Company{
		ID:                10,
		Title:             "测试集团",
		InherentRisk:      &oldRisk,
		InherentRiskGrade: &oldGrade,
	}
	r := buildRiskLog(company, 1, company.InherentRisk, 8.42, company.InherentRiskGrade, "橙", "每日定时计算", 20260428, now)

	if r.CompanyID != 10 || r.Type != 1 || r.RiskNum != 8.42 || r.Date != 20260428 {
		t.Errorf("scalar fields wrong: %+v", r)
	}
	if r.OldRiskNum == nil || *r.OldRiskNum != 12.5 {
		t.Errorf("old_risk_num not preserved: %+v", r.OldRiskNum)
	}
	if r.OldRiskGrade == nil || *r.OldRiskGrade != "黄" {
		t.Errorf("old_risk_grade not preserved: %+v", r.OldRiskGrade)
	}
	if r.RiskGrade == nil || *r.RiskGrade != "橙" {
		t.Errorf("risk_grade wrong: %+v", r.RiskGrade)
	}
	if r.Content == nil || *r.Content != "每日定时计算" {
		t.Errorf("content wrong: %+v", r.Content)
	}
	if r.CompanyName == nil || *r.CompanyName != "测试集团" {
		t.Errorf("company_name wrong: %+v", r.CompanyName)
	}
	if !r.CreatedAt.Equal(now) || !r.UpdatedAt.Equal(now) {
		t.Errorf("timestamps not set: %+v / %+v", r.CreatedAt, r.UpdatedAt)
	}
	// First-time write: previous values were null.
	first := buildRiskLog(&model.Company{ID: 11, Title: "新企业"}, 2, nil, 1.0, nil, "红", "x", 20260428, now)
	if first.OldRiskNum != nil || first.OldRiskGrade != nil {
		t.Errorf("nil-old values should be preserved as nil, got %+v / %+v", first.OldRiskNum, first.OldRiskGrade)
	}
}

// TestComputePhase1 builds a synthetic bulk cache and three companies, then
// checks both the per-company metric rows and the cross-company globalMax.
//
// Companies:
//  1. C1 — has personnel, certificate, security investments, accidents.
//  2. C2 — has personnel only, all aggregations zero (skipped).
//  3. C3 — no personnel at all (skipped almost entirely except cert lookup,
//     which itself needs personnel.security_officer > 0).
func TestComputePhase1(t *testing.T) {
	tz := time.UTC
	thisYear := 2026
	daysThisYear := 100
	createdLastYear := time.Date(2025, 6, 1, 0, 0, 0, 0, tz)
	createdSameYear := time.Date(2026, 1, 1, 0, 0, 0, 0, tz)

	companies := []model.Company{
		{ID: 1, Title: "C1", Acreage: ptrFloat(1000), Audit: model.Audit{CreatedAt: createdLastYear}},
		{ID: 2, Title: "C2", Acreage: ptrFloat(500), Audit: model.Audit{CreatedAt: createdSameYear}},
		{ID: 3, Title: "C3", Acreage: ptrFloat(200), Audit: model.Audit{CreatedAt: createdLastYear}},
	}

	bulk := &bulkData{
		personnel: map[int64]*model.Personnel{
			1: {TotalPeople: 100, SecurityOfficer: 4, HighSchoolDiploma: 60},
			2: {TotalPeople: 50, SecurityOfficer: 0},
			// C3: no personnel
		},
		cert: map[int64]*model.CertificateStatus{
			1: {AlreadyObtainPeople: ptrInt(2)},
		},
		secInv: map[int64]secInvAgg{
			1: {CompanyID: 1, BudgetSum: 200, IncomeSum: 100, Cnt: 1},
		},
		majHazProcTotal:  map[int64]int64{1: 3},
		hazProcTotal:     map[int64]int64{1: 2},
		hazChemStats:     map[int64]hazChemStats{1: {CompanyID: 1, Total: 5}},
		internalMeeting:  map[int64]int64{1: 10},
		internalTraining: map[int64]int64{1: 5},
		accidentStats:    map[int64]accidentStats{1: {CompanyID: 1, Total: 2}},
		emergencyTeam:    map[int64]float64{1: 80},
		emergencyDrill:   map[int64]int64{1: 4},
		jobManagement:    map[int64]int64{1: 6},
		occHaz:           map[int64]float64{1: 30},
	}

	state := &runState{bulk: bulk}
	state.computePhase1(companies, daysThisYear, thisYear, tz)

	c1 := state.riskMaxByCid[1]
	c2 := state.riskMaxByCid[2]
	c3 := state.riskMaxByCid[3]

	approx := func(name string, got, want float64) {
		t.Helper()
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("%s = %v, want %v", name, got, want)
		}
	}

	// C1 happy path.
	approx("c1 security_officer_scale", c1["security_officer_scale"], 4.0/100)
	approx("c1 already_obtain_people_scale", c1["already_obtain_people_scale"], 2.0/4)
	approx("c1 last_year_income_scale", c1["last_year_income_scale"], 200.0/100*1)
	approx("c1 acreage_scale", c1["acreage_scale"], 100.0/1000)
	approx("c1 hazardous_processes_scale", c1["hazardous_processes_scale"], 5.0/100)
	approx("c1 hazardous_chemicals_scale", c1["hazardous_chemicals_scale"], 5.0/100)
	approx("c1 internal_meeting_scale", c1["internal_meeting_scale"], 10.0/100*30)
	approx("c1 internal_training_scale", c1["internal_training_scale"], 5.0/100*30)
	approx("c1 accident_bulletin_scale", c1["accident_bulletin_scale"], 2.0/100*365)
	approx("c1 emergency_team_scale", c1["emergency_team_scale"], 80.0/100)
	// C1 created in 2025; thisYear=2026 → years=1 → drillCount/years = 4.
	approx("c1 emergency_drill_scale", c1["emergency_drill_scale"], 4.0/1)
	approx("c1 job_management_scale", c1["job_management_scale"], 6.0/100*30)
	approx("c1 health_checkup_scale", c1["health_checkup_scale"], 30.0/100)

	// C2: personnel exists but security_officer = 0 → all guarded metrics skipped.
	if v, ok := c2["security_officer_scale"]; ok && v != 0 {
		t.Errorf("c2 security_officer_scale should be skipped, got %v", v)
	}
	// acreage_scale needs total_people > 0 — C2 has 50, acreage 500, so present.
	approx("c2 acreage_scale", c2["acreage_scale"], 50.0/500)

	// C3: no personnel → most rows missing.
	if len(c3) != 0 {
		t.Errorf("c3 row should be empty (no personnel), got %v", c3)
	}

	// Global max should reflect the loudest C1 values across the 13 max keys.
	approx("globalMax security_officer_scale", state.globalMax["security_officer_scale"], 4.0/100)
	approx("globalMax acreage_scale", state.globalMax["acreage_scale"], 100.0/1000)
	approx("globalMax accident_bulletin_scale", state.globalMax["accident_bulletin_scale"], 2.0/100*365)
}

// TestComputePhase1NoCompanies guards the "empty input doesn't panic" edge.
func TestComputePhase1NoCompanies(t *testing.T) {
	state := &runState{bulk: &bulkData{personnel: map[int64]*model.Personnel{}}}
	state.computePhase1(nil, 1, 2026, time.UTC)
	if len(state.riskMaxByCid) != 0 {
		t.Errorf("expected empty riskMaxByCid, got %v", state.riskMaxByCid)
	}
	for _, v := range state.globalMax {
		if v != 0 {
			t.Errorf("expected globalMax zeros, got %v", state.globalMax)
		}
	}
}

func TestIntsCSV(t *testing.T) {
	cases := []struct {
		in   []int64
		want string
	}{
		{nil, ""},
		{[]int64{}, ""},
		{[]int64{42}, "42"},
		{[]int64{1, 2, 3}, "1,2,3"},
		{[]int64{-7, 0, 7}, "-7,0,7"},
	}
	for _, c := range cases {
		if got := intsCSV(c.in); got != c.want {
			t.Errorf("intsCSV(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}

// Sanity: weight tables stay self-consistent. If someone tweaks the maps
// without updating consumers, this catches a missing key fast.
func TestWeightTablesContainExpectedKeys(t *testing.T) {
	wantInherent := []string{
		"safety_architecture_value", "security_officer_scale", "already_obtain_people_scale",
		"last_year_income_scale", "funding_execution_scale", "production_standardization_max",
		"health_system_scale", "assess_detect_scale", "high_school_diploma_scale",
		"acreage_scale", "staff_scale", "safety_engineering_major_scale",
		"hazardous_processes_scale", "hazardous_process_scale", "hazardous_process_more_scale",
		"process_is_automatic_scale", "special_equipment_number_scale", "equipment_more_number_scale",
		"hazardous_chemical_scale", "hazardous_chemicals_number_scale", "major_hazard_max",
	}
	for _, k := range wantInherent {
		if _, ok := inherentWeights[k]; !ok {
			t.Errorf("inherentWeights missing key %q", k)
		}
	}
	if len(inherentWeights) != len(wantInherent) {
		t.Errorf("inherentWeights size mismatch: got %d, want %d", len(inherentWeights), len(wantInherent))
	}
	wantFloat := []string{
		"internal_meeting_scale", "internal_training_scale", "accident_bulletin_scale",
		"accident_bulletin_severe_max", "accident_report_url_scale", "emergency_team_scale",
		"emergency_drill_scale", "job_management_scale", "monthly_investigation_scale",
		"three_together_over_scale", "unit_sum_number_scale", "evaluation_unit_scale",
		"evaluation_unit_school_scale", "health_checkup_people_scale", "health_checkup_scale",
		"health_report_scale",
	}
	for _, k := range wantFloat {
		if _, ok := floatWeights[k]; !ok {
			t.Errorf("floatWeights missing key %q", k)
		}
	}
	if len(floatWeights) != len(wantFloat) {
		t.Errorf("floatWeights size mismatch: got %d, want %d", len(floatWeights), len(wantFloat))
	}
}

func ptrFloat(v float64) *float64 { return &v }
func ptrInt(v int) *int           { return &v }
func ptrStr(v string) *string     { return &v }
func ptrInt64(v int64) *int64     { return &v }

// TestComputeInherentMinimal: zero inputs everywhere except where the trait
// awards a default. Locks down the unconditional contributions:
//   - all normalized*/normalizedInverse helpers → 0 (val/max guard)
//   - major_hazard_scale → default 1.0 * weight (no level matched)
//   - everything else → 0 (conditional branches don't fire)
//
// Expected score = inherentWeights["major_hazard_max"] * 100 * 1.0
//                = 0.0198135 * 100 = 1.98135 → round2 = 1.98
//
// Grade = "红" (score < 5).
func TestComputeInherentMinimal(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	companies := []model.Company{
		{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}},
	}
	state := &runState{
		bulk:         &bulkData{}, // every map nil-safe (Go zero-value reads)
		riskMaxByCid: map[int64]map[string]float64{1: {}},
		globalMax:    map[string]float64{},
	}

	logs, riskLogs, updates := computeInherent(state, companies, now, "test", tz)

	if len(logs) != 1 || len(riskLogs) != 1 {
		t.Fatalf("want 1 log + 1 risk_log, got %d / %d", len(logs), len(riskLogs))
	}
	row := logs[0]
	expectMajor := inherentWeights["major_hazard_max"]
	if row.MajorHazardScale != expectMajor {
		t.Errorf("MajorHazardScale = %v, want %v (default branch)", row.MajorHazardScale, expectMajor)
	}
	// Every other field should be exactly 0.
	zeros := map[string]float64{
		"SafetyArchitectureValue":         row.SafetyArchitectureValue,
		"SecurityOfficerScale":            row.SecurityOfficerScale,
		"AlreadyObtainPeopleScale":        row.AlreadyObtainPeopleScale,
		"LastYearIncomeScale":             row.LastYearIncomeScale,
		"FundingExecutionScale":           row.FundingExecutionScale,
		"ProductionStandardizationScale":  row.ProductionStandardizationScale,
		"HealthSystemScale":               row.HealthSystemScale,
		"AssessDetectScale":               row.AssessDetectScale,
		"HighSchoolDiplomaScale":          row.HighSchoolDiplomaScale,
		"AcreageScale":                    row.AcreageScale,
		"StaffScale":                      row.StaffScale,
		"SafetyEngineeringMajorScale":     row.SafetyEngineeringMajorScale,
		"HazardousProcessesScale":         row.HazardousProcessesScale,
		"HazardousProcessScale":           row.HazardousProcessScale,
		"HazardousProcessMoreScale":       row.HazardousProcessMoreScale,
		"ProcessIsAutomaticScale":         row.ProcessIsAutomaticScale,
		"SpecialEquipmentNumberScale":     row.SpecialEquipmentNumberScale,
		"EquipmentMoreNumberScale":        row.EquipmentMoreNumberScale,
		"HazardousChemicalScale":          row.HazardousChemicalScale,
		"HazardousChemicalsNumberScale":   row.HazardousChemicalsNumberScale,
	}
	for name, v := range zeros {
		if v != 0 {
			t.Errorf("expected %s = 0 in minimal fixture, got %v", name, v)
		}
	}

	want := round2(expectMajor * 100)
	if updates[1].risk != want {
		t.Errorf("update.risk = %v, want %v", updates[1].risk, want)
	}
	if updates[1].grade != "红" {
		t.Errorf("update.grade = %q, want 红", updates[1].grade)
	}
	if riskLogs[0].Type != 1 || riskLogs[0].RiskNum != want {
		t.Errorf("riskLog wrong: type=%d, risk_num=%v", riskLogs[0].Type, riskLogs[0].RiskNum)
	}
}

// TestComputeFloatMinimal: zero inputs. Two unconditional defaults fire:
//   - accident_bulletin_severe_scale → default = 1.0 * weight
//   - health_report_scale → full weight (because count == 0 → "no report filed")
//
// Expected score = (severe_max + health_report_scale) * 100, round2.
//                = (0.058859415 + 0.056693) * 100 = 11.5552415 → 11.56
//
// Grade = "黄" (10 ≤ 11.56 < 15) — non-obvious result worth locking down.
func TestComputeFloatMinimal(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	companies := []model.Company{
		{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}},
	}
	state := &runState{
		bulk:         &bulkData{},
		riskMaxByCid: map[int64]map[string]float64{1: {}},
		globalMax:    map[string]float64{},
	}

	logs, riskLogs, updates := computeFloat(state, companies, now, "test", tz)
	if len(logs) != 1 || len(riskLogs) != 1 {
		t.Fatalf("want 1+1 rows, got %d / %d", len(logs), len(riskLogs))
	}
	row := logs[0]

	wantSevere := floatWeights["accident_bulletin_severe_max"]
	wantHealth := floatWeights["health_report_scale"]
	if row.AccidentBulletinSevereScale != wantSevere {
		t.Errorf("AccidentBulletinSevereScale = %v, want %v (default)", row.AccidentBulletinSevereScale, wantSevere)
	}
	if row.HealthReportScale != wantHealth {
		t.Errorf("HealthReportScale = %v, want %v (no-report-filed branch)", row.HealthReportScale, wantHealth)
	}
	if row.RiskSourceScale != 0 {
		t.Errorf("RiskSourceScale must be hard-zero, got %v", row.RiskSourceScale)
	}

	want := round2((wantSevere + wantHealth) * 100)
	if updates[1].risk != want {
		t.Errorf("update.risk = %v, want %v", updates[1].risk, want)
	}
	if updates[1].grade != "黄" {
		t.Errorf("update.grade = %q, want 黄 (current default fixture lands here)", updates[1].grade)
	}
	if riskLogs[0].Type != 2 {
		t.Errorf("riskLog type = %d, want 2", riskLogs[0].Type)
	}
}

// TestAccidentSeverityLadder locks down the priority order in the trait's
// `match (true)` ladder for accident_bulletin_severe_scale: otherCount > 0
// trumps everything (returns 0); then serious (0.3x); then light (0.6x); else
// default 1.0x. Each test case sets exactly the bucket needed.
func TestAccidentSeverityLadder(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	weight := floatWeights["accident_bulletin_severe_max"]

	cases := []struct {
		name  string
		stats accidentStats
		want  float64
	}{
		{"empty (default)", accidentStats{}, 1.0 * weight},
		{"light only → 0.6x", accidentStats{LightCnt: 1}, 0.6 * weight},
		{"serious only → 0.3x", accidentStats{SeriousCnt: 1}, 0.3 * weight},
		{"other only → 0", accidentStats{OtherCnt: 1}, 0},
		{"other beats serious", accidentStats{OtherCnt: 1, SeriousCnt: 1}, 0},
		{"serious beats light", accidentStats{SeriousCnt: 1, LightCnt: 1}, 0.3 * weight},
		{"all three present → other wins", accidentStats{OtherCnt: 1, SeriousCnt: 1, LightCnt: 1}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			companies := []model.Company{{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}}}
			state := &runState{
				bulk:         &bulkData{accidentStats: map[int64]accidentStats{1: c.stats}},
				riskMaxByCid: map[int64]map[string]float64{1: {}},
				globalMax:    map[string]float64{},
			}
			logs, _, _ := computeFloat(state, companies, now, "t", tz)
			if math.Abs(logs[0].AccidentBulletinSevereScale-c.want) > 1e-9 {
				t.Errorf("%s: AccidentBulletinSevereScale = %v, want %v", c.name, logs[0].AccidentBulletinSevereScale, c.want)
			}
		})
	}
}

// TestMajorHazardLadder locks down the priority order for major_hazard_scale
// across the 4 levels + default. Trait does Level1 → 0.6x, Level2 → 0.7x,
// Level3 → 0.8x, Level4 → 0.9x, default → 1.0x; first match wins, so Level1
// in a list also containing Level3 must produce the Level1 (lowest) result.
func TestMajorHazardLadder(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	weight := inherentWeights["major_hazard_max"]

	cases := []struct {
		name   string
		levels []string
		want   float64
	}{
		{"empty → 1.0x", nil, 1.0 * weight},
		{"unknown level → 1.0x", []string{"五级"}, 1.0 * weight},
		{"Level1 → 0.6x", []string{"一级"}, 0.6 * weight},
		{"Level2 → 0.7x", []string{"二级"}, 0.7 * weight},
		{"Level3 → 0.8x", []string{"三级"}, 0.8 * weight},
		{"Level4 → 0.9x", []string{"四级"}, 0.9 * weight},
		{"Level1 wins over Level3", []string{"三级", "一级"}, 0.6 * weight},
		{"Level2 wins over Level4", []string{"四级", "二级"}, 0.7 * weight},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			companies := []model.Company{{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}}}
			state := &runState{
				bulk:         &bulkData{majorHazard: map[int64][]string{1: c.levels}},
				riskMaxByCid: map[int64]map[string]float64{1: {}},
				globalMax:    map[string]float64{},
			}
			logs, _, _ := computeInherent(state, companies, now, "t", tz)
			if math.Abs(logs[0].MajorHazardScale-c.want) > 1e-9 {
				t.Errorf("%s: MajorHazardScale = %v, want %v", c.name, logs[0].MajorHazardScale, c.want)
			}
		})
	}
}

// TestProductionStandardizationLadder: Level1 → 0.9x, Level2 → 0.75x, Level3
// → 0.6x, anything else → 0 (no contribution).
func TestProductionStandardizationLadder(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	weight := inherentWeights["production_standardization_max"]

	cases := []struct {
		grade *string
		want  float64
	}{
		{nil, 0},
		{ptrStr("一级"), 0.9 * weight},
		{ptrStr("二级"), 0.75 * weight},
		{ptrStr("三级"), 0.6 * weight},
		{ptrStr("四级"), 0}, // unrecognized
	}
	for _, c := range cases {
		var name string
		if c.grade == nil {
			name = "nil grade"
		} else {
			name = *c.grade
		}
		t.Run(name, func(t *testing.T) {
			companies := []model.Company{{ID: 1, Title: "C", Audit: model.Audit{CreatedAt: now}}}
			state := &runState{
				bulk: &bulkData{
					productionStd: map[int64]*model.ProductionStandardization{
						1: {Grade: c.grade},
					},
				},
				riskMaxByCid: map[int64]map[string]float64{1: {}},
				globalMax:    map[string]float64{},
			}
			logs, _, _ := computeInherent(state, companies, now, "t", tz)
			if math.Abs(logs[0].ProductionStandardizationScale-c.want) > 1e-9 {
				t.Errorf("ProductionStandardizationScale = %v, want %v", logs[0].ProductionStandardizationScale, c.want)
			}
		})
	}
}

// TestComputeInherentIndustryWeight: same fixture, different enterprise_type_id.
// Confirms industry multiplier is applied to inherent score (and only inherent).
// score = sum * w * 100; with sum from minimal default = 0.0198135, w should
// scale linearly.
func TestComputeInherentIndustryWeight(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	cases := []struct {
		typeID *int64
		w      float64
	}{
		{nil, 1.0},
		{ptrInt64(1), 0.8},
		{ptrInt64(2), 0.6},
		{ptrInt64(3), 0.9},
		{ptrInt64(99), 1.0},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			companies := []model.Company{
				{ID: 1, Title: "C", EnterpriseTypeID: c.typeID, Audit: model.Audit{CreatedAt: now}},
			}
			state := &runState{
				bulk:         &bulkData{},
				riskMaxByCid: map[int64]map[string]float64{1: {}},
				globalMax:    map[string]float64{},
			}
			_, _, updates := computeInherent(state, companies, now, "t", tz)
			want := round2(inherentWeights["major_hazard_max"] * 100 * c.w)
			if math.Abs(updates[1].risk-want) > 1e-9 {
				t.Errorf("typeID=%v: risk = %v, want %v (sum * %v * 100, round2)", c.typeID, updates[1].risk, want, c.w)
			}
		})
	}
}

// TestComputeFloatNoIndustryMultiplier: a sanity guard against a regression
// where someone applies industryWeight() to the float compute by accident.
// Phase 3 uses sum * 100 only — no industry multiplier.
func TestComputeFloatNoIndustryMultiplier(t *testing.T) {
	now := time.Date(2026, 4, 28, 2, 0, 0, 0, time.UTC)
	tz := time.UTC
	id1, id3 := int64(1), int64(3)
	state := &runState{
		bulk:         &bulkData{},
		riskMaxByCid: map[int64]map[string]float64{},
		globalMax:    map[string]float64{},
	}
	mkCompany := func(id int64, typeID *int64) model.Company {
		return model.Company{ID: id, Title: "C", EnterpriseTypeID: typeID, Audit: model.Audit{CreatedAt: now}}
	}
	companies := []model.Company{
		mkCompany(1, nil),
		mkCompany(2, &id1),
		mkCompany(3, &id3),
	}
	for _, c := range companies {
		state.riskMaxByCid[c.ID] = map[string]float64{}
	}
	_, _, updates := computeFloat(state, companies, now, "t", tz)

	// All three should land on identical scores — float doesn't apply
	// industryWeight. If someone breaks this, the three values will differ.
	r1, r2, r3 := updates[1].risk, updates[2].risk, updates[3].risk
	if r1 != r2 || r1 != r3 {
		t.Errorf("float scores should be identical regardless of industry: %v / %v / %v", r1, r2, r3)
	}
}
