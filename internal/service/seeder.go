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
	"personnels.list", "personnels.create", "personnels.update",

	// Safety ledger (MVP)
	"special-equipment.list", "special-equipment.create", "special-equipment.update", "special-equipment.delete",
	"special-equipment-staff.list", "special-equipment-staff.create", "special-equipment-staff.update", "special-equipment-staff.delete",
	"dangerous-equipment.list", "dangerous-equipment.create", "dangerous-equipment.update", "dangerous-equipment.delete",
	"hazardous-chemicals.list", "hazardous-chemicals.create", "hazardous-chemicals.update", "hazardous-chemicals.delete",

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
			"base-data.view",
			"personnels.list", "personnels.create", "personnels.update",
			"special-equipment.list", "special-equipment.create", "special-equipment.update", "special-equipment.delete",
			"special-equipment-staff.list", "special-equipment-staff.create", "special-equipment-staff.update", "special-equipment-staff.delete",
			"dangerous-equipment.list", "dangerous-equipment.create", "dangerous-equipment.update", "dangerous-equipment.delete",
			"hazardous-chemicals.list", "hazardous-chemicals.create", "hazardous-chemicals.update", "hazardous-chemicals.delete",
			"upload.file", "upload.oss",
		},
	},
	{
		Name: "viewer",
		Permissions: []string{
			"users.list", "companies.list", "companies.tree", "base-data.view",
			"special-equipment.list", "dangerous-equipment.list", "hazardous-chemicals.list",
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

func seedBaseData(ctx context.Context, db *gorm.DB) error {
	for _, title := range SeededSpecialTypes {
		var existing model.SpecialType
		err := db.WithContext(ctx).Where("title = ?", title).First(&existing).Error
		if err == nil {
			continue
		}
		if err := db.WithContext(ctx).Create(&model.SpecialType{Title: title}).Error; err != nil {
			return err
		}
	}
	for _, title := range SeededEnterpriseTypes {
		var existing model.EnterpriseType
		err := db.WithContext(ctx).Where("title = ?", title).First(&existing).Error
		if err == nil {
			continue
		}
		if err := db.WithContext(ctx).Create(&model.EnterpriseType{Title: title}).Error; err != nil {
			return err
		}
	}
	for _, title := range SeededOperationalTypes {
		var existing model.OperationalType
		err := db.WithContext(ctx).Where("title = ?", title).First(&existing).Error
		if err == nil {
			continue
		}
		if err := db.WithContext(ctx).Create(&model.OperationalType{Title: title}).Error; err != nil {
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
		Username:               cred.Username,
		Password:               string(hashed),
		Nickname:               cred.Nickname,
		CompanyID:              nil,
		Status:                 UserStatusEnable,
		IsPlatformAdmin:        true,
		MustChangeCredentials:  true, // force first-login setup
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
