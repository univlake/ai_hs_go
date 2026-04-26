package model

// 会议培训：internal_meetings(+kids) / internal/external/other_trainings。
//
// 内部会议主表 + 参会人员子表（join_user_id → users.id），子表整体替换语义；
// 三类培训均无子表。training_date 是 DATE（用 model.Date 序列化为
// "YYYY-MM-DD"），meeting_date 是 TIMESTAMPTZ（model.DateTime 序列化为
// "YYYY-MM-DD HH:mm:ss"）。

type InternalMeeting struct {
	ID               int64     `gorm:"primaryKey;column:id" json:"id"`
	Name             *string   `gorm:"column:name" json:"name"`
	Theme            *string   `gorm:"column:theme" json:"theme"`
	Type             *string   `gorm:"column:type" json:"type"`
	MeetingDate      *DateTime `gorm:"column:meeting_date" json:"meeting_date"`
	Summary          *string   `gorm:"column:summary" json:"summary"`
	AnnexURL         *string   `gorm:"column:annex_url" json:"annex_url"`
	NoticeAnnexURL   *string   `gorm:"column:notice_annex_url" json:"notice_annex_url"`
	SummaryAnnexURL  *string   `gorm:"column:summary_annex_url" json:"summary_annex_url"`
	SignAnnexURL     *string   `gorm:"column:sign_annex_url" json:"sign_annex_url"`
	MaterialAnnexURL *string   `gorm:"column:material_annex_url" json:"material_annex_url"`
	PhotoAnnexURL    *string   `gorm:"column:photo_annex_url" json:"photo_annex_url"`
	CompanyID        int64     `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	Kids []InternalMeetingKid `gorm:"foreignKey:InternalMeetingID" json:"internal_meeting_kid,omitempty"`
}

func (InternalMeeting) TableName() string { return "internal_meetings" }

type InternalMeetingKid struct {
	ID                int64  `gorm:"primaryKey;column:id" json:"id"`
	InternalMeetingID int64  `gorm:"column:internal_meeting_id;not null" json:"internal_meeting_id"`
	JoinUserID        *int64 `gorm:"column:join_user_id" json:"join_user_id"`
	CompanyID         int64  `gorm:"column:company_id;not null" json:"company_id"`

	Audit

	JoinUser *User `gorm:"foreignKey:JoinUserID;references:ID" json:"join_user,omitempty"`
}

func (InternalMeetingKid) TableName() string { return "internal_meeting_kids" }

type InternalTraining struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Name             *string `gorm:"column:name" json:"name"`
	Comment          *string `gorm:"column:comment" json:"comment"`
	Object           *string `gorm:"column:object" json:"object"`
	Number           *int32  `gorm:"column:number" json:"number"`
	TrainingDate     *Date   `gorm:"column:training_date" json:"training_date"`
	TrainingType     *string `gorm:"column:training_type" json:"training_type"`
	Happening        *string `gorm:"column:happening" json:"happening"`
	AnnexURL         *string `gorm:"column:annex_url" json:"annex_url"`
	SummaryAnnexURL  *string `gorm:"column:summary_annex_url" json:"summary_annex_url"`
	SignAnnexURL     *string `gorm:"column:sign_annex_url" json:"sign_annex_url"`
	MaterialAnnexURL *string `gorm:"column:material_annex_url" json:"material_annex_url"`
	PhotoAnnexURL    *string `gorm:"column:photo_annex_url" json:"photo_annex_url"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (InternalTraining) TableName() string { return "internal_trainings" }

type ExternalTraining struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Name             *string `gorm:"column:name" json:"name"`
	Comment          *string `gorm:"column:comment" json:"comment"`
	Object           *string `gorm:"column:object" json:"object"`
	Number           *int32  `gorm:"column:number" json:"number"`
	TrainingDate     *Date   `gorm:"column:training_date" json:"training_date"`
	TrainingUnit     *string `gorm:"column:training_unit" json:"training_unit"`
	AnnexURL         *string `gorm:"column:annex_url" json:"annex_url"`
	SummaryAnnexURL  *string `gorm:"column:summary_annex_url" json:"summary_annex_url"`
	SignAnnexURL     *string `gorm:"column:sign_annex_url" json:"sign_annex_url"`
	MaterialAnnexURL *string `gorm:"column:material_annex_url" json:"material_annex_url"`
	PhotoAnnexURL    *string `gorm:"column:photo_annex_url" json:"photo_annex_url"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (ExternalTraining) TableName() string { return "external_trainings" }

type OtherTraining struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"id"`
	Name             *string `gorm:"column:name" json:"name"`
	Comment          *string `gorm:"column:comment" json:"comment"`
	Object           *string `gorm:"column:object" json:"object"`
	Number           *int32  `gorm:"column:number" json:"number"`
	TrainingDate     *Date   `gorm:"column:training_date" json:"training_date"`
	TaskContent      *string `gorm:"column:task_content" json:"task_content"`
	AnnexURL         *string `gorm:"column:annex_url" json:"annex_url"`
	SummaryAnnexURL  *string `gorm:"column:summary_annex_url" json:"summary_annex_url"`
	SignAnnexURL     *string `gorm:"column:sign_annex_url" json:"sign_annex_url"`
	MaterialAnnexURL *string `gorm:"column:material_annex_url" json:"material_annex_url"`
	PhotoAnnexURL    *string `gorm:"column:photo_annex_url" json:"photo_annex_url"`
	CompanyID        int64   `gorm:"column:company_id;not null" json:"company_id"`

	Audit
}

func (OtherTraining) TableName() string { return "other_trainings" }
