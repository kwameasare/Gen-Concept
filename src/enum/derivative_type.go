package enum

import (
	"encoding/json"
	
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
