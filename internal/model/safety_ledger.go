package model

type SpecialEquipment struct {
	ID               int64      `gorm:"primaryKey;column:id" json:"id"`
	EquipmentName    *string    `gorm:"column:equipment_name" json:"equipment_name"`
	SpecialTypeID    *int64     `gorm:"column:special_type_id" json:"special_type_id"`
	UtilizedLocation *string    `gorm:"column:utilized_location" json:"utilized_location"`
	NormModel1       *string    `gorm:"column:norm_model1" json:"norm_model1"`
	NormModel2       *string    `gorm:"column:norm_model2" json:"norm_model2"`
	NormModel3       *string    `gorm:"column:norm_model3" json:"norm_model3"`
	NormModel4       *string    `gorm:"column:norm_model4" json:"norm_model4"`
	SerialNumber     *string    `gorm:"column:serial_number" json:"serial_number"`
	UseTime          *Date `gorm:"column:use_time;type:date" json:"use_time"`
	ValidPeriod      *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	InspectUnit      *string    `gorm:"column:inspect_unit" json:"inspect_unit"`
	MaintainInfo     *string    `gorm:"column:maintain_info" json:"maintain_info"`
	Comment          *string    `gorm:"column:comment" json:"comment"`
	CompanyID        int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (SpecialEquipment) TableName() string { return "special_equipment" }

type SpecialEquipmentStaff struct {
	ID          int64      `gorm:"primaryKey;column:id" json:"id"`
	Name        *string    `gorm:"column:name" json:"name"`
	CertTitle   *string    `gorm:"column:cert_title" json:"cert_title"`
	CertNumber  *string    `gorm:"column:cert_number" json:"cert_number"`
	ValidPeriod *Date `gorm:"column:valid_period;type:date" json:"valid_period"`
	CompanyID   int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (SpecialEquipmentStaff) TableName() string { return "special_equipment_staff" }

type DangerousEquipment struct {
	ID               int64      `gorm:"primaryKey;column:id" json:"id"`
	EquipmentName    *string    `gorm:"column:equipment_name" json:"equipment_name"`
	UtilizedLocation *string    `gorm:"column:utilized_location" json:"utilized_location"`
	NormModel        *string    `gorm:"column:norm_model" json:"norm_model"`
	SerialNumber     *string    `gorm:"column:serial_number" json:"serial_number"`
	UseTime          *Date `gorm:"column:use_time;type:date" json:"use_time"`
	AutomationType   *string    `gorm:"column:automation_type" json:"automation_type"`
	MaintainInfo     *string    `gorm:"column:maintain_info" json:"maintain_info"`
	Comment          *string    `gorm:"column:comment" json:"comment"`
	CompanyID        int64      `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (DangerousEquipment) TableName() string { return "dangerous_equipment" }

type HazardousChemical struct {
	ID            int64   `gorm:"primaryKey;column:id" json:"id"`
	Name          *string `gorm:"column:name" json:"name"`
	Alias         *string `gorm:"column:alias" json:"alias"`
	IsSupervision int16   `gorm:"column:is_supervision;not null;default:0" json:"is_supervision"`
	IsToxin       int16   `gorm:"column:is_toxin;not null;default:0" json:"is_toxin"`
	IsExplosion   int16   `gorm:"column:is_explosion;not null;default:0" json:"is_explosion"`
	IsEasyExplode int16   `gorm:"column:is_easy_explode;not null;default:0" json:"is_easy_explode"`
	IsNarcotics   int16   `gorm:"column:is_narcotics;not null;default:0" json:"is_narcotics"`
	IsMonitor     int16   `gorm:"column:is_monitor;not null;default:0" json:"is_monitor"`
	IsMixing      int16   `gorm:"column:is_mixing;not null;default:0" json:"is_mixing"`
	CAS           *string `gorm:"column:cas" json:"cas"`
	Position      *string `gorm:"column:position" json:"position"`
	Temperature   *string `gorm:"column:temperature" json:"temperature"`
	Storage       *string `gorm:"column:storage" json:"storage"`
	YearUse       *string `gorm:"column:year_use" json:"year_use"`
	YearUseUnit   *string `gorm:"column:year_use_unit" json:"year_use_unit"`
	StorageUnit   *string `gorm:"column:storage_unit" json:"storage_unit"`
	Pressure      *string `gorm:"column:pressure" json:"pressure"`
	PressureUnit  *string `gorm:"column:pressure_unit" json:"pressure_unit"`
	FireClass     *string `gorm:"column:fire_class" json:"fire_class"`
	MatterState   *string `gorm:"column:matter_state" json:"matter_state"`
	MSDS          *string `gorm:"column:msds" json:"msds"`
	CompanyID     int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Ingredients []ChemicalIngredient `gorm:"foreignKey:HazardousChemicalID" json:"ingredients,omitempty"`
}

func (HazardousChemical) TableName() string { return "hazardous_chemicals" }

type ChemicalIngredient struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                *string `gorm:"column:name" json:"name"`
	Content             *string `gorm:"column:content" json:"content"`
	CAS                 *string `gorm:"column:cas" json:"cas"`
	HazardousChemicalID int64   `gorm:"column:hazardous_chemical_id;not null" json:"hazardous_chemical_id"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (ChemicalIngredient) TableName() string { return "chemical_ingredients" }
