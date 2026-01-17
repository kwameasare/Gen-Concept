package model

type Organization struct {
	BaseModel
	Name             string `gorm:"type:string;size:100;not null"`
	Description      string `gorm:"type:string;size:1000"`
	Domain           string `gorm:"type:string;size:100;unique"`
	SubscriptionPlan string `gorm:"type:string;size:50;default:'Free'"`
	Users            []User `gorm:"foreignKey:OrganizationID"`
	Teams            []Team `gorm:"foreignKey:OrganizationID"`
}
