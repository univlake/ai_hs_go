package platform

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
)

// Handler exposes platform-only operations available to is_platform_admin
// users. Today: impersonation entry/exit. Other cross-tenant tools (audit
// log, dictionaries) will land here as separate methods later.
type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// Impersonation strategy (path A from the plan):
//
//   - The platform admin's JWT continues to identify them. Authorisation on
//     business endpoints already short-circuits for is_platform_admin = true,
//     so the only behavioural switch is which company id their requests
//     scope to (X-Scope-Company-Id header).
//   - This endpoint validates the target company, writes an audit row, and
//     returns the company payload the frontend uses to update ScopeContext.
//     A platform admin may impersonate any company in the forest — root or
//     sub-company — so they can reproduce a specific user's view at any depth.
//   - Exiting impersonation writes a paired audit row.
//
// We deliberately do NOT mint a new JWT here — token swapping (path B) can
// be layered on later without breaking this contract.

type enterReq struct {
	TenantID int64 `json:"tenant_id" binding:"required"`
}

type enterResp struct {
	Tenant model.Company `json:"tenant"`
}

// Enter: POST /api/v1/platform/impersonate
func (h *Handler) Enter(c *gin.Context) {
	var req enterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	var tenant model.Company
	if err := h.db.WithContext(c.Request.Context()).First(&tenant, req.TenantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "目标企业不存在")
			return
		}
		response.Internal(c, "查询企业失败")
		return
	}
	if tenant.Status != 1 {
		response.Failed(c, http.StatusUnprocessableEntity, "目标企业未启用，无法化身")
		return
	}

	h.writeAudit(c, "platform.impersonate.enter", &tenant, nil)
	response.Success(c, enterResp{Tenant: tenant})
}

// Overview: GET /api/v1/platform/overview — aggregated cross-tenant counts
// for the platform-admin home screen.
type overviewResp struct {
	TenantsTotal     int64    `json:"tenants_total"`
	TenantsEnabled   int64    `json:"tenants_enabled"`
	TenantsDisabled  int64    `json:"tenants_disabled"`
	CompaniesTotal   int64    `json:"companies_total"`
	UsersTotal       int64    `json:"users_total"`
	UsersPlatform    int64    `json:"users_platform"`
	RolesTotal       int64    `json:"roles_total"`
	PermissionsTotal int64    `json:"permissions_total"`
	RecentEvents     []string `json:"recent_events"`
}

func (h *Handler) Overview(c *gin.Context) {
	ctx := c.Request.Context()
	out := overviewResp{}

	h.db.WithContext(ctx).Table("companies").Where("parent_id IS NULL AND deleted_at IS NULL").Count(&out.TenantsTotal)
	h.db.WithContext(ctx).Table("companies").Where("parent_id IS NULL AND status = 1 AND deleted_at IS NULL").Count(&out.TenantsEnabled)
	h.db.WithContext(ctx).Table("companies").Where("parent_id IS NULL AND status = 2 AND deleted_at IS NULL").Count(&out.TenantsDisabled)
	h.db.WithContext(ctx).Table("companies").Where("deleted_at IS NULL").Count(&out.CompaniesTotal)
	h.db.WithContext(ctx).Table("users").Where("deleted_at IS NULL").Count(&out.UsersTotal)
	h.db.WithContext(ctx).Table("users").Where("is_platform_admin = TRUE AND deleted_at IS NULL").Count(&out.UsersPlatform)
	h.db.WithContext(ctx).Table("roles").Count(&out.RolesTotal)
	h.db.WithContext(ctx).Table("permissions").Count(&out.PermissionsTotal)

	type row struct{ Description string }
	var rows []row
	h.db.WithContext(ctx).
		Table("activity_log").
		Where("log_name = ?", "platform").
		Order("id DESC").
		Limit(5).
		Find(&rows)
	for _, r := range rows {
		out.RecentEvents = append(out.RecentEvents, r.Description)
	}
	response.Success(c, out)
}

// Exit: POST /api/v1/platform/impersonate/exit
type exitReq struct {
	TenantID int64 `json:"tenant_id"`
}

func (h *Handler) Exit(c *gin.Context) {
	var req exitReq
	_ = c.ShouldBindJSON(&req) // body is optional — frontend may pass last tenant id for the audit row

	var tenant *model.Company
	if req.TenantID > 0 {
		var t model.Company
		if err := h.db.WithContext(c.Request.Context()).First(&t, req.TenantID).Error; err == nil {
			tenant = &t
		}
	}
	h.writeAudit(c, "platform.impersonate.exit", tenant, nil)
	response.Success(c, gin.H{"message": "已退出化身"})
}

// writeAudit logs a platform.* event into activity_log. We don't fail the
// request if the audit insert errors — best effort.
func (h *Handler) writeAudit(c *gin.Context, event string, subject *model.Company, extra map[string]any) {
	uid := middleware.CurrentUserID(c)
	desc := event
	if subject != nil {
		desc = event + " → " + strconv.FormatInt(subject.ID, 10) + " (" + subject.Title + ")"
	}
	props, _ := json.Marshal(extra)
	if extra == nil {
		props = []byte(`{}`)
	}
	row := &model.ActivityLog{
		LogName:          ptrStr("platform"),
		Description:      desc,
		Event:            ptrStr(event),
		CauserType:       ptrStr("App\\Models\\User"),
		CauserID:         &uid,
		Properties:       props,
		AttributeChanges: []byte(`{}`),
	}
	if subject != nil {
		row.SubjectType = ptrStr("App\\Models\\Company")
		row.SubjectID = &subject.ID
	}
	if err := h.db.WithContext(c.Request.Context()).Create(row).Error; err != nil {
		// Audit failure should not block the operation; log via gin error.
		_ = c.Error(err)
	}
}

func ptrStr(s string) *string { return &s }
