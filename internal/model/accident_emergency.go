package model

// 事故应急模块：事故快报 + 4 个应急子模块。集团视图（/bloc-accident-bulletins）
// 跟生产视图共用 AccidentBulletin 表，仅在 service 层按 company subtree 过滤。
//
// AccidentBulletin / AccidentBulletinKid    事故快报主表 + 伤亡人员明细（整体替换语义）
// EmergencyDrill                            应急演练
// EmergencySupplies                         应急物资（无子表）
// EmergencyTeam / EmergencyTeamKid          应急队伍 + 队伍成员明细（整体替换语义）

// AccidentBulletin — 事故快报主表。Kids 用 JSON key "kids" 跟 Laravel $appends=['kids'] 对齐，
// 这样前端 Index.tsx config 里的 `key: 'kids', type: 'sub-table'` 直接读得到。
type AccidentBulletin struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	Type                *string `gorm:"column:type" json:"type"`
	DirectCause         *string `gorm:"column:direct_cause" json:"direct_cause"`
	IndirectCause       *string `gorm:"column:indirect_cause" json:"indirect_cause"`
	AccidentLevel       *string `gorm:"column:accident_level" json:"accident_level"`
	Department          *string `gorm:"column:department" json:"department"`
	HappenTime          *Date   `gorm:"column:happen_time;type:date" json:"happen_time"`
	HappenArea          *string `gorm:"column:happen_area" json:"happen_area"`
	Content             *string `gorm:"column:content" json:"content"`
	DegreeInjury        *string `gorm:"column:degree_injury" json:"degree_injury"`
	DirectLossProperty  *string `gorm:"column:direct_loss_property" json:"direct_loss_property"`
	GaugeLossProperty   *string `gorm:"column:gauge_loss_property" json:"gauge_loss_property"`
	StaffDie            *int32  `gorm:"column:staff_die" json:"staff_die"`
	StaffHurt           *int32  `gorm:"column:staff_hurt" json:"staff_hurt"`
	NotStaffDie         *int32  `gorm:"column:not_staff_die" json:"not_staff_die"`
	NotStaffHurt        *int32  `gorm:"column:not_staff_hurt" json:"not_staff_hurt"`
	AccidentReportURL   *string `gorm:"column:accident_report_url" json:"accident_report_url"`
	AccidentPhotoURL    *string `gorm:"column:accident_photo_url" json:"accident_photo_url"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Kids    []AccidentBulletinKid `gorm:"foreignKey:AccidentBulletinID" json:"kids"`
	Company *Company              `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (AccidentBulletin) TableName() string { return "accident_bulletins" }

type AccidentBulletinKid struct {
	ID                 int64   `gorm:"primaryKey;column:id" json:"id"`
	Name               *string `gorm:"column:name" json:"name"`
	DegreeInjury       *string `gorm:"column:degree_injury" json:"degree_injury"`
	Department         *string `gorm:"column:department" json:"department"`
	Age                *int32  `gorm:"column:age" json:"age"`
	Sex                *string `gorm:"column:sex" json:"sex"`
	WorkType           *string `gorm:"column:work_type" json:"work_type"`
	WorkAge            *int32  `gorm:"column:work_age" json:"work_age"`
	EducationLevel     *string `gorm:"column:education_level" json:"education_level"`
	DaysOff            *int32  `gorm:"column:days_off" json:"days_off"`
	AccidentBulletinID int64   `gorm:"column:accident_bulletin_id;not null" json:"accident_bulletin_id"`
	CompanyID          int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (AccidentBulletinKid) TableName() string { return "accident_bulletin_kids" }

// EmergencyDrill — 应急演练。
type EmergencyDrill struct {
	ID                      int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                    *string `gorm:"column:name" json:"name"`
	DrillTime               *Date   `gorm:"column:drill_time;type:date" json:"drill_time"`
	DrillLocal              *string `gorm:"column:drill_local" json:"drill_local"`
	OrganizationDepartment  *string `gorm:"column:organization_department" json:"organization_department"`
	JoinDepartment          *string `gorm:"column:join_department" json:"join_department"`
	DrillNumber             *string `gorm:"column:drill_number" json:"drill_number"`
	DrillType               *string `gorm:"column:drill_type" json:"drill_type"`
	DrillForm               *string `gorm:"column:drill_form" json:"drill_form"`
	DrillKind               *string `gorm:"column:drill_kind" json:"drill_kind"`
	Event                   *string `gorm:"column:event" json:"event"`
	AnnexURL                *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID               int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (EmergencyDrill) TableName() string { return "emergency_drills" }

// EmergencySupplies — 应急物资（Laravel 模型保持复数命名 EmergencySupplies）。
type EmergencySupplies struct {
	ID         int64   `gorm:"primaryKey;column:id" json:"id"`
	Name       *string `gorm:"column:name" json:"name"`
	Number     *string `gorm:"column:number" json:"number"`
	SaveUser   *string `gorm:"column:save_user" json:"save_user"`
	Tel        *string `gorm:"column:tel" json:"tel"`
	Comment    *string `gorm:"column:comment" json:"comment"`
	CompanyID  int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (EmergencySupplies) TableName() string { return "emergency_supplies" }

// EmergencyTeam — 应急队伍主表。Kids JSON key "kids" 沿用 Laravel $appends=['kids']，
// 关联字段 emergency_team_id 跟 Laravel hasMany 第二参数对齐。
type EmergencyTeam struct {
	ID         int64   `gorm:"primaryKey;column:id" json:"id"`
	Name       *string `gorm:"column:name" json:"name"`
	TeamNumber *string `gorm:"column:team_number" json:"team_number"`
	SaveUser   *string `gorm:"column:save_user" json:"save_user"`
	Tel        *string `gorm:"column:tel" json:"tel"`
	CompanyID  int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Kids []EmergencyTeamKid `gorm:"foreignKey:EmergencyTeamID" json:"kids"`
}

func (EmergencyTeam) TableName() string { return "emergency_teams" }

type EmergencyTeamKid struct {
	ID              int64   `gorm:"primaryKey;column:id" json:"id"`
	Name            *string `gorm:"column:name" json:"name"`
	Department      *string `gorm:"column:department" json:"department"`
	Content         *string `gorm:"column:content" json:"content"`
	Tel             *string `gorm:"column:tel" json:"tel"`
	EmergencyTeamID int64   `gorm:"column:emergency_team_id;not null" json:"emergency_team_id"`
	CompanyID       int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (EmergencyTeamKid) TableName() string { return "emergency_team_kids" }
