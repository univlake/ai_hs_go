package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	svc  *service.AuthService
	rbac *service.RBACService
}

func NewHandler(svc *service.AuthService, rbac *service.RBACService) *Handler {
	return &Handler{svc: svc, rbac: rbac}
}

// PublicKey: GET /api/v1/public-key
func (h *Handler) PublicKey(c *gin.Context) {
	response.Success(c, gin.H{"public_key": h.svc.RSAPublicKey()})
}

// Login: POST /api/v1/login
type loginRequest struct {
	Username string `json:"username" binding:"required,min=1,max=64"`
	Password string `json:"password" binding:"required,min=1"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	out, err := h.svc.Login(c.Request.Context(), service.LoginInput{
		Username:          req.Username,
		EncryptedPassword: req.Password,
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{
		"access_token": out.AccessToken,
		"token_type":   out.TokenType,
		"expires_at":   out.ExpiresAt.Format("2006-01-02 15:04:05"),
	})
}

// Logout: POST /api/v1/logout
// JWT is stateless — without a Redis blacklist this is a client-side discard.
// Phase 4+ may introduce a denylist.
func (h *Handler) Logout(c *gin.Context) {
	response.Success(c, gin.H{"message": "退出成功"})
}

// UserInfo: GET /api/v1/user-info
func (h *Handler) UserInfo(c *gin.Context) {
	uid := middleware.CurrentUserID(c)
	if uid == 0 {
		response.Unauthorized(c, "")
		return
	}
	user, err := h.svc.GetUser(c.Request.Context(), uid)
	if err != nil {
		response.Internal(c, "查询用户失败")
		return
	}
	c.Set(middleware.CtxKeyMustChangeCreds, user.MustChangeCredentials)
	roles, perms, err := h.rbac.GetUserRolesAndPermissions(c.Request.Context(), user.ID)
	if err != nil {
		response.Internal(c, "查询权限失败")
		return
	}
	// 前端 authApi.userInfo 直接对 roles / permissions 调用 .map()，nil 切片会序列化成
	// JSON null 把前端打挂。统一兜底成 [] 让契约对前端可枚举类型保持一致。
	if roles == nil {
		roles = []string{}
	}
	if perms == nil {
		perms = []string{}
	}

	// Resolve the user's company so the frontend ScopeContext can bootstrap
	// without a second round-trip. Platform admins legitimately have
	// company_id = NULL — they get company = null in the response.
	var company any
	if user.CompanyID != nil && h.svc != nil {
		if co, err := h.svc.GetCompanyByID(c.Request.Context(), *user.CompanyID); err == nil && co != nil {
			company = gin.H{
				"id":         co.ID,
				"title":      co.Title,
				"is_bloc":    co.IsBloc,
				"is_produce": co.IsProduce,
				"parent_id":  co.ParentID,
				"root_id":    co.RootID,
				"status":     co.Status,
			}
		}
	}

	response.Success(c, gin.H{
		"id":                      user.ID,
		"username":                user.Username,
		"nickname":                user.Nickname,
		"company_id":              user.CompanyID,
		"company":                 company,
		"phone":                   user.Phone,
		"email":                   user.Email,
		"status":                  user.Status,
		"is_platform_admin":       user.IsPlatformAdmin,
		"must_change_credentials": user.MustChangeCredentials,
		"roles":                   roles,
		"permissions":             perms,
	})
}

// UpdateUserInfo: PUT /api/v1/user-info
type updateProfileRequest struct {
	Nickname  string  `json:"nickname" binding:"required,max=100"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email" binding:"omitempty,email"`
	HeaderImg *string `json:"header_img"`
	Username  *string `json:"username"`
}

func (h *Handler) UpdateUserInfo(c *gin.Context) {
	uid := middleware.CurrentUserID(c)
	if uid == 0 {
		response.Unauthorized(c, "")
		return
	}
	var req updateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.svc.UpdateProfile(c.Request.Context(), uid, service.ProfileUpdateInput{
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Email:     req.Email,
		HeaderImg: req.HeaderImg,
		Username:  req.Username,
	}); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "个人资料更新成功"})
}

// CheckUsernameAvailability: GET /api/v1/username-availability?username=xxx
func (h *Handler) CheckUsernameAvailability(c *gin.Context) {
	username := c.Query("username")
	uid := middleware.CurrentUserID(c)
	check := h.svc.CheckUsername(c.Request.Context(), username, uid)
	response.Success(c, check)
}

// FirstLoginSetup: POST /api/v1/first-login-setup
type firstLoginRequest struct {
	NewUsername          string `json:"new_username" binding:"required"`
	NewPassword          string `json:"new_password" binding:"required,min=6"`
	NewPasswordConfirmed string `json:"new_password_confirmation" binding:"required"`
}

func (h *Handler) FirstLoginSetup(c *gin.Context) {
	uid := middleware.CurrentUserID(c)
	if uid == 0 {
		response.Unauthorized(c, "")
		return
	}
	var req firstLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.NewPassword != req.NewPasswordConfirmed {
		response.Failed(c, http.StatusUnprocessableEntity, "两次输入的新密码不一致")
		return
	}
	if err := h.svc.FirstLoginSetup(c.Request.Context(), uid, service.FirstLoginInput{
		NewUsername: req.NewUsername,
		NewPassword: req.NewPassword,
	}); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "设置成功，请重新登录"})
}

// ResetPassword: POST /api/v1/reset-password
type resetPasswordRequest struct {
	OldPassword          string `json:"old_password" binding:"required"`
	NewPassword          string `json:"new_password" binding:"required,min=6"`
	NewPasswordConfirmed string `json:"new_password_confirmation" binding:"required"`
}

func (h *Handler) ResetPassword(c *gin.Context) {
	uid := middleware.CurrentUserID(c)
	if uid == 0 {
		response.Unauthorized(c, "")
		return
	}
	var req resetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.NewPassword != req.NewPasswordConfirmed {
		response.Failed(c, http.StatusUnprocessableEntity, "两次输入的新密码不一致")
		return
	}
	if err := h.svc.ResetPassword(c.Request.Context(), uid, service.ResetPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "密码修改成功"})
}

func writeAppError(c *gin.Context, err error) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeUnauthorized:
			response.Unauthorized(c, ae.Message)
		case errs.CodeForbidden:
			if ae.Reason != "" {
				response.FailedWithReason(c, http.StatusForbidden, ae.Message, ae.Reason)
			} else {
				response.Forbidden(c, ae.Message)
			}
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable, errs.CodeUsernameInvalid, errs.CodeUsernameReserved, errs.CodeUsernameTaken:
			if ae.Reason != "" {
				response.FailedWithReason(c, http.StatusUnprocessableEntity, ae.Message, ae.Reason)
			} else {
				response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
			}
		case errs.CodeConflict:
			response.Failed(c, http.StatusConflict, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, err.Error())
}
