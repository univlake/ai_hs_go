package model

import (
	"time"

	"github.com/google/uuid"
)

type ActivityLog struct {
	ID                int64      `gorm:"primaryKey;column:id" json:"id"`
	LogName           *string    `gorm:"column:log_name" json:"log_name"`
	Description       string     `gorm:"column:description;not null" json:"description"`
	SubjectType       *string    `gorm:"column:subject_type" json:"subject_type"`
	SubjectID         *int64     `gorm:"column:subject_id" json:"subject_id"`
	CauserType        *string    `gorm:"column:causer_type" json:"causer_type"`
	CauserID          *int64     `gorm:"column:causer_id" json:"causer_id"`
	Event             *string    `gorm:"column:event" json:"event"`
	Properties        []byte     `gorm:"column:properties;type:jsonb;default:'{}'" json:"properties"`
	AttributeChanges  []byte     `gorm:"column:attribute_changes;type:jsonb;default:'{}'" json:"attribute_changes"`
	BatchUUID         *uuid.UUID `gorm:"column:batch_uuid" json:"batch_uuid"`
	CreatedAt         time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (ActivityLog) TableName() string { return "activity_log" }
