package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	users *service.UserService
	rbac  *service.RBACService
}

func NewHandler(users *service.UserService, rbac *service.RBACService) *Handler {
	return &Handler{users: users, rbac: rbac}
}

// List: GET /api/v1/users
//
// Query params:
//
//	page, per_page             pagination
//	company_id                 explicit override (defaults to scope)
//	username / keyword         substring match on username/nickname/phone
//	status                     filter by status (1 enabled / 2 disabled)
//	role                       filter by role name (e.g. "tenant_admin")
func (h *Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	// Default company filter to the active scope unless an explicit
	// company_id query param is supplied. Scope resolution already enforces
	// the cross-tenant guard for non-platform-admins.
	var companyID *int64
	if v := c.Query("company_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			companyID = &id
		}
	} else if cid, ok := middleware.ScopeCompanyID(c); ok {
		companyID = &cid
	}

	var status *int16
	if v := c.Query("status"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			st := int16(n)
			status = &st
		}
	}

	keyword := c.Query("username")
	if keyword == "" {
		keyword = c.Query("keyword")
	}

	items, total, err := h.users.List(c.Request.Context(), service.UserListParams{
		Page: page, PerPage: per,
		CompanyID: companyID,
		Keyword:   keyword,
		Status:    status,
		Role:      c.Query("role"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}

	// Enrich each row with its role names so the UI can render badges
	// without a second round-trip.
	type listRow struct {
		*model.User
		Roles []string `json:"roles"`
	}
	enriched := make([]listRow, len(items))
	for i := range items {
		roles, _, _ := h.rbac.GetUserRolesAndPermissions(c.Request.Context(), items[i].ID)
		enriched[i] = listRow{User: &items[i], Roles: roles}
	}
	response.Page(c, enriched, page, per, total)
}

// Get: GET /api/v1/users/:id
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	user, err := h.users.Get(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, "用户不存在")
		return
	}
	roles, _, _ := h.rbac.GetUserRolesAndPermissions(c.Request.Context(), user.ID)
	response.Success(c, gin.H{"user": user, "roles": roles})
}

type createReq struct {
	Username       string  `json:"username" binding:"required"`
	Password       string  `json:"password" binding:"required,min=6"`
	Nickname       string  `json:"nickname" binding:"required,max=100"`
	CompanyID      *int64  `json:"company_id"`
	Phone          *string `json:"phone"`
	Email          *string `json:"email" binding:"omitempty,email"`
	Status         int16   `json:"status"`
	IsExpert       int16   `json:"is_expert"`
	HeaderImg      *string `json:"header_img"`
	DepartmentName *string `json:"department_name"`
	Role           string  `json:"role"` // single role name, optional
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	// Default company to the active scope so a tenant admin always creates
	// users inside their own tenant.
	if req.CompanyID == nil {
		if cid, ok := middleware.ScopeCompanyID(c); ok && cid > 0 {
			id := cid
			req.CompanyID = &id
		}
	}
	if req.CompanyID == nil {
		response.Failed(c, http.StatusForbidden, "当前账号未绑定企业，无法创建用户")
		return
	}
	user, err := h.users.Create(c.Request.Context(), service.CreateUserInput{
		Username:       req.Username,
		Password:       req.Password,
		Nickname:       req.Nickname,
		CompanyID:      req.CompanyID,
		Phone:          req.Phone,
		Email:          req.Email,
		Status:         req.Status,
		IsExpert:       req.IsExpert,
		HeaderImg:      req.HeaderImg,
		DepartmentName: req.DepartmentName,
		CreatedBy:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	if role := req.Role; role != "" {
		if err := h.rbac.AssignSingleRoleToUser(c.Request.Context(), user.ID, role); err != nil {
			writeAppError(c, err)
			return
		}
	}
	response.Created(c, user)
}

type updateReq struct {
	Nickname       *string `json:"nickname"`
	CompanyID      *int64  `json:"company_id"`
	Phone          *string `json:"phone"`
	Email          *string `json:"email" binding:"omitempty,email"`
	Status         *int16  `json:"status"`
	IsExpert       *int16  `json:"is_expert"`
	HeaderImg      *string `json:"header_img"`
	DepartmentName *string `json:"department_name"`
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	user, err := h.users.Update(c.Request.Context(), id, service.UpdateUserInput{
		Nickname:       req.Nickname,
		CompanyID:      req.CompanyID,
		Phone:          req.Phone,
		Email:          req.Email,
		Status:         req.Status,
		IsExpert:       req.IsExpert,
		HeaderImg:      req.HeaderImg,
		DepartmentName: req.DepartmentName,
		UpdatedBy:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, user)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.users.Delete(c.Request.Context(), id); err != nil {
		writeAppError(c, err)
		return
	}
	response.NoContent(c)
}

type assignRolesReq struct {
	Roles []string `json:"roles" binding:"required"`
}

// AssignRoles: POST /api/v1/users/:id/roles — sync the full role set.
func (h *Handler) AssignRoles(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req assignRolesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.rbac.AssignRolesToUser(c.Request.Context(), id, req.Roles); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "分配成功"})
}

// AssignRole: PUT /api/v1/users/:id/role — set a single role (Laravel parity).
type assignRoleReq struct {
	Role string `json:"role" binding:"required"`
}

func (h *Handler) AssignRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req assignRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.rbac.AssignSingleRoleToUser(c.Request.Context(), id, req.Role); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "分配成功"})
}

// ResetPassword: PUT /api/v1/users/:id/password — admin resets a child account.
type adminResetPasswordReq struct {
	Password string `json:"password" binding:"required,min=6"`
}

func (h *Handler) ResetPassword(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req adminResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.users.AdminResetPassword(
		c.Request.Context(),
		id,
		req.Password,
		middleware.CurrentUserID(c),
	); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "密码已重置"})
}

func writeAppError(c *gin.Context, err error) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable, errs.CodeUsernameInvalid:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeForbidden:
			response.Forbidden(c, ae.Message)
		case errs.CodeConflict:
			response.Failed(c, http.StatusConflict, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, err.Error())
}
