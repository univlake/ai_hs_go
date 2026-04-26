package repository

import (
	"context"

	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) *CompanyRepo {
	return &CompanyRepo{db: db}
}

func (r *CompanyRepo) DB() *gorm.DB { return r.db }

// DescendantIDs returns ids of `id` and every descendant within the same root.
// Uses recursive CTE; bounded depth check is enforced via WHERE root_id = same.
func (r *CompanyRepo) DescendantIDs(ctx context.Context, id int64) ([]int64, error) {
	const q = `
WITH RECURSIVE descendants AS (
    SELECT id, parent_id, root_id FROM companies WHERE id = ? AND deleted_at IS NULL
    UNION ALL
    SELECT c.id, c.parent_id, c.root_id FROM companies c
    JOIN descendants d ON c.parent_id = d.id
    WHERE c.root_id = d.root_id AND c.deleted_at IS NULL
)
SELECT id FROM descendants
`
	var ids []int64
	if err := r.db.WithContext(ctx).Raw(q, id).Scan(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

// IsAncestorOrSelf returns true when ancestor is `target` or one of its
// ancestors within the same root.
func (r *CompanyRepo) IsAncestorOrSelf(ctx context.Context, ancestor, target int64) (bool, error) {
	if ancestor == target {
		return true, nil
	}
	const q = `
WITH RECURSIVE ancestors AS (
    SELECT id, parent_id, root_id FROM companies WHERE id = ? AND deleted_at IS NULL
    UNION ALL
    SELECT c.id, c.parent_id, c.root_id FROM companies c
    JOIN ancestors a ON a.parent_id = c.id
    WHERE c.root_id = a.root_id AND c.deleted_at IS NULL
)
SELECT 1 FROM ancestors WHERE id = ? LIMIT 1
`
	var hit int
	err := r.db.WithContext(ctx).Raw(q, target, ancestor).Scan(&hit).Error
	if err != nil {
		return false, err
	}
	return hit == 1, nil
}

func (r *CompanyRepo) GetByID(ctx context.Context, id int64) (*model.Company, error) {
	var c model.Company
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// GetByIDWithRelations returns a company with EnterpriseType / OperationalType
// preloaded — the detail screen / current company info form needs both.
func (r *CompanyRepo) GetByIDWithRelations(ctx context.Context, id int64) (*model.Company, error) {
	var c model.Company
	if err := r.db.WithContext(ctx).
		Preload("EnterpriseType").
		Preload("OperationalType").
		First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// Forest returns every root company (parent_id IS NULL).
func (r *CompanyRepo) Forest(ctx context.Context) ([]model.Company, error) {
	var roots []model.Company
	if err := r.db.WithContext(ctx).Where("parent_id IS NULL").Order("id").Find(&roots).Error; err != nil {
		return nil, err
	}
	return roots, nil
}

// Subtree returns the entire subtree rooted at id.
func (r *CompanyRepo) Subtree(ctx context.Context, id int64) ([]model.Company, error) {
	ids, err := r.DescendantIDs(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, nil
	}
	var rows []model.Company
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Order("id").Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// Children returns direct children of parent.
func (r *CompanyRepo) Children(ctx context.Context, parentID int64) ([]model.Company, error) {
	var rows []model.Company
	if err := r.db.WithContext(ctx).Where("parent_id = ?", parentID).Order("id").Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
