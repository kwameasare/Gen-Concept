package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DisplayStatus int

const (
	Summary DisplayStatus = iota
	Detail
	Hide
	Show
)

// String method for pretty printing

func (d DisplayStatus) String() string {
	return [...]string{"Summary", "Detail", "Hide","Show"}[d]
}

// MarshalJSON for custom JSON encoding

func (d DisplayStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON for custom JSON decoding

func (d *DisplayStatus) UnmarshalJSON(data []byte) error {
	var displayStatusStr string
	if err := json.Unmarshal(data, &displayStatusStr); err != nil {
		return err
	}

	switch displayStatusStr {
	case "Summary":
		*d = Summary
	case "Detail":
		*d = Detail
	case "hide":
		*d = Hide
	default:
		*d = Show
	}

	return nil
}

// Implement the driver.Valuer interface
func (d DisplayStatus) Value() (driver.Value, error) {
	return d.String(), nil
}

// Implement the sql.Scanner interface
func (d *DisplayStatus) Scan(value interface{}) error {
	if value == nil {
		*d = Show // Default value or handle as needed
		return nil
	}

	var displayStatusStr string

	switch v := value.(type) {
	case string:
		displayStatusStr = v
	case []byte:
		displayStatusStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for DisplayStatus: %T", value)
	}

	switch displayStatusStr {
	case "Summary":
		*d = Summary
	case "Detail":
		*d = Detail
	case "Hide":
		*d = Hide
	case "Show":
		*d = Show
	default:
		return fmt.Errorf("invalid DisplayStatus: %s", displayStatusStr)
	}

	return nil
}