package model

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type BaseModel struct {
	ID uint `gorm:"primarykey"`
	Uuid uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt *time.Time `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  gorm.DeletedAt `gorm:"type:TIMESTAMP with time zone;null"`
	CreatedBy  uint            `gorm:"not null"`
	ModifiedBy *uint `gorm:"null"`
	DeletedBy  *uint `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId uint
	if value != nil {
		userId = uint(value.(float64))
	}
	if m.Uuid == uuid.Nil {
        m.Uuid = uuid.New()
    }
	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId uint
	if value != nil {
		userId = value.(uint)
	}
	now := time.Now().UTC()
	m.ModifiedAt = &now
	m.ModifiedBy = &userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId uint
	if value != nil {
		userId = value.(uint)
	}
	now := time.Now().UTC()
	m.DeletedAt = gorm.DeletedAt{Time: now}
	m.DeletedBy = &userId
	return
}
