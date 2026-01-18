package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type JourneyNode struct {
	BaseModel
	JourneyID    uuid.UUID      `gorm:"type:uuid;index;not null"`
	Type         string         `gorm:"size:50"`
	Label        string         `gorm:"size:255"`
	Level        string         `gorm:"size:20;index"` // HIGH, MEDIUM, LOW
	ParentNodeID *uuid.UUID     `gorm:"type:uuid;index"`
	BlueprintID  *uuid.UUID     `gorm:"type:uuid"`
	Metadata     datatypes.JSON `gorm:"type:jsonb"`
}

type JourneyEdge struct {
	BaseModel
	JourneyID uuid.UUID `gorm:"type:uuid;index;not null"`
	SourceID  uuid.UUID `gorm:"type:uuid;not null"`
	TargetID  uuid.UUID `gorm:"type:uuid;not null"`
	Label     string    `gorm:"size:50"`
}
