package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DataType int

const (
	String DataType = iota
	Int
	Float
	Bool
	DateTime
	Enum
	Collection
	Entity
)

// String method for pretty printing
func (d DataType) String() string {
	return [...]string{"String", "Int", "Float", "Bool", "DateTime", "Enum", "Collection", "Entity"}[d]
}

// MarshalJSON for custom JSON encoding

func (d DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON for custom JSON decoding

func (d *DataType) UnmarshalJSON(data []byte) error {
	var dataTypeStr string
	if err := json.Unmarshal(data, &dataTypeStr); err != nil {
		return err
	}

	switch dataTypeStr {
	case "String":
		*d = String
	case "Int":
		*d = Int
	case "Float":
		*d = Float
	case "Bool":
		*d = Bool
	case "DateTime":
		*d = DateTime
	case "Enum":
		*d = Enum
	case "Collection":
		*d = Collection
	case "Entity":
		*d = Entity
	default:
		return fmt.Errorf("invalid data type: %s", dataTypeStr)
	}

	return nil
}


// Implement the driver.Valuer interface
func (d DataType) Value() (driver.Value, error) {
	return d.String(), nil
}

// Implement the sql.Scanner interface
func (d *DataType) Scan(value interface{}) error {
	if value == nil {
		*d = String // Default value if needed
		return nil
	}

	var dataTypeStr string

	switch v := value.(type) {
	case string:
		dataTypeStr = v
	case []byte:
		dataTypeStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for DataType: %T", value)
	}

	switch dataTypeStr {
	case "String":
		*d = String
	case "Int":
		*d = Int
	case "Float":
		*d = Float
	case "Bool":
		*d = Bool
	case "DateTime":
		*d = DateTime
	case "Enum":
		*d = Enum
	case "Collection":
		*d = Collection
	case "Entity":
		*d = Entity
	default:
		return fmt.Errorf("invalid data type: %s", dataTypeStr)
	}

	return nil
}