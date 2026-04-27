package model

import (
	"time"

	"gorm.io/gorm"
)

// 考核评分模块数据模型 — 5 个维度的填报/评分 + 模板树。
//
//   daily_check_templates / daily_checks       日常目标管理工作（填报草稿 → 已提交 → 已评分）
//   accident_check_templates / accident_checks 安全生产事故考核（一次性评分）
//   add_points_templates / add_points          加分项（一次性评分）
//   system_checks                              活动参与及材料上报（compute 自动算）
//   norm_checks                                安全生产标准化（每企业每年单条）
//
// 三张模板表共用 (title, level, sort, year, pid, score_range) 形状，
// 通过 pid 自引用形成两层树（level=1 一级指标、level=2 二级）。
// 模板表是字典型数据，没有 user_id / last_user_id 审计列，只有 timestamps + softDelete
// （跟 Laravel 原表一致）；填报数据表才走完整的 Audit 嵌入。
// JSON 字段 `child` 由 service 层 attach（不走 GORM Preload）。

// timestampOnly — 模板表用：仅时间戳 + 软删，不含 user_id / last_user_id。
type timestampOnly struct {
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

// ── 日常检查（filing 流） ────────────────────────────────────────────

type DailyCheckTemplate struct {
	ID         int64  `gorm:"primaryKey;column:id" json:"id"`
	Title      string `gorm:"column:title;not null" json:"title"`
	Level      int    `gorm:"column:level;not null;default:1" json:"level"`
	Sort       int    `gorm:"column:sort;not null;default:0" json:"sort"`
	Year       int    `gorm:"column:year;not null;default:2022" json:"year"`
	Pid        int64  `gorm:"column:pid;not null;default:0" json:"pid"`
	ScoreRange string `gorm:"column:score_range;not null;default:'1'" json:"score_range"`

	timestampOnly

	Child []DailyCheckTemplate `gorm:"-" json:"child,omitempty"`
	Check *DailyCheck          `gorm:"-" json:"check,omitempty"`
}

func (DailyCheckTemplate) TableName() string { return "daily_check_templates" }

// DailyCheck — 注意 number 是 NUMERIC(8,1)；GORM string 反序列化保留小数。
type DailyCheck struct {
	ID                   int64   `gorm:"primaryKey;column:id" json:"id"`
	DailyCheckTemplateID int64   `gorm:"column:daily_check_template_id;not null" json:"daily_check_template_id"`
	Path                 *string `gorm:"column:path" json:"path"`
	Number               float64 `gorm:"column:number;not null;default:0;type:numeric(8,1)" json:"number"`
	Comment              *string `gorm:"column:comment" json:"comment"`
	Status               int     `gorm:"column:status;not null;default:0" json:"status"`
	Year                 int     `gorm:"column:year;not null;default:2022" json:"year"`
	CompanyID            int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (DailyCheck) TableName() string { return "daily_checks" }

// ── 事故检查 ────────────────────────────────────────────────────────

type AccidentCheckTemplate struct {
	ID         int64  `gorm:"primaryKey;column:id" json:"id"`
	Title      string `gorm:"column:title;not null" json:"title"`
	Level      int    `gorm:"column:level;not null;default:1" json:"level"`
	Sort       int    `gorm:"column:sort;not null;default:0" json:"sort"`
	Year       int    `gorm:"column:year;not null;default:2022" json:"year"`
	Pid        int64  `gorm:"column:pid;not null;default:0" json:"pid"`
	ScoreRange string `gorm:"column:score_range;not null;default:'1'" json:"score_range"`

	timestampOnly

	Child []AccidentCheckTemplate `gorm:"-" json:"child,omitempty"`
	Check *AccidentCheck          `gorm:"-" json:"check,omitempty"`
}

func (AccidentCheckTemplate) TableName() string { return "accident_check_templates" }

type AccidentCheck struct {
	ID                      int64   `gorm:"primaryKey;column:id" json:"id"`
	AccidentCheckTemplateID int64   `gorm:"column:accident_check_template_id;not null" json:"accident_check_template_id"`
	Path                    *string `gorm:"column:path" json:"path"`
	Number                  float64 `gorm:"column:number;not null;default:0;type:numeric(8,1)" json:"number"`
	Comment                 *string `gorm:"column:comment" json:"comment"`
	Year                    int     `gorm:"column:year;not null;default:2022" json:"year"`
	CompanyID               int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (AccidentCheck) TableName() string { return "accident_checks" }

// ── 加分项 ─────────────────────────────────────────────────────────

type AddPointsTemplate struct {
	ID         int64  `gorm:"primaryKey;column:id" json:"id"`
	Title      string `gorm:"column:title;not null" json:"title"`
	Level      int    `gorm:"column:level;not null;default:1" json:"level"`
	Sort       int    `gorm:"column:sort;not null;default:0" json:"sort"`
	Year       int    `gorm:"column:year;not null;default:2022" json:"year"`
	Pid        int64  `gorm:"column:pid;not null;default:0" json:"pid"`
	ScoreRange string `gorm:"column:score_range;not null;default:'1'" json:"score_range"`

	timestampOnly

	Child []AddPointsTemplate `gorm:"-" json:"child,omitempty"`
	Check *AddPoint           `gorm:"-" json:"check,omitempty"`
}

func (AddPointsTemplate) TableName() string { return "add_points_templates" }

type AddPoint struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	AddPointsTemplateID int64   `gorm:"column:add_points_template_id;not null" json:"add_points_template_id"`
	Path                *string `gorm:"column:path" json:"path"`
	Number              float64 `gorm:"column:number;not null;default:0;type:numeric(8,1)" json:"number"`
	Comment             *string `gorm:"column:comment" json:"comment"`
	Year                int     `gorm:"column:year;not null;default:2022" json:"year"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	AddPointsTemplate *AddPointsTemplate `gorm:"foreignKey:AddPointsTemplateID;references:ID" json:"addPointsTemplate,omitempty"`
}

func (AddPoint) TableName() string { return "add_points" }

// ── 系统考核（活动参与及材料上报） ─────────────────────────────────

type SystemCheck struct {
	ID         int64   `gorm:"primaryKey;column:id" json:"id"`
	NumberV1   float32 `gorm:"column:number_v1;not null;default:0" json:"number_v1"`
	NumberV2   float32 `gorm:"column:number_v2;not null;default:0" json:"number_v2"`
	NumberV3   float32 `gorm:"column:number_v3;not null;default:0" json:"number_v3"`
	NumberV4   float32 `gorm:"column:number_v4;not null;default:0" json:"number_v4"`
	NumberV5   float32 `gorm:"column:number_v5;not null;default:0" json:"number_v5"`
	Number     int     `gorm:"column:number;not null;default:0" json:"number"`
	Year       int     `gorm:"column:year;not null;default:2022" json:"year"`
	CompanyID  int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Company *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (SystemCheck) TableName() string { return "system_checks" }

// ── 标准化考核 ─────────────────────────────────────────────────────

type NormCheck struct {
	ID        int64   `gorm:"primaryKey;column:id" json:"id"`
	Path      *string `gorm:"column:path" json:"path"`
	Number    int     `gorm:"column:number;not null;default:0" json:"number"`
	Year      int     `gorm:"column:year;not null;default:2022" json:"year"`
	Type      string  `gorm:"column:type;not null" json:"type"`
	CompanyID int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Company *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (NormCheck) TableName() string { return "norm_checks" }
