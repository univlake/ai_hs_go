package risk

import (
	"context"
	"sort"
	"time"

	"ai_hs_go/internal/model"
)

// companyTree maps a company id → its direct (status=1) child ids. Built
// once per run; key never appears unless the company has at least one child.
type companyTree map[int64][]int64

// computeBlocs replaces the package-level stub. For each bloc (is_bloc=1,
// status=1), set inherent_risk + float_risk to the average of its DIRECT
// children only (not transitive descendants), then emit risk_logs deltas.
//
// Blocs are processed shallow-subtree-first so a parent bloc reads its
// child blocs' freshly-written averages — a bloc that is also is_produce
// has its Phase-2 produce score overwritten here. (Mirrors the same comment
// in app/Traits/CalculateRiskValue.php — load-bearing semantics.)
func (c *Calculator) computeBlocs(ctx context.Context, content string, now time.Time) error {
	db := c.db.WithContext(ctx)

	var blocs []model.Company
	if err := db.Where("is_bloc = ? AND status = ?", 1, 1).Find(&blocs).Error; err != nil {
		return err
	}
	if len(blocs) == 0 {
		return nil
	}

	tree, err := c.loadCompanyTree(ctx)
	if err != nil {
		return err
	}

	depths := make(map[int64]int, len(blocs))
	for i := range blocs {
		depths[blocs[i].ID] = subtreeDepth(blocs[i].ID, tree)
	}
	sort.SliceStable(blocs, func(i, j int) bool {
		return depths[blocs[i].ID] < depths[blocs[j].ID]
	})

	dateInt := dateInt(now, c.tz)
	var blocRiskLogs []model.RiskLog

	for i := range blocs {
		bloc := &blocs[i]
		childIDs := tree[bloc.ID]
		if len(childIDs) == 0 {
			continue
		}

		var children []model.Company
		if err := db.Where("id IN ? AND status = ? AND inherent_risk IS NOT NULL", childIDs, 1).
			Find(&children).Error; err != nil {
			return err
		}
		if len(children) == 0 {
			continue
		}

		var inherentSum, floatSum float64
		for _, child := range children {
			if child.InherentRisk != nil {
				inherentSum += *child.InherentRisk
			}
			if child.FloatRisk != nil {
				floatSum += *child.FloatRisk
			}
		}
		count := float64(len(children))
		inherentAvg := round2(inherentSum / count)
		floatAvg := round2(floatSum / count)
		inherentGrade := riskGrade(inherentAvg)
		floatGrade := riskGrade(floatAvg)

		blocRiskLogs = append(blocRiskLogs,
			buildRiskLog(bloc, 1, bloc.InherentRisk, inherentAvg, bloc.InherentRiskGrade, inherentGrade, content, dateInt, now),
			buildRiskLog(bloc, 2, bloc.FloatRisk, floatAvg, bloc.FloatRiskGrade, floatGrade, content, dateInt, now),
		)

		// Write per-bloc immediately so any shallower bloc that has this one
		// as a direct child reads the freshly-averaged value below.
		if err := db.Model(&model.Company{}).Where("id = ?", bloc.ID).
			Updates(map[string]any{
				"inherent_risk":       inherentAvg,
				"inherent_risk_grade": inherentGrade,
				"float_risk":          floatAvg,
				"float_risk_grade":    floatGrade,
			}).Error; err != nil {
			return err
		}
	}

	if len(blocRiskLogs) > 0 {
		if err := db.CreateInBatches(&blocRiskLogs, 500).Error; err != nil {
			return err
		}
	}
	return nil
}

// loadCompanyTree maps every active company's id → list of direct child ids
// in a single query. Replaces the legacy recursive Company::where('parent_id')
// which issued one query per descendant level per bloc.
func (c *Calculator) loadCompanyTree(ctx context.Context) (companyTree, error) {
	type row struct {
		ID       int64  `gorm:"column:id"`
		ParentID *int64 `gorm:"column:parent_id"`
	}
	var rows []row
	if err := c.db.WithContext(ctx).Table("companies").
		Select("id, parent_id").
		Where("status = ? AND deleted_at IS NULL", 1).
		Order("id").
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	tree := companyTree{}
	for _, r := range rows {
		if r.ParentID != nil && *r.ParentID > 0 {
			tree[*r.ParentID] = append(tree[*r.ParentID], r.ID)
		}
	}
	return tree, nil
}

// subtreeDepth returns the maximum depth of descendants below companyID. A
// leaf returns 0; a node whose only children are leaves returns 1; etc.
// Sorts blocs into a bottom-up processing order in computeBlocs.
func subtreeDepth(companyID int64, tree companyTree) int {
	children := tree[companyID]
	if len(children) == 0 {
		return 0
	}
	max := 0
	for _, child := range children {
		d := 1 + subtreeDepth(child, tree)
		if d > max {
			max = d
		}
	}
	return max
}
