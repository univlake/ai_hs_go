package middleware

import (
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/rbac"
	"ai_hs_go/internal/response"
)

// RequirePermission gates a route by a single Spatie-style permission name like
// "companies.list" — split into resource/action for the Casbin matcher.
//
// Platform admins (is_platform_admin=true) bypass the check entirely.
func RequirePermission(enforcer *casbin.Enforcer, permission string) gin.HandlerFunc {
	resource, action := splitPermission(permission)
	return func(c *gin.Context) {
		if IsPlatformAdmin(c) {
			c.Next()
			return
		}
		uid := CurrentUserID(c)
		if uid == 0 {
			response.Unauthorized(c, "")
			return
		}
		ok, err := enforcer.Enforce(rbac.SubjectForUser(uid), resource, action)
		if err != nil {
			response.Internal(c, "权限校验失败")
			return
		}
		if !ok {
			response.Forbidden(c, "无权执行此操作")
			return
		}
		c.Next()
	}
}

// RequireAnyPermission allows access when the user has at least one of the
// listed permissions.
func RequireAnyPermission(enforcer *casbin.Enforcer, permissions ...string) gin.HandlerFunc {
	type pair struct{ resource, action string }
	pairs := make([]pair, len(permissions))
	for i, p := range permissions {
		r, a := splitPermission(p)
		pairs[i] = pair{r, a}
	}
	return func(c *gin.Context) {
		if IsPlatformAdmin(c) {
			c.Next()
			return
		}
		uid := CurrentUserID(c)
		if uid == 0 {
			response.Unauthorized(c, "")
			return
		}
		sub := rbac.SubjectForUser(uid)
		for _, p := range pairs {
			if ok, _ := enforcer.Enforce(sub, p.resource, p.action); ok {
				c.Next()
				return
			}
		}
		response.Forbidden(c, "无权执行此操作")
	}
}

// splitPermission turns "companies.list" into ("companies", "list").
// Falls back to ("*", permission) when no dot is present.
func splitPermission(p string) (string, string) {
	idx := strings.LastIndex(p, ".")
	if idx < 0 {
		return "*", p
	}
	return p[:idx], p[idx+1:]
}
