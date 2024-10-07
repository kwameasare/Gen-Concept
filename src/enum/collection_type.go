package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type CollectionType int

const (
	List CollectionType = iota
	Set
	Map
	Array
	None
)

// String method for pretty printing
func (c CollectionType) String() string {
	return [...]string{"List", "Set", "Map", "Array","None"}[c]
}

// MarshalJSON for custom JSON encoding

func (c CollectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

// UnmarshalJSON for custom JSON decoding

func (c *CollectionType) UnmarshalJSON(data []byte) error {
	var collectionTypeStr string
	if err := json.Unmarshal(data, &collectionTypeStr); err != nil {
		return err
	}

	switch collectionTypeStr {
	case "List":
		*c = List
	case "Set":
		*c = Set
	case "Map":
		*c = Map
	case "Array":
		*c = Array
	default:
		*c = None
	}

	return nil
}

// Implement the driver.Valuer interface
func (c CollectionType) Value() (driver.Value, error) {
    return c.String(), nil
}

// Implement the sql.Scanner interface
func (c *CollectionType) Scan(value interface{}) error {
    if value == nil {
        *c = None
        return nil
    }

    var collectionTypeStr string
    switch v := value.(type) {
    case string:
        collectionTypeStr = v
    case []byte:
        collectionTypeStr = string(v)
    default:
        return fmt.Errorf("unsupported Scan type for CollectionType: %T", value)
    }

    switch collectionTypeStr {
    case "List":
        *c = List
    case "Set":
        *c = Set
    case "Map":
        *c = Map
    case "Array":
        *c = Array
    default:
        *c = None
    }

    return nil
}