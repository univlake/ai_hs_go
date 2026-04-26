package model

import (
	"time"

	"gorm.io/gorm"
)

// Audit holds the four-field audit columns shared by most domain tables.
// Use embedding instead of inheritance.
type Audit struct {
	UserID     int64          `gorm:"column:user_id;not null;default:0" json:"user_id"`
	LastUserID int64          `gorm:"column:last_user_id;not null;default:0" json:"last_user_id"`
	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}
