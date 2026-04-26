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
	majorhazardapi "ai_hs_go/internal/api/majorhazard"
	permissionapi "ai_hs_go/internal/api/permission"
	platformapi "ai_hs_go/internal/api/platform"
	processapi "ai_hs_go/internal/api/process"
	roleapi "ai_hs_go/internal/api/role"
	safetyarchapi "ai_hs_go/internal/api/safetyarchive"
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
	archSvc := service.NewSafetyArchitectureService(r.deps.DB)
	filingSvc := service.NewFilingPermissionService(r.deps.DB)
	buildingSvc := service.NewBuildingService(r.deps.DB)
	investmentSvc := service.NewSecurityInvestmentService(r.deps.DB)
	personnelSvc := service.NewPersonnelService(r.deps.DB)
	certStatusSvc := service.NewCertificateStatusService(r.deps.DB)
	productionStdSvc := service.NewProductionStandardizationService(r.deps.DB)
	ohsSvc := service.NewOccupationalHealthSafetyService(r.deps.DB)
	fundingExecSvc := service.NewFundingExecutionService(r.deps.DB)
	safetyMgmtStaffSvc := service.NewSafetyManagementStaffService(r.deps.DB)
	specialEquipStaffSvc := service.NewSpecialEquipmentStaffService(r.deps.DB)
	elimCtrlStaffSvc := service.NewEliminationControlStaffService(r.deps.DB)

	authH := authapi.NewHandler(authSvc, rbacSvc)
	userH := userapi.NewHandler(userSvc, rbacSvc)
	roleH := roleapi.NewHandler(rbacSvc)
	permissionH := permissionapi.NewHandler(rbacSvc)
	companyH := companyapi.NewHandler(companySvc)
	companyGroupH := companygroupapi.NewHandler(companyGroupSvc)
	tenantH := tenantapi.NewHandler(companySvc)
	ledgerSvc := service.NewSafetyLedgerService(r.deps.DB)
	ledgerH := safetyapi.NewHandler(ledgerSvc)
	majorHazardSvc := service.NewMajorHazardService(r.deps.DB)
	majorHazardH := majorhazardapi.NewHandler(majorHazardSvc)
	processSvc := service.NewProcessService(r.deps.DB)
	processH := processapi.NewHandler(processSvc)
	uploadH := uploadapi.NewHandler(r.deps.OSS)
	basedataH := basedataapi.NewHandler(dictSvc, checklistSvc, regLibSvc)
	safetyArchH := safetyarchapi.NewHandler(safetyarchapi.HandlerDeps{
		Architecture:         archSvc,
		Filing:               filingSvc,
		Building:             buildingSvc,
		Investment:           investmentSvc,
		Personnel:            personnelSvc,
		CertStatus:           certStatusSvc,
		ProductionStd:        productionStdSvc,
		OHS:                  ohsSvc,
		FundingExec:          fundingExecSvc,
		SafetyMgmtStaff:      safetyMgmtStaffSvc,
		SpecialEquipStaff:    specialEquipStaffSvc,
		EliminationCtrlStaff: elimCtrlStaffSvc,
		OSS:                  r.deps.OSS,
	})

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
		"hazards",
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

	// 安全档案 module — 5 sub-pages backed by 5 services in safetyarchive.Handler.
	business.GET("/safety-architectures", middleware.RequirePermission(r.deps.Enforcer, "safety-architectures.list"), safetyArchH.ArchitectureGet)
	business.POST("/safety-architectures", middleware.RequirePermission(r.deps.Enforcer, "safety-architectures.create"), safetyArchH.ArchitectureCreate)
	business.PUT("/safety-architectures/:id", middleware.RequirePermission(r.deps.Enforcer, "safety-architectures.update"), safetyArchH.ArchitectureUpdate)

	business.GET("/filing-permissions", middleware.RequirePermission(r.deps.Enforcer, "filing-permissions.list"), safetyArchH.FilingList)
	business.POST("/filing-permissions", middleware.RequirePermission(r.deps.Enforcer, "filing-permissions.create"), safetyArchH.FilingCreate)
	business.PUT("/filing-permissions/:id", middleware.RequirePermission(r.deps.Enforcer, "filing-permissions.update"), safetyArchH.FilingUpdate)
	business.DELETE("/filing-permissions/:id", middleware.RequirePermission(r.deps.Enforcer, "filing-permissions.delete"), safetyArchH.FilingDelete)

	business.GET("/buildings", middleware.RequirePermission(r.deps.Enforcer, "buildings.list"), safetyArchH.BuildingList)
	business.POST("/buildings", middleware.RequirePermission(r.deps.Enforcer, "buildings.create"), safetyArchH.BuildingCreate)
	business.PUT("/buildings/:id", middleware.RequirePermission(r.deps.Enforcer, "buildings.update"), safetyArchH.BuildingUpdate)
	business.DELETE("/buildings/:id", middleware.RequirePermission(r.deps.Enforcer, "buildings.delete"), safetyArchH.BuildingDelete)

	business.GET("/security-investments", middleware.RequirePermission(r.deps.Enforcer, "security-investments.list"), safetyArchH.InvestmentList)
	business.POST("/security-investments", middleware.RequirePermission(r.deps.Enforcer, "security-investments.create"), safetyArchH.InvestmentCreate)
	business.PUT("/security-investments/:id", middleware.RequirePermission(r.deps.Enforcer, "security-investments.update"), safetyArchH.InvestmentUpdate)
	business.DELETE("/security-investments/:id", middleware.RequirePermission(r.deps.Enforcer, "security-investments.delete"), safetyArchH.InvestmentDelete)

	business.GET("/personnels", middleware.RequirePermission(r.deps.Enforcer, "personnels.list"), safetyArchH.PersonnelList)
	business.GET("/personnels/:id", middleware.RequirePermission(r.deps.Enforcer, "personnels.list"), safetyArchH.PersonnelGet)
	business.POST("/personnels", middleware.RequirePermission(r.deps.Enforcer, "personnels.create"), safetyArchH.PersonnelCreate)
	business.PUT("/personnels/:id", middleware.RequirePermission(r.deps.Enforcer, "personnels.update"), safetyArchH.PersonnelUpdate)

	// Personnels page sub-tabs (持证情况)
	business.GET("/certificate-statuses", middleware.RequirePermission(r.deps.Enforcer, "certificate-statuses.list"), safetyArchH.CertStatusList)
	business.POST("/certificate-statuses", middleware.RequirePermission(r.deps.Enforcer, "certificate-statuses.create"), safetyArchH.CertStatusCreate)
	business.PUT("/certificate-statuses/:id", middleware.RequirePermission(r.deps.Enforcer, "certificate-statuses.update"), safetyArchH.CertStatusUpdate)

	// 安全人员管理 — 3 staff registries
	business.GET("/safety-management", middleware.RequirePermission(r.deps.Enforcer, "safety-management.list"), safetyArchH.SafetyMgmtStaffList)
	business.GET("/safety-management/template", middleware.RequirePermission(r.deps.Enforcer, "safety-management.list"), safetyArchH.SafetyMgmtStaffTemplate)
	business.POST("/safety-management/import", middleware.RequirePermission(r.deps.Enforcer, "safety-management.create"), safetyArchH.SafetyMgmtStaffImport)
	business.POST("/safety-management", middleware.RequirePermission(r.deps.Enforcer, "safety-management.create"), safetyArchH.SafetyMgmtStaffCreate)
	business.PUT("/safety-management/:id", middleware.RequirePermission(r.deps.Enforcer, "safety-management.update"), safetyArchH.SafetyMgmtStaffUpdate)
	business.DELETE("/safety-management/:id", middleware.RequirePermission(r.deps.Enforcer, "safety-management.delete"), safetyArchH.SafetyMgmtStaffDelete)

	business.GET("/special-equipment-staff", middleware.RequirePermission(r.deps.Enforcer, "special-equipment-staff.list"), safetyArchH.SpecialEquipStaffList)
	business.GET("/special-equipment-staff/template", middleware.RequirePermission(r.deps.Enforcer, "special-equipment-staff.list"), safetyArchH.SpecialEquipStaffTemplate)
	business.POST("/special-equipment-staff/import", middleware.RequirePermission(r.deps.Enforcer, "special-equipment-staff.create"), safetyArchH.SpecialEquipStaffImport)
	business.POST("/special-equipment-staff", middleware.RequirePermission(r.deps.Enforcer, "special-equipment-staff.create"), safetyArchH.SpecialEquipStaffCreate)
	business.PUT("/special-equipment-staff/:id", middleware.RequirePermission(r.deps.Enforcer, "special-equipment-staff.update"), safetyArchH.SpecialEquipStaffUpdate)
	business.DELETE("/special-equipment-staff/:id", middleware.RequirePermission(r.deps.Enforcer, "special-equipment-staff.delete"), safetyArchH.SpecialEquipStaffDelete)

	business.GET("/elimination-control-staff", middleware.RequirePermission(r.deps.Enforcer, "elimination-control-staff.list"), safetyArchH.ElimCtrlStaffList)
	business.GET("/elimination-control-staff/template", middleware.RequirePermission(r.deps.Enforcer, "elimination-control-staff.list"), safetyArchH.ElimCtrlStaffTemplate)
	business.POST("/elimination-control-staff/import", middleware.RequirePermission(r.deps.Enforcer, "elimination-control-staff.create"), safetyArchH.ElimCtrlStaffImport)
	business.POST("/elimination-control-staff", middleware.RequirePermission(r.deps.Enforcer, "elimination-control-staff.create"), safetyArchH.ElimCtrlStaffCreate)
	business.PUT("/elimination-control-staff/:id", middleware.RequirePermission(r.deps.Enforcer, "elimination-control-staff.update"), safetyArchH.ElimCtrlStaffUpdate)
	business.DELETE("/elimination-control-staff/:id", middleware.RequirePermission(r.deps.Enforcer, "elimination-control-staff.delete"), safetyArchH.ElimCtrlStaffDelete)

	// Qualifications page sub-tabs (per-company singletons)
	business.GET("/production-standardizations", middleware.RequirePermission(r.deps.Enforcer, "production-standardizations.list"), safetyArchH.ProductionStdGet)
	business.POST("/production-standardizations", middleware.RequirePermission(r.deps.Enforcer, "production-standardizations.create"), safetyArchH.ProductionStdCreate)
	business.PUT("/production-standardizations/:id", middleware.RequirePermission(r.deps.Enforcer, "production-standardizations.update"), safetyArchH.ProductionStdUpdate)

	business.GET("/occupational-health-safeties", middleware.RequirePermission(r.deps.Enforcer, "occupational-health-safeties.list"), safetyArchH.OHSGet)
	business.POST("/occupational-health-safeties", middleware.RequirePermission(r.deps.Enforcer, "occupational-health-safeties.create"), safetyArchH.OHSCreate)
	business.PUT("/occupational-health-safeties/:id", middleware.RequirePermission(r.deps.Enforcer, "occupational-health-safeties.update"), safetyArchH.OHSUpdate)

	// 安全投入 — 经费执行情况
	business.GET("/funding-executions", middleware.RequirePermission(r.deps.Enforcer, "funding-executions.list"), safetyArchH.FundingExecList)
	business.POST("/funding-executions", middleware.RequirePermission(r.deps.Enforcer, "funding-executions.create"), safetyArchH.FundingExecCreate)
	business.PUT("/funding-executions/:id", middleware.RequirePermission(r.deps.Enforcer, "funding-executions.update"), safetyArchH.FundingExecUpdate)
	business.DELETE("/funding-executions/:id", middleware.RequirePermission(r.deps.Enforcer, "funding-executions.delete"), safetyArchH.FundingExecDelete)

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

	// 重大危险源（major_hazards + major_constitutes）— 主表 CRUD，子表整体替换。
	business.GET("/major-hazards", middleware.RequirePermission(r.deps.Enforcer, "major-hazards.list"), majorHazardH.List)
	business.POST("/major-hazards", middleware.RequirePermission(r.deps.Enforcer, "major-hazards.create"), majorHazardH.Create)
	business.PUT("/major-hazards/:id", middleware.RequirePermission(r.deps.Enforcer, "major-hazards.update"), majorHazardH.Update)
	business.DELETE("/major-hazards/:id", middleware.RequirePermission(r.deps.Enforcer, "major-hazards.delete"), majorHazardH.Delete)

	// 工艺信息 — 三模块共用 ProcessService（重点监管 / 较大危险 / 有限空间）。
	business.GET("/major-hazardous-processes", middleware.RequirePermission(r.deps.Enforcer, "major-hazardous-processes.list"), processH.ListMajorHazardous)
	business.POST("/major-hazardous-processes", middleware.RequirePermission(r.deps.Enforcer, "major-hazardous-processes.create"), processH.CreateMajorHazardous)
	business.PUT("/major-hazardous-processes/:id", middleware.RequirePermission(r.deps.Enforcer, "major-hazardous-processes.update"), processH.UpdateMajorHazardous)
	business.DELETE("/major-hazardous-processes/:id", middleware.RequirePermission(r.deps.Enforcer, "major-hazardous-processes.delete"), processH.DeleteMajorHazardous)

	business.GET("/hazardous-processes", middleware.RequirePermission(r.deps.Enforcer, "hazardous-processes.list"), processH.ListHazardous)
	business.POST("/hazardous-processes", middleware.RequirePermission(r.deps.Enforcer, "hazardous-processes.create"), processH.CreateHazardous)
	business.PUT("/hazardous-processes/:id", middleware.RequirePermission(r.deps.Enforcer, "hazardous-processes.update"), processH.UpdateHazardous)
	business.DELETE("/hazardous-processes/:id", middleware.RequirePermission(r.deps.Enforcer, "hazardous-processes.delete"), processH.DeleteHazardous)

	business.GET("/space-works", middleware.RequirePermission(r.deps.Enforcer, "space-works.list"), processH.ListSpaceWork)
	business.POST("/space-works", middleware.RequirePermission(r.deps.Enforcer, "space-works.create"), processH.CreateSpaceWork)
	business.PUT("/space-works/:id", middleware.RequirePermission(r.deps.Enforcer, "space-works.update"), processH.UpdateSpaceWork)
	business.DELETE("/space-works/:id", middleware.RequirePermission(r.deps.Enforcer, "space-works.delete"), processH.DeleteSpaceWork)

	// Upload (OSS multipart)
	upload := business.Group("/upload-oss", middleware.RequirePermission(r.deps.Enforcer, "upload.oss"))
	upload.POST("/init", uploadH.Init)
	upload.POST("/sign-part", uploadH.SignPart)
	upload.POST("/complete", uploadH.Complete)
	upload.POST("/abort", uploadH.Abort)
	upload.GET("/signed", uploadH.SignedURL)

	// Upload (OSS browser-direct POST policy) — used by OssUploadField.
	// Different prefix (/upload, no -oss) to match the Laravel routes the
	// frontend already calls.
	uploadPolicy := business.Group("/upload", middleware.RequirePermission(r.deps.Enforcer, "upload.oss"))
	uploadPolicy.GET("/oss-signature", uploadH.PostPolicy)
	uploadPolicy.GET("/oss-sign", uploadH.SignByURL)

	// Dashboard placeholders for not-yet-ported modules. Empty success
	// payloads keep the SPA rendering until the real modules port (plan §10).
	mountDashboardStubs(business)

	// ping (smoke)
	business.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pong": true})
	})

	return engine
}
