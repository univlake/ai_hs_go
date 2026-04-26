package model

import (
	"time"

	"gorm.io/gorm"
)

// SimpleDictType is the shared shape every base-data dictionary uses on
// the wire (frontend `system/*-types/Index.tsx` is one CrudPage that reads
// id/title/sort across every dictionary). The Go side stores them as
// embeddable struct so each type still has its own GORM model + table name.
type SimpleDictType struct {
	ID        int64          `gorm:"primaryKey;column:id" json:"id"`
	Title     string         `gorm:"column:title;size:255;not null" json:"title"`
	Sort      int            `gorm:"column:sort;not null;default:0" json:"sort"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

type RecordType struct{ SimpleDictType }

func (RecordType) TableName() string { return "record_types" }

type ProcessType struct{ SimpleDictType }

func (ProcessType) TableName() string { return "process_types" }

type SpaceWorkType struct{ SimpleDictType }

func (SpaceWorkType) TableName() string { return "space_work_types" }

// SafetyManagementType is intentionally NOT embedded into SimpleDictType:
// GORM's Preload over a relation with an embedded primary key sometimes
// fails to flatten the FK reference (we hit this with
// SafetyManagementStaff.SafetyManagementType showing up nil despite
// `Preload("SafetyManagementType")`). A flat shape with the PK at the
// top level lets the FK reference resolve cleanly.
type SafetyManagementType struct {
	ID    int64  `gorm:"primaryKey;column:id" json:"id"`
	Title string `gorm:"column:title;size:255;not null" json:"title"`
	// Sort / timestamps exist in the table but are intentionally
	// dropped from the JSON output here — when this struct ships as a
	// nested relation on SafetyManagementStaff, the detail drawer
	// renders unknown sub-fields raw, and "sort: 0 / created_at: …"
	// is just noise. Dropdown queries go through DictRow (separate
	// flat shape) so they aren't affected.
	Sort      int            `gorm:"column:sort;not null;default:0" json:"-"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (SafetyManagementType) TableName() string { return "safety_management_types" }

type DangerousCraftType struct{ SimpleDictType }

func (DangerousCraftType) TableName() string { return "dangerous_craft_types" }

// RegulationLibraryType is multi-level (tree). pid points at the parent
// id (0 = top-level), all_name is the slash-joined full path
// ("安全生产法律 / 行政法规 / 安全许可证条例") used by the list display
// and by the regulation-libraries selector. company_id mirrors Laravel
// even though we currently keep this dictionary global (always 0).
type RegulationLibraryType struct {
	ID         int64          `gorm:"primaryKey;column:id" json:"id"`
	Title      string         `gorm:"column:title;size:255;not null" json:"title"`
	Pid        int64          `gorm:"column:pid;not null;default:0" json:"pid"`
	AllName    *string        `gorm:"column:all_name" json:"all_name"`
	Sort       int            `gorm:"column:sort;not null;default:0" json:"sort"`
	CompanyID  int64          `gorm:"column:company_id;not null;default:0" json:"company_id"`
	UserID     int64          `gorm:"column:user_id;not null;default:0" json:"user_id"`
	LastUserID int64          `gorm:"column:last_user_id;not null;default:0" json:"last_user_id"`
	CreatedAt  time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (RegulationLibraryType) TableName() string { return "regulation_library_types" }

// ChecklistLibrary is a tenant-scoped checklist asset (typically an
// uploaded Excel/PDF + preview image). Used by the 知识库 / 检查表 page.
type ChecklistLibrary struct {
	ID              int64   `gorm:"primaryKey;column:id" json:"id"`
	Title           string  `gorm:"column:title;size:255;not null" json:"title"`
	Status          int16   `gorm:"column:status;not null;default:1" json:"status"`
	PreviewAnnexURL *string `gorm:"column:preview_annex_url" json:"preview_annex_url"`
	AnnexURL        *string `gorm:"column:annex_url" json:"annex_url"`
	CompanyID       int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (ChecklistLibrary) TableName() string { return "checklist_libraries" }
