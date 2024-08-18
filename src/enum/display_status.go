package enum

import (
	"encoding/json"
	"fmt"
)

type DisplayStatus int

const (
	Summary DisplayStatus = iota
	Detail
	hide
)

// String method for pretty printing

func (d DisplayStatus) String() string {
	return [...]string{"Summary", "Detail", "hide"}[d]
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
		*d = hide
	default:
		return fmt.Errorf("invalid display status: %s", displayStatusStr)
	}

	return nil
}
