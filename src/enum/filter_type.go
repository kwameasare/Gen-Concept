package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type FilterType int

const (
	DateRange FilterType = iota
	FieldFilter
	TextSearch
	NumericRange
	GeoSpatial
	EnumFilter
	RegexFilter
	NullFilter
	CompositeFilter
	CustomFilter
)

func (f FilterType) String() string {
	names := [...]string{
		"DATE_RANGE",
		"FIELD",
		"TEXT_SEARCH",
		"NUMERIC_RANGE",
		"GEO_SPATIAL",
		"ENUM",
		"REGEX",
		"NULL",
		"COMPOSITE",
		"CUSTOM",
	}
	if f < DateRange || int(f) >= len(names) {
		return "UNKNOWN"
	}
	return names[f]
}

func (f FilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *FilterType) UnmarshalJSON(data []byte) error {
	var filterStr string
	if err := json.Unmarshal(data, &filterStr); err != nil {
		return err
	}

	switch filterStr {
	case "DATE_RANGE":
		*f = DateRange
	case "FIELD":
		*f = FieldFilter
	case "TEXT_SEARCH":
		*f = TextSearch
	case "NUMERIC_RANGE":
		*f = NumericRange
	case "GEO_SPATIAL":
		*f = GeoSpatial
	case "ENUM":
		*f = EnumFilter
	case "REGEX":
		*f = RegexFilter
	case "NULL":
		*f = NullFilter
	case "COMPOSITE":
		*f = CompositeFilter
	case "CUSTOM":
		*f = CustomFilter
	default:
		return fmt.Errorf("invalid FilterType: %s", filterStr)
	}

	return nil
}

func (f FilterType) Value() (driver.Value, error) {
	return f.String(), nil
}

func (f *FilterType) Scan(value interface{}) error {
	if value == nil {
		*f = DateRange
		return nil
	}

	var filterStr string

	switch v := value.(type) {
	case string:
		filterStr = v
	case []byte:
		filterStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for FilterType: %T", value)
	}

	switch filterStr {
	case "DATE_RANGE":
		*f = DateRange
	case "FIELD":
		*f = FieldFilter
	case "TEXT_SEARCH":
		*f = TextSearch
	case "NUMERIC_RANGE":
		*f = NumericRange
	case "GEO_SPATIAL":
		*f = GeoSpatial
	case "ENUM":
		*f = EnumFilter
	case "REGEX":
		*f = RegexFilter
	case "NULL":
		*f = NullFilter
	case "COMPOSITE":
		*f = CompositeFilter
	case "CUSTOM":
		*f = CustomFilter
	default:
		return fmt.Errorf("invalid FilterType: %s", filterStr)
	}

	return nil
}
