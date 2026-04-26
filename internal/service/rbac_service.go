package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/rbac"
)

const (
	UserModelType = "App\\Models\\User"
)

type RBACService struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
}

func NewRBACService(db *gorm.DB, enforcer *casbin.Enforcer) *RBACService {
	return &RBACService{db: db, enforcer: enforcer}
}

// ListPermissions returns all known permission names.
func (s *RBACService) ListPermissions(ctx context.Context) ([]model.Permission, error) {
	var perms []model.Permission
	if err := s.db.WithContext(ctx).Order("name").Find(&perms).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "查询权限失败", err)
	}
	return perms, nil
}

// builtinRoles cannot be deleted or renamed. Names align with the seeder.
var builtinRoles = map[string]bool{
	"platform_admin":    true,
	"tenant_root_admin": true,
	"tenant_admin":      true,
	"viewer":            true,
}

// RoleWithPermissions bundles a role and its assigned permission names — the
// shape the frontend table renders.
type RoleWithPermissions struct {
	Role        model.Role `json:"role"`
	Permissions []string   `json:"permissions"`
}

// ListRoles returns all defined roles, each with its assigned permission names.
func (s *RBACService) ListRoles(ctx context.Context) ([]RoleWithPermissions, error) {
	var roles []model.Role
	if err := s.db.WithContext(ctx).Order("name").Find(&roles).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "查询角色失败", err)
	}
	out := make([]RoleWithPermissions, len(roles))
	for i, r := range roles {
		var names []string
		if err := s.db.WithContext(ctx).Table("permissions").
			Select("permissions.name").
			Joins("JOIN role_has_permissions rhp ON rhp.permission_id = permissions.id").
			Where("rhp.role_id = ?", r.ID).
			Pluck("permissions.name", &names).Error; err != nil {
			return nil, errs.Wrap(errs.CodeInternal, "查询角色权限失败", err)
		}
		out[i] = RoleWithPermissions{Role: r, Permissions: names}
	}
	return out, nil
}

// GetRole returns one role + its permission names.
func (s *RBACService) GetRole(ctx context.Context, id int64) (*RoleWithPermissions, error) {
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, errs.New(errs.CodeNotFound, "角色不存在")
	}
	var names []string
	if err := s.db.WithContext(ctx).Table("permissions").
		Select("permissions.name").
		Joins("JOIN role_has_permissions rhp ON rhp.permission_id = permissions.id").
		Where("rhp.role_id = ?", role.ID).
		Pluck("permissions.name", &names).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "查询角色权限失败", err)
	}
	return &RoleWithPermissions{Role: role, Permissions: names}, nil
}

func (s *RBACService) CreateRole(ctx context.Context, name string, permissions []string) (*RoleWithPermissions, error) {
	role := &model.Role{Name: strings.TrimSpace(name), GuardName: "api"}
	if role.Name == "" {
		return nil, errs.New(errs.CodeUnprocessable, "角色名称不能为空")
	}
	if err := s.db.WithContext(ctx).Create(role).Error; err != nil {
		return nil, errs.Wrap(errs.CodeConflict, "创建角色失败（可能与现有同名）", err)
	}
	if len(permissions) > 0 {
		if err := s.AssignPermissionsToRole(ctx, role.ID, permissions); err != nil {
			return nil, err
		}
	}
	return s.GetRole(ctx, role.ID)
}

// UpdateRole renames and/or syncs permissions for an existing role. Built-in
// roles cannot be renamed but can have their permissions updated.
type UpdateRoleInput struct {
	Name        *string
	Permissions []string // nil means "leave alone"; empty slice means "clear"
}

func (s *RBACService) UpdateRole(ctx context.Context, id int64, in UpdateRoleInput) (*RoleWithPermissions, error) {
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, errs.New(errs.CodeNotFound, "角色不存在")
	}
	if in.Name != nil {
		newName := strings.TrimSpace(*in.Name)
		if newName == "" {
			return nil, errs.New(errs.CodeUnprocessable, "角色名称不能为空")
		}
		if newName != role.Name {
			if builtinRoles[role.Name] {
				return nil, errs.New(errs.CodeForbidden, "系统内置角色名称不可修改")
			}
			oldSub := rbac.SubjectForRole(role.Name)
			if err := s.db.WithContext(ctx).Model(&role).Update("name", newName).Error; err != nil {
				return nil, errs.Wrap(errs.CodeConflict, "重命名失败", err)
			}
			// Migrate casbin policies + groupings to the new role subject.
			policies, _ := s.enforcer.GetFilteredPolicy(0, oldSub)
			groupings, _ := s.enforcer.GetFilteredGroupingPolicy(1, oldSub)
			if _, err := s.enforcer.RemoveFilteredPolicy(0, oldSub); err != nil {
				return nil, errs.Wrap(errs.CodeInternal, "清除旧策略失败", err)
			}
			if _, err := s.enforcer.RemoveFilteredGroupingPolicy(1, oldSub); err != nil {
				return nil, errs.Wrap(errs.CodeInternal, "清除旧分组失败", err)
			}
			newSub := rbac.SubjectForRole(newName)
			for _, p := range policies {
				if len(p) >= 3 {
					_, _ = s.enforcer.AddPolicy(newSub, p[1], p[2])
				}
			}
			for _, g := range groupings {
				if len(g) >= 2 {
					_, _ = s.enforcer.AddGroupingPolicy(g[0], newSub)
				}
			}
			role.Name = newName
		}
	}
	if in.Permissions != nil {
		if err := s.AssignPermissionsToRole(ctx, role.ID, in.Permissions); err != nil {
			return nil, err
		}
	}
	return s.GetRole(ctx, role.ID)
}

func (s *RBACService) DeleteRole(ctx context.Context, id int64) error {
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return errs.New(errs.CodeNotFound, "角色不存在")
	}
	if builtinRoles[role.Name] {
		return errs.New(errs.CodeForbidden, "系统内置角色不可删除")
	}
	// Refuse delete when users still hold this role.
	var inUse int64
	if err := s.db.WithContext(ctx).Model(&model.ModelHasRole{}).
		Where("role_id = ?", role.ID).Count(&inUse).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "查询角色使用情况失败", err)
	}
	if inUse > 0 {
		return errs.New(errs.CodeConflict, "该角色下仍有用户，请先移除这些用户的此角色")
	}
	if err := s.db.WithContext(ctx).Delete(&role).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "删除角色失败", err)
	}
	if _, err := s.enforcer.RemoveFilteredGroupingPolicy(1, rbac.SubjectForRole(role.Name)); err != nil {
		return errs.Wrap(errs.CodeInternal, "清理角色策略失败", err)
	}
	if _, err := s.enforcer.RemoveFilteredPolicy(0, rbac.SubjectForRole(role.Name)); err != nil {
		return errs.Wrap(errs.CodeInternal, "清理权限策略失败", err)
	}
	return nil
}

// AssignPermissionsToRole replaces the role's permissions with the given list.
func (s *RBACService) AssignPermissionsToRole(ctx context.Context, roleID int64, permissionNames []string) error {
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, roleID).Error; err != nil {
		return errs.New(errs.CodeNotFound, "角色不存在")
	}

	var permissions []model.Permission
	if err := s.db.WithContext(ctx).Where("name IN ?", permissionNames).Find(&permissions).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "查询权限失败", err)
	}
	if len(permissions) != len(permissionNames) {
		return errs.New(errs.CodeUnprocessable, "存在未知权限")
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", role.ID).Delete(&model.RoleHasPermission{}).Error; err != nil {
			return err
		}
		links := make([]model.RoleHasPermission, len(permissions))
		for i, p := range permissions {
			links[i] = model.RoleHasPermission{RoleID: role.ID, PermissionID: p.ID}
		}
		if len(links) > 0 {
			if err := tx.Create(&links).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errs.Wrap(errs.CodeInternal, "更新角色权限失败", err)
	}

	// resync casbin: remove all p rules where v0=role and re-add
	roleSub := rbac.SubjectForRole(role.Name)
	if _, err := s.enforcer.RemoveFilteredPolicy(0, roleSub); err != nil {
		return errs.Wrap(errs.CodeInternal, "清除旧策略失败", err)
	}
	policies := make([][]string, 0, len(permissions))
	for _, p := range permissions {
		resource, action := splitPermissionName(p.Name)
		policies = append(policies, []string{roleSub, resource, action})
	}
	if len(policies) > 0 {
		if _, err := s.enforcer.AddPolicies(policies); err != nil {
			return errs.Wrap(errs.CodeInternal, "写入策略失败", err)
		}
	}
	return nil
}

// AssignSingleRoleToUser is a convenience for the common 1-user-1-role flow,
// equivalent to AssignRolesToUser with a single-element slice.
func (s *RBACService) AssignSingleRoleToUser(ctx context.Context, userID int64, roleName string) error {
	return s.AssignRolesToUser(ctx, userID, []string{roleName})
}

// AssignRolesToUser replaces the user's roles with the given list.
func (s *RBACService) AssignRolesToUser(ctx context.Context, userID int64, roleNames []string) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return errs.New(errs.CodeNotFound, "用户不存在")
	}
	var roles []model.Role
	if len(roleNames) > 0 {
		if err := s.db.WithContext(ctx).Where("name IN ?", roleNames).Find(&roles).Error; err != nil {
			return errs.Wrap(errs.CodeInternal, "查询角色失败", err)
		}
		if len(roles) != len(roleNames) {
			return errs.New(errs.CodeUnprocessable, "存在未知角色")
		}
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("model_type = ? AND model_id = ?", UserModelType, user.ID).
			Delete(&model.ModelHasRole{}).Error; err != nil {
			return err
		}
		links := make([]model.ModelHasRole, len(roles))
		for i, r := range roles {
			links[i] = model.ModelHasRole{RoleID: r.ID, ModelType: UserModelType, ModelID: user.ID}
		}
		if len(links) > 0 {
			if err := tx.Create(&links).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errs.Wrap(errs.CodeInternal, "更新用户角色失败", err)
	}

	userSub := rbac.SubjectForUser(user.ID)
	if _, err := s.enforcer.RemoveFilteredGroupingPolicy(0, userSub); err != nil {
		return errs.Wrap(errs.CodeInternal, "清除旧分组失败", err)
	}
	for _, r := range roles {
		if _, err := s.enforcer.AddGroupingPolicy(userSub, rbac.SubjectForRole(r.Name)); err != nil {
			return errs.Wrap(errs.CodeInternal, "写入分组失败", err)
		}
	}
	return nil
}

// GetUserRolesAndPermissions returns the user's role names and the union of
// their effective permissions.
func (s *RBACService) GetUserRolesAndPermissions(ctx context.Context, userID int64) ([]string, []string, error) {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, nil, errs.New(errs.CodeNotFound, "用户不存在")
	}
	if user.IsPlatformAdmin {
		// Platform admins get every defined permission as a courtesy for the
		// frontend `permissions` array — actual enforcement still bypasses Casbin.
		var allPerms []string
		if err := s.db.WithContext(ctx).Model(&model.Permission{}).Pluck("name", &allPerms).Error; err != nil {
			return nil, nil, err
		}
		return []string{"platform_admin"}, allPerms, nil
	}

	var roleIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.ModelHasRole{}).
		Where("model_type = ? AND model_id = ?", UserModelType, user.ID).
		Pluck("role_id", &roleIDs).Error; err != nil {
		return nil, nil, err
	}

	var roleNames []string
	if len(roleIDs) > 0 {
		if err := s.db.WithContext(ctx).Model(&model.Role{}).
			Where("id IN ?", roleIDs).Pluck("name", &roleNames).Error; err != nil {
			return nil, nil, err
		}
	}

	permsSet := map[string]struct{}{}
	for _, name := range roleNames {
		// Casbin returns flat (sub, obj, act) tuples; collect those whose sub
		// matches the role's mapped subject.
		filtered, err := s.enforcer.GetFilteredPolicy(0, rbac.SubjectForRole(name))
		if err != nil {
			return nil, nil, err
		}
		for _, row := range filtered {
			if len(row) < 3 {
				continue
			}
			permsSet[joinPermission(row[1], row[2])] = struct{}{}
		}
	}
	perms := make([]string, 0, len(permsSet))
	for p := range permsSet {
		perms = append(perms, p)
	}
	return roleNames, perms, nil
}

func splitPermissionName(name string) (string, string) {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[:i], name[i+1:]
		}
	}
	return "*", name
}

func joinPermission(resource, action string) string {
	if resource == "*" {
		return action
	}
	return fmt.Sprintf("%s.%s", resource, action)
}
