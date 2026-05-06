// Package service — 集团安全风险看板聚合服务。
//
// Single endpoint backing /api/v1/safety-dashboard/overview. Aggregates data
// from ~10 different domain tables (personnels, hidden_danger_infos, drills,
// etc.) for one company subtree (root_id self+descendants) and returns the
// shape the React dashboard expects (web/src/pages/safety-dashboard/types.ts).
//
// Aggregation scope: every query is bounded by company_id IN (subtree),
// resolved up-front via CompanyRepo.DescendantIDs. Empty subtree short-circuits
// to a zero-valued response so the SPA still paints (0/empty cards).
package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/model"
	"ai_hs_go/internal/repository"
)

type SafetyDashboardService struct {
	db          *gorm.DB
	companyRepo *repository.CompanyRepo
}

func NewSafetyDashboardService(db *gorm.DB, companyRepo *repository.CompanyRepo) *SafetyDashboardService {
	return &SafetyDashboardService{db: db, companyRepo: companyRepo}
}

// ── DTOs (JSON keys MUST match web/src/pages/safety-dashboard/types.ts) ────

type DashboardOverview struct {
	Meta               DashboardMeta              `json:"meta"`
	Personnel          DashboardPersonnel         `json:"personnel"`
	Training           []DashboardMonthlyTraining `json:"training"`
	Investment         []DashboardMonthlyValue    `json:"investment"`
	HazardYear         []DashboardHazardRow       `json:"hazardYear"`
	HazardMonth        []DashboardHazardRow       `json:"hazardMonth"`
	RiskPoints         []DashboardRiskPoint       `json:"riskPoints"`
	RiskLevels         []DashboardRiskLevelInfo   `json:"riskLevels"`
	IndustryRisk       []DashboardIndustrySlice   `json:"industryRisk"`
	StandardLevel      DashboardStandardLevel     `json:"standardLevel"`
	Drills             []DashboardDrillKpi        `json:"drills"`
	OccupationalHealth []DashboardOhRow           `json:"occupationalHealth"`
	Alerts             []DashboardAlertEntry      `json:"alerts"`
}

type DashboardMeta struct {
	Title     string `json:"title"`
	SafeDays  int    `json:"safeDays"`
	UpdatedAt string `json:"updatedAt"`
}

type DashboardPersonnel struct {
	Total         int                       `json:"total"`
	Composition   []DashboardCompositionRow `json:"composition"`
	Professionals []DashboardProfRow        `json:"professionals"`
}

type DashboardCompositionRow struct {
	Label string `json:"label"`
	Value int    `json:"value"`
	Color string `json:"color"` // primary | accent | warn — matches the SPA's color role
}

type DashboardProfRow struct {
	Label   string  `json:"label"`
	Value   int     `json:"value"`
	Percent float64 `json:"percent"`
}

type DashboardMonthlyTraining struct {
	Month    string `json:"month"`
	Sessions int    `json:"sessions"`
	People   int    `json:"people"`
}

type DashboardMonthlyValue struct {
	Month string  `json:"month"`
	Value float64 `json:"value"`
}

type DashboardHazardRow struct {
	Industry string  `json:"industry"`
	Found    int     `json:"found"`
	Fixed    int     `json:"fixed"`
	Rate     float64 `json:"rate"`
}

// RiskPoint contract uses lng/lat now (was x/y normalized in mock). The SPA's
// ECharts geo component projects lng/lat directly so the backend doesn't have
// to know the map's coordinate system.
type DashboardRiskPoint struct {
	City  string  `json:"city"`
	Lng   float64 `json:"lng"`
	Lat   float64 `json:"lat"`
	Level string  `json:"level"` // severe | major | normal | low
	Name  string  `json:"name"`
}

type DashboardRiskLevelInfo struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Color string `json:"color"`
}

type DashboardIndustrySlice struct {
	Name    string `json:"name"`
	Value   int    `json:"value"`
	Percent int    `json:"percent"`
	Color   string `json:"color"`
}

type DashboardStandardLevel struct {
	Total     int                          `json:"total"`
	Current   string                       `json:"current"`
	Breakdown []DashboardStandardBreakdown `json:"breakdown"`
}

type DashboardStandardBreakdown struct {
	Label string `json:"label"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

type DashboardDrillKpi struct {
	Label  string `json:"label"`
	Value  int    `json:"value"`
	Suffix string `json:"suffix"`
	Color  string `json:"color"` // primary | warn
}

type DashboardOhRow struct {
	Hazard string `json:"hazard"`
	People int    `json:"people"`
}

type DashboardAlertEntry struct {
	Time  string `json:"time"`
	Level string `json:"level"` // major | normal | low
	Text  string `json:"text"`
}

// ── Static lookups (mirror the SPA's mock palette) ────────────────────────

var staticRiskLevels = []DashboardRiskLevelInfo{
	{Key: "severe", Label: "重大风险", Color: "#ef4444"},
	{Key: "major", Label: "较大风险", Color: "#f97316"},
	{Key: "normal", Label: "一般风险", Color: "#eab308"},
	{Key: "low", Label: "低风险", Color: "#3b82f6"},
}

var industryColorPalette = []string{"#3b82f6", "#a855f7", "#ef4444", "#64748b", "#10b981", "#f59e0b"}

var standardGradeColors = map[string]string{
	"一级": "#eab308",
	"二级": "#3b82f6",
	"三级": "#a855f7",
	"未取得": "#64748b",
}

// risk_monitor_service.go grades scores into 红/橙/黄/蓝; map those to the
// RiskLevel keys the SPA understands. Anything else (empty / unrecognized)
// falls to "low" so the dot still draws on the map.
var gradeToLevel = map[string]string{
	"红": "severe",
	"橙": "major",
	"黄": "normal",
	"蓝": "low",
}

// ── Top-level entry point ─────────────────────────────────────────────────

func (s *SafetyDashboardService) Overview(ctx context.Context, companyID int64) (*DashboardOverview, error) {
	subtree, err := s.companyRepo.DescendantIDs(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("resolve subtree: %w", err)
	}

	out := &DashboardOverview{
		Meta: DashboardMeta{
			Title:     "杭实集团安全风险看板",
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		RiskLevels: staticRiskLevels,
	}

	if len(subtree) == 0 {
		// No companies in scope — return defaults so the SPA paints empties.
		out.Personnel.Composition = []DashboardCompositionRow{}
		out.Personnel.Professionals = []DashboardProfRow{}
		out.Training = []DashboardMonthlyTraining{}
		out.Investment = []DashboardMonthlyValue{}
		out.HazardYear = []DashboardHazardRow{}
		out.HazardMonth = []DashboardHazardRow{}
		out.RiskPoints = []DashboardRiskPoint{}
		out.IndustryRisk = []DashboardIndustrySlice{}
		out.StandardLevel.Breakdown = []DashboardStandardBreakdown{}
		out.Drills = []DashboardDrillKpi{}
		out.OccupationalHealth = []DashboardOhRow{}
		out.Alerts = []DashboardAlertEntry{}
		return out, nil
	}

	// Run aggregations sequentially. The dataset is bounded (max ~hundreds of
	// companies, low-thousands of rows per table) so single-connection
	// sequential queries are fine; parallelizing here would mostly buy
	// connection-pool churn.

	if v, err := s.computeSafeDays(ctx, subtree); err != nil {
		return nil, fmt.Errorf("safe days: %w", err)
	} else {
		out.Meta.SafeDays = v
	}

	if v, err := s.aggregatePersonnel(ctx, subtree); err != nil {
		return nil, fmt.Errorf("personnel: %w", err)
	} else {
		out.Personnel = v
	}

	if v, err := s.aggregateTraining12M(ctx, subtree); err != nil {
		return nil, fmt.Errorf("training: %w", err)
	} else {
		out.Training = v
	}

	if v, err := s.aggregateInvestment12M(ctx, subtree); err != nil {
		return nil, fmt.Errorf("investment: %w", err)
	} else {
		out.Investment = v
	}

	year, month, err := s.aggregateHazards(ctx, subtree)
	if err != nil {
		return nil, fmt.Errorf("hazards: %w", err)
	}
	out.HazardYear = year
	out.HazardMonth = month

	if v, err := s.queryRiskPoints(ctx, subtree); err != nil {
		return nil, fmt.Errorf("risk points: %w", err)
	} else {
		out.RiskPoints = v
	}

	if v, err := s.queryIndustryRisk(ctx, subtree); err != nil {
		return nil, fmt.Errorf("industry risk: %w", err)
	} else {
		out.IndustryRisk = v
	}

	if v, err := s.queryStandardLevel(ctx, subtree); err != nil {
		return nil, fmt.Errorf("standard level: %w", err)
	} else {
		out.StandardLevel = v
	}

	if v, err := s.queryDrills(ctx, subtree); err != nil {
		return nil, fmt.Errorf("drills: %w", err)
	} else {
		out.Drills = v
	}

	if v, err := s.queryOccupationalHealth(ctx, subtree); err != nil {
		return nil, fmt.Errorf("occupational health: %w", err)
	} else {
		out.OccupationalHealth = v
	}

	if v, err := s.queryRecentAlerts(ctx, subtree); err != nil {
		return nil, fmt.Errorf("alerts: %w", err)
	} else {
		out.Alerts = v
	}

	return out, nil
}

// ── Aggregations ──────────────────────────────────────────────────────────

// computeSafeDays returns days since the last accident_bulletin within the
// subtree. No bulletins → fall back to 0 (we don't claim a safety streak we
// can't substantiate).
func (s *SafetyDashboardService) computeSafeDays(ctx context.Context, subtree []int64) (int, error) {
	var last *time.Time
	err := s.db.WithContext(ctx).
		Model(&model.AccidentBulletin{}).
		Where("company_id IN ?", subtree).
		Select("MAX(created_at)").
		Scan(&last).Error
	if err != nil {
		return 0, err
	}
	if last == nil || last.IsZero() {
		return 0, nil
	}
	days := int(time.Since(*last).Hours() / 24)
	if days < 0 {
		return 0, nil
	}
	return days, nil
}

func (s *SafetyDashboardService) aggregatePersonnel(ctx context.Context, subtree []int64) (DashboardPersonnel, error) {
	var sums struct {
		TotalPeople            int64
		Staff                  int64
		Manager                int64
		Service                int64
		SafetyEngineer         int64
		SafetyEngineeringMajor int64
		FireEngineer           int64
		Other                  int64
	}
	err := s.db.WithContext(ctx).
		Model(&model.Personnel{}).
		Where("company_id IN ?", subtree).
		Select(`
			COALESCE(SUM(total_people),0)              AS total_people,
			COALESCE(SUM(staff),0)                     AS staff,
			COALESCE(SUM(manager),0)                   AS manager,
			COALESCE(SUM(service),0)                   AS service,
			COALESCE(SUM(safety_engineer),0)           AS safety_engineer,
			COALESCE(SUM(safety_engineering_major),0)  AS safety_engineering_major,
			COALESCE(SUM(fire_engineer),0)             AS fire_engineer,
			COALESCE(SUM(other),0)                     AS other`).
		Scan(&sums).Error
	if err != nil {
		return DashboardPersonnel{}, err
	}

	composition := []DashboardCompositionRow{
		{Label: "职工", Value: int(sums.Staff), Color: "primary"},
		{Label: "管理人员", Value: int(sums.Manager), Color: "accent"},
		{Label: "劳务派遣用工", Value: int(sums.Service), Color: "warn"},
	}

	profs := []struct {
		Label string
		Value int
	}{
		{"注册安全工程师", int(sums.SafetyEngineer)},
		{"安全工程专业", int(sums.SafetyEngineeringMajor)},
		{"注册消防工程师", int(sums.FireEngineer)},
		{"其他", int(sums.Other)},
	}
	totalProf := 0
	for _, p := range profs {
		totalProf += p.Value
	}
	professionals := make([]DashboardProfRow, len(profs))
	for i, p := range profs {
		percent := 0.0
		if totalProf > 0 {
			percent = float64(p.Value) / float64(totalProf) * 100
		}
		professionals[i] = DashboardProfRow{
			Label:   p.Label,
			Value:   p.Value,
			Percent: roundTo(percent, 1),
		}
	}

	return DashboardPersonnel{
		Total:         int(sums.TotalPeople),
		Composition:   composition,
		Professionals: professionals,
	}, nil
}

// aggregateTraining12M returns the trailing 12 months of training activity by
// unioning internal_trainings + external_trainings + other_trainings. Each
// session contributes 1 to sessions and Number (people count) to people.
func (s *SafetyDashboardService) aggregateTraining12M(ctx context.Context, subtree []int64) ([]DashboardMonthlyTraining, error) {
	now := time.Now()
	start := now.AddDate(0, -11, 0)
	startMonth := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.Local)

	type row struct {
		Bucket   string
		Sessions int64
		People   int64
	}
	const q = `
WITH unioned AS (
	SELECT training_date, COALESCE(number,0) AS number FROM internal_trainings
		WHERE company_id IN ? AND training_date >= ? AND deleted_at IS NULL
	UNION ALL
	SELECT training_date, COALESCE(number,0) FROM external_trainings
		WHERE company_id IN ? AND training_date >= ? AND deleted_at IS NULL
	UNION ALL
	SELECT training_date, COALESCE(number,0) FROM other_trainings
		WHERE company_id IN ? AND training_date >= ? AND deleted_at IS NULL
)
SELECT to_char(training_date, 'YY/MM') AS bucket,
       COUNT(*)          AS sessions,
       COALESCE(SUM(number),0) AS people
FROM unioned
WHERE training_date IS NOT NULL
GROUP BY bucket
ORDER BY bucket
`
	var rows []row
	err := s.db.WithContext(ctx).Raw(q,
		subtree, startMonth,
		subtree, startMonth,
		subtree, startMonth).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	idx := make(map[string]row, len(rows))
	for _, r := range rows {
		idx[r.Bucket] = r
	}
	return fillMonthly12(now, idx, func(b string, r row, ok bool) DashboardMonthlyTraining {
		if !ok {
			return DashboardMonthlyTraining{Month: b}
		}
		return DashboardMonthlyTraining{Month: b, Sessions: int(r.Sessions), People: int(r.People)}
	}), nil
}

// aggregateInvestment12M groups funding_executions.budget by month-of-creation.
// `Year` is stored as a YYYY string only — we lean on created_at for the month
// bucket, which matches how the rest of the dashboard surfaces "12-month"
// trend data.
func (s *SafetyDashboardService) aggregateInvestment12M(ctx context.Context, subtree []int64) ([]DashboardMonthlyValue, error) {
	now := time.Now()
	start := now.AddDate(0, -11, 0)
	startMonth := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.Local)

	type row struct {
		Bucket string
		Total  float64
	}
	const q = `
SELECT to_char(created_at, 'YY/MM') AS bucket,
       COALESCE(SUM(budget),0)      AS total
FROM funding_executions
WHERE company_id IN ? AND created_at >= ? AND deleted_at IS NULL
GROUP BY bucket
ORDER BY bucket
`
	var rows []row
	if err := s.db.WithContext(ctx).Raw(q, subtree, startMonth).Scan(&rows).Error; err != nil {
		return nil, err
	}
	idx := make(map[string]row, len(rows))
	for _, r := range rows {
		idx[r.Bucket] = r
	}
	return fillMonthly12(now, idx, func(b string, r row, ok bool) DashboardMonthlyValue {
		if !ok {
			return DashboardMonthlyValue{Month: b, Value: 0}
		}
		return DashboardMonthlyValue{Month: b, Value: roundTo(r.Total, 2)}
	}), nil
}

// aggregateHazards groups hidden_danger_infos by hidden_type, splitting into
// year-to-date vs current-month windows. fixed = correct_status=3 AND
// invest_status=2 (matches the rest of the codebase's done-flag convention).
// Industries beyond the canonical four (化工/机械/商贸租赁/其他) collapse into
// "其他行业" so the chart stays at 4 columns.
func (s *SafetyDashboardService) aggregateHazards(ctx context.Context, subtree []int64) (year, month []DashboardHazardRow, err error) {
	now := time.Now()
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	canonical := []string{"化工行业", "机械行业", "商贸租赁行业", "其他行业"}
	canonicalSet := map[string]struct{}{}
	for _, c := range canonical {
		canonicalSet[c] = struct{}{}
	}

	queryWindow := func(from time.Time) (map[string]*DashboardHazardRow, error) {
		type r struct {
			Industry string
			Found    int64
			Fixed    int64
		}
		var rows []r
		const q = `
SELECT
	CASE WHEN hidden_type IN ('化工行业','机械行业','商贸租赁行业') THEN hidden_type
	     ELSE '其他行业' END AS industry,
	COUNT(*) AS found,
	COUNT(*) FILTER (WHERE correct_status = 3 AND invest_status = 2) AS fixed
FROM hidden_danger_infos
WHERE check_company_id IN ? AND created_at >= ? AND deleted_at IS NULL
GROUP BY industry
`
		if err := s.db.WithContext(ctx).Raw(q, subtree, from).Scan(&rows).Error; err != nil {
			return nil, err
		}
		out := make(map[string]*DashboardHazardRow, len(rows))
		for _, x := range rows {
			rate := 0.0
			if x.Found > 0 {
				rate = roundTo(float64(x.Fixed)/float64(x.Found)*100, 1)
			}
			out[x.Industry] = &DashboardHazardRow{
				Industry: x.Industry,
				Found:    int(x.Found),
				Fixed:    int(x.Fixed),
				Rate:     rate,
			}
		}
		return out, nil
	}

	yearMap, err := queryWindow(yearStart)
	if err != nil {
		return nil, nil, err
	}
	monthMap, err := queryWindow(monthStart)
	if err != nil {
		return nil, nil, err
	}

	year = make([]DashboardHazardRow, 0, 4)
	month = make([]DashboardHazardRow, 0, 4)
	for _, name := range canonical {
		if r, ok := yearMap[name]; ok {
			year = append(year, *r)
		} else {
			year = append(year, DashboardHazardRow{Industry: name})
		}
		if r, ok := monthMap[name]; ok {
			month = append(month, *r)
		} else {
			month = append(month, DashboardHazardRow{Industry: name})
		}
	}
	return year, month, nil
}

func (s *SafetyDashboardService) queryRiskPoints(ctx context.Context, subtree []int64) ([]DashboardRiskPoint, error) {
	type row struct {
		ID                int64
		Title             string
		Longitude         *float64
		Dimension         *float64
		AreaTitle         *string
		InherentRiskGrade *string
	}
	const q = `
SELECT c.id, c.title, c.longitude, c.dimension,
       a.area_name AS area_title,
       c.inherent_risk_grade
FROM companies c
LEFT JOIN base_areas a ON a.area_id = c.area_id
WHERE c.id IN ? AND c.deleted_at IS NULL
  AND c.longitude IS NOT NULL AND c.dimension IS NOT NULL
ORDER BY c.id
`
	var rows []row
	if err := s.db.WithContext(ctx).Raw(q, subtree).Scan(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]DashboardRiskPoint, 0, len(rows))
	for _, r := range rows {
		if r.Longitude == nil || r.Dimension == nil {
			continue
		}
		level := "low"
		if r.InherentRiskGrade != nil {
			if mapped, ok := gradeToLevel[*r.InherentRiskGrade]; ok {
				level = mapped
			}
		}
		city := ""
		if r.AreaTitle != nil {
			city = *r.AreaTitle
		}
		out = append(out, DashboardRiskPoint{
			City:  city,
			Lng:   *r.Longitude,
			Lat:   *r.Dimension,
			Level: level,
			Name:  r.Title,
		})
	}
	return out, nil
}

// queryIndustryRisk groups subtree companies by enterprise type (joined
// through base_enterprise_types) and counts how many fall under each.
// Percent is rounded to int — the SPA renders it as `xx%` and doesn't need
// decimals.
func (s *SafetyDashboardService) queryIndustryRisk(ctx context.Context, subtree []int64) ([]DashboardIndustrySlice, error) {
	type row struct {
		Name  *string
		Count int64
	}
	const q = `
SELECT t.title AS name, COUNT(*) AS count
FROM companies c
LEFT JOIN enterprise_types t ON t.id = c.enterprise_type_id
WHERE c.id IN ? AND c.deleted_at IS NULL
GROUP BY t.title
ORDER BY count DESC
`
	var rows []row
	if err := s.db.WithContext(ctx).Raw(q, subtree).Scan(&rows).Error; err != nil {
		return nil, err
	}
	var total int64
	for _, r := range rows {
		total += r.Count
	}
	out := make([]DashboardIndustrySlice, 0, len(rows))
	for i, r := range rows {
		name := "其他"
		if r.Name != nil && *r.Name != "" {
			name = *r.Name
		}
		percent := 0
		if total > 0 {
			percent = int(r.Count * 100 / total)
		}
		color := industryColorPalette[i%len(industryColorPalette)]
		out = append(out, DashboardIndustrySlice{
			Name:    name,
			Value:   int(r.Count),
			Percent: percent,
			Color:   color,
		})
	}
	return out, nil
}

func (s *SafetyDashboardService) queryStandardLevel(ctx context.Context, subtree []int64) (DashboardStandardLevel, error) {
	type row struct {
		Grade *string
		Count int64
	}
	const q = `
SELECT grade, COUNT(*) AS count
FROM production_standardizations
WHERE company_id IN ? AND deleted_at IS NULL
GROUP BY grade
ORDER BY count DESC
`
	var rows []row
	if err := s.db.WithContext(ctx).Raw(q, subtree).Scan(&rows).Error; err != nil {
		return DashboardStandardLevel{}, err
	}
	var total int64
	current := ""
	breakdown := make([]DashboardStandardBreakdown, 0, len(rows))
	for i, r := range rows {
		grade := "未取得"
		if r.Grade != nil && *r.Grade != "" {
			grade = *r.Grade
		}
		if i == 0 {
			current = grade
		}
		color := standardGradeColors[grade]
		if color == "" {
			color = "#64748b"
		}
		breakdown = append(breakdown, DashboardStandardBreakdown{
			Label: grade,
			Value: int(r.Count),
			Color: color,
		})
		total += r.Count
	}
	return DashboardStandardLevel{
		Total:     int(total),
		Current:   current,
		Breakdown: breakdown,
	}, nil
}

// queryDrills classifies emergency_drills by drill_type ("消防演练" wins on
// substring match — covers "消防演练 / 综合应急消防演练" etc., everything else
// counts as 安全演练). people defaults to 0 when drill_number isn't a parseable
// integer (it's a *string in the schema, so this is best-effort).
func (s *SafetyDashboardService) queryDrills(ctx context.Context, subtree []int64) ([]DashboardDrillKpi, error) {
	type row struct {
		DrillType   *string
		DrillNumber *string
	}
	var rows []row
	err := s.db.WithContext(ctx).
		Model(&model.EmergencyDrill{}).
		Where("company_id IN ?", subtree).
		Select("drill_type", "drill_number").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	var safeSessions, safePeople, fireSessions, firePeople int
	for _, r := range rows {
		isFire := false
		if r.DrillType != nil && strings.Contains(*r.DrillType, "消防") {
			isFire = true
		}
		people := 0
		if r.DrillNumber != nil {
			if n, ok := parseLooseInt(*r.DrillNumber); ok {
				people = n
			}
		}
		if isFire {
			fireSessions++
			firePeople += people
		} else {
			safeSessions++
			safePeople += people
		}
	}

	return []DashboardDrillKpi{
		{Label: "安全演练人数", Value: safePeople, Suffix: "人", Color: "primary"},
		{Label: "安全演练次数", Value: safeSessions, Suffix: "次", Color: "primary"},
		{Label: "消防演练人数", Value: firePeople, Suffix: "人", Color: "warn"},
		{Label: "消防演练次数", Value: fireSessions, Suffix: "次", Color: "warn"},
	}, nil
}

func (s *SafetyDashboardService) queryOccupationalHealth(ctx context.Context, subtree []int64) ([]DashboardOhRow, error) {
	type row struct {
		Type   *string
		People int64
	}
	const q = `
SELECT type, COALESCE(SUM(number),0) AS people
FROM occupational_hazards
WHERE company_id IN ? AND deleted_at IS NULL
GROUP BY type
ORDER BY people DESC
LIMIT 10
`
	var rows []row
	if err := s.db.WithContext(ctx).Raw(q, subtree).Scan(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]DashboardOhRow, 0, len(rows))
	for _, r := range rows {
		hazard := "其他"
		if r.Type != nil && *r.Type != "" {
			hazard = *r.Type
		}
		out = append(out, DashboardOhRow{Hazard: hazard, People: int(r.People)})
	}
	return out, nil
}

// queryRecentAlerts pulls the last few activity_log rows that look alert-like.
// activity_log has no company_id column — we scope via causer_id → users.id →
// users.company_id, so the ticker only shows entries authored by users in the
// caller's subtree. We don't store an explicit `level`; derive one from the
// event verb / description text.
func (s *SafetyDashboardService) queryRecentAlerts(ctx context.Context, subtree []int64) ([]DashboardAlertEntry, error) {
	type row struct {
		Event       *string
		Description string
		CreatedAt   time.Time
	}
	const q = `
SELECT a.event, a.description, a.created_at
FROM activity_log a
JOIN users u ON u.id = a.causer_id
WHERE u.company_id IN ?
ORDER BY a.created_at DESC
LIMIT 7
`
	var rows []row
	if err := s.db.WithContext(ctx).Raw(q, subtree).Scan(&rows).Error; err != nil {
		// activity_log may not exist in some envs — degrade gracefully.
		return []DashboardAlertEntry{}, nil
	}
	out := make([]DashboardAlertEntry, 0, len(rows))
	for _, r := range rows {
		event := ""
		if r.Event != nil {
			event = *r.Event
		}
		text := r.Description
		if text == "" {
			text = event
		}
		level := "low"
		switch {
		case strings.Contains(event, "delete"), strings.Contains(text, "事故"), strings.Contains(text, "重大"):
			level = "major"
		case strings.Contains(event, "create"), strings.Contains(event, "update"), strings.Contains(text, "整改"):
			level = "normal"
		}
		out = append(out, DashboardAlertEntry{
			Time:  r.CreatedAt.Format("15:04"),
			Level: level,
			Text:  text,
		})
	}
	return out, nil
}

// ── helpers ────────────────────────────────────────────────────────────────

// fillMonthly12 produces exactly 12 monthly buckets ending at `now`, in YY/MM
// format. Missing months from the index map fall to the zero-valued cell so
// the SPA's bar/line chart never has gaps.
func fillMonthly12[R any, T any](now time.Time, index map[string]R, build func(bucket string, row R, ok bool) T) []T {
	out := make([]T, 0, 12)
	for i := 11; i >= 0; i-- {
		t := now.AddDate(0, -i, 0)
		bucket := t.Format("06/01")
		r, ok := index[bucket]
		out = append(out, build(bucket, r, ok))
	}
	return out
}

func roundTo(v float64, decimals int) float64 {
	mult := 1.0
	for i := 0; i < decimals; i++ {
		mult *= 10
	}
	return float64(int(v*mult+0.5)) / mult
}

// parseLooseInt accepts "12人" / " 12 " / "12.0" — anything that starts with
// digits — and returns the leading integer. Used for emergency_drills.
// drill_number which is a free-form string column.
func parseLooseInt(s string) (int, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, false
	}
	end := 0
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	if end == 0 {
		return 0, false
	}
	n := 0
	for i := 0; i < end; i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}
