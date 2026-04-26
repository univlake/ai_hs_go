package service

import (
	"context"
	"regexp"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

const (
	UsernameMinLen = 3
	UsernameMaxLen = 30
)

var (
	usernameFormatRE   = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9_\-.]{1,28}[a-zA-Z0-9])?$`)
	usernameReservedRE = regexp.MustCompile(`^c\d+_admin$`)
)

type UsernameCheck struct {
	Available bool   `json:"available"`
	Reason    string `json:"reason,omitempty"`
}

// CheckUsername mirrors UsernameValidator::check from the Laravel codebase.
// Pass excludeUserID = 0 to disable exclusion.
func CheckUsername(ctx context.Context, db *gorm.DB, username string, excludeUserID int64) UsernameCheck {
	username = strings.TrimSpace(username)
	if username == "" {
		return UsernameCheck{Reason: "用户名不能为空"}
	}
	if l := len([]rune(username)); l < UsernameMinLen || l > UsernameMaxLen {
		return UsernameCheck{Reason: "用户名长度须为 3–30 个字符"}
	}
	if !usernameFormatRE.MatchString(username) {
		return UsernameCheck{Reason: "只允许字母、数字、下划线、连字符、点，且首尾必须是字母或数字"}
	}
	if usernameReservedRE.MatchString(username) {
		return UsernameCheck{Reason: "该格式为系统保留，不可使用"}
	}

	q := db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username)
	if excludeUserID > 0 {
		q = q.Where("id != ?", excludeUserID)
	}
	var count int64
	if err := q.Count(&count).Error; err != nil {
		return UsernameCheck{Reason: "校验失败，请稍后再试"}
	}
	if count > 0 {
		return UsernameCheck{Reason: "该用户名已被使用"}
	}

	return UsernameCheck{Available: true}
}
