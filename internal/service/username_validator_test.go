package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockGorm(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	t.Helper()
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("sqlmock: %v", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open: %v", err)
	}
	return gormDB, mock, mockDB
}

func TestCheckUsername(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		exclude  int64
		setup    func(sqlmock.Sqlmock)
		want     UsernameCheck
		hasMock  bool
	}{
		{name: "empty", input: "", want: UsernameCheck{Reason: "用户名不能为空"}},
		{name: "too short", input: "ab", want: UsernameCheck{Reason: "用户名长度须为 3–30 个字符"}},
		{name: "too long", input: "abcdefghijabcdefghijabcdefghij1", want: UsernameCheck{Reason: "用户名长度须为 3–30 个字符"}},
		{name: "leading underscore", input: "_alice", want: UsernameCheck{Reason: "只允许字母、数字、下划线、连字符、点，且首尾必须是字母或数字"}},
		{name: "reserved", input: "c12_admin", want: UsernameCheck{Reason: "该格式为系统保留，不可使用"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gormDB, _, raw := newMockGorm(t)
			defer raw.Close()
			got := CheckUsername(context.Background(), gormDB, tc.input, tc.exclude)
			if got != tc.want {
				t.Fatalf("got %+v want %+v", got, tc.want)
			}
		})
	}
}
