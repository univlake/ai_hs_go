package service

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"ai_hs_go/internal/model"
	"ai_hs_go/internal/rbac"
)

// CorePermissions enumerates the MVP-scope permission strings. The full Laravel
// catalog has 322 entries; this seed covers what is needed for Phases 1–7.
// Subsequent migration phases (per /Users/zero/.claude/plans/...) extend this.
var CorePermissions = []string{
	// Auth & users
	"users.list", "users.create", "users.update", "users.delete",
	"roles.list", "roles.create", "roles.update", "roles.delete", "roles.assign",
	"permissions.list",

	// Companies / tenants
	"companies.list", "companies.create", "companies.update", "companies.delete",
	"companies.tree", "companies.review",
	"tenants.list", "tenants.create", "tenants.update", "tenants.delete",

	// Base data
	"base-data.view", "base-data.manage",
	"checklist-libraries.list", "checklist-libraries.create", "checklist-libraries.update", "checklist-libraries.delete",

	// Safety archive (5 sub-modules)
	"safety-architectures.list", "safety-architectures.create", "safety-architectures.update",
	"filing-permissions.list", "filing-permissions.create", "filing-permissions.update", "filing-permissions.delete",
	"buildings.list", "buildings.create", "buildings.update", "buildings.delete",
	"security-investments.list", "security-investments.create", "security-investments.update", "security-investments.delete",
	"personnels.list", "personnels.create", "personnels.update",

	// Safety archive sub-tabs (人员管理子页 + 安全资质子页 + 安全投入子页)
	"certificate-statuses.list", "certificate-statuses.create", "certificate-statuses.update",
	"safety-management.list", "safety-management.create", "safety-management.update", "safety-management.delete",
	"special-equipment-staff.list", "special-equipment-staff.create", "special-equipment-staff.update", "special-equipment-staff.delete",
	"elimination-control-staff.list", "elimination-control-staff.create", "elimination-control-staff.update", "elimination-control-staff.delete",
	"production-standardizations.list", "production-standardizations.create", "production-standardizations.update",
	"occupational-health-safeties.list", "occupational-health-safeties.create", "occupational-health-safeties.update",
	"funding-executions.list", "funding-executions.create", "funding-executions.update", "funding-executions.delete",

	// Safety ledger (MVP)
	"special-equipment.list", "special-equipment.create", "special-equipment.update", "special-equipment.delete",
	"special-equipment-staff.list", "special-equipment-staff.create", "special-equipment-staff.update", "special-equipment-staff.delete",
	"dangerous-equipment.list", "dangerous-equipment.create", "dangerous-equipment.update", "dangerous-equipment.delete",
	"hazardous-chemicals.list", "hazardous-chemicals.create", "hazardous-chemicals.update", "hazardous-chemicals.delete",
	"major-hazards.list", "major-hazards.create", "major-hazards.update", "major-hazards.delete",

	// 工艺信息（重点监管 / 较大危险 / 有限空间辨识）
	"major-hazardous-processes.list", "major-hazardous-processes.create", "major-hazardous-processes.update", "major-hazardous-processes.delete",
	"hazardous-processes.list", "hazardous-processes.create", "hazardous-processes.update", "hazardous-processes.delete",
	"space-works.list", "space-works.create", "space-works.update", "space-works.delete",

	// 第三方管理（设计单位 / 评价单位 / 技术服务单位）
	"design-units.list", "design-units.create", "design-units.update", "design-units.delete",
	"evaluation-units.list", "evaluation-units.create", "evaluation-units.update", "evaluation-units.delete",
	"skill-units.list", "skill-units.create", "skill-units.update", "skill-units.delete",

	// 会议培训（内部会议 / 内部、外部、相关方培训）
	"internal-meetings.list", "internal-meetings.create", "internal-meetings.update", "internal-meetings.delete",
	"internal-trainings.list", "internal-trainings.create", "internal-trainings.update", "internal-trainings.delete",
	"external-trainings.list", "external-trainings.create", "external-trainings.update", "external-trainings.delete",
	"other-trainings.list", "other-trainings.create", "other-trainings.update", "other-trainings.delete",

	// File upload
	"upload.file", "upload.oss",
}

// Roles seeded by default. platform_admin is special — middleware bypasses
// Casbin for is_platform_admin=true users, so this role is mostly cosmetic.
var SeededRoles = []SeededRole{
	{
		Name:        "platform_admin",
		Permissions: nil, // bypass — gets all anyway
	},
	{
		Name:        "tenant_root_admin",
		Permissions: CorePermissions, // every MVP permission, scope-bounded by middleware
	},
	{
		Name: "tenant_admin",
		Permissions: []string{
			"users.list", "users.create", "users.update",
			"companies.list", "companies.update", "companies.tree",
			"base-data.view", "base-data.manage",
			"checklist-libraries.list", "checklist-libraries.create", "checklist-libraries.update", "checklist-libraries.delete",
			"safety-architectures.list", "safety-architectures.create", "safety-architectures.update",
			"filing-permissions.list", "filing-permissions.create", "filing-permissions.update", "filing-permissions.delete",
			"buildings.list", "buildings.create", "buildings.update", "buildings.delete",
			"security-investments.list", "security-investments.create", "security-investments.update", "security-investments.delete",
			"personnels.list", "personnels.create", "personnels.update",
			"certificate-statuses.list", "certificate-statuses.create", "certificate-statuses.update",
			"safety-management.list", "safety-management.create", "safety-management.update", "safety-management.delete",
			"special-equipment-staff.list", "special-equipment-staff.create", "special-equipment-staff.update", "special-equipment-staff.delete",
			"elimination-control-staff.list", "elimination-control-staff.create", "elimination-control-staff.update", "elimination-control-staff.delete",
			"production-standardizations.list", "production-standardizations.create", "production-standardizations.update",
			"occupational-health-safeties.list", "occupational-health-safeties.create", "occupational-health-safeties.update",
			"funding-executions.list", "funding-executions.create", "funding-executions.update", "funding-executions.delete",
			"special-equipment.list", "special-equipment.create", "special-equipment.update", "special-equipment.delete",
			"special-equipment-staff.list", "special-equipment-staff.create", "special-equipment-staff.update", "special-equipment-staff.delete",
			"dangerous-equipment.list", "dangerous-equipment.create", "dangerous-equipment.update", "dangerous-equipment.delete",
			"hazardous-chemicals.list", "hazardous-chemicals.create", "hazardous-chemicals.update", "hazardous-chemicals.delete",
			"major-hazards.list", "major-hazards.create", "major-hazards.update", "major-hazards.delete",
			"major-hazardous-processes.list", "major-hazardous-processes.create", "major-hazardous-processes.update", "major-hazardous-processes.delete",
			"hazardous-processes.list", "hazardous-processes.create", "hazardous-processes.update", "hazardous-processes.delete",
			"space-works.list", "space-works.create", "space-works.update", "space-works.delete",
			"design-units.list", "design-units.create", "design-units.update", "design-units.delete",
			"evaluation-units.list", "evaluation-units.create", "evaluation-units.update", "evaluation-units.delete",
			"skill-units.list", "skill-units.create", "skill-units.update", "skill-units.delete",
			"internal-meetings.list", "internal-meetings.create", "internal-meetings.update", "internal-meetings.delete",
			"internal-trainings.list", "internal-trainings.create", "internal-trainings.update", "internal-trainings.delete",
			"external-trainings.list", "external-trainings.create", "external-trainings.update", "external-trainings.delete",
			"other-trainings.list", "other-trainings.create", "other-trainings.update", "other-trainings.delete",
			"upload.file", "upload.oss",
		},
	},
	{
		Name: "viewer",
		Permissions: []string{
			"users.list", "companies.list", "companies.tree", "base-data.view",
			"checklist-libraries.list",
			"safety-architectures.list", "filing-permissions.list", "buildings.list",
			"security-investments.list", "personnels.list",
			"certificate-statuses.list", "safety-management.list",
			"special-equipment-staff.list", "elimination-control-staff.list",
			"production-standardizations.list", "occupational-health-safeties.list",
			"funding-executions.list",
			"special-equipment.list", "dangerous-equipment.list", "hazardous-chemicals.list",
			"major-hazards.list",
			"major-hazardous-processes.list", "hazardous-processes.list", "space-works.list",
			"design-units.list", "evaluation-units.list", "skill-units.list",
			"internal-meetings.list", "internal-trainings.list", "external-trainings.list", "other-trainings.list",
		},
	},
}

type SeededRole struct {
	Name        string
	Permissions []string
}

// SuperAdminCredentials is what the seeder writes for the platform super admin.
type SuperAdminCredentials struct {
	Username string
	Password string // plaintext; hashed before insert
	Nickname string
}

// Seed populates roles, permissions, casbin policy, the platform super
// admin user, and base-data dictionaries (special_types …). Idempotent —
// safe to re-run.
func Seed(ctx context.Context, db *gorm.DB, enforcer *casbin.Enforcer, superAdmin SuperAdminCredentials) error {
	if err := seedPermissions(ctx, db); err != nil {
		return fmt.Errorf("seed permissions: %w", err)
	}
	if err := seedRoles(ctx, db, enforcer); err != nil {
		return fmt.Errorf("seed roles: %w", err)
	}
	if err := seedSuperAdmin(ctx, db, enforcer, superAdmin); err != nil {
		return fmt.Errorf("seed super admin: %w", err)
	}
	if err := seedBaseData(ctx, db); err != nil {
		return fmt.Errorf("seed base data: %w", err)
	}
	if err := enforcer.SavePolicy(); err != nil {
		return fmt.Errorf("save policy: %w", err)
	}
	return nil
}

// SpecialTypes match the Laravel project's seeder so frontend specs that
// hard-code spec-field configs by id (compressor / pipe / lift …) keep
// working unchanged.
var SeededSpecialTypes = []string{
	"锅炉",
	"压力容器(含气瓶)",
	"压力管道",
	"电梯",
	"起重机械",
	"(厂)场内专用车辆",
}

// SeededEnterpriseTypes / SeededOperationalTypes mirror Laravel's BaseDataSeeder
// so the enterprise create / edit dropdowns aren't empty out of the box.
var SeededEnterpriseTypes = []string{"机械", "化工", "商贸租赁", "其他"}
var SeededOperationalTypes = []string{"生产", "停工", "改造", "搬迁"}

// New base-data dictionary seeds (mirrors Laravel BaseDataSeeder + the
// RegulationLibraryTypeSeeder split file). The Laravel project ships these
// with no `sort` column for most tables; we keep sort=0 to preserve the
// natural insertion order, which the frontend's `ORDER BY sort, id` query
// then renders top-to-bottom in the same sequence as Laravel.
var SeededRecordTypes = []string{
	"安全生产许可证", "危化品安全使用许可证", "危化品登记证",
	"危化品使用备案", "重大危险源备案", "易制毒生产备案",
	"应急预案备案有效", "职业健康备案",
}

var SeededProcessTypes = []string{
	"高温高压（压力大于等于0.1Mpa，温度高于及等于100度）",
	"铸造工艺", "冲剪工艺", "吊装", "焊接",
	"压力容器操作", "喷涂工艺", "产生可燃爆粉尘的工艺",
	"产生有毒有害物质的工艺", "使用甲乙类危险化学品的工艺",
	"涉及涉氨制冷工艺", "其他",
}

var SeededSpaceWorkTypes = []string{
	"坑", "沟", "槽", "池", "井", "管道", "罐", "箱",
	"炉", "室", "仓(斗)", "塔(釜)", "其他", "反应器", "锅筒",
}

var SeededSafetyManagementTypes = []string{
	"安全负责人", "安管员", "生产管理", "设备管理",
	"质量管理", "人事管理", "其他管理",
}

var SeededDangerousCraftTypes = []string{
	"光气及光气化工艺", "电解工艺", "氯化工艺", "硝化工艺",
	"合成氨工艺", "裂解(裂化)工艺", "氟化工艺", "加氢工艺",
	"重氮化工艺", "氧化工艺", "过氧化工艺", "氨基化工艺",
	"磺化工艺", "聚合工艺", "烷基化工艺", "新型煤化工艺",
	"电石生产工艺", "偶氮化工艺",
}

// SeededHazards — 危险有害因素初始项。confined-space 页面里多选用。
var SeededHazards = []string{
	"易燃易爆气体或者其蒸汽",
	"有毒气体或者其蒸汽",
	"缺氧或富氧",
	"坍塌",
	"机械伤害",
	"酸碱腐蚀性化学品或者其蒸汽",
	"易挥发的危险化学品",
	"通风不良",
	"交叉作业",
	"高热高湿",
	"物体打击",
	"高处坠落",
	"触电",
	"灼烫",
}

// RegulationLibraryType seeds describe a 2-level tree: 5 top-level
// categories + 9 sub-items under 安全生产部门规章 + 15 sub-items under
// 标准规范 = 29 rows total. ParentTitle is matched against the top-level
// row's title to look up pid at insert time. AllName is set verbatim from
// the catalog the user provided (parent（child）format) — that same
// format is what the service's computeAllName produces for any
// user-created descendants going forward.
type seededRegLibType struct {
	Title       string
	ParentTitle string // "" for top-level
	AllName     string
}

var SeededRegulationLibraryTypes = []seededRegLibType{
	// Top-level
	{"安全生产部门规章", "", "安全生产部门规章"},
	{"安全生产法律", "", "安全生产法律"},
	{"安全生产行政法规", "", "安全生产行政法规"},
	{"标准规范", "", "标准规范"},
	{"地方管理条例", "", "地方管理条例"},

	// Children of 安全生产部门规章
	{"劳动防护用品", "安全生产部门规章", "安全生产部门规章（劳动防护用品）"},
	{"劳动卫生", "安全生产部门规章", "安全生产部门规章（劳动卫生）"},
	{"培训考核", "安全生产部门规章", "安全生产部门规章（培训考核）"},
	{"三同时", "安全生产部门规章", "安全生产部门规章（三同时）"},
	{"伤亡事故调查和处理", "安全生产部门规章", "安全生产部门规章（伤亡事故调查和处理）"},
	{"特种设备", "安全生产部门规章", "安全生产部门规章（特种设备）"},
	{"危险化学品", "安全生产部门规章", "安全生产部门规章（危险化学品）"},
	{"消防管理", "安全生产部门规章", "安全生产部门规章（消防管理）"},
	{"综合类", "安全生产部门规章", "安全生产部门规章（综合类）"},

	// Children of 标准规范
	{"安全技术综合", "标准规范", "标准规范（安全技术综合）"},
	{"电气方面", "标准规范", "标准规范（电气方面）"},
	{"防火防爆安全技术", "标准规范", "标准规范（防火防爆安全技术）"},
	{"机械方面", "标准规范", "标准规范（机械方面）"},
	{"交通", "标准规范", "标准规范（交通）"},
	{"劳动防护用品", "标准规范", "标准规范（劳动防护用品）"},
	{"劳动卫生", "标准规范", "标准规范（劳动卫生）"},
	{"喷涂企业", "标准规范", "标准规范（喷涂企业）"},
	{"涉氨企业", "标准规范", "标准规范（涉氨企业）"},
	{"涉燃爆粉尘", "标准规范", "标准规范（涉燃爆粉尘）"},
	{"生产设备安全技术", "标准规范", "标准规范（生产设备安全技术）"},
	{"危险化学品", "标准规范", "标准规范（危险化学品）"},
	{"消防综合", "标准规范", "标准规范（消防综合）"},
	{"冶炼企业", "标准规范", "标准规范（冶炼企业）"},
	{"有限空间", "标准规范", "标准规范（有限空间）"},
}

// seedDictTable upserts a slice of {title} rows into a base-data table.
// Idempotent — uses a SELECT-then-INSERT loop instead of ON CONFLICT so
// reseeding doesn't reset auto-increment ids that other rows already point
// at (e.g. SafetyEquipment.special_type_id).
func seedDictTable(ctx context.Context, db *gorm.DB, table string, titles []string, withSort bool) error {
	for i, title := range titles {
		var count int64
		if err := db.WithContext(ctx).Table(table).
			Where("title = ? AND deleted_at IS NULL", title).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		row := map[string]any{
			"title":      title,
			"created_at": gorm.Expr("NOW()"),
			"updated_at": gorm.Expr("NOW()"),
		}
		if withSort {
			row["sort"] = i
		}
		if err := db.WithContext(ctx).Table(table).Create(&row).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedBaseData(ctx context.Context, db *gorm.DB) error {
	if err := seedDictTable(ctx, db, "special_types", SeededSpecialTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "enterprise_types", SeededEnterpriseTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "operational_types", SeededOperationalTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "record_types", SeededRecordTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "process_types", SeededProcessTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "space_work_types", SeededSpaceWorkTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "safety_management_types", SeededSafetyManagementTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "dangerous_craft_types", SeededDangerousCraftTypes, false); err != nil {
		return err
	}
	if err := seedDictTable(ctx, db, "hazards", SeededHazards, false); err != nil {
		return err
	}
	if err := seedRegulationLibraryTypes(ctx, db); err != nil {
		return err
	}
	return nil
}

// seedRegulationLibraryTypes plants the 2-level catalog described by
// SeededRegulationLibraryTypes.
//
// Two passes: top-level rows go in first so child rows have their pid
// available via a title→id lookup. Each row is upserted (insert when
// missing, otherwise leave alone) keyed on (title, pid) so reseeding is
// idempotent. After all 29 entries are in, a cleanup pass hard-deletes
// any other regulation_library_types rows so the table mirrors the
// catalog exactly — important here because earlier seed runs put items
// like "劳动防护用品" at pid=0 that now belong under pid=parent.
func seedRegulationLibraryTypes(ctx context.Context, db *gorm.DB) error {
	type rlt struct {
		ID    int64
		Title string
		Pid   int64
	}

	upsert := func(title string, pid int64, allName string, sort int) error {
		var count int64
		if err := db.WithContext(ctx).Table("regulation_library_types").
			Where("title = ? AND pid = ? AND deleted_at IS NULL", title, pid).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			// Refresh all_name / sort + force company_id back to 0 so
			// rows accidentally tagged to a tenant by an earlier UI test
			// drop back to the global dictionary state.
			return db.WithContext(ctx).Table("regulation_library_types").
				Where("title = ? AND pid = ? AND deleted_at IS NULL", title, pid).
				Updates(map[string]any{
					"all_name":   allName,
					"sort":       sort,
					"company_id": 0,
					"updated_at": gorm.Expr("NOW()"),
				}).Error
		}
		row := map[string]any{
			"title":      title,
			"pid":        pid,
			"all_name":   allName,
			"sort":       sort,
			"created_at": gorm.Expr("NOW()"),
			"updated_at": gorm.Expr("NOW()"),
		}
		return db.WithContext(ctx).Table("regulation_library_types").Create(&row).Error
	}

	// Pass 1: top-level rows. sort = index across the whole list (Laravel
	// preserves the data-file order via foreach key).
	for i, e := range SeededRegulationLibraryTypes {
		if e.ParentTitle != "" {
			continue
		}
		if err := upsert(e.Title, 0, e.AllName, i); err != nil {
			return err
		}
	}

	// Build a title→id map for top-level rows (parent lookup). Children
	// expect to land under exactly one parent, so duplicate top-level
	// titles aren't supported in the seed catalog (and the data above
	// indeed has none).
	var roots []rlt
	if err := db.WithContext(ctx).Table("regulation_library_types").
		Select("id, title, pid").
		Where("pid = 0 AND deleted_at IS NULL").
		Find(&roots).Error; err != nil {
		return err
	}
	rootIDByTitle := make(map[string]int64, len(roots))
	for _, r := range roots {
		rootIDByTitle[r.Title] = r.ID
	}

	// Pass 2: children. Skip silently if the named parent is missing
	// (would only happen if the data above is internally inconsistent).
	for i, e := range SeededRegulationLibraryTypes {
		if e.ParentTitle == "" {
			continue
		}
		pid, ok := rootIDByTitle[e.ParentTitle]
		if !ok {
			continue
		}
		if err := upsert(e.Title, pid, e.AllName, i); err != nil {
			return err
		}
	}

	// Cleanup: collect (title, pid) keys present in the seed and
	// hard-delete every row not matching. Earlier seeds left orphans at
	// pid=0 with titles that now belong as children — purge them so the
	// list reflects the new catalog exactly.
	wanted := make(map[string]struct{}, len(SeededRegulationLibraryTypes))
	for _, e := range SeededRegulationLibraryTypes {
		var pid int64
		if e.ParentTitle != "" {
			pid = rootIDByTitle[e.ParentTitle]
		}
		wanted[fmt.Sprintf("%d|%s", pid, e.Title)] = struct{}{}
	}

	var existing []rlt
	if err := db.WithContext(ctx).Table("regulation_library_types").
		Select("id, title, pid").
		Where("deleted_at IS NULL").
		Find(&existing).Error; err != nil {
		return err
	}
	var orphanIDs []int64
	for _, row := range existing {
		key := fmt.Sprintf("%d|%s", row.Pid, row.Title)
		if _, ok := wanted[key]; !ok {
			orphanIDs = append(orphanIDs, row.ID)
		}
	}
	if len(orphanIDs) > 0 {
		// Hard delete (Unscoped) — soft-deleting would just leave the rows
		// blocking future re-seeds via the (title,pid) uniqueness-by-
		// existence check above.
		if err := db.WithContext(ctx).Table("regulation_library_types").
			Unscoped().
			Where("id IN ?", orphanIDs).
			Delete(nil).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedPermissions(ctx context.Context, db *gorm.DB) error {
	for _, name := range CorePermissions {
		var existing model.Permission
		err := db.WithContext(ctx).Where("name = ? AND guard_name = ?", name, "api").First(&existing).Error
		if err == nil {
			continue
		}
		if err := db.WithContext(ctx).Create(&model.Permission{Name: name, GuardName: "api"}).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedRoles(ctx context.Context, db *gorm.DB, enforcer *casbin.Enforcer) error {
	for _, sr := range SeededRoles {
		var role model.Role
		err := db.WithContext(ctx).Where("name = ? AND guard_name = ?", sr.Name, "api").First(&role).Error
		if err != nil {
			role = model.Role{Name: sr.Name, GuardName: "api"}
			if err := db.WithContext(ctx).Create(&role).Error; err != nil {
				return err
			}
		}

		// Wipe and re-add permissions for this role to keep it idempotent.
		if err := db.WithContext(ctx).Where("role_id = ?", role.ID).Delete(&model.RoleHasPermission{}).Error; err != nil {
			return err
		}
		if _, err := enforcer.RemoveFilteredPolicy(0, rbac.SubjectForRole(role.Name)); err != nil {
			return err
		}

		if len(sr.Permissions) == 0 {
			continue
		}

		var perms []model.Permission
		if err := db.WithContext(ctx).Where("name IN ? AND guard_name = ?", sr.Permissions, "api").Find(&perms).Error; err != nil {
			return err
		}
		links := make([]model.RoleHasPermission, len(perms))
		for i, p := range perms {
			links[i] = model.RoleHasPermission{RoleID: role.ID, PermissionID: p.ID}
		}
		if err := db.WithContext(ctx).Create(&links).Error; err != nil {
			return err
		}

		policies := make([][]string, 0, len(perms))
		roleSub := rbac.SubjectForRole(role.Name)
		for _, p := range perms {
			resource, action := splitPermissionName(p.Name)
			policies = append(policies, []string{roleSub, resource, action})
		}
		if _, err := enforcer.AddPolicies(policies); err != nil {
			return err
		}
	}
	return nil
}

func seedSuperAdmin(ctx context.Context, db *gorm.DB, enforcer *casbin.Enforcer, cred SuperAdminCredentials) error {
	if cred.Username == "" {
		cred.Username = "superadmin"
	}
	if cred.Nickname == "" {
		cred.Nickname = "平台管理员"
	}

	var existing model.User
	err := db.WithContext(ctx).Where("username = ?", cred.Username).First(&existing).Error
	if err == nil {
		// keep existing — never overwrite a deployed super admin.
		return ensureSuperAdminBindings(ctx, db, enforcer, existing.ID)
	}

	password := cred.Password
	if password == "" {
		password = "Aifix@2026"
	}
	hashed, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err2 != nil {
		return err2
	}
	user := model.User{
		Username:              cred.Username,
		Password:              string(hashed),
		Nickname:              cred.Nickname,
		CompanyID:             nil,
		Status:                UserStatusEnable,
		IsPlatformAdmin:       true,
		MustChangeCredentials: true, // force first-login setup
	}
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return ensureSuperAdminBindings(ctx, db, enforcer, user.ID)
}

func ensureSuperAdminBindings(ctx context.Context, db *gorm.DB, enforcer *casbin.Enforcer, userID int64) error {
	var role model.Role
	if err := db.WithContext(ctx).Where("name = ? AND guard_name = ?", "platform_admin", "api").First(&role).Error; err != nil {
		return err
	}
	if err := db.WithContext(ctx).
		Where("role_id = ? AND model_type = ? AND model_id = ?", role.ID, UserModelType, userID).
		FirstOrCreate(&model.ModelHasRole{RoleID: role.ID, ModelType: UserModelType, ModelID: userID}).Error; err != nil {
		return err
	}
	if _, err := enforcer.AddGroupingPolicy(rbac.SubjectForUser(userID), rbac.SubjectForRole(role.Name)); err != nil {
		return err
	}
	return nil
}
