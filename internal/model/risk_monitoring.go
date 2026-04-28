package model

// 风险监测栏目（集团视图）支撑模型。
//
// FocusCompany 关注企业关系（company_id 集团 → focus_id 关注企业）。
// 序列化时以 `company` 字段返回 focus_id 指向的企业资料，沿用 Laravel 关系名。
//
// RiskLog 集团下属企业每日风险指数快照（type=1 固有 / 2 浮动），date 为 YYYYMMDD 整数。

type FocusCompany struct {
	ID        int64  `gorm:"primaryKey;column:id" json:"id"`
	FocusID   *int64 `gorm:"column:focus_id" json:"focus_id"`
	CompanyID int64  `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	// Eager-loaded — Laravel 中 FocusCompany::company belongsTo(Company, 'focus_id')。
	// 列表接口仅返回必要字段：title + 风险等级 / 数值，前端 RiskMonitor 侧面板直接渲染。
	Company *Company `gorm:"foreignKey:FocusID;references:ID" json:"company,omitempty"`
}

func (FocusCompany) TableName() string { return "focus_companies" }

type RiskLog struct {
	ID            int64    `gorm:"primaryKey;column:id" json:"id"`
	Type          int16    `gorm:"column:type;not null;default:0" json:"type"`
	OldRiskNum    *float64 `gorm:"column:old_risk_num" json:"old_risk_num"`
	RiskNum       float64  `gorm:"column:risk_num;not null;default:0" json:"risk_num"`
	Date          int      `gorm:"column:date;not null;default:0" json:"date"`
	Content       *string  `gorm:"column:content" json:"content"`
	CompanyName   *string  `gorm:"column:company_name" json:"company_name"`
	OldRiskGrade  *string  `gorm:"column:old_risk_grade" json:"old_risk_grade"`
	RiskGrade     *string  `gorm:"column:risk_grade" json:"risk_grade"`
	CompanyID     int64    `gorm:"column:company_id;not null;default:0" json:"company_id"`

	Audit
}

func (RiskLog) TableName() string { return "risk_logs" }
