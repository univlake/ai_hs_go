package risk

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// scanCountMap / scanSumMap are thin wrappers over Scan() that translate
// rows into a map keyed by company_id. The full loadBulk method is excluded
// from the unit suite — its 30+ Postgres-specific queries (CAST(NULLIF…),
// CASE WHEN aggregations) would turn the test into a mirror of the SQL
// strings rather than an independent check. Verify those at the integration
// layer instead.

func TestScanCountMap(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	mock.ExpectQuery(`SELECT company_id, COUNT.*FROM "foo"`).WillReturnRows(
		sqlmock.NewRows([]string{"company_id", "cnt"}).
			AddRow(int64(1), int64(5)).
			AddRow(int64(2), int64(0)).
			AddRow(int64(3), int64(99)),
	)

	dst := map[int64]int64{}
	tx := gdb.Table("foo").Select("company_id, COUNT(*) AS cnt").Group("company_id")
	if err := scanCountMap(tx, dst); err != nil {
		t.Fatalf("scanCountMap: %v", err)
	}
	want := map[int64]int64{1: 5, 2: 0, 3: 99}
	if len(dst) != len(want) {
		t.Fatalf("size = %d, want %d (%v)", len(dst), len(want), dst)
	}
	for k, v := range want {
		if dst[k] != v {
			t.Errorf("dst[%d] = %d, want %d", k, dst[k], v)
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestScanCountMapEmpty(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	mock.ExpectQuery(`SELECT.*FROM "foo"`).WillReturnRows(
		sqlmock.NewRows([]string{"company_id", "cnt"}),
	)
	dst := map[int64]int64{}
	tx := gdb.Table("foo").Select("company_id, COUNT(*) AS cnt").Group("company_id")
	if err := scanCountMap(tx, dst); err != nil {
		t.Fatalf("empty result: %v", err)
	}
	if len(dst) != 0 {
		t.Errorf("dst should remain empty, got %v", dst)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestScanCountMapErrorPropagates(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	want := errors.New("boom")
	mock.ExpectQuery(`SELECT.*FROM "foo"`).WillReturnError(want)
	dst := map[int64]int64{}
	tx := gdb.Table("foo").Select("company_id, COUNT(*) AS cnt")
	if err := scanCountMap(tx, dst); err == nil {
		t.Fatal("want error, got nil")
	} else if !errors.Is(err, want) {
		// gorm wraps driver errors; compare on message rather than identity.
		if err.Error() != want.Error() {
			t.Errorf("err = %v, want %v", err, want)
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestScanSumMap(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	mock.ExpectQuery(`SELECT company_id, .*SUM.*FROM "bar"`).WillReturnRows(
		sqlmock.NewRows([]string{"company_id", "total"}).
			AddRow(int64(1), 12.5).
			AddRow(int64(7), 0.0),
	)

	dst := map[int64]float64{}
	tx := gdb.Table("bar").Select("company_id, COALESCE(SUM(amount),0) AS total").Group("company_id")
	if err := scanSumMap(tx, dst); err != nil {
		t.Fatalf("scanSumMap: %v", err)
	}
	if dst[1] != 12.5 {
		t.Errorf("dst[1] = %v, want 12.5", dst[1])
	}
	if dst[7] != 0 {
		t.Errorf("dst[7] = %v, want 0", dst[7])
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestScanSumMapErrorPropagates(t *testing.T) {
	gdb, mock, raw := newMockGorm(t)
	defer raw.Close()

	want := errors.New("kaboom")
	mock.ExpectQuery(`SELECT.*FROM "bar"`).WillReturnError(want)
	dst := map[int64]float64{}
	tx := gdb.Table("bar").Select("company_id, SUM(x) AS total")
	if err := scanSumMap(tx, dst); err == nil {
		t.Fatal("want error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
