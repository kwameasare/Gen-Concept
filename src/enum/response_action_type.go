package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ResponseActionType int

const (
	SetField ResponseActionType = iota
	Validation
	ProceedToNextStep
	TransformResponse
	UpdateContext
	StoreResponseData
	TriggerNotification
	InvokeAnotherAPI
	LogResponse
	ConditionalBranch
	AggregateData
)

func (r ResponseActionType) String() string {
	names := [...]string{
		"SET_FIELD",
		"VALIDATION",
		"PROCEED_TO_NEXT_STEP",
		"TRANSFORM_RESPONSE",
		"UPDATE_CONTEXT",
		"STORE_RESPONSE_DATA",
		"TRIGGER_NOTIFICATION",
		"INVOKE_ANOTHER_API",
		"LOG_RESPONSE",
		"CONDITIONAL_BRANCH",
		"AGGREGATE_DATA",
	}
	if r < SetField || int(r) >= len(names) {
		return "UNKNOWN"
	}
	return names[r]
}

func (r ResponseActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *ResponseActionType) UnmarshalJSON(data []byte) error {
	var actionStr string
	if err := json.Unmarshal(data, &actionStr); err != nil {
		return err
	}

	switch actionStr {
	case "SET_FIELD":
		*r = SetField
	case "VALIDATION":
		*r = Validation
	case "PROCEED_TO_NEXT_STEP":
		*r = ProceedToNextStep
	case "TRANSFORM_RESPONSE":
		*r = TransformResponse
	case "UPDATE_CONTEXT":
		*r = UpdateContext
	case "STORE_RESPONSE_DATA":
		*r = StoreResponseData
	case "TRIGGER_NOTIFICATION":
		*r = TriggerNotification
	case "INVOKE_ANOTHER_API":
		*r = InvokeAnotherAPI
	case "LOG_RESPONSE":
		*r = LogResponse
	case "CONDITIONAL_BRANCH":
		*r = ConditionalBranch
	case "AGGREGATE_DATA":
		*r = AggregateData
	default:
		return fmt.Errorf("invalid ResponseActionType: %s", actionStr)
	}

	return nil
}

func (r ResponseActionType) Value() (driver.Value, error) {
	return r.String(), nil
}

func (r *ResponseActionType) Scan(value interface{}) error {
	if value == nil {
		*r = SetField
		return nil
	}

	var actionStr string

	switch v := value.(type) {
	case string:
		actionStr = v
	case []byte:
		actionStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for ResponseActionType: %T", value)
	}

	switch actionStr {
	case "SET_FIELD":
		*r = SetField
	case "VALIDATION":
		*r = Validation
	case "PROCEED_TO_NEXT_STEP":
		*r = ProceedToNextStep
	case "TRANSFORM_RESPONSE":
		*r = TransformResponse
	case "UPDATE_CONTEXT":
		*r = UpdateContext
	case "STORE_RESPONSE_DATA":
		*r = StoreResponseData
	case "TRIGGER_NOTIFICATION":
		*r = TriggerNotification
	case "INVOKE_ANOTHER_API":
		*r = InvokeAnotherAPI
	case "LOG_RESPONSE":
		*r = LogResponse
	case "CONDITIONAL_BRANCH":
		*r = ConditionalBranch
	case "AGGREGATE_DATA":
		*r = AggregateData
	default:
		return fmt.Errorf("invalid ResponseActionType: %s", actionStr)
	}

	return nil
}
