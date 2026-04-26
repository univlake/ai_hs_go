package model

// 职业健康 + 消防安全九张业务表。OccupationalHazard / HealthFile 带 import/export，
// HealthSystem 是 per-company singleton（mirror OccupationalHealthSafety 写法）。
// HygieneEvaluation / FireSafety 关联 SkillUnit + SkillUnitKid（外键软引用，
// 已存在的 0017_third_party_module 创建的两张表）。

type OccupationalHazard struct {
	ID          int64   `gorm:"primaryKey;column:id" json:"id"`
	Name        *string `gorm:"column:name" json:"name"`
	Type        *string `gorm:"column:type" json:"type"`
	Number      *int32  `gorm:"column:number" json:"number"`
	IsDeclare   *string `gorm:"column:is_declare" json:"is_declare"`
	DeclareDate *Date   `gorm:"column:declare_date;type:date" json:"declare_date"`
	CompanyID   int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (OccupationalHazard) TableName() string { return "occupational_hazards" }

// HealthSystem — 职业健康体系，company 内 0~1 行。Laravel index 直接返回 first()。
type HealthSystem struct {
	ID                int64   `gorm:"primaryKey;column:id" json:"id"`
	CertificationTime *Date   `gorm:"column:certification_time;type:date" json:"certification_time"`
	AnnexURL          *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID         int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (HealthSystem) TableName() string { return "health_systems" }

type HealthFile struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Name             *string `gorm:"column:name" json:"name"`
	Number           *string `gorm:"column:number" json:"number"`
	Department       *string `gorm:"column:department" json:"department"`
	Post             *string `gorm:"column:post" json:"post"`
	CheckupDate      *Date   `gorm:"column:checkup_date;type:date" json:"checkup_date"`
	CheckupExpiry    *string `gorm:"column:checkup_expiry" json:"checkup_expiry"`
	ReviewSituation  *string `gorm:"column:review_situation" json:"review_situation"`
	HandlingOpinions *string `gorm:"column:handling_opinions" json:"handling_opinions"`
	ProcessResult    *string `gorm:"column:process_result" json:"process_result"`
	AnnexURL         *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (HealthFile) TableName() string { return "health_files" }

type HealthReport struct {
	ID           int64   `gorm:"primaryKey;column:id" json:"id"`
	FindDate     *Date   `gorm:"column:find_date;type:date" json:"find_date"`
	Name         *string `gorm:"column:name" json:"name"`
	Post         *string `gorm:"column:post" json:"post"`
	Type         *string `gorm:"column:type" json:"type"`
	IsTrue       *int16  `gorm:"column:is_true" json:"is_true"`
	Mechanism    *string `gorm:"column:mechanism" json:"mechanism"`
	IdentifyDate *Date   `gorm:"column:identify_date;type:date" json:"identify_date"`
	AnnexURL     *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID    int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (HealthReport) TableName() string { return "health_reports" }

type HealthCheckup struct {
	ID             int64   `gorm:"primaryKey;column:id" json:"id"`
	Type           *string `gorm:"column:type" json:"type"`
	Department     *string `gorm:"column:department" json:"department"`
	Number         *int32  `gorm:"column:number" json:"number"`
	CheckupDate    *Date   `gorm:"column:checkup_date;type:date" json:"checkup_date"`
	CheckupExpiry  *string `gorm:"column:checkup_expiry" json:"checkup_expiry"`
	Measure        *string `gorm:"column:measure" json:"measure"`
	AnnexURL       *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID      int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (HealthCheckup) TableName() string { return "health_checkups" }

// HygieneEvaluation 关联 skill_units / skill_unit_kids，前端 detail 抽屉直接渲染。
type HygieneEvaluation struct {
	ID             int64   `gorm:"primaryKey;column:id" json:"id"`
	Name           *string `gorm:"column:name" json:"name"`
	Type           *string `gorm:"column:type" json:"type"`
	SkillUnitID    *int64  `gorm:"column:skill_unit_id" json:"skill_unit_id"`
	SkillUnitKidID *int64  `gorm:"column:skill_unit_kid_id" json:"skill_unit_kid_id"`
	ValidPeriod    *Date   `gorm:"column:valid_period;type:date" json:"valid_period"`
	AnnexURL       *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID      int64   `gorm:"column:company_id;not null" json:"company_id"`

	SkillUnit    *SkillUnit    `gorm:"foreignKey:SkillUnitID;references:ID" json:"skill_unit,omitempty"`
	SkillUnitKid *SkillUnitKid `gorm:"foreignKey:SkillUnitKidID;references:ID" json:"skill_unit_kid,omitempty"`

	Audit
}

func (HygieneEvaluation) TableName() string { return "hygiene_evaluations" }

// FireSafety — 跟 HygieneEvaluation 同结构，valid_period 沿用 Laravel string 类型。
type FireSafety struct {
	ID             int64   `gorm:"primaryKey;column:id" json:"id"`
	Name           *string `gorm:"column:name" json:"name"`
	Type           *string `gorm:"column:type" json:"type"`
	SkillUnitID    *int64  `gorm:"column:skill_unit_id" json:"skill_unit_id"`
	SkillUnitKidID *int64  `gorm:"column:skill_unit_kid_id" json:"skill_unit_kid_id"`
	ValidPeriod    *string `gorm:"column:valid_period" json:"valid_period"`
	AnnexURL       *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID      int64   `gorm:"column:company_id;not null" json:"company_id"`

	SkillUnit    *SkillUnit    `gorm:"foreignKey:SkillUnitID;references:ID" json:"skill_unit,omitempty"`
	SkillUnitKid *SkillUnitKid `gorm:"foreignKey:SkillUnitKidID;references:ID" json:"skill_unit_kid,omitempty"`

	Audit
}

func (FireSafety) TableName() string { return "fire_safeties" }

type FireEquipment struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	EquipmentName    *string `gorm:"column:equipment_name" json:"equipment_name"`
	UtilizedLocation *string `gorm:"column:utilized_location" json:"utilized_location"`
	SpecialType      *string `gorm:"column:special_type" json:"special_type"`
	UseTime          *Date   `gorm:"column:use_time;type:date" json:"use_time"`
	AutomationType   *string `gorm:"column:automation_type" json:"automation_type"`
	MaintainInfo     *string `gorm:"column:maintain_info" json:"maintain_info"`
	Comment          *string `gorm:"column:comment" json:"comment"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (FireEquipment) TableName() string { return "fire_equipment" }

type FireFacility struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	EquipmentName    *string `gorm:"column:equipment_name" json:"equipment_name"`
	UtilizedLocation *string `gorm:"column:utilized_location" json:"utilized_location"`
	SpecialType      *string `gorm:"column:special_type" json:"special_type"`
	UseTime          *Date   `gorm:"column:use_time;type:date" json:"use_time"`
	AutomationType   *string `gorm:"column:automation_type" json:"automation_type"`
	MaintainInfo     *string `gorm:"column:maintain_info" json:"maintain_info"`
	Comment          *string `gorm:"column:comment" json:"comment"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (FireFacility) TableName() string { return "fire_facilities" }
