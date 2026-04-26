package role

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)


type Handler struct {
	rbac *service.RBACService
}

func NewHandler(rbac *service.RBACService) *Handler {
	return &Handler{rbac: rbac}
}

func (h *Handler) List(c *gin.Context) {
	roles, err := h.rbac.ListRoles(c.Request.Context())
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, roles)
}

// Get: GET /api/v1/roles/:id — single role + its permissions.
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.rbac.GetRole(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

type createReq struct {
	Name        string   `json:"name" binding:"required,max=100"`
	Permissions []string `json:"permissions"`
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	role, err := h.rbac.CreateRole(c.Request.Context(), req.Name, req.Permissions)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, role)
}

// Update: PUT /api/v1/roles/:id — rename + (optionally) sync permissions.
//
// Permissions uses a pointer slice so the caller can distinguish "leave
// alone" (omit field) from "clear all" (send `[]`). nil = unchanged.
type updateReq struct {
	Name        *string   `json:"name"`
	Permissions *[]string `json:"permissions"`
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
	in := service.UpdateRoleInput{Name: req.Name}
	if req.Permissions != nil {
		// Treat nil-inside-pointer (unlikely but defensive) as empty slice
		// so the service consistently sees a non-nil slice meaning "sync".
		if *req.Permissions == nil {
			in.Permissions = []string{}
		} else {
			in.Permissions = *req.Permissions
		}
	}
	row, err := h.rbac.UpdateRole(c.Request.Context(), id, in)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.rbac.DeleteRole(c.Request.Context(), id); err != nil {
		writeAppError(c, err)
		return
	}
	response.NoContent(c)
}

type assignPermsReq struct {
	Permissions []string `json:"permissions" binding:"required"`
}

func (h *Handler) AssignPermissions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req assignPermsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.rbac.AssignPermissionsToRole(c.Request.Context(), id, req.Permissions); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "权限已更新"})
}

func writeAppError(c *gin.Context, err error) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeConflict:
			response.Failed(c, http.StatusConflict, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, err.Error())
}
