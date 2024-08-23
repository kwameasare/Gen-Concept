package dto

import (
	"gen-concept-api/enum"
	"gen-concept-api/usecase/dto"
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

func ToUseCaseProject(from Project) dto.Project {
	return dto.Project{
		ProjectName:         from.ProjectName,
		ProjectDescription:  from.ProjectDescription,
		ProjectType:         from.ProjectType,
		IsMultiTenant:       from.IsMultiTenant,
		IsMultiLingual:      from.IsMultiLingual,
		ProgrammingLanguage: from.ProgrammingLanguage,
		Entities:            ToUsecaseEntities(from.Entities),
	}
}

func ToUsecaseEntities(from []Entity) []dto.Entity {
	var entities []dto.Entity
	for _, entity := range from {
		entities = append(entities, ToUseCaseEntity(entity))
	}
	return entities
}

func ToUseCaseEntity(from Entity) dto.Entity {
	return dto.Entity{
		EntityName:                 from.EntityName,
		EntityDescription:          from.EntityDescription,
		ImplementsRBAC:             from.ImplementsRBAC,
		IsAuthenticationRequired:   from.IsAuthenticationRequired,
		ImplementsAudit:            from.ImplementsAudit,
		ImplementsChangeManagement: from.ImplementsChangeManagement,
		IsReadOnly:                 from.IsReadOnly,
		IsIndependentEntity:        from.IsIndependentEntity,
		DependsOnEntities:          ToUseCaseDependsOnEntities(from.DependsOnEntities),
		Version:                    from.Version,
		IsBackendOnly:              from.IsBackendOnly,
		PreferredDB:                from.PreferredDB,
		ModeOfDBInteraction:        from.ModeOfDBInteraction,
		EntityFields:               ToUseCaseEntityFields(from.EntityFields),
	}
}

func ToUseCaseDependsOnEntities(from []DependsOnEntity) []dto.DependsOnEntity {
	var dependsOnEntities []dto.DependsOnEntity
	for _, dependsOnEntity := range from {
		dependsOnEntities = append(dependsOnEntities, ToUseCaseDependsOnEntity(dependsOnEntity))
	}
	return dependsOnEntities
}

func ToUseCaseDependsOnEntity(from DependsOnEntity) dto.DependsOnEntity {
	return dto.DependsOnEntity{
		EntityName:   from.EntityName,
		FieldName:    from.FieldName,
		RelationType: from.RelationType,
	}
}

func ToUseCaseEntityFields(from []EntityField) []dto.EntityField {
	var entityFields []dto.EntityField
	for _, entityField := range from {
		entityFields = append(entityFields, ToUseCaseEntityField(entityField))
	}
	return entityFields
}

func ToUseCaseEntityField(from EntityField) dto.EntityField {
	return dto.EntityField{
		FieldName:            from.FieldName,
		DisplayName:          from.DisplayName,
		FieldDescription:     from.FieldDescription,
		FieldType:            from.FieldType,
		IsMandatory:          from.IsMandatory,
		IsUnique:             from.IsUnique,
		IsReadOnly:           from.IsReadOnly,
		IsSensitive:          from.IsSensitive,
		IsEditable:           from.IsEditable,
		IsDerived:            from.IsDerived,
		IsCollection:         from.IsCollection,
		CollectionType:       from.CollectionType,
		IsEnum:               from.IsEnum,
		EnumValues:           from.EnumValues,
		DerivativeType:       from.DerivativeType,
		DerivativeExpression: from.DerivativeExpression,
		IsBackendOnly:        from.IsBackendOnly,
		DisplayStatus:        from.DisplayStatus,
		SampleData:           from.SampleData,
		InputValidation:      ToUseCaseInputValidation(from.InputValidation),
	}
}

func ToUseCaseInputValidation(from InputValidation) dto.InputValidation {
	return dto.InputValidation{
		Description:        from.Description,
		AbortOnFailure:     from.AbortOnFailure,
		CustomErrorMessage: from.CustomErrorMessage,
	}
}

func ToProjectResponse(from dto.Project) Project {
	return Project{
		ProjectName:         from.ProjectName,
		ProjectDescription:  from.ProjectDescription,
		ProjectType:         from.ProjectType,
		IsMultiTenant:       from.IsMultiTenant,
		IsMultiLingual:      from.IsMultiLingual,
		ProgrammingLanguage: from.ProgrammingLanguage,
		Entities:            ToEntitiesResponse(from.Entities),
	}
}

func ToEntitiesResponse(from []dto.Entity) []Entity {
	var entities []Entity
	for _, entity := range from {
		entities = append(entities, ToEntityResponse(entity))
	}
	return entities
}

func ToEntityResponse(from dto.Entity) Entity {
	return Entity{
		EntityName:                 from.EntityName,
		EntityDescription:          from.EntityDescription,
		ImplementsRBAC:             from.ImplementsRBAC,
		IsAuthenticationRequired:   from.IsAuthenticationRequired,
		ImplementsAudit:            from.ImplementsAudit,
		ImplementsChangeManagement: from.ImplementsChangeManagement,
		IsReadOnly:                 from.IsReadOnly,
		IsIndependentEntity:        from.IsIndependentEntity,
		DependsOnEntities:          ToDependsOnEntitiesResponse(from.DependsOnEntities),
		Version:                    from.Version,
		IsBackendOnly:              from.IsBackendOnly,
		PreferredDB:                from.PreferredDB,
		ModeOfDBInteraction:        from.ModeOfDBInteraction,
		EntityFields:               ToEntityFieldsResponse(from.EntityFields),
	}
}

func ToDependsOnEntitiesResponse(from []dto.DependsOnEntity) []DependsOnEntity {
	var dependsOnEntities []DependsOnEntity
	for _, dependsOnEntity := range from {
		dependsOnEntities = append(dependsOnEntities, ToDependsOnEntityResponse(dependsOnEntity))
	}
	return dependsOnEntities
}

func ToDependsOnEntityResponse(from dto.DependsOnEntity) DependsOnEntity {
	return DependsOnEntity{
		EntityName:   from.EntityName,
		FieldName:    from.FieldName,
		RelationType: from.RelationType,
	}
}

func ToEntityFieldsResponse(from []dto.EntityField) []EntityField {

	var entityFields []EntityField
	for _, entityField := range from {
		entityFields = append(entityFields, ToEntityFieldResponse(entityField))
	}
	return entityFields
}

func ToEntityFieldResponse(from dto.EntityField) EntityField {
	return EntityField{
		FieldName:            from.FieldName,
		DisplayName:          from.DisplayName,
		FieldDescription:     from.FieldDescription,
		FieldType:            from.FieldType,
		IsMandatory:          from.IsMandatory,
		IsUnique:             from.IsUnique,
		IsReadOnly:           from.IsReadOnly,
		IsSensitive:          from.IsSensitive,
		IsEditable:           from.IsEditable,
		IsDerived:            from.IsDerived,
		IsCollection:         from.IsCollection,
		CollectionType:       from.CollectionType,
		IsEnum:               from.IsEnum,
		EnumValues:           from.EnumValues,
		DerivativeType:       from.DerivativeType,
		DerivativeExpression: from.DerivativeExpression,
		IsBackendOnly:        from.IsBackendOnly,
		DisplayStatus:        from.DisplayStatus,
		SampleData:           from.SampleData,
		InputValidation:      ToInputValidationResponse(from.InputValidation),
	}
}

func ToInputValidationResponse(from dto.InputValidation) InputValidation {
	return InputValidation{
		Description:        from.Description,
		AbortOnFailure:     from.AbortOnFailure,
		CustomErrorMessage: from.CustomErrorMessage,
	}
}
