package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// Date is a thin wrapper around time.Time that JSON-serializes as
// "YYYY-MM-DD" instead of the full RFC 3339 timestamp.
//
// Use it on date-only columns (e.g. valid_period, filing_time, get_time)
// where the wall-clock time component would only confuse the user.
// gorm scans/saves the same way as time.Time — the only behaviour change
// is on JSON marshalling.
type Date time.Time

const dateLayout = "2006-01-02"

func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + time.Time(d).Format(dateLayout) + `"`), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		*d = Date(time.Time{})
		return nil
	}
	// Accept the Date format AND the RFC 3339 form Postgres can ship
	// (timestamps with time zone), so a re-saved record round-trips.
	for _, layout := range []string{dateLayout, time.RFC3339, time.RFC3339Nano, "2006-01-02 15:04:05"} {
		if t, err := time.Parse(layout, s); err == nil {
			*d = Date(t)
			return nil
		}
	}
	return fmt.Errorf("invalid date: %q", s)
}

func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}

func (d *Date) Scan(value any) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
	case []byte:
		t, err := time.Parse(dateLayout, string(v))
		if err != nil {
			return err
		}
		*d = Date(t)
	case string:
		t, err := time.Parse(dateLayout, v)
		if err != nil {
			return err
		}
		*d = Date(t)
	default:
		return fmt.Errorf("Date.Scan: unsupported type %T", value)
	}
	return nil
}

// DateTime is a thin wrapper around time.Time that JSON-serializes as
// "YYYY-MM-DD HH:mm:ss" — friendlier for table columns and detail
// drawers than the default RFC 3339. Used on audit timestamps
// (created_at / updated_at).
type DateTime time.Time

const dateTimeLayout = "2006-01-02 15:04:05"

func (d DateTime) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + time.Time(d).Format(dateTimeLayout) + `"`), nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		*d = DateTime(time.Time{})
		return nil
	}
	for _, layout := range []string{dateTimeLayout, time.RFC3339, time.RFC3339Nano, dateLayout} {
		if t, err := time.Parse(layout, s); err == nil {
			*d = DateTime(t)
			return nil
		}
	}
	return fmt.Errorf("invalid datetime: %q", s)
}

func (d DateTime) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}

func (d *DateTime) Scan(value any) error {
	if value == nil {
		*d = DateTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DateTime(v)
	case []byte:
		t, err := time.Parse(dateTimeLayout, string(v))
		if err != nil {
			return err
		}
		*d = DateTime(t)
	case string:
		t, err := time.Parse(dateTimeLayout, v)
		if err != nil {
			return err
		}
		*d = DateTime(t)
	default:
		return fmt.Errorf("DateTime.Scan: unsupported type %T", value)
	}
	return nil
}
