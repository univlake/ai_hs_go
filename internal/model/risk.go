package model

// 风险管理模块：风险源管理 / 隐患管理 / 作业管理 / 三同时管理 / 安全评价与检测。
//
// RiskDistributionMap     per-company singleton（风险分布图，前端 hazards.list 权限）
// MonthlyInvestigation    隐患排查主表（关联 reviews/views 两张子表，整体替换语义）
// MonthlyInvestigationReview / MonthlyInvestigationView   附件子表，company_id + audit
// JobManagement           作业管理（只读关联 EvaluationUnit）
// ThreeTogether           三同时管理（只读关联 DesignUnit / EvaluationUnit）
// AssessDetect            安全评价与检测（只读关联 SkillUnit / SkillUnitKid）

// RiskDistributionMap — 风险源管理。company 内 0~1 行，沿用 HealthSystem 的 singleton 写法。
type RiskDistributionMap struct {
	ID        int64   `gorm:"primaryKey;column:id" json:"id"`
	AnnexURL  *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID int64   `gorm:"column:company_id;not null" json:"company_id"`

	// Filled in at the handler boundary, never persisted — handler 在 GET
	// 路径上调用 oss.SignURL 生成有效期 1 小时的 GET 预签名 URL，前端 <img src> 可直接用。
	SignedAnnexURL *string `gorm:"-" json:"signed_annex_url,omitempty"`

	Audit
}

func (RiskDistributionMap) TableName() string { return "risk_distribution_maps" }

// MonthlyInvestigation — 隐患管理主表。Laravel 字段名 review_discription 是
// 拼写错误（discription→description），但前端表单 / Laravel 字段都已固化，沿用避免 hash 不一致。
type MonthlyInvestigation struct {
	ID                int64   `gorm:"primaryKey;column:id" json:"id"`
	Type              *string `gorm:"column:type" json:"type"`
	IsFireFighting    *int16  `gorm:"column:is_fire_fighting" json:"is_fire_fighting"`
	IsEmergency       *int16  `gorm:"column:is_emergency" json:"is_emergency"`
	Content           *string `gorm:"column:content" json:"content"`
	InvTime           *Date   `gorm:"column:inv_time;type:date" json:"inv_time"`
	InvName           *string `gorm:"column:inv_name" json:"inv_name"`
	Level             *string `gorm:"column:level" json:"level"`
	Proposal          *string `gorm:"column:proposal" json:"proposal"`
	Legalbasis        *string `gorm:"column:legalbasis" json:"legalbasis"`
	CorrectDate       *Date   `gorm:"column:correct_date;type:date" json:"correct_date"`
	ReviewDiscription *string `gorm:"column:review_discription" json:"review_discription"`
	ReviewDate        *Date   `gorm:"column:review_date;type:date" json:"review_date"`
	ReviewContent     *string `gorm:"column:review_content" json:"review_content"`
	ReviewUser        *string `gorm:"column:review_user" json:"review_user"`
	CorrectResult     *string `gorm:"column:correct_result" json:"correct_result"`
	IsGovern          *int16  `gorm:"column:is_govern" json:"is_govern"`
	GovernTitle       *string `gorm:"column:govern_title" json:"govern_title"`
	GovernSupplies    *string `gorm:"column:govern_supplies" json:"govern_supplies"`
	GovernUser        *string `gorm:"column:govern_user" json:"govern_user"`
	GovernDate        *string `gorm:"column:govern_date" json:"govern_date"`
	GovernText        *string `gorm:"column:govern_text" json:"govern_text"`
	GovernMoney       *string `gorm:"column:govern_money" json:"govern_money"`
	CompanyID         int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	// JSON 键名沿用 Laravel 关系名（驼峰），前端表单读写都用这两个 key。
	Reviews []MonthlyInvestigationReview `gorm:"foreignKey:MonthlyInvestigationID" json:"monthlyInvestigationReviews,omitempty"`
	Views   []MonthlyInvestigationView   `gorm:"foreignKey:MonthlyInvestigationID" json:"monthlyInvestigationViews,omitempty"`

	// 集团端 /bloc-monthly-investigations 列表 preload company，单企业 CRUD
	// 不主动加载 — omitempty 保证不会污染原有响应。
	Company *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (MonthlyInvestigation) TableName() string { return "monthly_investigations" }

type MonthlyInvestigationReview struct {
	ID                     int64   `gorm:"primaryKey;column:id" json:"id"`
	MonthlyInvestigationID int64   `gorm:"column:monthly_investigation_id;not null" json:"monthly_investigation_id"`
	AnnexURL               *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID              int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (MonthlyInvestigationReview) TableName() string { return "monthly_investigation_reviews" }

type MonthlyInvestigationView struct {
	ID                     int64   `gorm:"primaryKey;column:id" json:"id"`
	MonthlyInvestigationID int64   `gorm:"column:monthly_investigation_id;not null" json:"monthly_investigation_id"`
	AnnexURL               *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID              int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (MonthlyInvestigationView) TableName() string { return "monthly_investigation_views" }

// JobManagement — 作业管理。Laravel 表名 job_management（单数），保持一致。
type JobManagement struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Department       *string `gorm:"column:department" json:"department"`
	Content          *string `gorm:"column:content" json:"content"`
	Type             *string `gorm:"column:type" json:"type"`
	SpecialType      *string `gorm:"column:special_type" json:"special_type"`
	Position         *string `gorm:"column:position" json:"position"`
	JobTime          *Date   `gorm:"column:job_time;type:date" json:"job_time"`
	Unit             *string `gorm:"column:unit" json:"unit"`
	IsOutsourcing    *int16  `gorm:"column:is_outsourcing" json:"is_outsourcing"`
	EvaluationUnitID *int64  `gorm:"column:evaluation_unit_id" json:"evaluation_unit_id"`
	JobPeopleNumber  *int32  `gorm:"column:job_people_number" json:"job_people_number"`
	JobPeopleName    *string `gorm:"column:job_people_name" json:"job_people_name"`
	RiskAnalysis     *string `gorm:"column:risk_analysis" json:"risk_analysis"`
	RiskProtection   *string `gorm:"column:risk_protection" json:"risk_protection"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	EvaluationUnit *EvaluationUnit `gorm:"foreignKey:EvaluationUnitID;references:ID" json:"evaluation_unit,omitempty"`
}

func (JobManagement) TableName() string { return "job_management" }

// ThreeTogether — 三同时管理（设计/评价单位 readonly 关联）。
type ThreeTogether struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                *string `gorm:"column:name" json:"name"`
	ThreeTogetherType   *string `gorm:"column:three_together_type" json:"three_together_type"`
	EstablishTime       *Date   `gorm:"column:establish_time;type:date" json:"establish_time"`
	EstablishSection    *string `gorm:"column:establish_section" json:"establish_section"`
	Status              *string `gorm:"column:status" json:"status"`
	DesignUnitID        *int64  `gorm:"column:design_unit_id" json:"design_unit_id"`
	DesignAnnexURL      *string `gorm:"column:design_annex_url" json:"design_annex_url"`
	EvaluationUnitID    *int64  `gorm:"column:evaluation_unit_id" json:"evaluation_unit_id"`
	WriteTime           *Date   `gorm:"column:write_time;type:date" json:"write_time"`
	ConstructionComment *string `gorm:"column:construction_comment" json:"construction_comment"`
	SafetyRequirements  *string `gorm:"column:safety_requirements" json:"safety_requirements"`
	CompleteComment     *string `gorm:"column:complete_comment" json:"complete_comment"`
	Acceptance          *string `gorm:"column:acceptance" json:"acceptance"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	DesignUnit     *DesignUnit     `gorm:"foreignKey:DesignUnitID;references:ID" json:"design_unit,omitempty"`
	EvaluationUnit *EvaluationUnit `gorm:"foreignKey:EvaluationUnitID;references:ID" json:"evaluation_unit,omitempty"`
}

func (ThreeTogether) TableName() string { return "three_togethers" }

// AssessDetect — 安全评价与检测，结构跟 HygieneEvaluation 几乎一致，
// 只是 type 字段叫 assess_detect_type。
type AssessDetect struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Name             *string `gorm:"column:name;not null" json:"name"`
	AssessDetectType *string `gorm:"column:assess_detect_type" json:"assess_detect_type"`
	SkillUnitID      *int64  `gorm:"column:skill_unit_id" json:"skill_unit_id"`
	SkillUnitKidID   *int64  `gorm:"column:skill_unit_kid_id" json:"skill_unit_kid_id"`
	ValidPeriod      *Date   `gorm:"column:valid_period;type:date" json:"valid_period"`
	AnnexURL         *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	SkillUnit    *SkillUnit    `gorm:"foreignKey:SkillUnitID;references:ID" json:"skill_unit,omitempty"`
	SkillUnitKid *SkillUnitKid `gorm:"foreignKey:SkillUnitKidID;references:ID" json:"skill_unit_kid,omitempty"`
}

func (AssessDetect) TableName() string { return "assess_detects" }
