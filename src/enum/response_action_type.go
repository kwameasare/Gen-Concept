package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ResponseActionType int

const (
	SetField ResponseActionType = iota
	ResponseBusinessValidation
)

func (r ResponseActionType) String() string {
	return [...]string{"SET_FIELD", "BUSINESS_VALIDATION"}[r]
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
	case "BUSINESS_VALIDATION":
		*r = ResponseBusinessValidation
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
	case "BUSINESS_VALIDATION":
		*r = ResponseBusinessValidation
	default:
		return fmt.Errorf("invalid ResponseActionType: %s", actionStr)
	}

	return nil
}
