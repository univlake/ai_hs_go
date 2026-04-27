package model

// 通知公告 / 会议通知 / 信息稿管理 — 8 张表对应的模型。
//
// 命名跟 Laravel 模型一致，所以前端取数路径（office_notify_scope、meeting_notices、
// information_drafts ...）不需要改动。Kids 整体替换语义跟 risk / accident 模块一致：
// service 收到 nil 切片 = 不动子表，非 nil（含空切片）= 整体替换。
//
// OfficeNotify / OfficeNotifyScope        通知公告 - 任务型（带反馈截止 + 反馈附件）
// NewNotify / NewNotifyScope              通知公告 - 通知型（仅需阅读确认）
// Meeting / MeetingNotice / MeetingSignUp 会议通知（接收企业 kids + 报名子表）
// InformationDraft                        信息稿（next_reviewer + status 驱动上报/发布）

// ── 任务型通知 ──────────────────────────────────────────────────────

type OfficeNotify struct {
	ID            int64   `gorm:"primaryKey;column:id" json:"id"`
	Name          *string `gorm:"column:name" json:"name"`
	FeedbackTime  *Date   `gorm:"column:feedback_time;type:date" json:"feedback_time"`
	AlertOthers   *string `gorm:"column:alert_others" json:"alert_others"`
	Comment       *string `gorm:"column:comment" json:"comment"`
	AnnexURL      *string `gorm:"column:annex_url" json:"annex_url"`
	ExcelURL      *string `gorm:"column:excel_url" json:"excel_url"`
	CompanyID     int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	OfficeNotifyScope []OfficeNotifyScope `gorm:"foreignKey:OfficeNotifyID" json:"office_notify_scope"`
}

func (OfficeNotify) TableName() string { return "office_notifies" }

type OfficeNotifyScope struct {
	ID                int64   `gorm:"primaryKey;column:id" json:"id"`
	CompanyName       *string `gorm:"column:company_name" json:"company_name"`
	FeedbackTime      *Date   `gorm:"column:feedback_time;type:date" json:"feedback_time"`
	FeedbackStatus    string  `gorm:"column:feedback_status;not null;default:'未反馈'" json:"feedback_status"`
	FeedbackAnnexURL  *string `gorm:"column:feedback_annex_url" json:"feedback_annex_url"`
	OfficeNotifyID    int64   `gorm:"column:office_notify_id;not null" json:"office_notify_id"`
	CompanyID         int64   `gorm:"column:company_id;not null" json:"company_id"`
	CompanyGroupName  *string `gorm:"column:company_group_name" json:"company_group_name"`
	CompanyGroupID    *int64  `gorm:"column:company_group_id" json:"company_group_id"`

	Audit

	OfficeNotify *OfficeNotify `gorm:"foreignKey:OfficeNotifyID;references:ID" json:"office_notify,omitempty"`
}

func (OfficeNotifyScope) TableName() string { return "office_notify_scopes" }

// ── 通知型通知 ──────────────────────────────────────────────────────

type NewNotify struct {
	ID           int64   `gorm:"primaryKey;column:id" json:"id"`
	Name         *string `gorm:"column:name" json:"name"`
	AlertOthers  *string `gorm:"column:alert_others" json:"alert_others"`
	Comment      *string `gorm:"column:comment" json:"comment"`
	AnnexURL     *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID    int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	NewNotifyScope []NewNotifyScope `gorm:"foreignKey:NewNotifyID" json:"new_notify_scope"`
}

func (NewNotify) TableName() string { return "new_notifies" }

type NewNotifyScope struct {
	ID                int64   `gorm:"primaryKey;column:id" json:"id"`
	CompanyName       *string `gorm:"column:company_name" json:"company_name"`
	ReadTime          *Date   `gorm:"column:read_time;type:date" json:"read_time"`
	ReadStatus        string  `gorm:"column:read_status;not null;default:'未读'" json:"read_status"`
	NewNotifyID       int64   `gorm:"column:new_notify_id;not null" json:"new_notify_id"`
	CompanyID         int64   `gorm:"column:company_id;not null" json:"company_id"`
	CompanyGroupName  *string `gorm:"column:company_group_name" json:"company_group_name"`
	CompanyGroupID    *int64  `gorm:"column:company_group_id" json:"company_group_id"`

	Audit

	NewNotify *NewNotify `gorm:"foreignKey:NewNotifyID;references:ID" json:"new_notify,omitempty"`
}

func (NewNotifyScope) TableName() string { return "new_notify_scopes" }

// ── 会议通知 ────────────────────────────────────────────────────────
//
// MeetingDate 用 DateTime（Laravel 字段是 timestamp，前端 datetime-local 输入）。
// MeetingNotice / MeetingSignUp 跟 Laravel hasMany 默认外键 (meeting_id) 对齐。

type Meeting struct {
	ID                int64     `gorm:"primaryKey;column:id" json:"id"`
	Name              *string   `gorm:"column:name" json:"name"`
	MeetingDate       *DateTime `gorm:"column:meeting_date" json:"meeting_date"`
	SignUpDate        *DateTime `gorm:"column:sign_up_date" json:"sign_up_date"`
	SignEndDate       *DateTime `gorm:"column:sign_end_date" json:"sign_end_date"`
	AlertOthers       *string   `gorm:"column:alert_others" json:"alert_others"`
	Theme             *string   `gorm:"column:theme" json:"theme"`
	Place             *string   `gorm:"column:place" json:"place"`
	Summary           *string   `gorm:"column:summary" json:"summary"`
	AnnexURL          *string   `gorm:"column:annex_url" json:"annex_url"`
	SummaryAnnexURL   *string   `gorm:"column:summary_annex_url" json:"summary_annex_url"`
	SignAnnexURL      *string   `gorm:"column:sign_annex_url" json:"sign_annex_url"`
	MaterialAnnexURL  *string   `gorm:"column:material_annex_url" json:"material_annex_url"`
	PhotoAnnexURL     *string   `gorm:"column:photo_annex_url" json:"photo_annex_url"`
	CompanyID         int64     `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	MeetingNotices []MeetingNotice  `gorm:"foreignKey:MeetingID" json:"meeting_notices"`
	MeetingSignUps []MeetingSignUp  `gorm:"foreignKey:MeetingID" json:"meeting_sign_ups,omitempty"`
	Company        *Company         `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (Meeting) TableName() string { return "meetings" }

type MeetingNotice struct {
	ID                int64     `gorm:"primaryKey;column:id" json:"id"`
	CompanyName       *string   `gorm:"column:company_name" json:"company_name"`
	MeetingID         int64     `gorm:"column:meeting_id;not null" json:"meeting_id"`
	CompanyID         int64     `gorm:"column:company_id;not null" json:"company_id"`
	CompanyGroupID    *int64    `gorm:"column:company_group_id" json:"company_group_id"`
	CompanyGroupName  *string   `gorm:"column:company_group_name" json:"company_group_name"`
	NotifyTime        *DateTime `gorm:"column:notify_time" json:"notify_time"`
	ConfirmTime       *DateTime `gorm:"column:confirm_time" json:"confirm_time"`

	Audit

	Meeting *Meeting `gorm:"foreignKey:MeetingID;references:ID" json:"meeting,omitempty"`
	Company *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (MeetingNotice) TableName() string { return "meeting_notices" }

type MeetingSignUp struct {
	ID           int64   `gorm:"primaryKey;column:id" json:"id"`
	CompanyName  *string `gorm:"column:company_name" json:"company_name"`
	Name         *string `gorm:"column:name" json:"name"`
	Sex          *string `gorm:"column:sex" json:"sex"`
	JobTitle     *string `gorm:"column:job_title" json:"job_title"`
	Tel          *string `gorm:"column:tel" json:"tel"`
	CardNumber   *string `gorm:"column:card_number" json:"card_number"`
	IsDriving    *string `gorm:"column:is_driving" json:"is_driving"`
	MeetingID    int64   `gorm:"column:meeting_id;not null" json:"meeting_id"`
	CompanyID    int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (MeetingSignUp) TableName() string { return "meeting_sign_ups" }

// ── 信息稿 ──────────────────────────────────────────────────────────

type InformationDraft struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Name             string  `gorm:"column:name;not null" json:"name"`
	CompanyName      string  `gorm:"column:company_name;not null" json:"company_name"`
	DraftDepartment  *string `gorm:"column:draft_department" json:"draft_department"`
	DraftOthers      *string `gorm:"column:draft_others" json:"draft_others"`
	Comment          *string `gorm:"column:comment" json:"comment"`
	AnnexURL         *string `gorm:"column:annex_url" json:"annex_url"`
	NextReviewer     *int64  `gorm:"column:next_reviewer" json:"next_reviewer"`
	Status           int     `gorm:"column:status;not null;default:0" json:"status"`
	IsTop            int     `gorm:"column:is_top;not null;default:0" json:"is_top"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (InformationDraft) TableName() string { return "information_drafts" }
