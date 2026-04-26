package companygroup

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	svc *service.CompanyGroupService
}

func NewHandler(svc *service.CompanyGroupService) *Handler {
	return &Handler{svc: svc}
}

// List: GET /api/v1/company-groups?company_id=X
func (h *Handler) List(c *gin.Context) {
	cidRaw := c.Query("company_id")
	if cidRaw == "" {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	cid, err := strconv.ParseInt(cidRaw, 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 company_id")
		return
	}
	rows, err := h.svc.List(c.Request.Context(), cid)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, rows)
}

type createReq struct {
	Name      string `json:"name" binding:"required,max=255"`
	CompanyID int64  `json:"company_id" binding:"required"`
}

// Create: POST /api/v1/company-groups
func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	g, err := h.svc.Create(c.Request.Context(), service.CreateCompanyGroupInput{
		Name:      req.Name,
		CompanyID: req.CompanyID,
		UserID:    middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, g)
}

type updateReq struct {
	Name string `json:"name" binding:"required,max=255"`
}

// Update: PUT /api/v1/company-groups/:id
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
	g, err := h.svc.Update(c.Request.Context(), id, req.Name, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, g)
}

// Delete: DELETE /api/v1/company-groups/:id
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

type syncKidsReq struct {
	// Frontend uses `company_ids` to mirror Laravel.
	CompanyIDs []int64 `json:"company_ids"`
}

// SyncKids: POST /api/v1/company-groups/:id/sync-kids
func (h *Handler) SyncKids(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req syncKidsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.svc.SyncKids(c.Request.Context(), id, req.CompanyIDs); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"message": "已更新成员"})
}

func writeAppError(c *gin.Context, err error) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeForbidden:
			response.Forbidden(c, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, err.Error())
}
