package model

type MajorHazard struct {
	ID              int64   `gorm:"primaryKey;column:id" json:"id"`
	Name            *string `gorm:"column:name" json:"name"`
	Level           *string `gorm:"column:level" json:"level"`
	Position        *string `gorm:"column:position" json:"position"`
	RValue          *string `gorm:"column:r_value" json:"r_value"`
	IsSIS           int16   `gorm:"column:is_sis;not null;default:0" json:"is_sis"`
	SILLevel        *string `gorm:"column:sil_level" json:"sil_level"`
	DetectParameter *string `gorm:"column:detect_parameter" json:"detect_parameter"`
	CompanyID       int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	MajorConstitutes []MajorConstitute `gorm:"foreignKey:MajorHazardID" json:"major_constitutes,omitempty"`
}

func (MajorHazard) TableName() string { return "major_hazards" }

type MajorConstitute struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                *string `gorm:"column:name" json:"name"`
	HazardousChemicalID *int64  `gorm:"column:hazardous_chemical_id" json:"hazardous_chemical_id"`
	Number              *string `gorm:"column:number" json:"number"`
	MajorHazardID       int64   `gorm:"column:major_hazard_id;not null" json:"major_hazard_id"`
	CompanyID           int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (MajorConstitute) TableName() string { return "major_constitutes" }
