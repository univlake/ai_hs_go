package api

import (
	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
)

// mountDashboardStubs registers placeholder endpoints used by the Dashboard
// page that belong to feature modules not yet ported from Laravel.
//
// They return empty success payloads so the SPA renders cleanly. Each stub is
// listed in the migration manifest (see plan §10) — replace the corresponding
// stub with a real handler when its module gets ported.
func mountDashboardStubs(g *gin.RouterGroup) {
	emptyArray := func(c *gin.Context) { response.Success(c, []any{}) }

	for _, p := range []string{
		// Expert inspection notifications (manifest §10.2 — controllers ExpertInspectionsController, ExpertCompany)
		"/expert-inspection-notifications/company",
		"/expert-inspection-notifications/expert",
		// Meeting / Office / NewNotify pending notifications (manifest §10.2 H. Notifications)
		"/meeting-notifications/pending",
		"/office-notify-notifications/pending",
		"/new-notify-notifications/pending",
		// Information drafts (manifest §10.2 H. Notifications)
		"/information-drafts/dashboard/pending",
		"/information-drafts/dashboard/readings",
	} {
		g.GET(p, emptyArray)
	}

	// Confirm / mark-read POST endpoints — accept and 200 silently so the UI
	// doesn't strand. They become real handlers once their module ports.
	confirmStub := func(c *gin.Context) { response.Success(c, gin.H{"message": "ok"}) }
	g.POST("/expert-inspection-companies/:id/confirm", confirmStub)
	g.POST("/expert-inspection-users/:id/confirm", confirmStub)
	g.POST("/meeting-notices/:id/confirm", confirmStub)
	g.POST("/new-notify-scopes/:id/read", confirmStub)
}

// IsStubPath returns whether a path was served by the dashboard stubs.
// (Reserved for instrumentation / metrics; kept inline so future code can
// distinguish placeholder hits from real ones.)
func IsStubPath(path string) bool {
	switch path {
	case "/api/v1/expert-inspection-notifications/company",
		"/api/v1/expert-inspection-notifications/expert",
		"/api/v1/meeting-notifications/pending",
		"/api/v1/office-notify-notifications/pending",
		"/api/v1/new-notify-notifications/pending",
		"/api/v1/information-drafts/dashboard/pending",
		"/api/v1/information-drafts/dashboard/readings":
		return true
	}
	return false
}

// suppressStubsActivityLog is a tiny middleware stub-list checker — kept
// alongside the stubs themselves so that activity-log noise filtering and
// stub registration evolve together.
func suppressStubsActivityLog(_ *gin.Context) bool {
	// Kept as a hook for the activity_log middleware to consult once we wire
	// stub-aware skip logic. Currently the middleware ignores GET requests
	// already, so noise is bounded.
	_ = middleware.CtxKeyUserID
	return false
}
