package enum

import (
	"encoding/json"
	
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
