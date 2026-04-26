package model

// MajorHazardousProcess — 重点监管危险工艺。
//
// 与 hazardous_processes 互不相干（一个对应"重点监管类"，一个对应"较大危险类"），
// 仅外键依赖完全相同：dangerous_craft_types / hazardous_chemicals / dangerous_equipment。
type MajorHazardousProcess struct {
	ID                   int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                 *string `gorm:"column:name" json:"name"`
	DangerousCraftTypeID *int64  `gorm:"column:dangerous_craft_type_id" json:"dangerous_craft_type_id"`
	Position             *string `gorm:"column:position" json:"position"`
	HazardousChemicalID  *int64  `gorm:"column:hazardous_chemical_id" json:"hazardous_chemical_id"`
	DangerousEquipmentID *int64  `gorm:"column:dangerous_equipment_id" json:"dangerous_equipment_id"`
	OperatingPeopleNum   *string `gorm:"column:operating_people_num" json:"operating_people_num"`
	Temperature          *string `gorm:"column:temperature" json:"temperature"`
	Pressure             *string `gorm:"column:pressure" json:"pressure"`
	PressureUnit         *string `gorm:"column:pressure_unit" json:"pressure_unit"`
	IsSIS                int16   `gorm:"column:is_sis;not null;default:0" json:"is_sis"`
	SILLevel             *string `gorm:"column:sil_level" json:"sil_level"`
	DetectParameter      *string `gorm:"column:detect_parameter" json:"detect_parameter"`
	CompanyID            int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	DangerousCraftType *DangerousCraftType `gorm:"foreignKey:DangerousCraftTypeID;references:ID" json:"dangerous_craft_type,omitempty"`
	HazardousChemical  *HazardousChemical  `gorm:"foreignKey:HazardousChemicalID;references:ID" json:"hazardous_chemical,omitempty"`
	DangerousEquipment *DangerousEquipment `gorm:"foreignKey:DangerousEquipmentID;references:ID" json:"dangerous_equipment,omitempty"`
}

func (MajorHazardousProcess) TableName() string { return "major_hazardous_processes" }

// HazardousProcess — 较大危险工艺。
type HazardousProcess struct {
	ID                   int64   `gorm:"primaryKey;column:id" json:"id"`
	ProcessTypeID        *int64  `gorm:"column:process_type_id" json:"process_type_id"`
	Position             *string `gorm:"column:position" json:"position"`
	HazardousChemicalID  *int64  `gorm:"column:hazardous_chemical_id" json:"hazardous_chemical_id"`
	DangerousEquipmentID *int64  `gorm:"column:dangerous_equipment_id" json:"dangerous_equipment_id"`
	OtherEquipmentName   *string `gorm:"column:other_equipment_name" json:"other_equipment_name"`
	OperatingPeopleNum   *string `gorm:"column:operating_people_num" json:"operating_people_num"`
	IsOnlineTest         int16   `gorm:"column:is_online_test;not null;default:0" json:"is_online_test"`
	IsAutomatic          int16   `gorm:"column:is_automatic;not null;default:0" json:"is_automatic"`
	IsControl            int16   `gorm:"column:is_control;not null;default:0" json:"is_control"`
	IsDust               int16   `gorm:"column:is_dust;not null;default:0" json:"is_dust"`
	CompanyID            int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	ProcessType        *ProcessType        `gorm:"foreignKey:ProcessTypeID;references:ID" json:"process_type,omitempty"`
	HazardousChemical  *HazardousChemical  `gorm:"foreignKey:HazardousChemicalID;references:ID" json:"hazardous_chemical,omitempty"`
	DangerousEquipment *DangerousEquipment `gorm:"foreignKey:DangerousEquipmentID;references:ID" json:"dangerous_equipment,omitempty"`
}

func (HazardousProcess) TableName() string { return "hazardous_processes" }

// SpaceWork — 有限空间辨识。HazardID 是逗号分隔的 hazards.id 字符串
// （前端 toggleHazard 拼装），保持 Laravel 原本的多对多 inline 方案，
// 不在 DB 层面建外键。
type SpaceWork struct {
	ID              int64   `gorm:"primaryKey;column:id" json:"id"`
	Name            *string `gorm:"column:name" json:"name"`
	Position        *string `gorm:"column:position" json:"position"`
	SpaceWorkTypeID *int64  `gorm:"column:space_work_type_id" json:"space_work_type_id"`
	Size            *string `gorm:"column:size" json:"size"`
	PeopleNum       *int32  `gorm:"column:people_num" json:"people_num"`
	HazardID        *string `gorm:"column:hazard_id" json:"hazard_id"`
	WordContent     *string `gorm:"column:word_content" json:"word_content"`
	CompanyID       int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	SpaceWorkType *SpaceWorkType `gorm:"foreignKey:SpaceWorkTypeID;references:ID" json:"space_work_type,omitempty"`
}

func (SpaceWork) TableName() string { return "space_works" }
