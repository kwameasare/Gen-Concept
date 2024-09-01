package enum

import (
	"encoding/json"
)

type CollectionItemType int

const (
	StringType CollectionItemType = iota
	IntType
	FloatType
	BoolType
	DateTimeType
	EnumType
	NestedCollectionType
	OtherEntityType
	NoType
)

// String method for pretty printing
func (d CollectionItemType) String() string {
	return [...]string{"String", "Int", "Float", "Bool", "DateTime", "Enum", "Nested Collection", "Entity","No Type"}[d]
}

// MarshalJSON for custom JSON encoding

func (d CollectionItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON for custom JSON decoding

func (d *CollectionItemType) UnmarshalJSON(data []byte) error {
	var dataTypeStr string
	if err := json.Unmarshal(data, &dataTypeStr); err != nil {
		return err
	}

	switch dataTypeStr {
	case "String":
		*d = StringType
	case "Int":
		*d = IntType
	case "Float":
		*d = FloatType
	case "Bool":
		*d = BoolType
	case "DateTime":
		*d = DateTimeType
	case "Enum":
		*d = EnumType
	case "Nested Collection":
		*d = NestedCollectionType
	case "Entity":
		*d = OtherEntityType
	default:
		*d = NoType
	}

	return nil
}
