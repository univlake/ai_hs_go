package middleware

import (
	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/response"
)

// RequirePlatformAdmin gates a route to platform admins only (used for /tenants).
func RequirePlatformAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsPlatformAdmin(c) {
			response.Forbidden(c, "仅平台管理员可访问")
			return
		}
		c.Next()
	}
}
