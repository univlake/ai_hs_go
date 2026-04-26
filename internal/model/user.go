package model

type User struct {
	ID                     int64   `gorm:"primaryKey;column:id" json:"id"`
	Username               string  `gorm:"column:username;size:30;not null" json:"username"`
	Password               string  `gorm:"column:password;size:255;not null" json:"-"`
	CompanyID              *int64  `gorm:"column:company_id" json:"company_id"`
	Nickname               string  `gorm:"column:nickname;size:100;not null;default:''" json:"nickname"`
	HeaderImg              *string `gorm:"column:header_img" json:"header_img"`
	Phone                  *string `gorm:"column:phone" json:"phone"`
	DepartmentName         *string `gorm:"column:department_name" json:"department_name"`
	Email                  *string `gorm:"column:email" json:"email"`
	IsExpert               int16   `gorm:"column:is_expert;not null;default:0" json:"is_expert"`
	RoleID                 int64   `gorm:"column:role_id;not null;default:0" json:"role_id"`
	Status                 int16   `gorm:"column:status;not null;default:1" json:"status"`
	IsPlatformAdmin        bool    `gorm:"column:is_platform_admin;not null;default:false" json:"is_platform_admin"`
	MustChangeCredentials  bool    `gorm:"column:must_change_credentials;not null;default:false" json:"must_change_credentials"`
	WeappOpenID            *string `gorm:"column:weapp_openid" json:"weapp_openid"`
	WeixinSessionKey       *string `gorm:"column:weixin_session_key" json:"-"`
	LoginTimes             *string `gorm:"column:login_times" json:"login_times"`

	Audit
}

func (User) TableName() string { return "users" }
