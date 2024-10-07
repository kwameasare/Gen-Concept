package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DerivativeType int

const (
	Arithmetic DerivativeType = iota
	Formula
	Concatenation
	Runtime
	NotDerived
)

// String method for pretty printing
func (d DerivativeType) String() string {
	return [...]string{"Arithmetic", "Formula", "Concatenation", "Runtime","Not Derived"}[d]
}

// MarshalJSON for custom JSON encoding

func (d DerivativeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON for custom JSON decoding

func (d *DerivativeType) UnmarshalJSON(data []byte) error {
	var derivativeTypeStr string
	if err := json.Unmarshal(data, &derivativeTypeStr); err != nil {
		return err
	}

	switch derivativeTypeStr {
	case "Arithmetic":
		*d = Arithmetic
	case "Formula":
		*d = Formula
	case "Concatenation":
		*d = Concatenation
	case "Runtime":
		*d = Runtime
	default:
		*d = NotDerived
	}

	return nil
}

// Implement the driver.Valuer interface
func (d DerivativeType) Value() (driver.Value, error) {
	return d.String(), nil
}

// Implement the sql.Scanner interface
func (d *DerivativeType) Scan(value interface{}) error {
	if value == nil {
		*d = NotDerived
		return nil
	}

	var derivativeTypeStr string

	switch v := value.(type) {
	case string:
		derivativeTypeStr = v
	case []byte:
		derivativeTypeStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for DerivativeType: %T", value)
	}

	switch derivativeTypeStr {
	case "Arithmetic":
		*d = Arithmetic
	case "Formula":
		*d = Formula
	case "Concatenation":
		*d = Concatenation
	case "Runtime":
		*d = Runtime
	case "Not Derived":
		*d = NotDerived
	default:
		return fmt.Errorf("invalid DerivativeType: %s", derivativeTypeStr)
	}

	return nil
}