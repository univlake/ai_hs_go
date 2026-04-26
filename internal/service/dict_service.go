package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
)

// DictService backs every base-data 字典 endpoint. All 9 dictionary tables
// share the same {id, title, sort, timestamps, deleted_at} shape and the
// same CRUD semantics, so they ride a single service driven by a
// whitelisted table name. Whitelisting matters: the URL `:name` parameter
// flows directly into the SQL FROM clause, so accepting unbound names
// would let any client read/write arbitrary tables.
type DictService struct {
	db *gorm.DB
}

func NewDictService(db *gorm.DB) *DictService {
	return &DictService{db: db}
}

// dictTableForSlug returns the SQL table name for a frontend URL segment
// (e.g. "enterprise-types" → "enterprise_types"). Returns empty string
// when the slug isn't a known dictionary — caller must treat that as 404
// to avoid leaking which tables exist.
func dictTableForSlug(slug string) string {
	switch slug {
	case "enterprise-types",
		"operational-types",
		"record-types",
		"process-types",
		"special-types",
		"space-work-types",
		"safety-management-types",
		"dangerous-craft-types":
		// regulation-library-types is intentionally NOT in this list — its
		// table has extra hierarchy columns (pid / all_name / audit) and
		// flows through RegulationLibraryTypeService.
		return strings.ReplaceAll(slug, "-", "_")
	}
	return ""
}

// IsKnownDict reports whether `slug` resolves to a dictionary table.
// Handlers use this to bail out early with 404 before touching the DB.
func (s *DictService) IsKnownDict(slug string) bool {
	return dictTableForSlug(slug) != ""
}

// DictRow is the on-the-wire shape — kept narrow so future per-table
// columns (e.g. safety_management_types.is_cert) don't accidentally leak
// across dictionaries through a shared struct.
type DictRow struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DictListParams struct {
	Page    int
	PerPage int
	Title   string // fuzzy match
}

// List returns rows for a dictionary table. Always paginated — the
// frontend CrudPage reads `data.data.data` so a flat array would surface
// as an empty list.
func (s *DictService) List(ctx context.Context, slug string, p DictListParams) ([]DictRow, int64, error) {
	table := dictTableForSlug(slug)
	if table == "" {
		return nil, 0, errs.New(errs.CodeNotFound, "未知字典")
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Table(table).Where("deleted_at IS NULL")
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ?", "%"+t+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []DictRow
	if err := q.
		Select("id", "title", "sort", "created_at", "updated_at").
		Order("sort ASC, id ASC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// Get returns a single row.
func (s *DictService) Get(ctx context.Context, slug string, id int64) (*DictRow, error) {
	table := dictTableForSlug(slug)
	if table == "" {
		return nil, errs.New(errs.CodeNotFound, "未知字典")
	}
	var row DictRow
	err := s.db.WithContext(ctx).Table(table).
		Select("id", "title", "sort", "created_at", "updated_at").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.New(errs.CodeNotFound, "记录不存在")
	}
	if err != nil {
		return nil, err
	}
	return &row, nil
}

type DictMutateInput struct {
	Title string
	Sort  int
}

// Create inserts a new row. Returns the persisted record.
//
// Implementation note: we use a Raw INSERT ... RETURNING because
// `db.Table(name).Create(&map[string]any{...})` does not reliably populate
// the auto-increment id back into the map across PG drivers, and a follow
// up SELECT-by-title isn't safe (titles can collide). RETURNING is one
// round-trip and pulls every column we need for the response.
func (s *DictService) Create(ctx context.Context, slug string, in DictMutateInput) (*DictRow, error) {
	table := dictTableForSlug(slug)
	if table == "" {
		return nil, errs.New(errs.CodeNotFound, "未知字典")
	}
	if strings.TrimSpace(in.Title) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "标题不能为空")
	}
	q := fmt.Sprintf(
		`INSERT INTO %s (title, sort, created_at, updated_at)
		 VALUES (?, ?, NOW(), NOW())
		 RETURNING id, title, sort, created_at, updated_at`,
		table,
	)
	var row DictRow
	if err := s.db.WithContext(ctx).Raw(q, in.Title, in.Sort).Scan(&row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return &row, nil
}

// Update partially patches the row.
func (s *DictService) Update(ctx context.Context, slug string, id int64, in DictMutateInput) (*DictRow, error) {
	table := dictTableForSlug(slug)
	if table == "" {
		return nil, errs.New(errs.CodeNotFound, "未知字典")
	}
	if strings.TrimSpace(in.Title) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "标题不能为空")
	}
	res := s.db.WithContext(ctx).Table(table).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]any{
			"title":      in.Title,
			"sort":       in.Sort,
			"updated_at": gorm.Expr("NOW()"),
		})
	if res.Error != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, errs.New(errs.CodeNotFound, "记录不存在")
	}
	return s.Get(ctx, slug, id)
}

// Delete soft-deletes the row (sets deleted_at).
func (s *DictService) Delete(ctx context.Context, slug string, id int64) error {
	table := dictTableForSlug(slug)
	if table == "" {
		return errs.New(errs.CodeNotFound, "未知字典")
	}
	res := s.db.WithContext(ctx).Table(table).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("NOW()"))
	if res.Error != nil {
		return errs.Wrap(errs.CodeInternal, "删除失败", res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "记录不存在")
	}
	return nil
}

