package model

import "time"

type Permission struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	GuardName string    `gorm:"column:guard_name;not null;default:'api'" json:"guard_name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Permission) TableName() string { return "permissions" }

type Role struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	GuardName string    `gorm:"column:guard_name;not null;default:'api'" json:"guard_name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Role) TableName() string { return "roles" }

type RoleHasPermission struct {
	PermissionID int64 `gorm:"primaryKey;column:permission_id"`
	RoleID       int64 `gorm:"primaryKey;column:role_id"`
}

func (RoleHasPermission) TableName() string { return "role_has_permissions" }

type ModelHasRole struct {
	RoleID    int64  `gorm:"primaryKey;column:role_id"`
	ModelType string `gorm:"primaryKey;column:model_type"`
	ModelID   int64  `gorm:"primaryKey;column:model_id"`
}

func (ModelHasRole) TableName() string { return "model_has_roles" }

type ModelHasPermission struct {
	PermissionID int64  `gorm:"primaryKey;column:permission_id"`
	ModelType    string `gorm:"primaryKey;column:model_type"`
	ModelID      int64  `gorm:"primaryKey;column:model_id"`
}

func (ModelHasPermission) TableName() string { return "model_has_permissions" }
