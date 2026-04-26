package model


// SafetyArchitecture is the per-company 安全架构 record (max 1 per
// company). Stores OSS paths for thumbnail + main annex; the handler
// pre-signs them on read.
type SafetyArchitecture struct {
	ID           int64   `gorm:"primaryKey;column:id" json:"id"`
	ThumbnailURL *string `gorm:"column:thumbnail_url" json:"thumbnail_url"`
	AnnexURL     *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID    *int64  `gorm:"column:company_id" json:"company_id"`

	// Filled in at the handler boundary, never persisted.
	SignedAnnexURL *string `gorm:"-" json:"signed_annex_url,omitempty"`

	Audit
}

func (SafetyArchitecture) TableName() string { return "safety_architectures" }

// FilingPermission is one 备案/许可 record. record_type_id is a soft FK
// into the record_types dictionary (set null on dictionary delete).
type FilingPermission struct {
	ID                int64      `gorm:"primaryKey;column:id" json:"id"`
	RecordTypeID      *int64     `gorm:"column:record_type_id" json:"record_type_id"`
	IssuingDepartment *string    `gorm:"column:issuing_department" json:"issuing_department"`
	FilingRange       *string    `gorm:"column:filing_range" json:"filing_range"`
	FilingTime        *Date `gorm:"column:filing_time;type:date" json:"filing_time"`
	ValidPeriod       *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	Comment           *string    `gorm:"column:comment" json:"comment"`
	AnnexURL          *string    `gorm:"column:annex_url" json:"annex_url"`
	CompanyID         int64      `gorm:"column:company_id;not null" json:"company_id"`

	// Eager-loaded on list/detail so the frontend can render the type
	// label without a second round-trip.
	RecordType *RecordType `gorm:"foreignKey:RecordTypeID;references:ID" json:"record_type,omitempty"`

	Audit
}

func (FilingPermission) TableName() string { return "filing_permissions" }

// Building is one 建筑物及构筑物 row. company_id allows null in the
// Laravel schema (the migration omits NOT NULL) so leaders can leave
// drafts; we mirror that to avoid breaking existing data flows.
type Building struct {
	ID        int64   `gorm:"primaryKey;column:id" json:"id"`
	Title     *string `gorm:"column:title" json:"title"`
	Structure *string `gorm:"column:structure" json:"structure"`
	Area      *string `gorm:"column:area" json:"area"`
	FireClass *string `gorm:"column:fire_class" json:"fire_class"`
	CompanyID *int64  `gorm:"column:company_id" json:"company_id"`

	Audit
}

func (Building) TableName() string { return "buildings" }

// SecurityInvestment is one 安全投入 line in the yearly plan. budget +
// last_year_income are NUMERIC(14,2); GORM scans them as *float64 which
// is fine for the fairly small magnitudes seen in this domain.
type SecurityInvestment struct {
	ID              int64    `gorm:"primaryKey;column:id" json:"id"`
	Year            *string  `gorm:"column:year" json:"year"`
	Type            *string  `gorm:"column:type" json:"type"`
	Title           *string  `gorm:"column:title" json:"title"`
	Content         *string  `gorm:"column:content" json:"content"`
	Budget          *float64 `gorm:"column:budget" json:"budget"`
	LastYearIncome  *float64 `gorm:"column:last_year_income" json:"last_year_income"`
	CompanyID       int64    `gorm:"column:company_id;not null" json:"company_id"`

	// Computed at list/detail time from budget / last_year_income.
	// Empty when last_year_income is missing or zero — never persisted.
	Scale string `gorm:"-" json:"scale,omitempty"`

	Audit
}

func (SecurityInvestment) TableName() string { return "security_investments" }

// CertificateStatus is one row in the 持证情况 panel — type partitions
// the rows into 三类人员 (1-3) and 消防控制室 (4-6) buckets that the
// frontend renders as separate tabs.
type CertificateStatus struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                *string `gorm:"column:name" json:"name"`
	Type                int16   `gorm:"column:type;not null;default:1" json:"type"`
	ShouldObtainPeople  *int    `gorm:"column:should_obtain_people" json:"should_obtain_people"`
	AlreadyObtainPeople *int    `gorm:"column:already_obtain_people" json:"already_obtain_people"`
	NoObtainPeople      *string `gorm:"column:no_obtain_people" json:"no_obtain_people"`
	TrainPeople         *int    `gorm:"column:train_people" json:"train_people"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (CertificateStatus) TableName() string { return "certificate_statuses" }

// ProductionStandardization is per-company 0~1 row, captures the level
// + validity of the company's safety standardization certificate.
type ProductionStandardization struct {
	ID          int64      `gorm:"primaryKey;column:id" json:"id"`
	Grade       *string    `gorm:"column:grade" json:"grade"`
	GetTime     *Date `gorm:"column:get_time;type:date" json:"get_time"`
	ValidPeriod *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	AnnexURL    *string    `gorm:"column:annex_url" json:"annex_url"`
	CompanyID   int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (ProductionStandardization) TableName() string { return "production_standardizations" }

// OccupationalHealthSafety mirrors ProductionStandardization but tracks
// the OHSAS / ISO 45001 system specifically (different cert body /
// standard edition).
type OccupationalHealthSafety struct {
	ID                int64      `gorm:"primaryKey;column:id" json:"id"`
	CertificationBody *string    `gorm:"column:certification_body" json:"certification_body"`
	StandardEdition   *string    `gorm:"column:standard_edition" json:"standard_edition"`
	GetTime           *Date `gorm:"column:get_time;type:date" json:"get_time"`
	ValidPeriod       *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	AnnexURL          *string    `gorm:"column:annex_url" json:"annex_url"`
	CompanyID         int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (OccupationalHealthSafety) TableName() string { return "occupational_health_safeties" }

// FundingExecution links to a SecurityInvestment plan and records actual
// expenditure. The list endpoint computes a `proportion` field on the
// fly (this row's budget / sum of all securityInvestments for the year).
type FundingExecution struct {
	ID                   int64    `gorm:"primaryKey;column:id" json:"id"`
	Year                 *string  `gorm:"column:year" json:"year"`
	SecurityInvestmentID *int64   `gorm:"column:security_investment_id" json:"security_investment_id"`
	SimpleContent        *string  `gorm:"column:simple_content" json:"simple_content"`
	Content              *string  `gorm:"column:content" json:"content"`
	Budget               *float64 `gorm:"column:budget" json:"budget"`
	CompanyID            int64    `gorm:"column:company_id;not null" json:"company_id"`

	// Eager-loaded plan + computed share — both are surfaces only on
	// list/detail (not stored).
	SecurityInvestment *SecurityInvestment `gorm:"foreignKey:SecurityInvestmentID;references:ID" json:"security_investment,omitempty"`
	Proportion         string              `gorm:"-" json:"proportion,omitempty"`

	Audit
}

func (FundingExecution) TableName() string { return "funding_executions" }

// SafetyManagementStaff is one row of 安全管理人员 (the staff itself,
// not the type dictionary). type_id points at safety_management_types;
// the relation is eager-loaded on list/detail so the frontend's
// "人员类型" column / detail drawer can render the title without a
// second round-trip.
type SafetyManagementStaff struct {
	ID                     int64      `gorm:"primaryKey;column:id" json:"id"`
	Name                   *string    `gorm:"column:name" json:"name"`
	SafetyManagementTypeID int64      `gorm:"column:safety_management_type_id;not null;default:1" json:"safety_management_type_id"`
	Tel                    *string    `gorm:"column:tel" json:"tel"`
	CertTitle              *string    `gorm:"column:cert_title" json:"cert_title"`
	CertNumber             *string    `gorm:"column:cert_number" json:"cert_number"`
	ValidPeriod            *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	CompanyID              int64      `gorm:"column:company_id;not null" json:"company_id"`

	SafetyManagementType *SafetyManagementType `gorm:"foreignKey:SafetyManagementTypeID;references:ID" json:"safety_management_type,omitempty"`

	Audit
}

func (SafetyManagementStaff) TableName() string { return "safety_management_staff" }

// EliminationControlStaff is the 消防控制室人员 台账. Laravel added a
// `type` column later but the frontend never surfaces it, so we keep
// the DB column (forward compatible) but hide it from the JSON
// response — otherwise the detail drawer renders an unlabeled raw
// integer the user has no business seeing.
type EliminationControlStaff struct {
	ID          int64      `gorm:"primaryKey;column:id" json:"id"`
	Name        *string    `gorm:"column:name" json:"name"`
	Type        int16      `gorm:"column:type;not null;default:1" json:"-"`
	CertTitle   *string    `gorm:"column:cert_title" json:"cert_title"`
	CertNumber  *string    `gorm:"column:cert_number" json:"cert_number"`
	ValidPeriod *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	CompanyID   int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (EliminationControlStaff) TableName() string { return "elimination_control_staff" }

