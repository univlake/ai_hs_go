// Package apiutil holds small helpers shared across handler subpackages
// (risk, expert, healthfire, etc.). Keeping these in one place avoids the
// per-module copies that previously drifted out of sync (e.g. parseDateTime
// supporting different input layouts in different modules).
package apiutil

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
)

var dateLayouts = []string{
	"2006-01-02",
	"2006/01/02",
	"2006/1/2",
	time.RFC3339,
}

// Union of layouts previously accepted across statistics/notification/expert/
// meetingtraining handlers. Listed roughly most-specific-first so common cases
// hit early.
var dateTimeLayouts = []string{
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04", // <input type="datetime-local">
	"2006-01-02 15:04",
	"2006-01-02",
	time.RFC3339,
	time.RFC3339Nano,
}

func ParseDate(p *string) *model.Date {
	if p == nil || *p == "" {
		return nil
	}
	for _, f := range dateLayouts {
		if t, err := time.Parse(f, *p); err == nil {
			d := model.Date(t)
			return &d
		}
	}
	return nil
}

func ParseDateTime(p *string) *model.DateTime {
	if p == nil || *p == "" {
		return nil
	}
	for _, f := range dateTimeLayouts {
		if t, err := time.Parse(f, *p); err == nil {
			dt := model.DateTime(t)
			return &dt
		}
	}
	return nil
}

func AddNullableString(m map[string]any, key string, v *string) {
	if v != nil {
		m[key] = v
	}
}

func FirstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// ScopeCompanyOrFail returns the caller's scoped company_id, or writes a 400
// response and returns 0 when the scope is missing. The cid==0 guard is
// defensive against future middleware changes — current ScopeCompanyID
// returns (0,false) on miss.
func ScopeCompanyOrFail(c *gin.Context) int64 {
	cid, ok := middleware.ScopeCompanyID(c)
	if !ok || cid == 0 {
		response.BadRequest(c, "缺少 X-Scope-Company-Id 或 company_id")
		return 0
	}
	return cid
}

// ScopeOrCompanyID prefers an explicit ?company_id= query param (group / admin
// flows) and falls back to the request scope. Returns 0 when neither resolves.
func ScopeOrCompanyID(c *gin.Context) int64 {
	if v := c.Query("company_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			return id
		}
	}
	if cid, ok := middleware.ScopeCompanyID(c); ok {
		return cid
	}
	return 0
}
