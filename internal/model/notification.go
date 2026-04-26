package model

import "time"

type NewNotify struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Title     string    `gorm:"column:title;not null" json:"title"`
	Body      *string   `gorm:"column:body" json:"body"`
	Data      []byte    `gorm:"column:data;type:jsonb;default:'{}'" json:"data"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (NewNotify) TableName() string { return "new_notifies" }

type NewNotifyScope struct {
	ID        int64      `gorm:"primaryKey;column:id" json:"id"`
	NotifyID  int64      `gorm:"column:notify_id;not null" json:"notify_id"`
	UserID    int64      `gorm:"column:user_id;not null" json:"user_id"`
	IsRead    bool       `gorm:"column:is_read;not null;default:false" json:"is_read"`
	ReadAt    *time.Time `gorm:"column:read_at" json:"read_at"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (NewNotifyScope) TableName() string { return "new_notify_scopes" }
