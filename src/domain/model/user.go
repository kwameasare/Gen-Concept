package model

type User struct {
	BaseModel
	Username       string `gorm:"type:string;size:20;not null;unique"`
	FirstName      string `gorm:"type:string;size:15;null"`
	LastName       string `gorm:"type:string;size:25;null"`
	MobileNumber   string `gorm:"type:string;size:15;null;unique;default:null"`
	Email          string `gorm:"type:string;size:64;null;unique;default:null"`
	Password       string `gorm:"type:string;size:64;not null"`
	Enabled        bool   `gorm:"default:true"`
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserRoles      *[]UserRole
	Teams          []Team `gorm:"many2many:team_users"`
}

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:10;not null,unique"`
	UserRoles *[]UserRole
}

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId uint
	RoleId uint
}
