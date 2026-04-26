package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

// ActivityLog records every successful write action (POST/PUT/PATCH/DELETE)
// after the handler completes. Subject is parsed from the path; details
// (attribute_changes) are best-effort — the existing GORM hooks layer is the
// preferred source for attribute_changes, this middleware records the audit
// trail with description + path + causer for high-level reporting.
func ActivityLog(db *gorm.DB, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		mutating := method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE"
		c.Next()
		if !mutating {
			return
		}
		status := c.Writer.Status()
		if status >= 400 {
			return
		}
		subjectType, subjectID := parseSubject(c.Request.URL.Path)
		causerID := CurrentUserID(c)
		event := strings.ToLower(method)
		log := &model.ActivityLog{
			LogName:     ptrString("audit"),
			Description: c.Request.Method + " " + c.Request.URL.Path,
			SubjectType: ptrStringIfSet(subjectType),
			SubjectID:   subjectID,
			CauserType:  ptrStringIfSet("App\\Models\\User"),
			CauserID:    int64Ptr(causerID),
			Event:       &event,
			Properties:  []byte(`{}`),
			AttributeChanges: []byte(`{}`),
		}
		if err := db.WithContext(c.Request.Context()).Create(log).Error; err != nil {
			logger.Warn("activity log insert", zap.Error(err))
		}
	}
}

func parseSubject(path string) (string, *int64) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	// expect /api/v1/<resource>/<id>
	if len(parts) < 3 {
		return "", nil
	}
	if parts[0] != "api" || parts[1] != "v1" {
		return "", nil
	}
	resource := parts[2]
	if len(parts) >= 4 {
		if id := parseInt(parts[3]); id > 0 {
			return resource, &id
		}
	}
	return resource, nil
}

func parseInt(s string) int64 {
	var n int64
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int64(c-'0')
	}
	return n
}

func ptrString(s string) *string { return &s }
func ptrStringIfSet(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func int64Ptr(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}
