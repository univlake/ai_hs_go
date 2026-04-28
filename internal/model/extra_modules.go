package model

// 剩余模块的简单实体集合：
//
//   ChangeManage         变更管理
//   RegulationLibrary    法律法规库（多对一指向 RegulationLibraryType，eager 加载为 regulation_type 字段，跟 Laravel 关系名一致）
//   TrainingFile         培训资料-文件 + 配套下载流水 TrainingFileLog
//   TrainingVideo        培训资料-视频 + 配套观看流水 TrainingVideoLog
//   InherentRiskLog      固有风险评分快照（21 维度）
//   FloatRiskLog         浮动风险评分快照（17 维度）

// ChangeManage — 变更管理。多个 *_date 字段都用 *Date，与 process / risk 模块对齐。
type ChangeManage struct {
	ID                       int64   `gorm:"primaryKey;column:id" json:"id"`
	Title                    *string `gorm:"column:title" json:"title"`
	DepartmentName           *string `gorm:"column:department_name" json:"department_name"`
	ApplyDate                *Date   `gorm:"column:apply_date;type:date" json:"apply_date"`
	Reason                   *string `gorm:"column:reason" json:"reason"`
	ApplyLeaderOpinion       *string `gorm:"column:apply_leader_opinion" json:"apply_leader_opinion"`
	ApplyLeaderOpinionDate   *Date   `gorm:"column:apply_leader_opinion_date;type:date" json:"apply_leader_opinion_date"`
	RiskAnalyze              *string `gorm:"column:risk_analyze" json:"risk_analyze"`
	ManageLeaderOpinion      *string `gorm:"column:manage_leader_opinion" json:"manage_leader_opinion"`
	ManageLeaderOpinionDate  *Date   `gorm:"column:manage_leader_opinion_date;type:date" json:"manage_leader_opinion_date"`
	Approve                  *string `gorm:"column:approve" json:"approve"`
	ApproveDate              *Date   `gorm:"column:approve_date;type:date" json:"approve_date"`
	AcceptanceInfo           *string `gorm:"column:acceptance_info" json:"acceptance_info"`
	AcceptanceInfoDate       *Date   `gorm:"column:acceptance_info_date;type:date" json:"acceptance_info_date"`
	CompanyID                int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (ChangeManage) TableName() string { return "change_manages" }

// RegulationLibrary — 法律法规库。Laravel 关系名 regulationType (camelCase JSON key)，前端
// /risk/risk-warning 之外的页面（regulation/libraries）也按这个 key 渲染类型徽标，沿用。
type RegulationLibrary struct {
	ID                       int64   `gorm:"primaryKey;column:id" json:"id"`
	Title                    string  `gorm:"column:title;not null" json:"title"`
	RegulationLibraryTypeID  *int64  `gorm:"column:regulation_library_type_id" json:"regulation_library_type_id"`
	EstablishedDate          *Date   `gorm:"column:established_date;type:date" json:"established_date"`
	EffectiveDate            *Date   `gorm:"column:effective_date;type:date" json:"effective_date"`
	PromulgatingDepartment   *string `gorm:"column:promulgating_department" json:"promulgating_department"`
	StandardEncoding         *string `gorm:"column:standard_encoding" json:"standard_encoding"`
	IdentifyDate             *Date   `gorm:"column:identify_date;type:date" json:"identify_date"`
	Status                   *string `gorm:"column:status" json:"status"`
	AnnexURL                 *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID                int64   `gorm:"column:company_id;not null;default:0" json:"company_id"`

	Audit

	RegulationType *RegulationLibraryType `gorm:"foreignKey:RegulationLibraryTypeID;references:ID" json:"regulation_type,omitempty"`
}

func (RegulationLibrary) TableName() string { return "regulation_libraries" }

// TrainingFile — 培训资料·文件。show 接口签名 file_url 后返回字符串供前端 window.open。
type TrainingFile struct {
	ID            int64   `gorm:"primaryKey;column:id" json:"id"`
	Title         *string `gorm:"column:title" json:"title"`
	FileURL       *string `gorm:"column:file_url" json:"file_url"`
	TimesDownload int32   `gorm:"column:times_download;not null;default:0" json:"times_download"`
	CompanyID     int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (TrainingFile) TableName() string { return "training_files" }

type TrainingFileLog struct {
	ID             int64 `gorm:"primaryKey;column:id" json:"id"`
	TrainingFileID int64 `gorm:"column:training_file_id;not null" json:"training_file_id"`
	CompanyID      int64 `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (TrainingFileLog) TableName() string { return "training_file_logs" }

// TrainingVideo — 培训资料·视频。`introduce` / `lector` / `duration` 跟 Laravel 字段名对齐。
type TrainingVideo struct {
	ID           int64   `gorm:"primaryKey;column:id" json:"id"`
	Title        *string `gorm:"column:title" json:"title"`
	Introduce    *string `gorm:"column:introduce" json:"introduce"`
	ThumbURL     *string `gorm:"column:thumb_url" json:"thumb_url"`
	VideoURL     *string `gorm:"column:video_url" json:"video_url"`
	TimesWatched int32   `gorm:"column:times_watched;not null;default:0" json:"times_watched"`
	Duration     *string `gorm:"column:duration" json:"duration"`
	Lector       *string `gorm:"column:lector" json:"lector"`
	CompanyID    int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (TrainingVideo) TableName() string { return "training_videos" }

type TrainingVideoLog struct {
	ID              int64 `gorm:"primaryKey;column:id" json:"id"`
	TrainingVideoID int64 `gorm:"column:training_video_id;not null" json:"training_video_id"`
	CompanyID       int64 `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (TrainingVideoLog) TableName() string { return "training_video_logs" }

// InherentRiskLog — 固有风险评分（21 维度）。/risks/portrait 取最新一行。
type InherentRiskLog struct {
	ID                              int64   `gorm:"primaryKey;column:id" json:"id"`
	SafetyArchitectureValue         float64 `gorm:"column:safety_architecture_value;not null;default:0" json:"safety_architecture_value"`
	SecurityOfficerScale            float64 `gorm:"column:security_officer_scale;not null;default:0" json:"security_officer_scale"`
	AlreadyObtainPeopleScale        float64 `gorm:"column:already_obtain_people_scale;not null;default:0" json:"already_obtain_people_scale"`
	LastYearIncomeScale             float64 `gorm:"column:last_year_income_scale;not null;default:0" json:"last_year_income_scale"`
	FundingExecutionScale           float64 `gorm:"column:funding_execution_scale;not null;default:0" json:"funding_execution_scale"`
	ProductionStandardizationScale  float64 `gorm:"column:production_standardization_scale;not null;default:0" json:"production_standardization_scale"`
	HealthSystemScale               float64 `gorm:"column:health_system_scale;not null;default:0" json:"health_system_scale"`
	AssessDetectScale               float64 `gorm:"column:assess_detect_scale;not null;default:0" json:"assess_detect_scale"`
	HighSchoolDiplomaScale          float64 `gorm:"column:high_school_diploma_scale;not null;default:0" json:"high_school_diploma_scale"`
	AcreageScale                    float64 `gorm:"column:acreage_scale;not null;default:0" json:"acreage_scale"`
	StaffScale                      float64 `gorm:"column:staff_scale;not null;default:0" json:"staff_scale"`
	SafetyEngineeringMajorScale     float64 `gorm:"column:safety_engineering_major_scale;not null;default:0" json:"safety_engineering_major_scale"`
	HazardousProcessesScale         float64 `gorm:"column:hazardous_processes_scale;not null;default:0" json:"hazardous_processes_scale"`
	HazardousProcessScale           float64 `gorm:"column:hazardous_process_scale;not null;default:0" json:"hazardous_process_scale"`
	HazardousProcessMoreScale       float64 `gorm:"column:hazardous_process_more_scale;not null;default:0" json:"hazardous_process_more_scale"`
	ProcessIsAutomaticScale         float64 `gorm:"column:process_is_automatic_scale;not null;default:0" json:"process_is_automatic_scale"`
	SpecialEquipmentNumberScale     float64 `gorm:"column:special_equipment_number_scale;not null;default:0" json:"special_equipment_number_scale"`
	EquipmentMoreNumberScale        float64 `gorm:"column:equipment_more_number_scale;not null;default:0" json:"equipment_more_number_scale"`
	HazardousChemicalScale          float64 `gorm:"column:hazardous_chemical_scale;not null;default:0" json:"hazardous_chemical_scale"`
	HazardousChemicalsNumberScale   float64 `gorm:"column:hazardous_chemicals_number_scale;not null;default:0" json:"hazardous_chemicals_number_scale"`
	MajorHazardScale                float64 `gorm:"column:major_hazard_scale;not null;default:0" json:"major_hazard_scale"`
	CompanyID                       int64   `gorm:"column:company_id;not null;default:0" json:"company_id"`

	Audit
}

func (InherentRiskLog) TableName() string { return "inherent_risk_logs" }

// FloatRiskLog — 浮动风险评分（17 维度）。
type FloatRiskLog struct {
	ID                            int64   `gorm:"primaryKey;column:id" json:"id"`
	InternalMeetingScale          float64 `gorm:"column:internal_meeting_scale;not null;default:0" json:"internal_meeting_scale"`
	InternalTrainingScale         float64 `gorm:"column:internal_training_scale;not null;default:0" json:"internal_training_scale"`
	AccidentBulletinScale         float64 `gorm:"column:accident_bulletin_scale;not null;default:0" json:"accident_bulletin_scale"`
	AccidentBulletinSevereScale   float64 `gorm:"column:accident_bulletin_severe_scale;not null;default:0" json:"accident_bulletin_severe_scale"`
	AccidentReportURLScale        float64 `gorm:"column:accident_report_url_scale;not null;default:0" json:"accident_report_url_scale"`
	EmergencyTeamScale            float64 `gorm:"column:emergency_team_scale;not null;default:0" json:"emergency_team_scale"`
	EmergencyDrillScale           float64 `gorm:"column:emergency_drill_scale;not null;default:0" json:"emergency_drill_scale"`
	JobManagementScale            float64 `gorm:"column:job_management_scale;not null;default:0" json:"job_management_scale"`
	RiskSourceScale               float64 `gorm:"column:risk_source_scale;not null;default:0" json:"risk_source_scale"`
	MonthlyInvestigationScale     float64 `gorm:"column:monthly_investigation_scale;not null;default:0" json:"monthly_investigation_scale"`
	ThreeTogetherOverScale        float64 `gorm:"column:three_together_over_scale;not null;default:0" json:"three_together_over_scale"`
	UnitSumNumberScale            float64 `gorm:"column:unit_sum_number_scale;not null;default:0" json:"unit_sum_number_scale"`
	EvaluationUnitScale           float64 `gorm:"column:evaluation_unit_scale;not null;default:0" json:"evaluation_unit_scale"`
	EvaluationUnitSchoolScale     float64 `gorm:"column:evaluation_unit_school_scale;not null;default:0" json:"evaluation_unit_school_scale"`
	HealthCheckupPeopleScale      float64 `gorm:"column:health_checkup_people_scale;not null;default:0" json:"health_checkup_people_scale"`
	HealthCheckupScale            float64 `gorm:"column:health_checkup_scale;not null;default:0" json:"health_checkup_scale"`
	HealthReportScale             float64 `gorm:"column:health_report_scale;not null;default:0" json:"health_report_scale"`
	CompanyID                     int64   `gorm:"column:company_id;not null;default:0" json:"company_id"`

	Audit
}

func (FloatRiskLog) TableName() string { return "float_risk_logs" }
