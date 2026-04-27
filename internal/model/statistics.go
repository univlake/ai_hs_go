package model

// 数据统计模块（集团视图 + 生产视图共享）— 月度数据上报 / 年度数据上报。
//
// MonthStatistic         月度数据上报主表（每企业每月一行；30 个数值统计列 + 1 个文本说明）
// MonthStatisticControl  治理计划详情子表（整体替换语义；nil = 不动，非 nil = 全替换）
// MonthStatisticAccident 当月事故详情子表（整体替换语义）
// YearStatistic          年度数据上报主表（每企业每年一行）
// YearStatisticKid       三类人员 / 消防控制室人员 持证情况子表（整体替换语义）
//
// 集团视图（pid 过滤直接下级）跟生产视图（company_id 过滤本企业）共用同一张表，
// 仅在 service 层根据请求参数挑选过滤维度。

type MonthStatistic struct {
	ID                              int64   `gorm:"primaryKey;column:id" json:"id"`
	AverageEmployee                 int     `gorm:"column:average_employee;not null;default:0" json:"average_employee"`
	SafetyMeetingCount              int     `gorm:"column:safety_meeting_count;not null;default:0" json:"safety_meeting_count"`
	FireMeetingCount                int     `gorm:"column:fire_meeting_count;not null;default:0" json:"fire_meeting_count"`
	SafetyCheckCount                int     `gorm:"column:safety_check_count;not null;default:0" json:"safety_check_count"`
	FireCheckCount                  int     `gorm:"column:fire_check_count;not null;default:0" json:"fire_check_count"`
	SafetyFunding                   int     `gorm:"column:safety_funding;not null;default:0" json:"safety_funding"`
	FireFunding                     int     `gorm:"column:fire_funding;not null;default:0" json:"fire_funding"`
	HydratingSeedlings              int     `gorm:"column:hydrating_seedlings;not null;default:0" json:"hydrating_seedlings"`
	RetrofitElectrical              int     `gorm:"column:retrofit_electrical;not null;default:0" json:"retrofit_electrical"`
	FireExtinguisher                int     `gorm:"column:fire_extinguisher;not null;default:0" json:"fire_extinguisher"`
	SmartFire                       int     `gorm:"column:smart_fire;not null;default:0" json:"smart_fire"`
	SafetyTrainingCount             int     `gorm:"column:safety_training_count;not null;default:0" json:"safety_training_count"`
	SafetyTrainingNumber            int     `gorm:"column:safety_training_number;not null;default:0" json:"safety_training_number"`
	NewPeopleTrainingNumber         int     `gorm:"column:new_people_training_number;not null;default:0" json:"new_people_training_number"`
	FireTrainingCount               int     `gorm:"column:fire_training_count;not null;default:0" json:"fire_training_count"`
	FireTrainingNumber              int     `gorm:"column:fire_training_number;not null;default:0" json:"fire_training_number"`
	SafetyEmergencyCount            int     `gorm:"column:safety_emergency_count;not null;default:0" json:"safety_emergency_count"`
	SafetyEmergencyNumber           int     `gorm:"column:safety_emergency_number;not null;default:0" json:"safety_emergency_number"`
	FireEmergencyCount              int     `gorm:"column:fire_emergency_count;not null;default:0" json:"fire_emergency_count"`
	FireEmergencyNumber             int     `gorm:"column:fire_emergency_number;not null;default:0" json:"fire_emergency_number"`
	SomeHiddenDanger                int     `gorm:"column:some_hidden_danger;not null;default:0" json:"some_hidden_danger"`
	SomeRectifyHiddenDanger         int     `gorm:"column:some_rectify_hidden_danger;not null;default:0" json:"some_rectify_hidden_danger"`
	NoRectifyInfo                   *string `gorm:"column:no_rectify_info" json:"no_rectify_info"`
	BigHiddenDanger                 int     `gorm:"column:big_hidden_danger;not null;default:0" json:"big_hidden_danger"`
	BigRectifyHiddenDanger          int     `gorm:"column:big_rectify_hidden_danger;not null;default:0" json:"big_rectify_hidden_danger"`
	GovernHiddenDangerCount         int     `gorm:"column:govern_hidden_danger_count;not null;default:0" json:"govern_hidden_danger_count"`
	FireFightingDanger              int     `gorm:"column:fire_fighting_danger;not null;default:0" json:"fire_fighting_danger"`
	FireFightingHiddenDanger        int     `gorm:"column:fire_fighting_hidden_danger;not null;default:0" json:"fire_fighting_hidden_danger"`
	FireFightingHiddenDangerRectify int     `gorm:"column:fire_fighting_hidden_danger_rectify;not null;default:0" json:"fire_fighting_hidden_danger_rectify"`
	AccidentBulletinCount           int     `gorm:"column:accident_bulletin_count;not null;default:0" json:"accident_bulletin_count"`
	CompanyID                       int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	// JSON 键沿用 Laravel HasMany 关系名（驼峰），前端表单读写都用这两个 key。
	MonthStatisticControl  []MonthStatisticControl  `gorm:"foreignKey:MonthStatisticID" json:"monthStatisticControl,omitempty"`
	MonthStatisticAccident []MonthStatisticAccident `gorm:"foreignKey:MonthStatisticID" json:"monthStatisticAccident,omitempty"`
	Company                *Company                 `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (MonthStatistic) TableName() string { return "month_statistics" }

type MonthStatisticControl struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	GovernTitle      *string `gorm:"column:govern_title" json:"govern_title"`
	GovernSupplies   *string `gorm:"column:govern_supplies" json:"govern_supplies"`
	GovernUser       *string `gorm:"column:govern_user" json:"govern_user"`
	GovernDate       *string `gorm:"column:govern_date" json:"govern_date"`
	GovernText       *string `gorm:"column:govern_text" json:"govern_text"`
	GovernMoney      *string `gorm:"column:govern_money" json:"govern_money"`
	MonthStatisticID int64   `gorm:"column:month_statistic_id;not null" json:"month_statistic_id"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (MonthStatisticControl) TableName() string { return "month_statistic_controls" }

type MonthStatisticAccident struct {
	ID                 int64   `gorm:"primaryKey;column:id" json:"id"`
	HappenTime         *Date   `gorm:"column:happen_time;type:date" json:"happen_time"`
	Department         *string `gorm:"column:department" json:"department"`
	Name               *string `gorm:"column:name" json:"name"`
	Sex                *string `gorm:"column:sex" json:"sex"`
	Age                *int    `gorm:"column:age" json:"age"`
	WorkType           *string `gorm:"column:work_type" json:"work_type"`
	WorkAge            *int    `gorm:"column:work_age" json:"work_age"`
	EducationLevel     *string `gorm:"column:education_level" json:"education_level"`
	DegreeInjury       *string `gorm:"column:degree_injury" json:"degree_injury"`
	Type               *string `gorm:"column:type" json:"type"`
	DirectCause        *string `gorm:"column:direct_cause" json:"direct_cause"`
	Content            *string `gorm:"column:content" json:"content"`
	DaysOff            *int    `gorm:"column:days_off" json:"days_off"`
	DirectLossProperty *string `gorm:"column:direct_loss_property" json:"direct_loss_property"`
	MonthStatisticID   int64   `gorm:"column:month_statistic_id;not null" json:"month_statistic_id"`
	CompanyID          int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (MonthStatisticAccident) TableName() string { return "month_statistic_accidents" }

// YearStatistic — 年度数据上报主表。Laravel 把 safety_funding / fire_fighting
// 存成字符串（保留前端用户原始输入，比如带千分位、单位），Go 端沿用相同语义；
// statistics 汇总接口在 service 层用 ::numeric 强转后再 SUM。
type YearStatistic struct {
	ID               int64     `gorm:"primaryKey;column:id" json:"id"`
	CompanyName      *string   `gorm:"column:company_name" json:"company_name"`
	BaseArea         *string   `gorm:"column:base_area" json:"base_area"`
	EnterpriseType   *string   `gorm:"column:enterprise_type" json:"enterprise_type"`
	SafetyFunding    *string   `gorm:"column:safety_funding" json:"safety_funding"`
	FireFighting     *string   `gorm:"column:fire_fighting" json:"fire_fighting"`
	ComplianceTime   *DateTime `gorm:"column:compliance_time" json:"compliance_time"`
	RevaluationTime  *DateTime `gorm:"column:revaluation_time" json:"revaluation_time"`
	CompanyID        int64     `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	YearStatisticKid []YearStatisticKid `gorm:"foreignKey:YearStatisticID" json:"yearStatisticKid,omitempty"`
	Company          *Company           `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (YearStatistic) TableName() string { return "year_statistics" }

type YearStatisticKid struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                *string `gorm:"column:name" json:"name"`
	Type                int     `gorm:"column:type;not null;default:1" json:"type"`
	ShouldObtainPeople  *int    `gorm:"column:should_obtain_people" json:"should_obtain_people"`
	AlreadyObtainPeople *int    `gorm:"column:already_obtain_people" json:"already_obtain_people"`
	NoObtainPeople      *string `gorm:"column:no_obtain_people" json:"no_obtain_people"`
	TrainPeople         *int    `gorm:"column:train_people" json:"train_people"`
	YearStatisticID     int64   `gorm:"column:year_statistic_id;not null" json:"year_statistic_id"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (YearStatisticKid) TableName() string { return "year_statistic_kids" }
