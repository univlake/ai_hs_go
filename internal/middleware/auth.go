package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/auth"
	"ai_hs_go/internal/response"
)

// RequireAuth parses Bearer token, populates request-scoped auth fields.
func RequireAuth(jwtSvc *auth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("Authorization")
		if raw == "" {
			response.Unauthorized(c, "未提供登录凭据")
			return
		}
		parts := strings.SplitN(raw, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Unauthorized(c, "登录凭据格式不正确")
			return
		}
		claims, err := jwtSvc.Parse(parts[1])
		if err != nil {
			response.Unauthorized(c, "登录已失效，请重新登录")
			return
		}

		c.Set(CtxKeyUserID, claims.UserID)
		c.Set(CtxKeyUsername, claims.Username)
		if claims.CompanyID != nil {
			c.Set(CtxKeyCompanyID, *claims.CompanyID)
		}
		c.Set(CtxKeyIsPlatformAdmin, claims.IsPlatformAdmin)
		c.Set(CtxKeyMustChangeCreds, claims.MustChangeCredentials)
		c.Next()
	}
}

// RequireFreshCredentials short-circuits any business endpoint when the user
// still has must_change_credentials=true. The frontend uses /api/v1/user-info
// to detect the flag and redirect to /first-login-setup; this is a safety net
// for direct API calls.
func RequireFreshCredentials() gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, ok := c.Get(CtxKeyMustChangeCreds); ok {
			if must, _ := v.(bool); must {
				response.FailedWithReason(c, 403, "请先完成首次登录设置", "must_change_credentials")
				return
			}
		}
		c.Next()
	}
}

// CurrentUserID is a small helper for handlers that already passed RequireAuth.
func CurrentUserID(c *gin.Context) int64 {
	if v, ok := c.Get(CtxKeyUserID); ok {
		if id, ok := v.(int64); ok {
			return id
		}
	}
	return 0
}

func IsPlatformAdmin(c *gin.Context) bool {
	if v, ok := c.Get(CtxKeyIsPlatformAdmin); ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}
