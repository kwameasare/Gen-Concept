package model

import (
	"gen-concept-api/enum"

	"github.com/google/uuid"
)

type Project struct {
	BaseModel
	ProjectName        string          `gorm:"unique;not null;size:150"`
	ProjectDescription string          `gorm:"size:1000"`
	ProjectType        enum.ProjectType `gorm:"type:varchar(20)"`
	IsMultiTenant      bool
	IsMultiLingual     bool
	Entities           []Entity `gorm:"foreignKey:ProjectUuid;references:Uuid"`
}

type Entity struct {
	BaseModel
	EntityName                 string `gorm:"not null;size:150"`
	EntityDescription          string `gorm:"size:1000"`
	ProjectUuid                uuid.UUID
	Project                    Project  `gorm:"foreignKey:ProjectUuid;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ImplementsRBAC             bool
	IsAuthenticationRequired   bool
	ImplementsAudit            bool
	ImplementsChangeManagement bool
	IsReadOnly                 bool
	IsIndependentEntity        bool
	DependsOnEntities          []DependsOnEntity `gorm:"foreignKey:EntityID"`
	Version                    string            `gorm:"size:50"`
	IsBackendOnly              bool
	PreferredDB                enum.PreferredDB
	ModeOfDBInteraction        enum.DbInteraction
	EntityFields               []EntityField `gorm:"foreignKey:EntityID"`
}

type DependsOnEntity struct {
	BaseModel
	EntityName   string `gorm:"not null;size:150"`
	EntityID     uint
	Entity	   Entity `gorm:"foreignKey:EntityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	FieldName    string `gorm:"not null;size:250"`
	RelationType enum.RelationType `gorm:"type:varchar(20)"`
}

type EntityField struct {
	BaseModel
	FieldName            string `gorm:"not null;size:250"`
	DisplayName          string `gorm:"size:250"`
	FieldDescription     string `gorm:"size:1000"`
	EntityID             uint
	Entity               Entity `gorm:"foreignKey:EntityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	FieldType            enum.DataType `gorm:"type:varchar(20)"`
	IsMandatory          bool
	IsUnique             bool
	IsReadOnly           bool
	IsSensitive          bool
	IsEditable           bool
	IsDerived            bool
	IsCollection         bool
	CollectionType       enum.CollectionType `gorm:"type:varchar(20)"`
	CollectionItemType   enum.CollectionItemType `gorm:"type:varchar(20)"`
	NestedCollectionItemType   enum.CollectionItemType `gorm:"type:varchar(20)"`
	CollectionEntity     string `gorm:"size:150"`
	IsEnum               bool
	EnumValues           []string `gorm:"type:json"`
	DerivativeType       enum.DerivativeType `gorm:"type:varchar(20)"`
	DerivativeExpression string
	IsBackendOnly        bool
	DisplayStatus        enum.DisplayStatus
	SampleData           string          `gorm:"size:1000"`
	InputValidations      []InputValidation `gorm:"foreignKey:EntityFieldID"`
}

type InputValidation struct {
	BaseModel
	Description        string `gorm:"size:1000"`
	EntityFieldID            uint
	EntityField               EntityField `gorm:"foreignKey:EntityFieldID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AbortOnFailure     bool
	CustomErrorMessage string `gorm:"size:1000"`
}
