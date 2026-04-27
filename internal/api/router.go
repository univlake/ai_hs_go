package api

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	accidentemergencyapi "ai_hs_go/internal/api/accidentemergency"
	assessmentapi "ai_hs_go/internal/api/assessment"
	authapi "ai_hs_go/internal/api/auth"
	basedataapi "ai_hs_go/internal/api/basedata"
	blocexpertapi "ai_hs_go/internal/api/blocexpert"
	companyapi "ai_hs_go/internal/api/company"
	companygroupapi "ai_hs_go/internal/api/companygroup"
	expertapi "ai_hs_go/internal/api/expert"
	healthfireapi "ai_hs_go/internal/api/healthfire"
	hiddendangerapi "ai_hs_go/internal/api/hiddendanger"
	majorhazardapi "ai_hs_go/internal/api/majorhazard"
	meetingtrainingapi "ai_hs_go/internal/api/meetingtraining"
	notificationapi "ai_hs_go/internal/api/notification"
	permissionapi "ai_hs_go/internal/api/permission"
	platformapi "ai_hs_go/internal/api/platform"
	processapi "ai_hs_go/internal/api/process"
	riskapi "ai_hs_go/internal/api/risk"
	roleapi "ai_hs_go/internal/api/role"
	safetyarchapi "ai_hs_go/internal/api/safetyarchive"
	safetyapi "ai_hs_go/internal/api/safetyledger"
	tenantapi "ai_hs_go/internal/api/tenant"
	thirdpartyapi "ai_hs_go/internal/api/thirdparty"
	todoapi "ai_hs_go/internal/api/todo"
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
	// 注入 rbacSvc 让 CompanyService.Create / CreateRoot 自动给派生 c{id}_admin
	// 分配 tenant_admin / tenant_root_admin 角色（Laravel 行为对齐）。
	companySvc := service.NewCompanyService(r.deps.DB, companyRepo).WithRBAC(rbacSvc)
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
	thirdPartySvc := service.NewThirdPartyService(r.deps.DB)
	thirdPartyH := thirdpartyapi.NewHandler(thirdPartySvc)
	meetingTrainingSvc := service.NewMeetingTrainingService(r.deps.DB)
	meetingTrainingH := meetingtrainingapi.NewHandler(meetingTrainingSvc)
	healthFireSvc := service.NewHealthFireService(r.deps.DB)
	healthFireH := healthfireapi.NewHandler(healthFireSvc)
	riskSvc := service.NewRiskService(r.deps.DB)
	riskH := riskapi.NewHandler(riskSvc, r.deps.OSS)
	accidentEmergencySvc := service.NewAccidentEmergencyService(r.deps.DB, companyRepo)
	accidentEmergencyH := accidentemergencyapi.NewHandler(accidentEmergencySvc)
	notificationSvc := service.NewNotificationService(r.deps.DB, companyRepo)
	notificationH := notificationapi.NewHandler(notificationSvc)
	// 检查通知单 docx 模板与生成文件落盘根目录跟 Laravel 对齐：
	//   templates: configs/templates/enterprise_export.docx (随项目分发)
	//   storage : storage/app/inspection-notices/YYYY/MM/   (运行时自动建)
	expertSvc := service.NewExpertService(r.deps.DB, "./configs/templates/enterprise_export.docx", "./storage/app")
	rectifySvc := service.NewRectificationSheetService(r.deps.DB, r.deps.OSS, "./storage/app")
	expertH := expertapi.NewHandler(expertSvc, rectifySvc)
	blocExpertH := blocexpertapi.NewHandler(expertSvc)
	todoH := todoapi.NewHandler()
	hiddenDangerSvc := service.NewHiddenDangerService(r.deps.DB)
	hiddenDangerH := hiddendangerapi.NewHandler(hiddenDangerSvc)
	// 考核评分模块 — 模板路径下面的三份 docx 跟 Laravel public/check_*.docx 对齐：
	//   configs/templates/check_bloc.docx
	//   configs/templates/check_company.docx
	//   configs/templates/check_company_new.docx
	// 落盘根目录 storage/app/check_word/YYYY/MM/<filename>。
	assessmentSvc := service.NewAssessmentService(r.deps.DB)
	assessmentH := assessmentapi.NewHandler(assessmentSvc, "./configs/templates", "./storage/app")
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
	business.POST("/special-equipment/copy", middleware.RequirePermission(r.deps.Enforcer, "special-equipment.create"), ledgerH.CopySpecialEquipment)

	business.GET("/dangerous-equipment", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.list"), ledgerH.ListDangerousEquipment)
	business.POST("/dangerous-equipment", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.create"), ledgerH.CreateDangerousEquipment)
	business.PUT("/dangerous-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.update"), ledgerH.UpdateDangerousEquipment)
	business.DELETE("/dangerous-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.delete"), ledgerH.DeleteDangerousEquipment)
	business.GET("/dangerous-equipment/template", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.list"), ledgerH.TemplateDangerousEquipment)
	business.GET("/dangerous-equipment/export", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.list"), ledgerH.ExportDangerousEquipment)
	business.POST("/dangerous-equipment/import", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.create"), ledgerH.ImportDangerousEquipment)
	business.POST("/dangerous-equipment/copy", middleware.RequirePermission(r.deps.Enforcer, "dangerous-equipment.create"), ledgerH.CopyDangerousEquipment)

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

	// 第三方管理 — 三模块共用 ThirdPartyService（设计单位 / 评价单位 / 技术服务单位）。
	business.GET("/design-units", middleware.RequirePermission(r.deps.Enforcer, "design-units.list"), thirdPartyH.ListDesignUnit)
	business.POST("/design-units", middleware.RequirePermission(r.deps.Enforcer, "design-units.create"), thirdPartyH.CreateDesignUnit)
	business.PUT("/design-units/:id", middleware.RequirePermission(r.deps.Enforcer, "design-units.update"), thirdPartyH.UpdateDesignUnit)
	business.DELETE("/design-units/:id", middleware.RequirePermission(r.deps.Enforcer, "design-units.delete"), thirdPartyH.DeleteDesignUnit)

	business.GET("/evaluation-units", middleware.RequirePermission(r.deps.Enforcer, "evaluation-units.list"), thirdPartyH.ListEvaluationUnit)
	business.POST("/evaluation-units", middleware.RequirePermission(r.deps.Enforcer, "evaluation-units.create"), thirdPartyH.CreateEvaluationUnit)
	business.PUT("/evaluation-units/:id", middleware.RequirePermission(r.deps.Enforcer, "evaluation-units.update"), thirdPartyH.UpdateEvaluationUnit)
	business.DELETE("/evaluation-units/:id", middleware.RequirePermission(r.deps.Enforcer, "evaluation-units.delete"), thirdPartyH.DeleteEvaluationUnit)

	business.GET("/skill-units", middleware.RequirePermission(r.deps.Enforcer, "skill-units.list"), thirdPartyH.ListSkillUnit)
	business.GET("/skill-units/:id", middleware.RequirePermission(r.deps.Enforcer, "skill-units.list"), thirdPartyH.GetSkillUnit)
	business.POST("/skill-units", middleware.RequirePermission(r.deps.Enforcer, "skill-units.create"), thirdPartyH.CreateSkillUnit)
	business.PUT("/skill-units/:id", middleware.RequirePermission(r.deps.Enforcer, "skill-units.update"), thirdPartyH.UpdateSkillUnit)
	business.DELETE("/skill-units/:id", middleware.RequirePermission(r.deps.Enforcer, "skill-units.delete"), thirdPartyH.DeleteSkillUnit)

	// 会议培训 — 四模块共用 MeetingTrainingService（内部会议 / 内部、外部、相关方培训）。
	business.GET("/internal-meetings", middleware.RequirePermission(r.deps.Enforcer, "internal-meetings.list"), meetingTrainingH.ListInternalMeeting)
	business.GET("/internal-meetings/:id", middleware.RequirePermission(r.deps.Enforcer, "internal-meetings.list"), meetingTrainingH.GetInternalMeeting)
	business.POST("/internal-meetings", middleware.RequirePermission(r.deps.Enforcer, "internal-meetings.create"), meetingTrainingH.CreateInternalMeeting)
	business.PUT("/internal-meetings/:id", middleware.RequirePermission(r.deps.Enforcer, "internal-meetings.update"), meetingTrainingH.UpdateInternalMeeting)
	business.DELETE("/internal-meetings/:id", middleware.RequirePermission(r.deps.Enforcer, "internal-meetings.delete"), meetingTrainingH.DeleteInternalMeeting)

	business.GET("/internal-trainings", middleware.RequirePermission(r.deps.Enforcer, "internal-trainings.list"), meetingTrainingH.ListInternalTraining)
	business.GET("/internal-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "internal-trainings.list"), meetingTrainingH.GetInternalTraining)
	business.POST("/internal-trainings", middleware.RequirePermission(r.deps.Enforcer, "internal-trainings.create"), meetingTrainingH.CreateInternalTraining)
	business.PUT("/internal-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "internal-trainings.update"), meetingTrainingH.UpdateInternalTraining)
	business.DELETE("/internal-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "internal-trainings.delete"), meetingTrainingH.DeleteInternalTraining)

	business.GET("/external-trainings", middleware.RequirePermission(r.deps.Enforcer, "external-trainings.list"), meetingTrainingH.ListExternalTraining)
	business.GET("/external-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "external-trainings.list"), meetingTrainingH.GetExternalTraining)
	business.POST("/external-trainings", middleware.RequirePermission(r.deps.Enforcer, "external-trainings.create"), meetingTrainingH.CreateExternalTraining)
	business.PUT("/external-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "external-trainings.update"), meetingTrainingH.UpdateExternalTraining)
	business.DELETE("/external-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "external-trainings.delete"), meetingTrainingH.DeleteExternalTraining)

	business.GET("/other-trainings", middleware.RequirePermission(r.deps.Enforcer, "other-trainings.list"), meetingTrainingH.ListOtherTraining)
	business.GET("/other-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "other-trainings.list"), meetingTrainingH.GetOtherTraining)
	business.POST("/other-trainings", middleware.RequirePermission(r.deps.Enforcer, "other-trainings.create"), meetingTrainingH.CreateOtherTraining)
	business.PUT("/other-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "other-trainings.update"), meetingTrainingH.UpdateOtherTraining)
	business.DELETE("/other-trainings/:id", middleware.RequirePermission(r.deps.Enforcer, "other-trainings.delete"), meetingTrainingH.DeleteOtherTraining)

	// 职业健康 + 消防安全 — 5 + 1（小模块）+ 3 模块共用 HealthFireService。
	// 静态后缀 (/template /export /import /copy) 必须先于 /:id 注册，避免命中通配。
	business.GET("/occupational-hazards", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.list"), healthFireH.ListOccupationalHazard)
	business.GET("/occupational-hazards-template", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.list"), healthFireH.TemplateOccupationalHazard)
	business.GET("/occupational-hazards-export", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.list"), healthFireH.ExportOccupationalHazard)
	business.POST("/occupational-hazards-import", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.create"), healthFireH.ImportOccupationalHazard)
	business.GET("/occupational-hazards/:id", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.list"), healthFireH.GetOccupationalHazard)
	business.POST("/occupational-hazards", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.create"), healthFireH.CreateOccupationalHazard)
	business.PUT("/occupational-hazards/:id", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.update"), healthFireH.UpdateOccupationalHazard)
	business.DELETE("/occupational-hazards/:id", middleware.RequirePermission(r.deps.Enforcer, "occupational-hazards.delete"), healthFireH.DeleteOccupationalHazard)

	business.GET("/health-systems", middleware.RequirePermission(r.deps.Enforcer, "health-systems.list"), healthFireH.GetHealthSystem)
	business.POST("/health-systems", middleware.RequirePermission(r.deps.Enforcer, "health-systems.create"), healthFireH.CreateHealthSystem)
	business.PUT("/health-systems/:id", middleware.RequirePermission(r.deps.Enforcer, "health-systems.update"), healthFireH.UpdateHealthSystem)

	business.GET("/health-files", middleware.RequirePermission(r.deps.Enforcer, "health-files.list"), healthFireH.ListHealthFile)
	business.GET("/health-files-template", middleware.RequirePermission(r.deps.Enforcer, "health-files.list"), healthFireH.TemplateHealthFile)
	business.GET("/health-files-export", middleware.RequirePermission(r.deps.Enforcer, "health-files.list"), healthFireH.ExportHealthFile)
	business.POST("/health-files-import", middleware.RequirePermission(r.deps.Enforcer, "health-files.create"), healthFireH.ImportHealthFile)
	business.GET("/health-files/:id", middleware.RequirePermission(r.deps.Enforcer, "health-files.list"), healthFireH.GetHealthFile)
	business.POST("/health-files", middleware.RequirePermission(r.deps.Enforcer, "health-files.create"), healthFireH.CreateHealthFile)
	business.PUT("/health-files/:id", middleware.RequirePermission(r.deps.Enforcer, "health-files.update"), healthFireH.UpdateHealthFile)
	business.DELETE("/health-files/:id", middleware.RequirePermission(r.deps.Enforcer, "health-files.delete"), healthFireH.DeleteHealthFile)

	business.GET("/health-reports", middleware.RequirePermission(r.deps.Enforcer, "health-reports.list"), healthFireH.ListHealthReport)
	business.GET("/health-reports/:id", middleware.RequirePermission(r.deps.Enforcer, "health-reports.list"), healthFireH.GetHealthReport)
	business.POST("/health-reports", middleware.RequirePermission(r.deps.Enforcer, "health-reports.create"), healthFireH.CreateHealthReport)
	business.PUT("/health-reports/:id", middleware.RequirePermission(r.deps.Enforcer, "health-reports.update"), healthFireH.UpdateHealthReport)
	business.DELETE("/health-reports/:id", middleware.RequirePermission(r.deps.Enforcer, "health-reports.delete"), healthFireH.DeleteHealthReport)

	business.GET("/health-checkups", middleware.RequirePermission(r.deps.Enforcer, "health-checkups.list"), healthFireH.ListHealthCheckup)
	business.GET("/health-checkups/:id", middleware.RequirePermission(r.deps.Enforcer, "health-checkups.list"), healthFireH.GetHealthCheckup)
	business.POST("/health-checkups", middleware.RequirePermission(r.deps.Enforcer, "health-checkups.create"), healthFireH.CreateHealthCheckup)
	business.PUT("/health-checkups/:id", middleware.RequirePermission(r.deps.Enforcer, "health-checkups.update"), healthFireH.UpdateHealthCheckup)
	business.DELETE("/health-checkups/:id", middleware.RequirePermission(r.deps.Enforcer, "health-checkups.delete"), healthFireH.DeleteHealthCheckup)

	business.GET("/hygiene-evaluations", middleware.RequirePermission(r.deps.Enforcer, "hygiene-evaluations.list"), healthFireH.ListHygieneEvaluation)
	business.GET("/hygiene-evaluations/:id", middleware.RequirePermission(r.deps.Enforcer, "hygiene-evaluations.list"), healthFireH.GetHygieneEvaluation)
	business.POST("/hygiene-evaluations", middleware.RequirePermission(r.deps.Enforcer, "hygiene-evaluations.create"), healthFireH.CreateHygieneEvaluation)
	business.PUT("/hygiene-evaluations/:id", middleware.RequirePermission(r.deps.Enforcer, "hygiene-evaluations.update"), healthFireH.UpdateHygieneEvaluation)
	business.DELETE("/hygiene-evaluations/:id", middleware.RequirePermission(r.deps.Enforcer, "hygiene-evaluations.delete"), healthFireH.DeleteHygieneEvaluation)

	business.GET("/fire-safeties", middleware.RequirePermission(r.deps.Enforcer, "fire-safeties.list"), healthFireH.ListFireSafety)
	business.GET("/fire-safeties/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-safeties.list"), healthFireH.GetFireSafety)
	business.POST("/fire-safeties", middleware.RequirePermission(r.deps.Enforcer, "fire-safeties.create"), healthFireH.CreateFireSafety)
	business.PUT("/fire-safeties/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-safeties.update"), healthFireH.UpdateFireSafety)
	business.DELETE("/fire-safeties/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-safeties.delete"), healthFireH.DeleteFireSafety)

	business.GET("/fire-equipment", middleware.RequirePermission(r.deps.Enforcer, "fire-equipment.list"), healthFireH.ListFireEquipment)
	business.POST("/fire-equipment/copy", middleware.RequirePermission(r.deps.Enforcer, "fire-equipment.create"), healthFireH.CopyFireEquipment)
	business.GET("/fire-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-equipment.list"), healthFireH.GetFireEquipment)
	business.POST("/fire-equipment", middleware.RequirePermission(r.deps.Enforcer, "fire-equipment.create"), healthFireH.CreateFireEquipment)
	business.PUT("/fire-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-equipment.update"), healthFireH.UpdateFireEquipment)
	business.DELETE("/fire-equipment/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-equipment.delete"), healthFireH.DeleteFireEquipment)

	business.GET("/fire-facilities", middleware.RequirePermission(r.deps.Enforcer, "fire-facilities.list"), healthFireH.ListFireFacility)
	business.POST("/fire-facilities/copy", middleware.RequirePermission(r.deps.Enforcer, "fire-facilities.create"), healthFireH.CopyFireFacility)
	business.GET("/fire-facilities/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-facilities.list"), healthFireH.GetFireFacility)
	business.POST("/fire-facilities", middleware.RequirePermission(r.deps.Enforcer, "fire-facilities.create"), healthFireH.CreateFireFacility)
	business.PUT("/fire-facilities/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-facilities.update"), healthFireH.UpdateFireFacility)
	business.DELETE("/fire-facilities/:id", middleware.RequirePermission(r.deps.Enforcer, "fire-facilities.delete"), healthFireH.DeleteFireFacility)

	// 风险管理 — 5 模块共用 RiskService（风险源 / 隐患 / 作业 / 三同时 / 安全评价）。
	// 风险分布图是 per-company singleton（沿用 health-systems 的 GET/POST/PUT 写法）。
	business.GET("/risk-distribution-maps", middleware.RequirePermission(r.deps.Enforcer, "risk-distribution-maps.list"), riskH.GetRiskDistributionMap)
	business.POST("/risk-distribution-maps", middleware.RequirePermission(r.deps.Enforcer, "risk-distribution-maps.create"), riskH.CreateRiskDistributionMap)
	business.PUT("/risk-distribution-maps/:id", middleware.RequirePermission(r.deps.Enforcer, "risk-distribution-maps.update"), riskH.UpdateRiskDistributionMap)

	business.GET("/monthly-investigations", middleware.RequirePermission(r.deps.Enforcer, "monthly-investigations.list"), riskH.ListMonthlyInvestigation)
	business.GET("/monthly-investigations/:id", middleware.RequirePermission(r.deps.Enforcer, "monthly-investigations.list"), riskH.GetMonthlyInvestigation)
	business.POST("/monthly-investigations", middleware.RequirePermission(r.deps.Enforcer, "monthly-investigations.create"), riskH.CreateMonthlyInvestigation)
	business.PUT("/monthly-investigations/:id", middleware.RequirePermission(r.deps.Enforcer, "monthly-investigations.update"), riskH.UpdateMonthlyInvestigation)
	business.DELETE("/monthly-investigations/:id", middleware.RequirePermission(r.deps.Enforcer, "monthly-investigations.delete"), riskH.DeleteMonthlyInvestigation)

	business.GET("/job-management", middleware.RequirePermission(r.deps.Enforcer, "job-management.list"), riskH.ListJobManagement)
	business.GET("/job-management/:id", middleware.RequirePermission(r.deps.Enforcer, "job-management.list"), riskH.GetJobManagement)
	business.POST("/job-management", middleware.RequirePermission(r.deps.Enforcer, "job-management.create"), riskH.CreateJobManagement)
	business.PUT("/job-management/:id", middleware.RequirePermission(r.deps.Enforcer, "job-management.update"), riskH.UpdateJobManagement)
	business.DELETE("/job-management/:id", middleware.RequirePermission(r.deps.Enforcer, "job-management.delete"), riskH.DeleteJobManagement)

	business.GET("/three-togethers", middleware.RequirePermission(r.deps.Enforcer, "three-togethers.list"), riskH.ListThreeTogether)
	business.GET("/three-togethers/:id", middleware.RequirePermission(r.deps.Enforcer, "three-togethers.list"), riskH.GetThreeTogether)
	business.POST("/three-togethers", middleware.RequirePermission(r.deps.Enforcer, "three-togethers.create"), riskH.CreateThreeTogether)
	business.PUT("/three-togethers/:id", middleware.RequirePermission(r.deps.Enforcer, "three-togethers.update"), riskH.UpdateThreeTogether)
	business.DELETE("/three-togethers/:id", middleware.RequirePermission(r.deps.Enforcer, "three-togethers.delete"), riskH.DeleteThreeTogether)

	business.GET("/assess-detects", middleware.RequirePermission(r.deps.Enforcer, "assess-detects.list"), riskH.ListAssessDetect)
	business.GET("/assess-detects/:id", middleware.RequirePermission(r.deps.Enforcer, "assess-detects.list"), riskH.GetAssessDetect)
	business.POST("/assess-detects", middleware.RequirePermission(r.deps.Enforcer, "assess-detects.create"), riskH.CreateAssessDetect)
	business.PUT("/assess-detects/:id", middleware.RequirePermission(r.deps.Enforcer, "assess-detects.update"), riskH.UpdateAssessDetect)
	business.DELETE("/assess-detects/:id", middleware.RequirePermission(r.deps.Enforcer, "assess-detects.delete"), riskH.DeleteAssessDetect)

	// 事故应急 — 5 模块共用 AccidentEmergencyService。
	// 集团事故快报（/bloc-accident-bulletins）跟生产视图共用 accident_bulletins 表，
	// 仅在 service 层按 company subtree 过滤。生产视图的事故上报与集团视图的事故上报
	// 共用同一份 /accident-bulletins CRUD（前端 management/accidents/report 直接复用 accidentReportConfig）。
	business.GET("/accident-bulletins", middleware.RequirePermission(r.deps.Enforcer, "accident-bulletins.list"), accidentEmergencyH.ListAccidentBulletin)
	business.GET("/accident-bulletins/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-bulletins.list"), accidentEmergencyH.GetAccidentBulletin)
	business.POST("/accident-bulletins", middleware.RequirePermission(r.deps.Enforcer, "accident-bulletins.create"), accidentEmergencyH.CreateAccidentBulletin)
	business.PUT("/accident-bulletins/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-bulletins.update"), accidentEmergencyH.UpdateAccidentBulletin)
	business.DELETE("/accident-bulletins/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-bulletins.delete"), accidentEmergencyH.DeleteAccidentBulletin)

	// 集团事故快报：/annual-stats 必须先于 通配的 :id 注册（这里没有 :id 路由，
	// 但保持习惯）。bloc 路由 GET-only，写操作走 production /accident-bulletins。
	business.GET("/bloc-accident-bulletins/annual-stats", middleware.RequirePermission(r.deps.Enforcer, "bloc-accident-bulletins.list"), accidentEmergencyH.BlocAccidentAnnualStats)
	business.GET("/bloc-accident-bulletins", middleware.RequirePermission(r.deps.Enforcer, "bloc-accident-bulletins.list"), accidentEmergencyH.ListBlocAccidentBulletin)

	business.GET("/emergency-drills", middleware.RequirePermission(r.deps.Enforcer, "emergency-drills.list"), accidentEmergencyH.ListEmergencyDrill)
	business.GET("/emergency-drills/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-drills.list"), accidentEmergencyH.GetEmergencyDrill)
	business.POST("/emergency-drills", middleware.RequirePermission(r.deps.Enforcer, "emergency-drills.create"), accidentEmergencyH.CreateEmergencyDrill)
	business.PUT("/emergency-drills/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-drills.update"), accidentEmergencyH.UpdateEmergencyDrill)
	business.DELETE("/emergency-drills/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-drills.delete"), accidentEmergencyH.DeleteEmergencyDrill)

	business.GET("/emergency-supplies", middleware.RequirePermission(r.deps.Enforcer, "emergency-supplies.list"), accidentEmergencyH.ListEmergencySupplies)
	business.GET("/emergency-supplies/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-supplies.list"), accidentEmergencyH.GetEmergencySupplies)
	business.POST("/emergency-supplies", middleware.RequirePermission(r.deps.Enforcer, "emergency-supplies.create"), accidentEmergencyH.CreateEmergencySupplies)
	business.PUT("/emergency-supplies/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-supplies.update"), accidentEmergencyH.UpdateEmergencySupplies)
	business.DELETE("/emergency-supplies/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-supplies.delete"), accidentEmergencyH.DeleteEmergencySupplies)

	business.GET("/emergency-teams", middleware.RequirePermission(r.deps.Enforcer, "emergency-teams.list"), accidentEmergencyH.ListEmergencyTeam)
	business.GET("/emergency-teams/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-teams.list"), accidentEmergencyH.GetEmergencyTeam)
	business.POST("/emergency-teams", middleware.RequirePermission(r.deps.Enforcer, "emergency-teams.create"), accidentEmergencyH.CreateEmergencyTeam)
	business.PUT("/emergency-teams/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-teams.update"), accidentEmergencyH.UpdateEmergencyTeam)
	business.DELETE("/emergency-teams/:id", middleware.RequirePermission(r.deps.Enforcer, "emergency-teams.delete"), accidentEmergencyH.DeleteEmergencyTeam)

	// 通知公告 / 会议通知 / 信息稿 — 三模块共用 NotificationService。
	// 生产视图（receiveOnly 前端）跟集团视图共用同一组路由，前端切换 receiveOnly
	// 即可隐藏写入控件，service 内 list 同时返回「我发起的 + 我接收的」两类记录。
	//
	// 任务型通知（office_notifies）
	business.GET("/office-notifies", middleware.RequirePermission(r.deps.Enforcer, "office-notifies.list"), notificationH.ListOfficeNotify)
	business.GET("/office-notifies/:id", middleware.RequirePermission(r.deps.Enforcer, "office-notifies.list"), notificationH.GetOfficeNotify)
	business.POST("/office-notifies", middleware.RequirePermission(r.deps.Enforcer, "office-notifies.create"), notificationH.CreateOfficeNotify)
	business.PUT("/office-notifies/:id", middleware.RequirePermission(r.deps.Enforcer, "office-notifies.update"), notificationH.UpdateOfficeNotify)
	business.DELETE("/office-notifies/:id", middleware.RequirePermission(r.deps.Enforcer, "office-notifies.delete"), notificationH.DeleteOfficeNotify)
	business.POST("/office-notify-scopes/:id/feedback", middleware.RequirePermission(r.deps.Enforcer, "office-notifies.list"), notificationH.SubmitOfficeNotifyFeedback)
	business.GET("/office-notify-notifications/pending", notificationH.PendingOfficeNotifyScopes)

	// 通知型通知（new_notifies）
	business.GET("/new-notifies", middleware.RequirePermission(r.deps.Enforcer, "new-notifies.list"), notificationH.ListNewNotify)
	business.GET("/new-notifies/:id", middleware.RequirePermission(r.deps.Enforcer, "new-notifies.list"), notificationH.GetNewNotify)
	business.POST("/new-notifies", middleware.RequirePermission(r.deps.Enforcer, "new-notifies.create"), notificationH.CreateNewNotify)
	business.PUT("/new-notifies/:id", middleware.RequirePermission(r.deps.Enforcer, "new-notifies.update"), notificationH.UpdateNewNotify)
	business.DELETE("/new-notifies/:id", middleware.RequirePermission(r.deps.Enforcer, "new-notifies.delete"), notificationH.DeleteNewNotify)
	business.POST("/new-notify-scopes/:id/read", notificationH.MarkNewNotifyRead)
	business.GET("/new-notify-notifications/pending", notificationH.PendingNewNotifyScopes)

	// 会议通知（meetings + meeting_notices + meeting_sign_ups）
	business.GET("/meetings", middleware.RequirePermission(r.deps.Enforcer, "meetings.list"), notificationH.ListMeeting)
	business.GET("/meetings/:id", middleware.RequirePermission(r.deps.Enforcer, "meetings.list"), notificationH.GetMeeting)
	business.POST("/meetings", middleware.RequirePermission(r.deps.Enforcer, "meetings.create"), notificationH.CreateMeeting)
	business.PUT("/meetings/:id", middleware.RequirePermission(r.deps.Enforcer, "meetings.update"), notificationH.UpdateMeeting)
	business.DELETE("/meetings/:id", middleware.RequirePermission(r.deps.Enforcer, "meetings.delete"), notificationH.DeleteMeeting)
	business.POST("/meeting-notices/:id/notify", middleware.RequirePermission(r.deps.Enforcer, "meetings.update"), notificationH.NotifyMeetingNotice)
	business.POST("/meeting-notices/:id/confirm", notificationH.ConfirmMeetingNotice)
	business.GET("/meeting-notifications/pending", notificationH.PendingMeetingNotices)

	business.GET("/meeting-sign-ups", middleware.RequirePermission(r.deps.Enforcer, "meeting-sign-ups.list"), notificationH.ListMeetingSignUp)
	business.GET("/meeting-sign-ups/:id", middleware.RequirePermission(r.deps.Enforcer, "meeting-sign-ups.list"), notificationH.GetMeetingSignUp)
	business.POST("/meeting-sign-ups", middleware.RequirePermission(r.deps.Enforcer, "meeting-sign-ups.create"), notificationH.CreateMeetingSignUp)
	business.PUT("/meeting-sign-ups/:id", middleware.RequirePermission(r.deps.Enforcer, "meeting-sign-ups.update"), notificationH.UpdateMeetingSignUp)
	business.DELETE("/meeting-sign-ups/:id", middleware.RequirePermission(r.deps.Enforcer, "meeting-sign-ups.delete"), notificationH.DeleteMeetingSignUp)

	// 信息稿管理（information_drafts）— 集团视图（received scope）跟生产视图（published scope）
	// 共享同一组路由，前端通过 ?scope= 切换可见集合。dashboard/* 给仪表板「最新动态」+「待办」用。
	business.GET("/information-drafts/dashboard/readings", notificationH.DashboardInformationDraftReadings)
	business.GET("/information-drafts/dashboard/pending", notificationH.DashboardInformationDraftPending)
	business.GET("/information-drafts", middleware.RequirePermission(r.deps.Enforcer, "information-drafts.list"), notificationH.ListInformationDraft)
	business.GET("/information-drafts/:id", middleware.RequirePermission(r.deps.Enforcer, "information-drafts.list"), notificationH.GetInformationDraft)
	business.POST("/information-drafts", middleware.RequirePermission(r.deps.Enforcer, "information-drafts.create"), notificationH.CreateInformationDraft)
	business.PUT("/information-drafts/:id", middleware.RequirePermission(r.deps.Enforcer, "information-drafts.update"), notificationH.UpdateInformationDraft)
	business.DELETE("/information-drafts/:id", middleware.RequirePermission(r.deps.Enforcer, "information-drafts.delete"), notificationH.DeleteInformationDraft)

	// 专家检查 + 专家库 + 待办事项 — 跟 Laravel routes/api/statistics-reporting.php 对齐。
	//
	// expert-inspections.{list,create,update,delete} 三个层级（计划 / 受检企业 / 专家指派）
	// 共享一组权限。仪表板通知、my-checks、my-assignments、confirm 不走权限中间件，
	// 登录态即可（跟 Laravel 一致），具体可见性通过 service 层的 (company_id / user_id)
	// 过滤兜底。
	business.GET("/expert-inspections", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.list"), expertH.ListInspection)
	business.GET("/expert-inspections/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.list"), expertH.GetInspection)
	business.GET("/expert-inspection-companies", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.list"), expertH.ListCompany)
	business.GET("/expert-inspection-users", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.list"), expertH.ListUser)
	business.POST("/expert-inspections", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.create"), expertH.CreateInspection)
	business.POST("/expert-inspection-companies", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.create"), expertH.CreateCompany)
	business.POST("/expert-inspection-users", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.create"), expertH.CreateUser)
	business.PUT("/expert-inspections/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.update"), expertH.UpdateInspection)
	business.PUT("/expert-inspection-companies/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.update"), expertH.UpdateCompany)
	business.PUT("/expert-inspection-users/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.update"), expertH.UpdateUser)
	business.DELETE("/expert-inspections/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.delete"), expertH.DeleteInspection)
	business.DELETE("/expert-inspection-companies/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.delete"), expertH.DeleteCompany)
	business.DELETE("/expert-inspection-users/:id", middleware.RequirePermission(r.deps.Enforcer, "expert-inspections.delete"), expertH.DeleteUser)
	// 通知 / 确认 / 我的检查 — 不走权限，仅鉴权。
	business.GET("/expert-inspection-notifications/company", expertH.CompanyNotifications)
	business.GET("/expert-inspection-notifications/expert", expertH.ExpertNotifications)
	business.GET("/my-check-inspections", expertH.MyCheckInspections)
	business.GET("/my-expert-assignments", expertH.MyExpertAssignments)
	business.POST("/expert-inspection-companies/:id/confirm", expertH.ConfirmCompany)
	business.POST("/expert-inspection-users/:id/confirm", expertH.ConfirmUser)
	// 检查通知单 docx 生成 / 下载（计划方生成；计划方 + 已指派专家可下载）。
	business.POST("/expert-inspection-companies/:id/inspection-notice", expertH.GenerateInspectionNotice)
	business.GET("/expert-inspection-companies/:id/inspection-notice", expertH.DownloadInspectionNotice)
	// 整改通知单 xlsx 生成 / 下载（计划方 + 已指派专家可生成；受检企业也可下载）。
	business.POST("/expert-inspection-companies/:id/rectification-sheet", expertH.GenerateRectificationSheet)
	business.GET("/expert-inspection-companies/:id/rectification-sheet", expertH.DownloadRectificationSheet)
	// 隐患分类统计 xlsx 导出（仅本企业可导出本企业的统计）。
	business.GET("/hazard-statistics", expertH.ExportHazardStatistics)

	// 集团专家库（bloc-experts）— 三个权限点：list / create / delete。
	business.GET("/bloc-experts", middleware.RequirePermission(r.deps.Enforcer, "bloc-experts.list"), blocExpertH.List)
	business.GET("/bloc-experts-list", middleware.RequirePermission(r.deps.Enforcer, "bloc-experts.list"), blocExpertH.AvailableList)
	business.GET("/bloc-experts/:id", middleware.RequirePermission(r.deps.Enforcer, "bloc-experts.list"), blocExpertH.Get)
	business.POST("/bloc-experts", middleware.RequirePermission(r.deps.Enforcer, "bloc-experts.create"), blocExpertH.Create)
	business.DELETE("/bloc-experts/:id", middleware.RequirePermission(r.deps.Enforcer, "bloc-experts.delete"), blocExpertH.Delete)

	// 待办事项 — 仍是占位实现，Laravel 也尚未接入聚合查询；保持两端契约一致。
	business.GET("/todos", middleware.RequirePermission(r.deps.Enforcer, "todos.list"), todoH.Index)
	business.GET("/todos-list-table", middleware.RequirePermission(r.deps.Enforcer, "todos.list"), todoH.ListTable)

	// 隐患信息（hidden_danger_infos）— 专家检查现场登记的隐患三态流转：登记 → 整改 → 复查/归档。
	// rectify / confirm-rectification / archive 不走权限中间件，由前端按角色控制可见性。
	business.GET("/hidden-danger-infos", middleware.RequirePermission(r.deps.Enforcer, "hidden-danger-infos.list"), hiddenDangerH.List)
	business.GET("/hidden-danger-infos/:id", middleware.RequirePermission(r.deps.Enforcer, "hidden-danger-infos.list"), hiddenDangerH.Get)
	business.POST("/hidden-danger-infos", middleware.RequirePermission(r.deps.Enforcer, "hidden-danger-infos.create"), hiddenDangerH.Create)
	business.PUT("/hidden-danger-infos/:id", middleware.RequirePermission(r.deps.Enforcer, "hidden-danger-infos.update"), hiddenDangerH.Update)
	business.DELETE("/hidden-danger-infos/:id", middleware.RequirePermission(r.deps.Enforcer, "hidden-danger-infos.delete"), hiddenDangerH.Delete)
	business.POST("/hidden-danger-infos/:id/rectify", hiddenDangerH.Rectify)
	business.POST("/hidden-danger-infos/:id/confirm-rectification", hiddenDangerH.ConfirmRectification)
	business.POST("/hidden-danger-infos/:id/archive", hiddenDangerH.Archive)

	// ── 考核评分模块（10 个 controller 合并） ──────────────────────────────
	// 跟 Laravel routes/api/inspections.php 对齐：5 维度 filing/scoring + 模板树 + 集团汇总。
	// 静态路由（/submissions /start-form /score /...）必须先于 /:id 注册避免 path 冲突。

	// 日常检查模板（CRUD）
	business.GET("/daily-check-templates", middleware.RequirePermission(r.deps.Enforcer, "daily-check-templates.list"), assessmentH.ListDailyCheckTemplates)
	business.GET("/daily-check-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "daily-check-templates.list"), assessmentH.GetDailyCheckTemplate)
	business.POST("/daily-check-templates", middleware.RequirePermission(r.deps.Enforcer, "daily-check-templates.create"), assessmentH.CreateDailyCheckTemplate)
	business.PUT("/daily-check-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "daily-check-templates.update"), assessmentH.UpdateDailyCheckTemplate)
	business.DELETE("/daily-check-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "daily-check-templates.delete"), assessmentH.DeleteDailyCheckTemplate)

	// 日常检查（filing：草稿→已提交→已评分）
	business.GET("/daily-checks/submissions", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.ListDailySubmissions)
	business.GET("/daily-checks/start-form", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.StartDailyForm)
	business.GET("/daily-checks/submission", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.ShowDailySubmission)
	business.POST("/daily-checks/save-draft", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.create"), assessmentH.SaveDailyDraft)
	business.POST("/daily-checks/complete", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.create"), assessmentH.CompleteDaily)
	business.POST("/daily-checks/score", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.update"), assessmentH.ScoreDaily)
	business.DELETE("/daily-checks/destroy-form", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.delete"), assessmentH.DestroyDailyForm)
	business.GET("/daily-checks", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.ListDailyChecks)
	business.GET("/daily-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.GetDailyCheck)

	// 系统检查（活动参与及材料上报）
	business.GET("/system-checks/submissions", middleware.RequirePermission(r.deps.Enforcer, "system-checks.list"), assessmentH.ListSystemSubmissions)
	business.POST("/system-checks/compute", middleware.RequirePermission(r.deps.Enforcer, "system-checks.create"), assessmentH.ComputeSystemCheck)
	business.GET("/system-checks", middleware.RequirePermission(r.deps.Enforcer, "system-checks.list"), assessmentH.ListSystemChecks)
	business.GET("/system-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "system-checks.list"), assessmentH.GetSystemCheck)
	business.POST("/system-checks", middleware.RequirePermission(r.deps.Enforcer, "system-checks.create"), assessmentH.CreateSystemCheck)
	business.PUT("/system-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "system-checks.update"), assessmentH.UpdateSystemCheck)
	business.DELETE("/system-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "system-checks.delete"), assessmentH.DeleteSystemCheck)

	// 标准化检查（大中型企业安全生产标准化）
	business.GET("/norm-checks/submissions", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.list"), assessmentH.ListNormSubmissions)
	business.GET("/norm-checks/template-link", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.list"), assessmentH.NormTemplateLink)
	business.POST("/norm-checks/save-score", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.create"), assessmentH.SaveNormScore)
	business.GET("/norm-checks", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.list"), assessmentH.ListNormChecks)
	business.GET("/norm-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.list"), assessmentH.GetNormCheck)
	business.POST("/norm-checks", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.create"), assessmentH.CreateNormCheck)
	business.PUT("/norm-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.update"), assessmentH.UpdateNormCheck)
	business.DELETE("/norm-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "norm-checks.delete"), assessmentH.DeleteNormCheck)

	// 事故检查模板
	business.GET("/accident-check-templates", middleware.RequirePermission(r.deps.Enforcer, "accident-check-templates.list"), assessmentH.ListAccidentCheckTemplates)
	business.GET("/accident-check-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-check-templates.list"), assessmentH.GetAccidentCheckTemplate)
	business.POST("/accident-check-templates", middleware.RequirePermission(r.deps.Enforcer, "accident-check-templates.create"), assessmentH.CreateAccidentCheckTemplate)
	business.PUT("/accident-check-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-check-templates.update"), assessmentH.UpdateAccidentCheckTemplate)
	business.DELETE("/accident-check-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-check-templates.delete"), assessmentH.DeleteAccidentCheckTemplate)

	// 事故检查
	business.GET("/accident-checks/submissions", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.list"), assessmentH.ListAccidentSubmissions)
	business.GET("/accident-checks/start-form", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.list"), assessmentH.StartAccidentForm)
	business.POST("/accident-checks/save-score", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.create"), assessmentH.SaveAccidentScore)
	business.GET("/accident-checks", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.list"), assessmentH.ListAccidentChecks)
	business.GET("/accident-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.list"), assessmentH.GetAccidentCheck)
	business.POST("/accident-checks", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.create"), assessmentH.CreateAccidentCheck)
	business.PUT("/accident-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.update"), assessmentH.UpdateAccidentCheck)
	business.DELETE("/accident-checks/:id", middleware.RequirePermission(r.deps.Enforcer, "accident-checks.delete"), assessmentH.DeleteAccidentCheck)

	// 加分模板
	business.GET("/add-points-templates", middleware.RequirePermission(r.deps.Enforcer, "add-points-templates.list"), assessmentH.ListAddPointsTemplates)
	business.GET("/add-points-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "add-points-templates.list"), assessmentH.GetAddPointsTemplate)
	business.POST("/add-points-templates", middleware.RequirePermission(r.deps.Enforcer, "add-points-templates.create"), assessmentH.CreateAddPointsTemplate)
	business.PUT("/add-points-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "add-points-templates.update"), assessmentH.UpdateAddPointsTemplate)
	business.DELETE("/add-points-templates/:id", middleware.RequirePermission(r.deps.Enforcer, "add-points-templates.delete"), assessmentH.DeleteAddPointsTemplate)

	// 加分项
	business.GET("/add-points/submissions", middleware.RequirePermission(r.deps.Enforcer, "add-points.list"), assessmentH.ListAddPointsSubmissions)
	business.GET("/add-points/start-form", middleware.RequirePermission(r.deps.Enforcer, "add-points.list"), assessmentH.StartAddPointsForm)
	business.POST("/add-points/save-score", middleware.RequirePermission(r.deps.Enforcer, "add-points.create"), assessmentH.SaveAddPointsScore)
	business.GET("/add-points", middleware.RequirePermission(r.deps.Enforcer, "add-points.list"), assessmentH.ListAddPoints)
	business.GET("/add-points/:id", middleware.RequirePermission(r.deps.Enforcer, "add-points.list"), assessmentH.GetAddPoint)
	business.POST("/add-points", middleware.RequirePermission(r.deps.Enforcer, "add-points.create"), assessmentH.CreateAddPoint)
	business.PUT("/add-points/:id", middleware.RequirePermission(r.deps.Enforcer, "add-points.update"), assessmentH.UpdateAddPoint)
	business.DELETE("/add-points/:id", middleware.RequirePermission(r.deps.Enforcer, "add-points.delete"), assessmentH.DeleteAddPoint)

	// 考核汇总（6 维度合算）
	business.GET("/assessment-summary", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.Summary)
	business.GET("/assessment-summary/export", middleware.RequirePermission(r.deps.Enforcer, "daily-checks.list"), assessmentH.ExportSummary)

	// 考核模板年度管理（年度并集 + 一键拷贝）
	business.GET("/assessment-templates/years", assessmentH.AssessmentYears)
	business.POST("/assessment-templates/copy", middleware.RequirePermission(r.deps.Enforcer, "daily-check-templates.create"), assessmentH.CopyAssessmentTemplates)

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
