package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type OperatorType int

const (
	Equals OperatorType = iota
)

func (o OperatorType) String() string {
	return [...]string{"EQUALS"}[o]
}

func (o OperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o *OperatorType) UnmarshalJSON(data []byte) error {
	var opStr string
	if err := json.Unmarshal(data, &opStr); err != nil {
		return err
	}

	switch opStr {
	case "EQUALS":
		*o = Equals
	default:
		return fmt.Errorf("invalid OperatorType: %s", opStr)
	}

	return nil
}

func (o OperatorType) Value() (driver.Value, error) {
	return o.String(), nil
}

func (o *OperatorType) Scan(value interface{}) error {
	if value == nil {
		*o = Equals
		return nil
	}

	var opStr string

	switch v := value.(type) {
	case string:
		opStr = v
	case []byte:
		opStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for OperatorType: %T", value)
	}

	switch opStr {
	case "EQUALS":
		*o = Equals
	default:
		return fmt.Errorf("invalid OperatorType: %s", opStr)
	}

	return nil
}