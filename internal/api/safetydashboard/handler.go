// Package safetydashboard handles the 集团安全风险看板 read endpoint.
//
// One route, one handler — Overview returns the entire dashboard payload for
// the caller's scoped company subtree. The shape is documented in the SPA's
// types.ts; service layer builds the same struct keyed for JSON.
package safetydashboard

import (
	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/api/apiutil"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	svc *service.SafetyDashboardService
}

func NewHandler(svc *service.SafetyDashboardService) *Handler {
	return &Handler{svc: svc}
}

// Overview — GET /api/v1/safety-dashboard/overview
//
// Requires X-Scope-Company-Id (resolved by the ResolveScope middleware) so the
// caller's view of "their company subtree" matches the rest of the 集团视图
// pages. Permission gate is intentionally absent — every field below is
// already drawn from a permission-gated table, so an unprivileged caller just
// sees zeroes for what they can't access.
func (h *Handler) Overview(c *gin.Context) {
	companyID := apiutil.ScopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	out, err := h.svc.Overview(c.Request.Context(), companyID)
	if err != nil {
		response.Internal(c, "数据汇总失败")
		return
	}
	response.Success(c, out)
}
