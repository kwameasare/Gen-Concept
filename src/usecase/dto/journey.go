package dto

import (
	"gen-concept-api/enum" // Update this import path to the correct one

	"github.com/google/uuid"
)

type Journey struct {
	UUID                uuid.UUID                `json:"uuid"`
	ProjectUUID         uuid.UUID                `json:"projectUUID"`
	ProgrammingLanguage enum.ProgrammingLanguage `json:"programmingLanguage"`
	EntityJourneys      []EntityJourney          `json:"entityJourneys"`
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
	BackendJourney  []JourneyStep      `json:"backendJourney"`
	Filters         []Filter           `json:"filters,omitempty"`
	Sort            []Sort             `json:"sort,omitempty"`
}

type JourneyStep struct {
	UUID            uuid.UUID                  `json:"uuid"`
	Index           int                        `json:"index"`
	Type            string                     `json:"type"`
	Description     string                     `json:"description,omitempty"`
	FieldsInvolved  []FieldInvolved            `json:"fieldsInvolved,omitempty"`
	Condition       string                     `json:"condition,omitempty"`
	AbortOnFail     bool                       `json:"abortOnFail,omitempty"`
	Error           string                     `json:"error,omitempty"`
	Curl            string                     `json:"curl,omitempty"`
	SampleResponse  string                     `json:"sampleResponse,omitempty"`
	Retry           bool                       `json:"retry,omitempty"`
	RetryCount      int                        `json:"retryCount,omitempty"`
	RetryInterval   int                        `json:"retryInterval,omitempty"`
	RetryConditions []RetryCondition           `json:"retryConditions,omitempty"`
	ResponseActions []ResponseAction           `json:"responseActions,omitempty"`
	DBAction        enum.DbActionType          `json:"dbAction,omitempty"`
	CacheAction     enum.CacheActionType       `json:"cacheAction,omitempty"`
	FailOnNotFound  bool                       `json:"failOnNotFound,omitempty"`
	Channels        []enum.NotificationChannel `json:"channels,omitempty"`
	Message         string                     `json:"message,omitempty"`
	Recipients      []string                   `json:"recipients,omitempty"`

	// Hierarchical fields
	SubSteps []JourneyStep `json:"subSteps,omitempty"`
	Level    string        `json:"level,omitempty"`
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
	UUID                 uuid.UUID               `json:"uuid"`
	Index                int                     `json:"index"`
	Type                 enum.ResponseActionType `json:"type"`
	FieldID              string                  `json:"fieldId,omitempty"`
	Value                string                  `json:"value,omitempty"`
	Description          string                  `json:"description,omitempty"`
	FieldsInvolved       []ResFieldInvolved      `json:"fieldsInvolved,omitempty"`
	Condition            string                  `json:"condition,omitempty"`
	AbortOnFail          bool                    `json:"abortOnFail,omitempty"`
	Error                string                  `json:"error,omitempty"`
	NestedResponseAction *ResponseAction         `json:"nestedResponseAction,omitempty"`
}
type ResFieldInvolved struct {
	UUID   uuid.UUID `json:"uuid"`
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Source string    `json:"source,omitempty"`
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
