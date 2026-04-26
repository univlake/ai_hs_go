package api

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	authapi "ai_hs_go/internal/api/auth"
	basedataapi "ai_hs_go/internal/api/basedata"
	companyapi "ai_hs_go/internal/api/company"
	companygroupapi "ai_hs_go/internal/api/companygroup"
	permissionapi "ai_hs_go/internal/api/permission"
	platformapi "ai_hs_go/internal/api/platform"
	roleapi "ai_hs_go/internal/api/role"
	safetyapi "ai_hs_go/internal/api/safetyledger"
	tenantapi "ai_hs_go/internal/api/tenant"
	uploadapi "ai_hs_go/internal/api/upload"
	userapi "ai_hs_go/internal/api/user"
	"ai_hs_go/internal/auth"
	"ai_hs_go/internal/config"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/repository"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/storage"
)

type Deps struct {
	Config   *config.Config
	Logger   *zap.Logger
	DB       *gorm.DB
	RSA      *auth.RSAService
	JWT      *auth.JWTService
	Enforcer *casbin.Enforcer
	OSS      *storage.Client // optional; nil disables OSS endpoints
}

type Router struct {
	deps Deps
}

func New(deps Deps) *Router {
	return &Router{deps: deps}
}

func (r *Router) Engine() *gin.Engine {
	if r.deps.Config.App.Env == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(middleware.Recovery(r.deps.Logger))
	engine.Use(middleware.Logger(r.deps.Logger))
	engine.Use(middleware.CORS(r.deps.Config.CORS.AllowedOrigins))

	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.mountSPA(engine)

	companyRepo := repository.NewCompanyRepo(r.deps.DB)

	authSvc := service.NewAuthService(r.deps.DB, r.deps.JWT, r.deps.RSA)
	rbacSvc := service.NewRBACService(r.deps.DB, r.deps.Enforcer)
	userSvc := service.NewUserService(r.deps.DB)
	companySvc := service.NewCompanyService(r.deps.DB, companyRepo)
	companyGroupSvc := service.NewCompanyGroupService(r.deps.DB)
	dictSvc := service.NewDictService(r.deps.DB)
	checklistSvc := service.NewChecklistLibraryService(r.deps.DB)
	regLibSvc := service.NewRegulationLibraryTypeService(r.deps.DB)

	authH := authapi.NewHandler(authSvc, rbacSvc)
	userH := userapi.NewHandler(userSvc, rbacSvc)
	roleH := roleapi.NewHandler(rbacSvc)
	permissionH := permissionapi.NewHandler(rbacSvc)
	companyH := companyapi.NewHandler(companySvc)
	companyGroupH := companygroupapi.NewHandler(companyGroupSvc)
	tenantH := tenantapi.NewHandler(companySvc)
	ledgerSvc := service.NewSafetyLedgerService(r.deps.DB)
	ledgerH := safetyapi.NewHandler(ledgerSvc)
	uploadH := uploadapi.NewHandler(r.deps.OSS)
	basedataH := basedataapi.NewHandler(dictSvc, checklistSvc, regLibSvc)

	v1 := engine.Group("/api/v1")

	// Public, NOT rate-limited — login pages may fan out a handful of these
	// during boot (key fetch + scope context + permission probe). Behind a
	// dev proxy every request looks like the same client IP, so a tight
	// guest quota would lock everyone out.
	v1.GET("/public-key", authH.PublicKey)
	v1.GET("/companies-public", companyH.PublicList)

	// Everything else inherits the standard quota.
	v1.Use(middleware.APIRateLimit(r.deps.Config.RateLimit.APIPerMinute, r.deps.Config.RateLimit.GuestPerMinute))

	// Login throttled separately
	loginGroup := v1.Group("")
	loginGroup.Use(middleware.LoginRateLimit(r.deps.Config.RateLimit.LoginPerMinute))
	loginGroup.POST("/login", authH.Login)

	// Authenticated — account-self routes always reachable
	authed := v1.Group("")
	authed.Use(middleware.RequireAuth(r.deps.JWT))

	authed.POST("/logout", authH.Logout)
	authed.GET("/user-info", authH.UserInfo)
	authed.PUT("/user-info", authH.UpdateUserInfo)
	authed.GET("/username-availability", authH.CheckUsernameAvailability)
	authed.POST("/first-login-setup", authH.FirstLoginSetup)
	authed.POST("/reset-password", authH.ResetPassword)

	// Business routes — must have completed first-login-setup + scope resolution
	business := authed.Group("")
	business.Use(middleware.RequireFreshCredentials())
	business.Use(middleware.ResolveScope(companyRepo))
	business.Use(middleware.ActivityLog(r.deps.DB, r.deps.Logger))

	// Users
	business.GET("/users", middleware.RequirePermission(r.deps.Enforcer, "users.list"), userH.List)
	business.GET("/users/:id", middleware.RequirePermission(r.deps.Enforcer, "users.list"), userH.Get)
	business.POST("/users", middleware.RequirePermission(r.deps.Enforcer, "users.create"), userH.Create)
	business.PUT("/users/:id", middleware.RequirePermission(r.deps.Enforcer, "users.update"), userH.Update)
	business.DELETE("/users/:id", middleware.RequirePermission(r.deps.Enforcer, "users.delete"), userH.Delete)
	business.POST("/users/:id/roles", middleware.RequirePermission(r.deps.Enforcer, "roles.assign"), userH.AssignRoles)
	business.PUT("/users/:id/role", middleware.RequirePermission(r.deps.Enforcer, "roles.assign"), userH.AssignRole)
	business.PUT("/users/:id/password", middleware.RequirePermission(r.deps.Enforcer, "users.update"), userH.ResetPassword)

	// Roles
	business.GET("/roles", middleware.RequirePermission(r.deps.Enforcer, "roles.list"), roleH.List)
	business.GET("/roles/:id", middleware.RequirePermission(r.deps.Enforcer, "roles.list"), roleH.Get)
	business.POST("/roles", middleware.RequirePermission(r.deps.Enforcer, "roles.create"), roleH.Create)
	business.PUT("/roles/:id", middleware.RequirePermission(r.deps.Enforcer, "roles.update"), roleH.Update)
	business.DELETE("/roles/:id", middleware.RequirePermission(r.deps.Enforcer, "roles.delete"), roleH.Delete)
	business.POST("/roles/:id/permissions", middleware.RequirePermission(r.deps.Enforcer, "roles.update"), roleH.AssignPermissions)

	// Permissions
	business.GET("/permissions", middleware.RequirePermission(r.deps.Enforcer, "permissions.list"), permissionH.List)

	// Base data dictionaries — 9 type tables share one generic handler set
	// dispatched by URL slug. Reads are gated by base-data.view; mutations
	// by base-data.manage to mirror Laravel routes/api/base-data.php.
	dictSlugs := []string{
		"enterprise-types",
		"operational-types",
		"record-types",
		"process-types",
		"special-types",
		"space-work-types",
		"safety-management-types",
		"dangerous-craft-types",
		// regulation-library-types is registered separately below — its
		// table carries pid/all_name and needs the dedicated service that
		// keeps the tree consistent.
	}
	for _, slug := range dictSlugs {
		s := slug
		business.GET("/"+s, middleware.RequirePermission(r.deps.Enforcer, "base-data.view"), basedataH.DictList(s))
		business.POST("/"+s, middleware.RequirePermission(r.deps.Enforcer, "base-data.manage"), basedataH.DictCreate(s))
		business.PUT("/"+s+"/:id", middleware.RequirePermission(r.deps.Enforcer, "base-data.manage"), basedataH.DictUpdate(s))
		business.DELETE("/"+s+"/:id", middleware.RequirePermission(r.deps.Enforcer, "base-data.manage"), basedataH.DictDelete(s))
	}

	// Regulation library types — multi-level (pid + all_name).
	business.GET("/regulation-library-types", middleware.RequirePermission(r.deps.Enforcer, "base-data.view"), basedataH.RegLibTypeList)
	business.POST("/regulation-library-types", middleware.RequirePermission(r.deps.Enforcer, "base-data.manage"), basedataH.RegLibTypeCreate)
	business.PUT("/regulation-library-types/:id", middleware.RequirePermission(r.deps.Enforcer, "base-data.manage"), basedataH.RegLibTypeUpdate)
	business.DELETE("/regulation-library-types/:id", middleware.RequirePermission(r.deps.Enforcer, "base-data.manage"), basedataH.RegLibTypeDelete)

	// Checklist libraries — separate from the dictionary CRUD because each
	// row carries title + status + annex URLs + tenant scope.
	business.GET("/checklist-libraries", middleware.RequirePermission(r.deps.Enforcer, "checklist-libraries.list"), basedataH.ChecklistList)
	business.POST("/checklist-libraries", middleware.RequirePermission(r.deps.Enforcer, "checklist-libraries.create"), basedataH.ChecklistCreate)
	business.PUT("/checklist-libraries/:id", middleware.RequirePermission(r.deps.Enforcer, "checklist-libraries.update"), basedataH.ChecklistUpdate)
	business.DELETE("/checklist-libraries/:id", middleware.RequirePermission(r.deps.Enforcer, "checklist-libraries.delete"), basedataH.ChecklistDelete)

	// Companies
	business.GET("/companies", middleware.RequirePermission(r.deps.Enforcer, "companies.list"), companyH.List)
	business.POST("/companies", middleware.RequirePermission(r.deps.Enforcer, "companies.create"), companyH.Create)
	business.GET("/companies/:id", middleware.RequirePermission(r.deps.Enforcer, "companies.list"), companyH.Get)
	business.PUT("/companies/:id", middleware.RequirePermission(r.deps.Enforcer, "companies.update"), companyH.Update)
	business.DELETE("/companies/:id", middleware.RequirePermission(r.deps.Enforcer, "companies.delete"), companyH.Delete)
	business.PATCH("/companies/:id/toggle-status", middleware.RequirePermission(r.deps.Enforcer, "companies.update"), companyH.ToggleStatus)
	business.POST("/companies/:id/review", middleware.RequirePermission(r.deps.Enforcer, "companies.review"), companyH.Review)
	business.GET("/companies-review-list", middleware.RequirePermission(r.deps.Enforcer, "companies.review"), companyH.ReviewList)
	business.GET("/companies-tree", middleware.RequirePermission(r.deps.Enforcer, "companies.tree"), companyH.Tree)
	business.GET("/companies-forest", companyH.Forest) // platform-admin gated inside handler

	// Company groups (集团视图自定义分组)
	business.GET("/company-groups", middleware.RequirePermission(r.deps.Enforcer, "companies.list"), companyGroupH.List)
	business.POST("/company-groups", middleware.RequirePermission(r.deps.Enforcer, "companies.update"), companyGroupH.Create)
	business.PUT("/company-groups/:id", middleware.RequirePermission(r.deps.Enforcer, "companies.update"), companyGroupH.Update)
	business.DELETE("/company-groups/:id", middleware.RequirePermission(r.deps.Enforcer, "companies.update"), companyGroupH.Delete)
	business.POST("/company-groups/:id/sync-kids", middleware.RequirePermission(r.deps.Enforcer, "companies.update"), companyGroupH.SyncKids)

	// Tenants (platform-admin)
	tenants := business.Group("/tenants", middleware.RequirePlatformAdmin())
	tenants.GET("", tenantH.List)
	tenants.POST("", tenantH.Create)
	tenants.PATCH("/:id/toggle-status", tenantH.ToggleStatus)

	// Platform-admin only impersonation endpoints. The actual scope switch
	// is driven by the X-Scope-Company-Id header on subsequent requests; this
	// route validates the target + writes audit rows so platform-admin
	// activity is traceable.
	platformH := platformapi.NewHandler(r.deps.DB)
	platform := business.Group("/platform", middleware.RequirePlatformAdmin())
	platform.POST("/impersonate", platformH.Enter)
	platform.POST("/impersonate/exit", platformH.Exit)
	platform.GET("/overview", platformH.Overview)

	// Safety ledger
	business.GET("/special-equipment", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.list"), ledgerH.ListSpecialEquipment)
	business.POST("/special-equipment", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.create"), ledgerH.CreateSpecialEquipment)
	business.PUT("/special-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.update"), ledgerH.UpdateSpecialEquipment)
	business.DELETE("/special-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.delete"), ledgerH.DeleteSpecialEquipment)
	business.GET("/special-equipment/template", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.list"), ledgerH.TemplateSpecialEquipment)
	business.GET("/special-equipment/export", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.list"), ledgerH.ExportSpecialEquipment)
	business.POST("/special-equipment/import", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.create"), ledgerH.ImportSpecialEquipment)

	business.GET("/dangerous-equipment", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.list"), ledgerH.ListDangerousEquipment)
	business.POST("/dangerous-equipment", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.create"), ledgerH.CreateDangerousEquipment)
	business.PUT("/dangerous-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.update"), ledgerH.UpdateDangerousEquipment)
	business.DELETE("/dangerous-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.delete"), ledgerH.DeleteDangerousEquipment)
	business.GET("/dangerous-equipment/template", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.list"), ledgerH.TemplateDangerousEquipment)
	business.GET("/dangerous-equipment/export", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.list"), ledgerH.ExportDangerousEquipment)
	business.POST("/dangerous-equipment/import", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.create"), ledgerH.ImportDangerousEquipment)

	business.GET("/hazardous-chemicals", middleware.RequirePermission(r.deps.Enforcer, "hazardous-chemicals.list"), ledgerH.ListHazardousChemicals)
	business.POST("/hazardous-chemicals", middleware.RequirePermission(r.deps.Enforcer, "hazardous-chemicals.create"), ledgerH.CreateHazardousChemical)
	business.PUT("/hazardous-chemicals/:id", middleware.RequirePermission(r.deps.Enforcer, "hazardous-chemicals.update"), ledgerH.UpdateHazardousChemical)
	business.DELETE("/hazardous-chemicals/:id", middleware.RequirePermission(r.deps.Enforcer, "hazardous-chemicals.delete"), ledgerH.DeleteHazardousChemical)

	// Upload (OSS multipart)
	upload := business.Group("/upload-oss", middleware.RequirePermission(r.deps.Enforcer, "upload.oss"))
	upload.POST("/init", uploadH.Init)
	upload.POST("/sign-part", uploadH.SignPart)
	upload.POST("/complete", uploadH.Complete)
	upload.POST("/abort", uploadH.Abort)
	upload.GET("/signed", uploadH.SignedURL)

	// Dashboard placeholders for not-yet-ported modules. Empty success
	// payloads keep the SPA rendering until the real modules port (plan §10).
	mountDashboardStubs(business)

	// ping (smoke)
	business.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pong": true})
	})

	return engine
}
