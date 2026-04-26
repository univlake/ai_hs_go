package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/repository"
	"ai_hs_go/internal/response"
)

// ResolveScope resolves the X-Scope-Company-Id header (or `company_id` query
// param fallback) and validates the active user can see that company.
//
// Behavior:
//   - Platform admin: any company id passes; if none provided, scope stays empty.
//   - Regular user: company_id must be a descendant of the user's company,
//     within the same root_id forest.
//   - Missing header for regular user: defaults to user's own company id.
func ResolveScope(repo *repository.CompanyRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("X-Scope-Company-Id")
		if raw == "" {
			raw = c.Query("company_id")
		}
		var targetID int64
		if raw != "" {
			parsed, err := strconv.ParseInt(raw, 10, 64)
			if err != nil || parsed <= 0 {
				response.BadRequest(c, "X-Scope-Company-Id 无效")
				return
			}
			targetID = parsed
		}

		if IsPlatformAdmin(c) {
			if targetID != 0 {
				c.Set(CtxKeyScopeCompanyID, targetID)
			}
			c.Next()
			return
		}

		userCompany, ok := c.Get(CtxKeyCompanyID)
		var ownCompanyID int64
		if ok {
			if v, ok := userCompany.(int64); ok {
				ownCompanyID = v
			}
		}
		if ownCompanyID == 0 {
			response.Forbidden(c, "当前账号未绑定企业")
			return
		}
		if targetID == 0 {
			targetID = ownCompanyID
		}

		// Cross-tenant check: target must be self-or-descendant of own company.
		ok2, err := repo.IsAncestorOrSelf(c.Request.Context(), ownCompanyID, targetID)
		if err != nil {
			response.Internal(c, "范围校验失败")
			return
		}
		if !ok2 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"code":    http.StatusForbidden,
				"message": "无权访问该企业数据",
				"reason":  "cross_tenant",
			})
			return
		}
		c.Set(CtxKeyScopeCompanyID, targetID)
		c.Next()
	}
}

func ScopeCompanyID(c *gin.Context) (int64, bool) {
	if v, ok := c.Get(CtxKeyScopeCompanyID); ok {
		if id, ok := v.(int64); ok {
			return id, true
		}
	}
	return 0, false
}
