package model

import (
	"gen-concept-api/enum"
)

type Project struct {
	ProjectName         string `gorm:"unique,not null;size:150"`
	ProjectDescription  string `gorm:"size:1000"`
	ProjectType         enum.ProjectType
	IsMultiTenant       bool
	IsMultiLingual      bool
	ProgrammingLanguage enum.ProgrammingLanguage
	Entities            []Entity `gorm:"foreignKey:ProjectId"`
	BaseModel
}

type Entity struct {
	EntityName                 string `gorm:"not null;size:150"`
	EntityDescription          string `gorm:"size:1000"`
	ProjectId                  int
	ImplementsRBAC             bool
	IsAuthenticationRequired   bool
	ImplementsAudit            bool
	ImplementsChangeManagement bool
	IsReadOnly                 bool
	IsIndependentEntity        bool
	DependsOnEntities          []DependsOnEntity `gorm:"foreignKey:EntityId"`
	Version                    string            `gorm:"size:50"`
	IsBackendOnly              bool
	PreferredDB                enum.PreferredDB
	ModeOfDBInteraction        enum.DbInteraction
	EntityFields               []EntityField `gorm:"foreignKey:EntityId"`
	BaseModel
}

type DependsOnEntity struct {
	EntityName   string `gorm:"not null;size:150"`
	EntityId     int
	FieldName    string `gorm:"not null;size:250"`
	RelationType enum.RelationType
	BaseModel
}

type EntityField struct {
	FieldName            string `gorm:"not null;size:250"`
	DisplayName          string `gorm:"size:250"`
	FieldDescription     string `gorm:"size:1000"`
	EntityId             int
	FieldType            enum.DataType
	IsMandatory          bool
	IsUnique             bool
	IsReadOnly           bool
	IsSensitive          bool
	IsEditable           bool
	IsDerived            bool
	IsCollection         bool
	CollectionType       enum.CollectionType
	IsEnum               bool
	EnumValues           []string
	DerivativeType       enum.DerivativeType
	DerivativeExpression string
	IsBackendOnly        bool
	DisplayStatus        enum.DisplayStatus
	SampleData           string          `gorm:"size:1000"`
	InputValidation      InputValidation `gorm:"foreignKey:FieldId"`
	BaseModel
}

type InputValidation struct {
	Description        string `gorm:"size:1000"`
	FieldId            int
	AbortOnFailure     bool
	CustomErrorMessage string `gorm:"size:1000"`
	BaseModel
}
