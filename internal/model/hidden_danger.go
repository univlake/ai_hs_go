package model

import "time"

// 隐患信息 + 隐患附件 — 专家检查模块底层数据。
//
// HiddenDangerInfo：检查现场登记的隐患，单条贯穿三态：登记 → 整改 → 复查/归档。
//   correct_status: 1 待整改 / 2 已整改待确认 / 3 已确认
//   invest_status:  1 处理中 / 2 已归档
//
// HiddenUploadAnnex：每条隐患可挂三类附件（is_review）：
//   1 = 排查证据 hiddenExamineAnnex
//   2 = 复查证据 hiddenReviewAnnex
//   3 = 整改证据 hiddenRectifyAnnex

type HiddenDangerInfo struct {
	ID                  int64      `gorm:"primaryKey;column:id" json:"id"`
	ExpertInspectionID  *int64     `gorm:"column:expert_inspection_id" json:"expert_inspection_id"`
	ExpertCompanyID     *int64     `gorm:"column:expert_company_id" json:"expert_company_id"`
	CheckCompanyID      *int64     `gorm:"column:check_company_id" json:"check_company_id"`
	Area                string     `gorm:"column:area;not null" json:"area"`
	Session             string     `gorm:"column:session;not null;default:'第一轮'" json:"session"`
	Description         string     `gorm:"column:description;not null" json:"description"`
	HiddenType          string     `gorm:"column:hidden_type;not null" json:"hidden_type"`
	HiddenLevel         string     `gorm:"column:hidden_level;not null" json:"hidden_level"`
	Legalises           *string    `gorm:"column:legalises" json:"legalises"`
	Proposal            *string    `gorm:"column:proposal" json:"proposal"`
	CorrectDeadline     *Date      `gorm:"column:correct_deadline;type:date" json:"correct_deadline"`
	InvestStatus        int        `gorm:"column:invest_status;not null;default:1" json:"invest_status"`
	ZgUserID            *int64     `gorm:"column:zg_user_id" json:"zg_user_id"`
	CorrectStatus       int        `gorm:"column:correct_status;not null;default:1" json:"correct_status"`
	CorrectDate         *Date      `gorm:"column:correct_date;type:date" json:"correct_date"`
	ReviewDescription   *string    `gorm:"column:review_description" json:"review_description"`
	ReviewUserID        *int64     `gorm:"column:review_user_id" json:"review_user_id"`
	ReviewTime          *time.Time `gorm:"column:review_time" json:"review_time"`
	ArchiveUserID       *int64     `gorm:"column:archive_user_id" json:"archive_user_id"`
	Comment             *string    `gorm:"column:comment" json:"comment"`
	CompanyID           int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	CheckCompany       *Company           `gorm:"foreignKey:CheckCompanyID;references:ID" json:"check_company,omitempty"`
	HiddenExamineAnnex []HiddenUploadAnnex `gorm:"-" json:"hidden_examine_annex,omitempty"`
	HiddenReviewAnnex  []HiddenUploadAnnex `gorm:"-" json:"hidden_review_annex,omitempty"`
	HiddenRectifyAnnex []HiddenUploadAnnex `gorm:"-" json:"hidden_rectify_annex,omitempty"`
}

func (HiddenDangerInfo) TableName() string { return "hidden_danger_infos" }

// HiddenUploadAnnex — Audit 不沿用：原表只有 user_id（创建者）+ timestamps，没有 last_user_id /
// soft-delete 字段。这里手写 5 字段保持跟 Laravel 一致。
type HiddenUploadAnnex struct {
	ID                  int64     `gorm:"primaryKey;column:id" json:"id"`
	HiddenDangerInfoID  int64     `gorm:"column:hidden_danger_info_id;not null" json:"hidden_danger_info_id"`
	Geolocation         string    `gorm:"column:geolocation;not null;default:''" json:"geolocation"`
	IsReview            int       `gorm:"column:is_review;not null;default:1" json:"is_review"`
	AnnexURL            string    `gorm:"column:annex_url;not null" json:"annex_url"`
	UserID              int64     `gorm:"column:user_id;not null;default:0" json:"user_id"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (HiddenUploadAnnex) TableName() string { return "hidden_upload_annexes" }
