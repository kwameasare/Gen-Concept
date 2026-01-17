package model

type Team struct {
	BaseModel
	Name           string `gorm:"size:100;not null"`
	Description    string `gorm:"size:1000"`
	OrganizationID uint   `gorm:"not null"`
	Organization   Organization
	Users          []User `gorm:"many2many:team_users"`
}
