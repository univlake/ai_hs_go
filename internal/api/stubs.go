package api

import (
	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/middleware"
)

// mountDashboardStubs registers placeholder endpoints used by the Dashboard
// page that belong to feature modules not yet ported from Laravel.
//
// They return empty success payloads so the SPA renders cleanly. Each stub is
// listed in the migration manifest (see plan §10) — replace the corresponding
// stub with a real handler when its module gets ported.
//
// 已下线的 stubs（在通知模块迁移后由真实 handler 接管）:
//
//	/meeting-notifications/pending     → notificationH.PendingMeetingNotices
//	/office-notify-notifications/pending → notificationH.PendingOfficeNotifyScopes
//	/new-notify-notifications/pending  → notificationH.PendingNewNotifyScopes
//	/information-drafts/dashboard/*    → notificationH.DashboardInformationDraft*
//	/meeting-notices/:id/confirm       → notificationH.ConfirmMeetingNotice
//	/new-notify-scopes/:id/read        → notificationH.MarkNewNotifyRead
//
// 已下线的 stubs（在专家检查模块迁移后由真实 handler 接管）:
//
//	/expert-inspection-notifications/company → expertH.CompanyNotifications
//	/expert-inspection-notifications/expert  → expertH.ExpertNotifications
//	/expert-inspection-companies/:id/confirm → expertH.ConfirmCompany
//	/expert-inspection-users/:id/confirm     → expertH.ConfirmUser
//
// 已下线的 stubs（在隐患信息模块迁移后由真实 handler 接管）:
//
//	/hidden-danger-infos (GET/POST/PUT/DELETE) → hiddenDangerH.{List,Create,Update,Delete}
//	/hidden-danger-infos/:id/{rectify,confirm-rectification,archive} → hiddenDangerH.*
//
// 已下线的 stubs（在整改通知单 / 隐患分类统计 Excel 生成接入后由真实 handler 接管）:
//
//	/expert-inspection-companies/:id/rectification-sheet (POST/GET) → expertH.{Generate,Download}RectificationSheet
//	/hazard-statistics → expertH.ExportHazardStatistics
//
// 当前没有剩余的 stubs；保留函数以便后续模块按相同模式追加。
func mountDashboardStubs(_ *gin.RouterGroup) {}

// IsStubPath returns whether a path was served by the dashboard stubs.
// (Reserved for instrumentation / metrics; kept inline so future code can
// distinguish placeholder hits from real ones.)
func IsStubPath(_ string) bool { return false }

// suppressStubsActivityLog is a tiny middleware stub-list checker — kept
// alongside the stubs themselves so that activity-log noise filtering and
// stub registration evolve together.
func suppressStubsActivityLog(_ *gin.Context) bool {
	_ = middleware.CtxKeyUserID
	return false
}
