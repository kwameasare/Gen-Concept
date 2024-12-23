package dto

import (
	"fmt"
	"gen-concept-api/enum"
	"gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

type Project struct {
	ProjectName         string                   `json:"projectName"`
	Uuid 			  uuid.UUID                   `json:"uuid"`
	ProjectDescription  string                   `json:"projectDescription"`
	ProjectType         enum.ProjectType         `json:"projectType"`
	IsMultiTenant       bool                     `json:"isMultiTenant"`
	IsMultiLingual      bool                     `json:"isMultiLingual"`
	Entities            []Entity                 `json:"entities"`
}
func (p Project) Validate() error {
	if p.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}


	for _, entity := range p.Entities {
		if err := entity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e Entity) Validate() error {
	if e.EntityName == "" {
		return fmt.Errorf("entity name is required")
	}

	for _, field := range e.EntityFields {
		if err := field.Validate(); err != nil {
			return err
		}
	}
	return nil
}
func (ef EntityField) Validate() error {
	if ef.FieldName == "" {
		return fmt.Errorf("field name is required")
	}

	if ef.IsCollection {
		if ef.CollectionType == enum.None{
			return fmt.Errorf("collection type is required when isCollection is true")
		}
		if ef.CollectionItemType == enum.NoType {
			return fmt.Errorf("collection item type is required when isCollection is true")
		}
	}

	if ef.IsEnum {
		if len(ef.EnumValues) == 0 {
			return fmt.Errorf("enum values are required when isEnum is true")
		}
	}

	if ef.IsDerived {
		if ef.DerivativeType == enum.NotDerived {
			return fmt.Errorf("derivative type is required when isDerived is true")
		}
		if ef.DerivativeExpression == "" {
			return fmt.Errorf("derivative expression is required when isDerived is true")
		}
	}

	for _, validation := range ef.InputValidations {
		if err := validation.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (iv InputValidation) Validate() error {
	if iv.Description == "" {
		return fmt.Errorf("description is required")
	}
	return nil
}


type Entity struct {
	EntityName                 string             `json:"entityName"`
	Uuid 					  uuid.UUID          `json:"uuid"`
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
	Uuid 					  uuid.UUID          `json:"uuid"`
	FieldName    string            `json:"fieldName"`
	RelationType enum.RelationType `json:"relationType"`
}

type EntityField struct {
	FieldName            string              `json:"fieldName"`
	Uuid 					  uuid.UUID          `json:"uuid"`
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
	CollectionItemType	 enum.CollectionItemType `json:"collectionItemType"`
	NestedCollectionItemType enum.CollectionItemType `json:"nestedCollectionItemType"`
	CollectionEntity    string              `json:"collectionEntity"`
	IsEnum               bool                `json:"isEnum"`
	EnumValues           []string            `json:"enumValues"`
	DerivativeType       enum.DerivativeType `json:"derivativeType"`
	DerivativeExpression string              `json:"derivativeExpression"`
	IsBackendOnly        bool                `json:"isBackendOnly"`
	DisplayStatus        enum.DisplayStatus  `json:"displayStatus"`
	SampleData           string              `json:"sampleData"`
	InputValidations      []InputValidation     `json:"inputValidations"`
}

type InputValidation struct {
	Description        string `json:"description"`
	Uuid 					  uuid.UUID          `json:"uuid"`
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
		CollectionItemType:   from.CollectionItemType,
		NestedCollectionItemType: from.NestedCollectionItemType,
		CollectionEntity:     from.CollectionEntity,
		IsEnum:               from.IsEnum,
		EnumValues:           from.EnumValues,
		DerivativeType:       from.DerivativeType,
		DerivativeExpression: from.DerivativeExpression,
		IsBackendOnly:        from.IsBackendOnly,
		DisplayStatus:        from.DisplayStatus,
		SampleData:           from.SampleData,
		InputValidations:      ToUseCaseInputValidation(from.InputValidations),
	}
}

func ToUseCaseInputValidation(from []InputValidation) []dto.InputValidation {
	
	validations:= []dto.InputValidation{}
	for _, inputValidation := range from {
		validations = append(validations, dto.InputValidation{
			Description:        inputValidation.Description,
			AbortOnFailure:     inputValidation.AbortOnFailure,
			CustomErrorMessage: inputValidation.CustomErrorMessage,
		})
	}
	return validations
}

func ToProjectResponse(from dto.Project) Project {
	return Project{
		ProjectName:         from.ProjectName,
		Uuid:                from.Uuid,
		ProjectDescription:  from.ProjectDescription,
		ProjectType:         from.ProjectType,
		IsMultiTenant:       from.IsMultiTenant,
		IsMultiLingual:      from.IsMultiLingual,
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
		Uuid:                       from.Uuid,
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
		Uuid:         from.Uuid,
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
		Uuid:                 from.Uuid,
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
		CollectionItemType:  from.CollectionItemType,
		NestedCollectionItemType: from.NestedCollectionItemType,
		CollectionEntity:     from.CollectionEntity,
		IsEnum:               from.IsEnum,
		EnumValues:           from.EnumValues,
		DerivativeType:       from.DerivativeType,
		DerivativeExpression: from.DerivativeExpression,
		IsBackendOnly:        from.IsBackendOnly,
		DisplayStatus:        from.DisplayStatus,
		SampleData:           from.SampleData,
		InputValidations:      ToInputValidationResponse(from.InputValidations),
	}
}

func ToInputValidationResponse(from []dto.InputValidation) []InputValidation {
	validations:= []InputValidation{}
	for _, inputValidation := range from {
		validations = append(validations, InputValidation{
			Description:        inputValidation.Description,
			Uuid:               inputValidation.Uuid,
			AbortOnFailure:     inputValidation.AbortOnFailure,
			CustomErrorMessage: inputValidation.CustomErrorMessage,
		})
	}
	return validations
}
