package model

// 第三方管理：design_units / evaluation_units / skill_units(+kids)。
//
// 都是单租户内自管的"外部单位档案"，互不关联。SkillUnit 多了一张子表
// skill_unit_kids 用于记录多个资质明细（前端 kids 数组每次提交都全量发上来）。

type DesignUnit struct {
	ID                int64   `gorm:"primaryKey;column:id" json:"id"`
	Name              *string `gorm:"column:name" json:"name"`
	QualificationName *string `gorm:"column:qualification_name" json:"qualification_name"`
	Contact           *string `gorm:"column:contact" json:"contact"`
	ContactTel        *string `gorm:"column:contact_tel" json:"contact_tel"`
	AnnexURL          *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID         int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (DesignUnit) TableName() string { return "design_units" }

type EvaluationUnit struct {
	ID                    int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                  *string `gorm:"column:name" json:"name"`
	IsQualification       int16   `gorm:"column:is_qualification;not null;default:2" json:"is_qualification"`
	QualificationName     *string `gorm:"column:qualification_name" json:"qualification_name"`
	QualificationAnnexURL *string `gorm:"column:qualification_annex_url" json:"qualification_annex_url"`
	EvaluationArea        *string `gorm:"column:evaluation_area" json:"evaluation_area"`
	Contact               *string `gorm:"column:contact" json:"contact"`
	ContactTel            *string `gorm:"column:contact_tel" json:"contact_tel"`
	IsAgreement           int16   `gorm:"column:is_agreement;not null;default:2" json:"is_agreement"`
	AgreementAnnexURL     *string `gorm:"column:agreement_annex_url" json:"agreement_annex_url"`
	SupervisionURL        *string `gorm:"column:supervision_url" json:"supervision_url"`
	AnnexURL              *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID             int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (EvaluationUnit) TableName() string { return "evaluation_units" }

type SkillUnit struct {
	ID              int64   `gorm:"primaryKey;column:id" json:"id"`
	Name            *string `gorm:"column:name" json:"name"`
	ServingType     *string `gorm:"column:serving_type" json:"serving_type"`
	IsQualification int16   `gorm:"column:is_qualification;not null;default:2" json:"is_qualification"`
	EvaluationArea  *string `gorm:"column:evaluation_area" json:"evaluation_area"`
	Contact         *string `gorm:"column:contact" json:"contact"`
	ContactTel      *string `gorm:"column:contact_tel" json:"contact_tel"`
	CompanyID       int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Kids []SkillUnitKid `gorm:"foreignKey:SkillUnitID" json:"kids,omitempty"`
}

func (SkillUnit) TableName() string { return "skill_units" }

type SkillUnitKid struct {
	ID                    int64   `gorm:"primaryKey;column:id" json:"id"`
	QualificationName     *string `gorm:"column:qualification_name" json:"qualification_name"`
	QualificationAnnexURL *string `gorm:"column:qualification_annex_url" json:"qualification_annex_url"`
	SkillUnitID           int64   `gorm:"column:skill_unit_id;not null" json:"skill_unit_id"`
	CompanyID             int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (SkillUnitKid) TableName() string { return "skill_unit_kids" }
