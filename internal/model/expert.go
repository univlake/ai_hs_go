package model

// 专家检查 + 专家库 — 4 张表对应的模型。
//
// 命名跟 Laravel 模型一致，所以前端取数路径（expert_companies、expert_inspection_users、
// bloc_expert_users ...）不需要改动。
//
// 三方协作：
//   ExpertInspection (检查计划，由计划方 company_id 创建)
//     └─ has_many ExpertCompany (受检企业 + 时间表)
//          └─ has_many ExpertInspectionUser (该次检查指派的专家清单)
//   BlocExpertUser (专家库 — 集团把系统用户登记为可用专家)
//
// Preload 关系跟 Laravel 路由层一致：
//   - GET /expert-inspections          .with([expertCompanies])
//   - GET /expert-inspections/:id      .with([expertCompanies.checkCompany, expertCompanies.expertInspectionUsers.expertUser])
//   - GET /expert-inspection-companies .with([checkCompany, expertInspectionUsers.expertUser])
//
// 时间字段：notify_company_time / check_company_time / confirm_time / notify_time 都是
// TIMESTAMPTZ；前端通过 datetime-local 输入「2026-04-21T09:55」字符串，handler 里
// parseDateTime 解析后落库。

// ── 检查计划 ────────────────────────────────────────────────────────

type ExpertInspection struct {
	ID          int64   `gorm:"primaryKey;column:id" json:"id"`
	Title       string  `gorm:"column:title;not null" json:"title"`
	Description *string `gorm:"column:description" json:"description"`
	Comment     *string `gorm:"column:comment" json:"comment"`
	StartDate   *Date   `gorm:"column:start_date;type:date" json:"start_date"`
	EndDate     *Date   `gorm:"column:end_date;type:date" json:"end_date"`
	CompanyID   int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	ExpertCompanies []ExpertCompany `gorm:"foreignKey:ExpertInspectionID" json:"expert_companies,omitempty"`
}

func (ExpertInspection) TableName() string { return "expert_inspections" }

// ── 受检企业 + 时间表 ───────────────────────────────────────────────

type ExpertCompany struct {
	ID                    int64     `gorm:"primaryKey;column:id" json:"id"`
	ExpertInspectionID    int64     `gorm:"column:expert_inspection_id;not null" json:"expert_inspection_id"`
	CheckCompanyID        int64     `gorm:"column:check_company_id;not null" json:"check_company_id"`
	NotifyCompanyTime     *DateTime `gorm:"column:notify_company_time" json:"notify_company_time"`
	CheckCompanyTime      *DateTime `gorm:"column:check_company_time" json:"check_company_time"`
	ConfirmTime           *DateTime `gorm:"column:confirm_time" json:"confirm_time"`
	InspectionNoticeURL   *string   `gorm:"column:inspection_notice_url" json:"inspection_notice_url"`
	RectificationSheetURL *string   `gorm:"column:rectification_sheet_url" json:"rectification_sheet_url"`
	Status                int       `gorm:"column:status;not null;default:1" json:"status"`
	CompanyID             int64     `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	ExpertInspection      *ExpertInspection      `gorm:"foreignKey:ExpertInspectionID;references:ID" json:"expert_inspection,omitempty"`
	CheckCompany          *Company               `gorm:"foreignKey:CheckCompanyID;references:ID" json:"check_company,omitempty"`
	ExpertInspectionUsers []ExpertInspectionUser `gorm:"foreignKey:ExpertCompanyID" json:"expert_inspection_users,omitempty"`
}

func (ExpertCompany) TableName() string { return "expert_companies" }

// ── 专家指派 ────────────────────────────────────────────────────────

type ExpertInspectionUser struct {
	ID              int64     `gorm:"primaryKey;column:id" json:"id"`
	ExpertUserID    int64     `gorm:"column:expert_user_id;not null" json:"expert_user_id"`
	ExpertCompanyID int64     `gorm:"column:expert_company_id;not null" json:"expert_company_id"`
	NotifyTime      *DateTime `gorm:"column:notify_time" json:"notify_time"`
	Status          int       `gorm:"column:status;not null;default:1" json:"status"`
	ConfirmTime     *DateTime `gorm:"column:confirm_time" json:"confirm_time"`
	Comment         *string   `gorm:"column:comment" json:"comment"`
	CompanyID       int64     `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	ExpertCompany *ExpertCompany `gorm:"foreignKey:ExpertCompanyID;references:ID" json:"expert_company,omitempty"`
	ExpertUser    *User          `gorm:"foreignKey:ExpertUserID;references:ID" json:"expert_user,omitempty"`
}

func (ExpertInspectionUser) TableName() string { return "expert_inspection_users" }

// ── 集团专家库 ──────────────────────────────────────────────────────

type BlocExpertUser struct {
	ID           int64 `gorm:"primaryKey;column:id" json:"id"`
	ExpertUserID int64 `gorm:"column:expert_user_id;not null" json:"expert_user_id"`
	CompanyID    int64 `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Company *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
	User    *User    `gorm:"foreignKey:ExpertUserID;references:ID" json:"user,omitempty"`
}

func (BlocExpertUser) TableName() string { return "bloc_expert_users" }
