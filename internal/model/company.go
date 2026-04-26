package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseArea struct {
	AreaID             int64   `gorm:"primaryKey;column:area_id" json:"area_id"`
	StatisticalAreaID  *int64  `gorm:"column:statistical_area_id" json:"statistical_area_id"`
	AreaName           string  `gorm:"column:area_name;size:100;not null" json:"area_name"`
	AreaParentID       *int64  `gorm:"column:area_parent_id" json:"area_parent_id"`
	AreaSort           *int16  `gorm:"column:area_sort" json:"area_sort"`
	AllName            *string `gorm:"column:allname" json:"allname"`
	AllChildrenIDs     *string `gorm:"column:all_childrens_id" json:"all_childrens_id"`
	ParentIDs          *string `gorm:"column:parent_ids" json:"parent_ids"`
	ParentIDsWithSelf  *string `gorm:"column:parent_ids_withself" json:"parent_ids_withself"`
}

func (BaseArea) TableName() string { return "base_areas" }

type EnterpriseType struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"column:title;not null" json:"title"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (EnterpriseType) TableName() string { return "enterprise_types" }

type OperationalType struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"column:title;not null" json:"title"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (OperationalType) TableName() string { return "operational_types" }

type Company struct {
	ID                  int64    `gorm:"primaryKey;column:id" json:"id"`
	Title               string   `gorm:"column:title;not null" json:"title"`
	EnterpriseTypeID    *int64   `gorm:"column:enterprise_type_id" json:"enterprise_type_id"`
	OperationalTypeID   int64    `gorm:"column:operational_type_id;not null;default:1" json:"operational_type_id"`
	AreaID              *int64   `gorm:"column:area_id" json:"area_id"`
	Address             *string  `gorm:"column:address" json:"address"`
	Established         *time.Time `gorm:"column:established;type:date" json:"established"`
	CreditCode          *string  `gorm:"column:credit_code" json:"credit_code"`
	BusinessScope       *string  `gorm:"column:business_scope" json:"business_scope"`
	CompanyIntroduction *string  `gorm:"column:company_introduction" json:"company_introduction"`
	Longitude           *float64 `gorm:"column:longitude" json:"longitude"`
	Dimension           *float64 `gorm:"column:dimension" json:"dimension"`
	GeographicRange     *string  `gorm:"column:geographic_range" json:"geographic_range"`
	Acreage             *float64 `gorm:"column:acreage" json:"acreage"`
	SuperviseName       *string  `gorm:"column:supervise_name" json:"supervise_name"`
	SuperviseTel        *string  `gorm:"column:supervise_tel" json:"supervise_tel"`
	SuperviseFix        *string  `gorm:"column:supervise_fix" json:"supervise_fix"`
	SuperviseEmail      *string  `gorm:"column:supervise_email" json:"supervise_email"`
	SupervisePosition   *string  `gorm:"column:supervise_position" json:"supervise_position"`
	Status              int16    `gorm:"column:status;not null;default:0" json:"status"`
	ParentID            *int64   `gorm:"column:parent_id" json:"parent_id"`
	RootID              int64    `gorm:"column:root_id;<-:create" json:"root_id"`
	IsProduce           int16    `gorm:"column:is_produce;not null;default:0" json:"is_produce"`
	IsBloc              int16    `gorm:"column:is_bloc;not null;default:0" json:"is_bloc"`
	InherentRisk        *float64 `gorm:"column:inherent_risk" json:"inherent_risk"`
	InherentRiskGrade   *string  `gorm:"column:inherent_risk_grade" json:"inherent_risk_grade"`
	FloatRisk           *float64 `gorm:"column:float_risk" json:"float_risk"`
	FloatRiskGrade      *string  `gorm:"column:float_risk_grade" json:"float_risk_grade"`

	// Virtual JSON-only field — Laravel exposes the parent column as `pid`
	// and most legacy frontend pages still read `c.pid`. We keep the DB
	// column as `parent_id` (cleaner) and mirror it into `pid` on
	// serialization via AfterFind. Don't write to this directly.
	Pid *int64 `gorm:"-" json:"pid"`

	// Eager-loaded relations (only filled when service requests them).
	// Frontend pages read these on company list / current-info screens.
	EnterpriseType  *EnterpriseType  `gorm:"foreignKey:EnterpriseTypeID;references:ID" json:"enterprise_type,omitempty"`
	OperationalType *OperationalType `gorm:"foreignKey:OperationalTypeID;references:ID" json:"operational_type,omitempty"`

	// Self-relation to the immediate parent. Approval list previews this
	// so reviewers know who the new company will sit under.
	Parent *Company `gorm:"foreignKey:ParentID;references:ID" json:"parent,omitempty"`

	Audit
}

func (Company) TableName() string { return "companies" }

// AfterFind keeps the Laravel-era `pid` JSON field in sync with `parent_id`.
// Runs on Find / First / Take, so list, detail, and tree responses all
// get the alias for free.
func (c *Company) AfterFind(_ *gorm.DB) error {
	c.Pid = c.ParentID
	return nil
}

// CompanyGroup is a tenant-defined collection of companies — used by
//集团视图 (bloc view) to aggregate risk / stats across hand-picked
// children rather than the natural tree.
type CompanyGroup struct {
	ID        int64  `gorm:"primaryKey;column:id" json:"id"`
	Name      string `gorm:"column:name;not null" json:"name"`
	CompanyID int64  `gorm:"column:company_id;not null" json:"company_id"`

	// Eager-loaded members. Frontend-facing JSON is intentionally minimal:
	// `kid_companies` is the flat company list the page renders as chips,
	// matching Laravel's `with('kidCompanies')` shape.
	Kids         []CompanyGroupKid `gorm:"foreignKey:CompanyGroupID" json:"-"`
	KidCompanies []Company         `gorm:"-" json:"kid_companies"`

	Audit
}

func (CompanyGroup) TableName() string { return "company_groups" }

// CompanyGroupKid is the join row connecting a CompanyGroup to one
// member company. Pure association — no per-row author (the parent
// CompanyGroup carries that), so we only track the timestamps + soft
// delete and skip the full Audit embed (the column set must match
// migration 0009 exactly or GORM-generated INSERTs error with
// "column user_id of relation company_group_kids does not exist").
type CompanyGroupKid struct {
	ID             int64          `gorm:"primaryKey;column:id" json:"id"`
	CompanyGroupID int64          `gorm:"column:company_group_id;not null" json:"company_group_id"`
	CompanyID      int64          `gorm:"column:company_id;not null" json:"company_id"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	// Optional eager-loaded company snapshot; included so the grouping page
	// can render the member's title / status without an extra round-trip.
	Company *Company `gorm:"foreignKey:CompanyID;references:ID" json:"company,omitempty"`
}

func (CompanyGroupKid) TableName() string { return "company_group_kids" }

type Personnel struct {
	ID                     int64 `gorm:"primaryKey;column:id" json:"id"`
	TotalPeople            int   `gorm:"column:total_people;not null;default:0" json:"total_people"`
	Staff                  int   `gorm:"column:staff;not null;default:0" json:"staff"`
	Service                int   `gorm:"column:service;not null;default:0" json:"service"`
	Manager                int   `gorm:"column:manager;not null;default:0" json:"manager"`
	HighSchoolDiploma      int   `gorm:"column:high_school_diploma;not null;default:0" json:"high_school_diploma"`
	SecurityOfficer        int   `gorm:"column:security_officer;not null;default:0" json:"security_officer"`
	SafetyEngineer         int   `gorm:"column:safety_engineer;not null;default:0" json:"safety_engineer"`
	SafetyEngineeringMajor int   `gorm:"column:safety_engineering_major;not null;default:0" json:"safety_engineering_major"`
	FireEngineer           int   `gorm:"column:fire_engineer;not null;default:0" json:"fire_engineer"`
	Other                  int   `gorm:"column:other;not null;default:0" json:"other"`
	CompanyID              int64 `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (Personnel) TableName() string { return "personnels" }
