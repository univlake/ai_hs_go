package risk

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/zap"
)

// newCalc returns a Calculator wired to a sqlmock-backed gorm.
// Caller must defer cleanup() to close the underlying *sql.DB.
func newCalc(t *testing.T) (*Calculator, sqlmock.Sqlmock, func()) {
	t.Helper()
	gdb, mock, raw := newMockGorm(t)
	cleanup := func() { _ = raw.Close() }
	return &Calculator{db: gdb, logger: zap.NewNop(), tz: time.UTC}, mock, cleanup
}

func TestLoadCompanyTreeEmpty(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id", "parent_id"})
	mock.ExpectQuery(`SELECT id, parent_id FROM "companies"`).WillReturnRows(rows)

	tree, err := c.loadCompanyTree(context.Background())
	if err != nil {
		t.Fatalf("loadCompanyTree: %v", err)
	}
	if len(tree) != 0 {
		t.Errorf("want empty tree, got %v", tree)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestLoadCompanyTreeBuildsEdges: rows with parent_id form the edge map; rows
// with NULL or zero parent_id do not appear as keys (root nodes are implicit).
func TestLoadCompanyTreeBuildsEdges(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id", "parent_id"}).
		AddRow(1, nil).      // root, no parent
		AddRow(2, int64(1)). // 1 → 2
		AddRow(3, int64(1)). // 1 → 3
		AddRow(4, int64(2)). // 2 → 4
		AddRow(5, int64(0))  // parent_id == 0 ignored (matches the > 0 guard)
	mock.ExpectQuery(`SELECT id, parent_id FROM "companies"`).WillReturnRows(rows)

	tree, err := c.loadCompanyTree(context.Background())
	if err != nil {
		t.Fatalf("loadCompanyTree: %v", err)
	}
	want := companyTree{1: {2, 3}, 2: {4}}
	if len(tree) != len(want) {
		t.Fatalf("tree size = %d, want %d (got %v)", len(tree), len(want), tree)
	}
	for k, v := range want {
		got := tree[k]
		if len(got) != len(v) {
			t.Errorf("tree[%d] = %v, want %v", k, got, v)
			continue
		}
		for i := range v {
			if got[i] != v[i] {
				t.Errorf("tree[%d][%d] = %d, want %d", k, i, got[i], v[i])
			}
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestComputeBlocsNoBlocs: zero blocs returned → no further DB activity.
func TestComputeBlocsNoBlocs(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(is_bloc`).WillReturnRows(rows)

	if err := c.computeBlocs(context.Background(), "test", time.Now()); err != nil {
		t.Fatalf("computeBlocs: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestComputeBlocsAveraging: one bloc with two children — averaged inherent +
// float written back, and two risk_logs (type 1 + 2) inserted.
//
// Children: inherent (10, 20) → avg 15.00 → "蓝"; float (4, 6) → avg 5.00 → "橙".
func TestComputeBlocsAveraging(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	now := time.Date(2026, 4, 28, 8, 0, 0, 0, time.UTC)
	companyCols := companyAllColumns()

	// Step 1 — bloc query returns one bloc.
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(is_bloc`).WillReturnRows(
		sqlmock.NewRows(companyCols).AddRow(blocRowValues(100, "集团", 1, nil, ptrFloat(0), ptrStr("红"), ptrFloat(0), ptrStr("红"))...),
	)
	// Step 2 — loadCompanyTree.
	mock.ExpectQuery(`SELECT id, parent_id FROM "companies"`).WillReturnRows(
		sqlmock.NewRows([]string{"id", "parent_id"}).
			AddRow(100, nil).
			AddRow(11, int64(100)).
			AddRow(12, int64(100)),
	)
	// Step 3 — children-of-bloc query returns 2 producing companies.
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(id IN`).WillReturnRows(
		sqlmock.NewRows(companyCols).
			AddRow(blocRowValues(11, "C1", 0, ptrInt64(100), ptrFloat(10), ptrStr("黄"), ptrFloat(4), ptrStr("红"))...).
			AddRow(blocRowValues(12, "C2", 0, ptrInt64(100), ptrFloat(20), ptrStr("蓝"), ptrFloat(6), ptrStr("橙"))...),
	)
	// Step 4 — UPDATE bloc with averaged risk + grade. GORM wraps Updates()
	// in a default transaction; map iteration order in the SET clause is
	// non-deterministic, so we only assert structure.
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "companies" SET .*"float_risk".*WHERE id = \$\d+`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), int64(100)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	// Step 5 — INSERT risk_logs (2 rows: type=1 inherent, type=2 float).
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "risk_logs"`).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2),
	)
	mock.ExpectCommit()

	if err := c.computeBlocs(context.Background(), "test", now); err != nil {
		t.Fatalf("computeBlocs: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestComputeBlocsShallowFirst: nested blocs (outer 100 → inner 200 → leaves).
// Verifies inner is written BEFORE outer's child query runs (sqlmock enforces
// expectation ordering, so setting them in shallow-first order locks down the
// "parent reads child's freshly-written average" semantic).
func TestComputeBlocsShallowFirst(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	now := time.Date(2026, 4, 28, 8, 0, 0, 0, time.UTC)
	cols := companyAllColumns()

	// Bloc query returns outer first then inner — verifies the sort step
	// reorders so the deeper-tree bloc (outer) lands later.
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(is_bloc`).WillReturnRows(
		sqlmock.NewRows(cols).
			AddRow(blocRowValues(100, "外", 1, nil, ptrFloat(0), ptrStr("红"), ptrFloat(0), ptrStr("红"))...).
			AddRow(blocRowValues(200, "内", 1, ptrInt64(100), ptrFloat(0), ptrStr("红"), ptrFloat(0), ptrStr("红"))...),
	)
	// Tree: 100 → 200 → {21, 22}. Outer's only direct child is the inner bloc.
	mock.ExpectQuery(`SELECT id, parent_id FROM "companies"`).WillReturnRows(
		sqlmock.NewRows([]string{"id", "parent_id"}).
			AddRow(100, nil).
			AddRow(200, int64(100)).
			AddRow(21, int64(200)).
			AddRow(22, int64(200)),
	)

	// — Inner (200) processed first: load its leaf children, write its average.
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(id IN`).WillReturnRows(
		sqlmock.NewRows(cols).
			AddRow(blocRowValues(21, "L1", 0, ptrInt64(200), ptrFloat(8), ptrStr("橙"), ptrFloat(2), ptrStr("红"))...).
			AddRow(blocRowValues(22, "L2", 0, ptrInt64(200), ptrFloat(12), ptrStr("黄"), ptrFloat(4), ptrStr("红"))...),
	)
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "companies" SET .* WHERE id = \$\d+`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), int64(200)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// — Outer (100): child query reads inner with its FRESH averaged values.
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(id IN`).WillReturnRows(
		sqlmock.NewRows(cols).
			AddRow(blocRowValues(200, "内", 1, ptrInt64(100), ptrFloat(10), ptrStr("黄"), ptrFloat(3), ptrStr("红"))...),
	)
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "companies" SET .* WHERE id = \$\d+`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), int64(100)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "risk_logs"`).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2).AddRow(3).AddRow(4),
	)
	mock.ExpectCommit()

	if err := c.computeBlocs(context.Background(), "test", now); err != nil {
		t.Fatalf("computeBlocs: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestComputeBlocsSkipsLeafBloc: a bloc whose tree entry is empty is skipped
// (no child query, no update). Guards the `len(childIDs) == 0` continue.
func TestComputeBlocsSkipsLeafBloc(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	cols := companyAllColumns()
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(is_bloc`).WillReturnRows(
		sqlmock.NewRows(cols).AddRow(blocRowValues(50, "孤集团", 1, nil, ptrFloat(0), ptrStr("红"), ptrFloat(0), ptrStr("红"))...),
	)
	mock.ExpectQuery(`SELECT id, parent_id FROM "companies"`).WillReturnRows(
		sqlmock.NewRows([]string{"id", "parent_id"}).AddRow(50, nil),
	)
	// No child query, no UPDATE, no INSERT — just the two SELECTs above.

	if err := c.computeBlocs(context.Background(), "test", time.Now()); err != nil {
		t.Fatalf("computeBlocs: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// TestComputeBlocsAllChildrenFiltered: child query returns 0 rows (all children
// have inherent_risk IS NULL or status != 1). Bloc is skipped, no risk_logs
// emitted for it. Guards the second `len(children) == 0` continue.
func TestComputeBlocsAllChildrenFiltered(t *testing.T) {
	c, mock, cleanup := newCalc(t)
	defer cleanup()

	cols := companyAllColumns()
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(is_bloc`).WillReturnRows(
		sqlmock.NewRows(cols).AddRow(blocRowValues(70, "G", 1, nil, ptrFloat(0), ptrStr("红"), ptrFloat(0), ptrStr("红"))...),
	)
	mock.ExpectQuery(`SELECT id, parent_id FROM "companies"`).WillReturnRows(
		sqlmock.NewRows([]string{"id", "parent_id"}).
			AddRow(70, nil).
			AddRow(71, int64(70)),
	)
	// Child query returns no rows → bloc is skipped, no UPDATE, no INSERT.
	mock.ExpectQuery(`SELECT .* FROM "companies" WHERE \(id IN`).WillReturnRows(
		sqlmock.NewRows(cols),
	)

	if err := c.computeBlocs(context.Background(), "test", time.Now()); err != nil {
		t.Fatalf("computeBlocs: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

// companyAllColumns lists columns in the order returned by GORM's SELECT *
// for model.Company. Pulled from model/company.go field order.
func companyAllColumns() []string {
	return []string{
		"id", "title", "enterprise_type_id", "operational_type_id", "area_id",
		"address", "established", "credit_code", "business_scope", "company_introduction",
		"longitude", "dimension", "geographic_range", "acreage",
		"supervise_name", "supervise_tel", "supervise_fix", "supervise_email", "supervise_position",
		"status", "parent_id", "root_id", "is_produce", "is_bloc",
		"inherent_risk", "inherent_risk_grade", "float_risk", "float_risk_grade",
		"user_id", "last_user_id", "created_at", "updated_at", "deleted_at",
	}
}

// blocRowValues fills a row matching companyAllColumns() with sensible defaults
// for fields that don't matter to the test, leaving the salient ones overridable.
func blocRowValues(id int64, title string, isBloc int16, parentID *int64,
	inherentRisk *float64, inherentGrade *string,
	floatRisk *float64, floatGrade *string,
) []driver.Value {
	return []driver.Value{
		id, title, nil, int64(1), nil,
		nil, nil, nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil, nil, nil,
		int16(1), parentID, int64(0), int16(0), isBloc,
		inherentRisk, inherentGrade, floatRisk, floatGrade,
		int64(0), int64(0), time.Now(), time.Now(), nil,
	}
}
