package tenant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

// Handler exposes tenant-level (一级企业) operations available only to the
// platform super admin.
type Handler struct {
	svc *service.CompanyService
}

func NewHandler(svc *service.CompanyService) *Handler {
	return &Handler{svc: svc}
}

// List: GET /api/v1/tenants
// Returns each forest root.
func (h *Handler) List(c *gin.Context) {
	roots, err := h.svc.Repo().Forest(c.Request.Context())
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, roots)
}

type createReq struct {
	Title             string  `json:"title" binding:"required,max=255"`
	EnterpriseTypeID  *int64  `json:"enterprise_type_id"`
	OperationalTypeID int64   `json:"operational_type_id"`
	AreaID            *int64  `json:"area_id"`
	Address           *string `json:"address"`
	CreditCode        *string `json:"credit_code"`
	BusinessScope     *string `json:"business_scope"`
	IsProduce         int16   `json:"is_produce"`
	IsBloc            int16   `json:"is_bloc"`
	InitialPassword   string  `json:"initial_password"`
}

// Create: POST /api/v1/tenants
func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	in := service.CreateCompanyInput{
		Title:             req.Title,
		EnterpriseTypeID:  req.EnterpriseTypeID,
		OperationalTypeID: req.OperationalTypeID,
		AreaID:            req.AreaID,
		Address:           req.Address,
		CreditCode:        req.CreditCode,
		BusinessScope:     req.BusinessScope,
		IsProduce:         req.IsProduce,
		IsBloc:            req.IsBloc,
		CreatedBy:         middleware.CurrentUserID(c),
	}
	tenant, admin, err := h.svc.CreateRoot(c.Request.Context(), in, req.InitialPassword)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, gin.H{
		"tenant": tenant,
		"admin": gin.H{
			"id":       admin.ID,
			"username": admin.Username,
			"must_change_credentials": admin.MustChangeCredentials,
		},
	})
}

// ToggleStatus: PATCH /api/v1/tenants/:id/toggle-status
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.ToggleStatus(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
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
