package enum

import (
	"encoding/json"
	"fmt"
)

type DbInteraction int

const (
	ORM DbInteraction = iota
	DBFunction
	RawSQL
	NoSql
)

// String method for pretty printing
func (d DbInteraction) String() string {
	return [...]string{"ORM", "DB FUNCTION", "RAW SQL, No SQL"}[d]
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
		return fmt.Errorf("invalid db interaction: %s", dbInteractionStr)
	}

	return nil
}
