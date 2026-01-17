package model

import (
	"gen-concept-api/enum"

	"github.com/google/uuid"
)

type Journey struct {
	BaseModel
	ProjectUUID         uuid.UUID `gorm:"type:uuid;not null"`
	ProgrammingLanguage enum.ProgrammingLanguage
	ReturnEntityID      string          `gorm:"size:150"`
	EntityJourneys      []EntityJourney `gorm:"foreignKey:JourneyID"`
}

type EntityJourney struct {
	BaseModel
	EntityID   string `gorm:"not null;size:150"`
	EntityName string `gorm:"not null;size:150"`
	JourneyID  uint
	Journey    Journey     `gorm:"foreignKey:JourneyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Operations []Operation `gorm:"foreignKey:EntityJourneyID"`
}

type Operation struct {
	BaseModel
	Type            enum.OperationType `gorm:"type:varchar(50)"`
	Name            string             `gorm:"not null;size:150"`
	Description     string             `gorm:"size:1000"`
	EntityJourneyID uint
	EntityJourney   EntityJourney `gorm:"foreignKey:EntityJourneyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FrontendJourney []interface{} `gorm:"type:text;serializer:json"`
	BackendJourney  []JourneyStep `gorm:"foreignKey:OperationID"`
	Filters         []Filter      `gorm:"foreignKey:OperationID"`
	Sort            []Sort        `gorm:"foreignKey:OperationID"`
}

type JourneyStep struct {
	BaseModel
	Index           int             `gorm:"not null"`
	Type            string          `gorm:"type:varchar(50)"`
	Description     string          `gorm:"size:1000"`
	FieldsInvolved  []FieldInvolved `gorm:"foreignKey:JourneyStepID"`
	Condition       string          `gorm:"size:1000"`
	AbortOnFail     bool
	Error           string `gorm:"size:1000"`
	Curl            string `gorm:"size:1000"`
	SampleResponse  string `gorm:"size:1000"`
	Retry           bool
	RetryCount      int
	RetryInterval   int
	RetryConditions []RetryCondition     `gorm:"foreignKey:JourneyStepID"`
	ResponseActions []ResponseAction     `gorm:"foreignKey:JourneyStepID"`
	DBAction        enum.DbActionType    `gorm:"type:varchar(50)"`
	CacheAction     enum.CacheActionType `gorm:"type:varchar(50)"`
	FailOnNotFound  bool
	Channels        []enum.NotificationChannel `gorm:"type:text;serializer:json"`
	Message         string                     `gorm:"size:1000"`
	Recipients      []string                   `gorm:"type:text;serializer:json"`
	OperationID     uint
	Operation       Operation `gorm:"foreignKey:OperationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type FieldInvolved struct {
	BaseModel
	ID            string `gorm:"not null;size:150"`
	Name          string `gorm:"not null;size:150"`
	Source        string `gorm:"size:150"`
	JourneyStepID uint
	JourneyStep   JourneyStep `gorm:"foreignKey:JourneyStepID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type RetryCondition struct {
	BaseModel
	Condition     string `gorm:"not null;size:1000"`
	Error         string `gorm:"size:1000"`
	JourneyStepID uint
	JourneyStep   JourneyStep `gorm:"foreignKey:JourneyStepID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ResponseAction struct {
	BaseModel
	Index                  int                     `gorm:"not null"`
	Type                   enum.ResponseActionType `gorm:"type:varchar(50)"`
	FieldID                string                  `gorm:"size:150"`
	Value                  string                  `gorm:"size:1000"`
	Description            string                  `gorm:"size:1000"`
	FieldsInvolved         []ResFieldInvolved      `gorm:"foreignKey:ResponseActionID"`
	Condition              string                  `gorm:"size:1000"`
	AbortOnFail            bool
	Error                  string `gorm:"size:1000"`
	JourneyStepID          uint
	JourneyStep            JourneyStep `gorm:"foreignKey:JourneyStepID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	NestedResponseActionID *uint
	NestedResponseAction   *ResponseAction `gorm:"foreignKey:NestedResponseActionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type ResFieldInvolved struct {
	BaseModel
	ID               string `gorm:"not null;size:150"`
	Name             string `gorm:"not null;size:150"`
	Source           string `gorm:"size:150"`
	ResponseActionID uint
	ResponseAction   ResponseAction `gorm:"foreignKey:ResponseActionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
type Filter struct {
	BaseModel
	Name        string            `gorm:"not null;size:150"`
	Type        enum.FilterType   `gorm:"type:varchar(50)"`
	FieldID     string            `gorm:"not null;size:150"`
	MaxRange    *Range            `gorm:"foreignKey:FilterID"`
	MinRange    *Range            `gorm:"foreignKey:FilterID"`
	Error       string            `gorm:"size:1000"`
	Operator    enum.OperatorType `gorm:"type:varchar(50)"`
	OperationID uint
	Operation   Operation `gorm:"foreignKey:OperationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Range struct {
	BaseModel
	Value    int    `gorm:"not null"`
	Unit     string `gorm:"size:50"`
	FilterID uint
	Filter   Filter `gorm:"foreignKey:FilterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Sort struct {
	BaseModel
	FieldID     string `gorm:"not null;size:150"`
	OperationID uint
	Operation   Operation `gorm:"foreignKey:OperationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
