package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type OperationType int

const (
	Create OperationType = iota
	Read
	Update
	Delete
	Custom
	ReadById
)

func (o OperationType) String() string {
	return [...]string{"CREATE", "READ", "UPDATE", "DELETE", "CUSTOM", "READ_BY_ID"}[o]
}

func (o OperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o *OperationType) UnmarshalJSON(data []byte) error {
	var opStr string
	if err := json.Unmarshal(data, &opStr); err != nil {
		return err
	}

	switch opStr {
	case "CREATE":
		*o = Create
	case "READ":
		*o = Read
	case "UPDATE":
		*o = Update
	case "DELETE":
		*o = Delete
	case "CUSTOM":
		*o = Custom
	case "READ_BY_ID":
		*o = ReadById
	default:
		return fmt.Errorf("invalid OperationType: %s", opStr)
	}

	return nil
}

func (o OperationType) Value() (driver.Value, error) {
	return o.String(), nil
}

func (o *OperationType) Scan(value interface{}) error {
	if value == nil {
		*o = Create
		return nil
	}

	var opStr string

	switch v := value.(type) {
	case string:
		opStr = v
	case []byte:
		opStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for OperationType: %T", value)
	}

	switch opStr {
	case "CREATE":
		*o = Create
	case "READ":
		*o = Read
	case "UPDATE":
		*o = Update
	case "DELETE":
		*o = Delete
	case "CUSTOM":
		*o = Custom
	case "READ_BY_ID":
		*o = ReadById
	default:
		return fmt.Errorf("invalid OperationType: %s", opStr)
	}

	return nil
}
