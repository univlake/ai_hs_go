package model

import (
	"time"

	"gorm.io/gorm"
)

// SpecialType is the catalog used by SpecialEquipment.special_type_id.
// Seeded with the 6 standard categories that match the Laravel project so
// the legacy "spec field config" (norm_model1..4 → 设计压力/容积/管径/...)
// addressed by id stays consistent across both stacks.
type SpecialType struct {
	ID        int64          `gorm:"primaryKey;column:id" json:"id"`
	Title     string         `gorm:"column:title;size:255;not null" json:"title"`
	Sort      int            `gorm:"column:sort;not null;default:0" json:"sort"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (SpecialType) TableName() string { return "special_types" }
