package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DbInteraction int

const (
	ORM DbInteraction = iota
	DBFunction
	RawSQL
	NoSql
	NoDB
)

// String method for pretty printing
func (d DbInteraction) String() string {
	return [...]string{"ORM", "DB FUNCTION", "RAW SQL, No SQL", "No DB"}[d]
}

// MarshalJSON for custom JSON encoding
func (d DbInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON for custom JSON decoding
func (d *DbInteraction) UnmarshalJSON(data []byte) error {
	var dbInteractionStr string
	if err := json.Unmarshal(data, &dbInteractionStr); err != nil {
		return err
	}

	switch dbInteractionStr {
	case "ORM":
		*d = ORM
	case "DB FUNCTION":
		*d = DBFunction
	case "RAW SQL":
		*d = RawSQL
	case "No SQL":
		*d = NoSql
	default:
		*d = NoDB
	}

	return nil
}


// Implement the driver.Valuer interface
func (d DbInteraction) Value() (driver.Value, error) {
	return d.String(), nil
}

// Implement the sql.Scanner interface
func (d *DbInteraction) Scan(value interface{}) error {
	if value == nil {
		*d = NoDB
		return nil
	}

	var dbInteractionStr string

	switch v := value.(type) {
	case string:
		dbInteractionStr = v
	case []byte:
		dbInteractionStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for DbInteraction: %T", value)
	}

	switch dbInteractionStr {
	case "ORM":
		*d = ORM
	case "DB FUNCTION":
		*d = DBFunction
	case "RAW SQL":
		*d = RawSQL
	case "No SQL":
		*d = NoSql
	case "No DB":
		*d = NoDB
	default:
		return fmt.Errorf("invalid DbInteraction type: %s", dbInteractionStr)
	}

	return nil
}