package enum

import (
	"encoding/json"
	"fmt"
	"database/sql/driver"
)

type Conformity int

const (
	Standard Conformity = iota
	CustomConformity
	Unknown
)

// String method for pretty printing
func (c Conformity) String() string {
	return [...]string{"Standard", "Custom", "Unknown"}[c]
}

// MarshalJSON for custom JSON encoding
func (c Conformity) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON for custom JSON decoding
func (c *Conformity) UnmarshalJSON(data []byte) error {
	var conformityStr string
	if err := json.Unmarshal(data, &conformityStr); err != nil {
		return err
	}

	switch conformityStr {
	case "Standard":
		*c = Standard
	case "Custom":
		*c = CustomConformity
	default:
		*c = Unknown
	}

	return nil
}

// Implement the driver.Valuer interface
func (c Conformity) Value() (driver.Value, error) {
	return c.String(), nil
}

// Implement the sql.Scanner interface
func (c *Conformity) Scan(value interface{}) error {
	if value == nil {
		*c = Unknown
		return nil
	}

	var conformityStr string
	switch v := value.(type) {
	case string:
		conformityStr = v
	case []byte:
		conformityStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for Conformity: %T", value)
	}

	switch conformityStr {
	case "Standard":
		*c = Standard
	case "Custom":
		*c = CustomConformity
	default:
		*c = Unknown
	}

	return nil
}