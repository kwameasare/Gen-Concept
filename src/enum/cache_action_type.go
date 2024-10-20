package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type CacheActionType int

const (
	CacheRead CacheActionType = iota
	CacheInsert
)

func (c CacheActionType) String() string {
	return [...]string{"READ", "INSERT"}[c]
}

func (c CacheActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CacheActionType) UnmarshalJSON(data []byte) error {
	var actionStr string
	if err := json.Unmarshal(data, &actionStr); err != nil {
		return err
	}

	switch actionStr {
	case "READ":
		*c = CacheRead
	case "INSERT":
		*c = CacheInsert
	default:
		return fmt.Errorf("invalid CacheActionType: %s", actionStr)
	}

	return nil
}

func (c CacheActionType) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c *CacheActionType) Scan(value interface{}) error {
	if value == nil {
		*c = CacheRead
		return nil
	}

	var actionStr string

	switch v := value.(type) {
	case string:
		actionStr = v
	case []byte:
		actionStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for CacheActionType: %T", value)
	}

	switch actionStr {
	case "READ":
		*c = CacheRead
	case "INSERT":
		*c = CacheInsert
	default:
		return fmt.Errorf("invalid CacheActionType: %s", actionStr)
	}

	return nil
}