package dto

import (
	"gen-concept-api/enum" 
	"gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

type Journey struct {
	UUID           uuid.UUID       `json:"uuid"`
	ProjectUUID    uuid.UUID       `json:"projectUUID"`
	EntityJourneys []EntityJourney `json:"entityJourneys"`
}

type EntityJourney struct {
	UUID       uuid.UUID   `json:"uuid"`
	EntityID   string      `json:"entityId"`
	EntityName string      `json:"entityName"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	UUID            uuid.UUID          `json:"uuid"`
	Type            enum.OperationType `json:"type"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	FrontendJourney []interface{}      `json:"frontendJourney"`
	BackendJourney  []BackendJourney   `json:"backendJourney"`
	Filters         []Filter           `json:"filters,omitempty"`
	Sort            []Sort             `json:"sort,omitempty"`
}

type BackendJourney struct {
	UUID            uuid.UUID                   `json:"uuid"`
	Index           int                         `json:"index"`
	Type            enum.BackendJourneyStepType `json:"type"`
	Description     string                      `json:"description,omitempty"`
	FieldsInvolved  []FieldInvolved             `json:"fieldsInvolved,omitempty"`
	Condition       string                      `json:"condition,omitempty"`
	AbortOnFail     bool                        `json:"abortOnFail,omitempty"`
	Error           string                      `json:"error,omitempty"`
	Curl            string                      `json:"curl,omitempty"`
	SampleResponse  string                      `json:"sampleResponse,omitempty"`
	Retry           bool                        `json:"retry,omitempty"`
	RetryCount      int                         `json:"retryCount,omitempty"`
	RetryInterval   int                         `json:"retryInterval,omitempty"`
	RetryConditions []RetryCondition            `json:"retryConditions,omitempty"`
	ResponseActions []ResponseAction            `json:"responseActions,omitempty"`
	DBAction        enum.DbActionType           `json:"dbAction,omitempty"`
	Channels        []enum.NotificationChannel  `json:"channels,omitempty"`
	Message         string                      `json:"message,omitempty"`
	Recipients      []string                    `json:"recipients,omitempty"`
}

type FieldInvolved struct {
	UUID   uuid.UUID `json:"uuid"`
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Source string    `json:"source,omitempty"`
}

type RetryCondition struct {
	UUID      uuid.UUID `json:"uuid"`
	Condition string    `json:"condition"`
	Error     string    `json:"error"`
}

type ResponseAction struct {
	UUID           uuid.UUID               `json:"uuid"`
	Index          int                     `json:"index"`
	Type           enum.ResponseActionType `json:"type"`
	FieldID        string                  `json:"fieldId,omitempty"`
	Value          string                  `json:"value,omitempty"`
	Description    string                  `json:"description,omitempty"`
	FieldsInvolved []ResFieldInvolved      `json:"fieldsInvolved,omitempty"`
	Condition      string                  `json:"condition,omitempty"`
	AbortOnFail    bool                    `json:"abortOnFail,omitempty"`
	Error          string                  `json:"error,omitempty"`
}
type ResFieldInvolved struct {
	UUID             uuid.UUID `json:"uuid"`
	ID               string    `gorm:"not null;size:150"`
	Name             string    `gorm:"not null;size:150"`
	Source           string    `gorm:"size:150"`
	ResponseActionID uint
	ResponseAction   ResponseAction `gorm:"foreignKey:JourneyStepID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
type Filter struct {
	UUID     uuid.UUID         `json:"uuid"`
	Name     string            `json:"name"`
	Type     enum.FilterType   `json:"type"`
	FieldID  string            `json:"fieldId"`
	MaxRange *Range            `json:"maxRange,omitempty"`
	MinRange *Range            `json:"minRange,omitempty"`
	Error    string            `json:"error,omitempty"`
	Operator enum.OperatorType `json:"operator,omitempty"`
}

type Range struct {
	UUID  uuid.UUID `json:"uuid"`
	Value int       `json:"value"`
	Unit  string    `json:"unit"`
}

type Sort struct {
	UUID    uuid.UUID `json:"uuid"`
	FieldID string    `json:"fieldId"`
}

func (j *Journey) FromUsecaseJourneyDTO(ucJourney *dto.Journey) {
	j.UUID = ucJourney.UUID
	j.ProjectUUID = ucJourney.ProjectUUID
	for _, ucEntityJourney := range ucJourney.EntityJourneys {
		var entityJourney EntityJourney
		entityJourney.FromUsecaseEntityJourneyDTO(&ucEntityJourney)
		j.EntityJourneys = append(j.EntityJourneys, entityJourney)
	}
}

func (ej *EntityJourney) FromUsecaseEntityJourneyDTO(ucEntityJourney *dto.EntityJourney) {
	ej.UUID = ucEntityJourney.UUID
	ej.EntityID = ucEntityJourney.EntityID
	ej.EntityName = ucEntityJourney.EntityName
	for _, ucOperation := range ucEntityJourney.Operations {
		var operation Operation
		operation.FromUsecaseOperationDTO(&ucOperation)
		ej.Operations = append(ej.Operations, operation)
	}
}

func (op *Operation) FromUsecaseOperationDTO(ucOperation *dto.Operation) {
	op.UUID = ucOperation.UUID
	op.Type = ucOperation.Type
	op.Name = ucOperation.Name
	op.Description = ucOperation.Description
	op.FrontendJourney = ucOperation.FrontendJourney
	for _, ucBackendJourney := range ucOperation.BackendJourney {
		var backendJourney BackendJourney
		backendJourney.FromUsecaseBackendJourneyDTO(&ucBackendJourney)
		op.BackendJourney = append(op.BackendJourney, backendJourney)
	}
	for _, ucFilter := range ucOperation.Filters {
		var filter Filter
		filter.FromUsecaseFilterDTO(&ucFilter)
		op.Filters = append(op.Filters, filter)
	}
	for _, ucSort := range ucOperation.Sort {
		var sort Sort
		sort.FromUsecaseSortDTO(&ucSort)
		op.Sort = append(op.Sort, sort)
	}
}

func (bj *BackendJourney) FromUsecaseBackendJourneyDTO(ucBackendJourney *dto.JourneyStep) {
	bj.UUID = ucBackendJourney.UUID
	bj.Index = ucBackendJourney.Index
	bj.Type = ucBackendJourney.Type
	bj.Description = ucBackendJourney.Description
	bj.Condition = ucBackendJourney.Condition
	bj.AbortOnFail = ucBackendJourney.AbortOnFail
	bj.Error = ucBackendJourney.Error
	bj.Curl = ucBackendJourney.Curl
	bj.SampleResponse = ucBackendJourney.SampleResponse
	bj.Retry = ucBackendJourney.Retry
	bj.RetryCount = ucBackendJourney.RetryCount
	bj.RetryInterval = ucBackendJourney.RetryInterval
	bj.DBAction = ucBackendJourney.DBAction
	bj.Channels = make([]enum.NotificationChannel, len(ucBackendJourney.Channels))
	copy(bj.Channels, ucBackendJourney.Channels)
	bj.Message = ucBackendJourney.Message
	bj.Recipients = ucBackendJourney.Recipients
	for _, ucFieldInvolved := range ucBackendJourney.FieldsInvolved {
		var fieldInvolved FieldInvolved
		fieldInvolved.FromUsecaseFieldInvolvedDTO(&ucFieldInvolved)
		bj.FieldsInvolved = append(bj.FieldsInvolved, fieldInvolved)
	}
	for _, ucRetryCondition := range ucBackendJourney.RetryConditions {
		var retryCondition RetryCondition
		retryCondition.FromUsecaseRetryConditionDTO(&ucRetryCondition)
		bj.RetryConditions = append(bj.RetryConditions, retryCondition)
	}
	for _, ucResponseAction := range ucBackendJourney.ResponseActions {
		var responseAction ResponseAction
		responseAction.FromUsecaseResponseActionDTO(&ucResponseAction)
		bj.ResponseActions = append(bj.ResponseActions, responseAction)
	}
}

func (fi *FieldInvolved) FromUsecaseFieldInvolvedDTO(ucFieldInvolved *dto.FieldInvolved) {
	fi.UUID = ucFieldInvolved.UUID
	fi.ID = ucFieldInvolved.ID
	fi.Name = ucFieldInvolved.Name
	fi.Source = ucFieldInvolved.Source
}

func (fi *ResFieldInvolved) FromUsecaseResponseActionFieldInvolvedDTO(ucFieldInvolved *dto.ResFieldInvolved) {
	fi.UUID = ucFieldInvolved.UUID
	fi.ID = ucFieldInvolved.ID
	fi.Name = ucFieldInvolved.Name
	fi.Source = ucFieldInvolved.Source
}

func (rc *RetryCondition) FromUsecaseRetryConditionDTO(ucRetryCondition *dto.RetryCondition) {
	rc.UUID = ucRetryCondition.UUID
	rc.Condition = ucRetryCondition.Condition
	rc.Error = ucRetryCondition.Error
}

func (ra *ResponseAction) FromUsecaseResponseActionDTO(ucResponseAction *dto.ResponseAction) {
	ra.UUID = ucResponseAction.UUID
	ra.Index = ucResponseAction.Index
	ra.Type = ucResponseAction.Type
	ra.FieldID = ucResponseAction.FieldID
	ra.Value = ucResponseAction.Value
	ra.Description = ucResponseAction.Description
	ra.Condition = ucResponseAction.Condition
	ra.AbortOnFail = ucResponseAction.AbortOnFail
	ra.Error = ucResponseAction.Error
	for _, ucFieldInvolved := range ucResponseAction.FieldsInvolved {
		var fieldInvolved ResFieldInvolved
		fieldInvolved.FromUsecaseResponseActionFieldInvolvedDTO(&ucFieldInvolved)
		ra.FieldsInvolved = append(ra.FieldsInvolved, fieldInvolved)
	}
}

func (f *Filter) FromUsecaseFilterDTO(ucFilter *dto.Filter) {
	f.UUID = ucFilter.UUID
	f.Name = ucFilter.Name
	f.Type = ucFilter.Type
	f.FieldID = ucFilter.FieldID
	f.Error = ucFilter.Error
	f.Operator = ucFilter.Operator
	if ucFilter.MaxRange != nil {
		var maxRange Range
		maxRange.FromUsecaseRangeDTO(ucFilter.MaxRange)
		f.MaxRange = &maxRange
	}
	if ucFilter.MinRange != nil {
		var minRange Range
		minRange.FromUsecaseRangeDTO(ucFilter.MinRange)
		f.MinRange = &minRange
	}
}

func (r *Range) FromUsecaseRangeDTO(ucRange *dto.Range) {
	r.UUID = ucRange.UUID
	r.Value = ucRange.Value
	r.Unit = ucRange.Unit
}

func (s *Sort) FromUsecaseSortDTO(ucSort *dto.Sort) {
	s.UUID = ucSort.UUID
	s.FieldID = ucSort.FieldID
}

func (j *Journey) ToUsecaseJourneyDTO() *dto.Journey {
	ucJourney := &dto.Journey{
		UUID:        j.UUID,
		ProjectUUID: j.ProjectUUID,
	}
	for _, entityJourney := range j.EntityJourneys {
		ucEntityJourney := entityJourney.ToUsecaseEntityJourneyDTO()
		ucJourney.EntityJourneys = append(ucJourney.EntityJourneys, *ucEntityJourney)
	}
	return ucJourney
}

func (ej *EntityJourney) ToUsecaseEntityJourneyDTO() *dto.EntityJourney {
	ucEntityJourney := &dto.EntityJourney{
		UUID:       ej.UUID,
		EntityID:   ej.EntityID,
		EntityName: ej.EntityName,
	}
	for _, operation := range ej.Operations {
		ucOperation := operation.ToUsecaseOperationDTO()
		ucEntityJourney.Operations = append(ucEntityJourney.Operations, *ucOperation)
	}
	return ucEntityJourney
}

func (op *Operation) ToUsecaseOperationDTO() *dto.Operation {
	ucOperation := &dto.Operation{
		UUID:            op.UUID,
		Type:            op.Type,
		Name:            op.Name,
		Description:     op.Description,
		FrontendJourney: op.FrontendJourney,
	}
	for _, backendJourney := range op.BackendJourney {
		ucBackendJourney := backendJourney.ToUsecaseBackendJourneyDTO()
		ucOperation.BackendJourney = append(ucOperation.BackendJourney, *ucBackendJourney)
	}
	for _, filter := range op.Filters {
		ucFilter := filter.ToUsecaseFilterDTO()
		ucOperation.Filters = append(ucOperation.Filters, *ucFilter)
	}
	for _, sort := range op.Sort {
		ucSort := sort.ToUsecaseSortDTO()
		ucOperation.Sort = append(ucOperation.Sort, *ucSort)
	}
	return ucOperation
}

func (bj *BackendJourney) ToUsecaseBackendJourneyDTO() *dto.JourneyStep {
	ucBackendJourney := &dto.JourneyStep{
		UUID:           bj.UUID,
		Index:          bj.Index,
		Type:           bj.Type,
		Description:    bj.Description,
		Condition:      bj.Condition,
		AbortOnFail:    bj.AbortOnFail,
		Error:          bj.Error,
		Curl:           bj.Curl,
		SampleResponse: bj.SampleResponse,
		Retry:          bj.Retry,
		RetryCount:     bj.RetryCount,
		RetryInterval:  bj.RetryInterval,
		DBAction:       bj.DBAction,
		Channels:       make([]enum.NotificationChannel, len(bj.Channels)),
		Message:        bj.Message,
		Recipients:     bj.Recipients,
	}
	copy(ucBackendJourney.Channels, bj.Channels)
	for _, fieldInvolved := range bj.FieldsInvolved {
		ucFieldInvolved := fieldInvolved.ToUsecaseFieldInvolvedDTO()
		ucBackendJourney.FieldsInvolved = append(ucBackendJourney.FieldsInvolved, *ucFieldInvolved)
	}
	for _, retryCondition := range bj.RetryConditions {
		ucRetryCondition := retryCondition.ToUsecaseRetryConditionDTO()
		ucBackendJourney.RetryConditions = append(ucBackendJourney.RetryConditions, *ucRetryCondition)
	}
	for _, responseAction := range bj.ResponseActions {
		ucResponseAction := responseAction.ToUsecaseResponseActionDTO()
		ucBackendJourney.ResponseActions = append(ucBackendJourney.ResponseActions, *ucResponseAction)
	}
	return ucBackendJourney
}

func (fi *FieldInvolved) ToUsecaseFieldInvolvedDTO() *dto.FieldInvolved {
	return &dto.FieldInvolved{
		UUID:   fi.UUID,
		ID:     fi.ID,
		Name:   fi.Name,
		Source: fi.Source,
	}
}
func (fi *ResFieldInvolved) ToUsecaseResFieldInvolvedDTO() *dto.ResFieldInvolved {
	return &dto.ResFieldInvolved{
		UUID:   fi.UUID,
		ID:     fi.ID,
		Name:   fi.Name,
		Source: fi.Source,
	}
}

func (rc *RetryCondition) ToUsecaseRetryConditionDTO() *dto.RetryCondition {
	return &dto.RetryCondition{
		UUID:      rc.UUID,
		Condition: rc.Condition,
		Error:     rc.Error,
	}
}

func (ra *ResponseAction) ToUsecaseResponseActionDTO() *dto.ResponseAction {
	ucResponseAction := &dto.ResponseAction{
		UUID:        ra.UUID,
		Index:       ra.Index,
		Type:        ra.Type,
		FieldID:     ra.FieldID,
		Value:       ra.Value,
		Description: ra.Description,
		Condition:   ra.Condition,
		AbortOnFail: ra.AbortOnFail,
		Error:       ra.Error,
	}
	for _, fieldInvolved := range ra.FieldsInvolved {
		ucFieldInvolved := fieldInvolved.ToUsecaseResFieldInvolvedDTO()
		ucResponseAction.FieldsInvolved = append(ucResponseAction.FieldsInvolved, *ucFieldInvolved)
	}
	return ucResponseAction
}

func (f *Filter) ToUsecaseFilterDTO() *dto.Filter {
	ucFilter := &dto.Filter{
		UUID:     f.UUID,
		Name:     f.Name,
		Type:     f.Type,
		FieldID:  f.FieldID,
		Error:    f.Error,
		Operator: f.Operator,
	}
	if f.MaxRange != nil {
		ucMaxRange := f.MaxRange.ToUsecaseRangeDTO()
		ucFilter.MaxRange = ucMaxRange
	}
	if f.MinRange != nil {
		ucMinRange := f.MinRange.ToUsecaseRangeDTO()
		ucFilter.MinRange = ucMinRange
	}
	return ucFilter
}

func (r *Range) ToUsecaseRangeDTO() *dto.Range {
	return &dto.Range{
		UUID:  r.UUID,
		Value: r.Value,
		Unit:  r.Unit,
	}
}

func (s *Sort) ToUsecaseSortDTO() *dto.Sort {
	return &dto.Sort{
		UUID:    s.UUID,
		FieldID: s.FieldID,
	}
}
