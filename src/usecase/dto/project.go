package dto

import (
	"gen-concept-api/enum"
)

type Project struct {
	ProjectName         string                   `json:"projectName"`
	ProjectDescription  string                   `json:"projectDescription"`
	ProjectType         enum.ProjectType         `json:"projectType"`
	IsMultiTenant       bool                     `json:"isMultiTenant"`
	IsMultiLingual      bool                     `json:"isMultiLingual"`
	ProgrammingLanguage enum.ProgrammingLanguage `json:"programmingLanguage"`
	Entities            []Entity                 `json:"entities"`
}

type Entity struct {
	EntityName                 string             `json:"entityName"`
	EntityDescription          string             `json:"entityDescription"`
	ImplementsRBAC             bool               `json:"implementsRBAC"`
	IsAuthenticationRequired   bool               `json:"isAuthenticationRequired"`
	ImplementsAudit            bool               `json:"implementsAudit"`
	ImplementsChangeManagement bool               `json:"implementsChangeManagement"`
	IsReadOnly                 bool               `json:"isReadOnly"`
	IsIndependentEntity        bool               `json:"isIndependentEntity"`
	DependsOnEntities          []DependsOnEntity  `json:"dependsOnEntities"`
	Version                    string             `json:"version"`
	IsBackendOnly              bool               `json:"isBackendOnly"`
	PreferredDB                enum.PreferredDB   `json:"preferredDB"`
	ModeOfDBInteraction        enum.DbInteraction `json:"modeOfDBInteraction"`
	EntityFields               []EntityField      `json:"entityFields"`
}

type DependsOnEntity struct {
	EntityName   string            `json:"entityName"`
	FieldName    string            `json:"fieldName"`
	RelationType enum.RelationType `json:"relationType"`
}

type EntityField struct {
	FieldName            string              `json:"fieldName"`
	DisplayName          string              `json:"displayName"`
	FieldDescription     string              `json:"fieldDescription"`
	FieldType            enum.DataType       `json:"fieldType"`
	IsMandatory          bool                `json:"isMandatory"`
	IsUnique             bool                `json:"isUnique"`
	IsReadOnly           bool                `json:"isReadOnly"`
	IsSensitive          bool                `json:"isSensitive"`
	IsEditable           bool                `json:"isEditable"`
	IsDerived            bool                `json:"isDerived"`
	IsCollection         bool                `json:"isCollection"`
	CollectionType       enum.CollectionType `json:"collectionType"`
	IsEnum               bool                `json:"isEnum"`
	EnumValues           []string            `json:"enumValues"`
	DerivativeType       enum.DerivativeType `json:"derivativeType"`
	DerivativeExpression string              `json:"derivativeExpression"`
	IsBackendOnly        bool                `json:"isBackendOnly"`
	DisplayStatus        enum.DisplayStatus  `json:"displayStatus"`
	SampleData           string              `json:"sampleData"`
	InputValidation      InputValidation     `json:"inputValidation"`
}

type InputValidation struct {
	Description        string `json:"description"`
	AbortOnFailure     bool   `json:"abortOnFailure"`
	CustomErrorMessage string `json:"customErrorMessage"`
}
