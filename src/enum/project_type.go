package enum

import (
	"encoding/json"
	"fmt"
)

type ProjectType int

const (
	Enterprise ProjectType = iota
	Website
)

// String method for pretty printing
func (p ProjectType) String() string {
	return [...]string{"Enterprise", "Website"}[p]
}

// MarshalJSON for custom JSON encoding
func (p ProjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// UnmarshalJSON for custom JSON decoding
func (p *ProjectType) UnmarshalJSON(data []byte) error {
	var projectTypeStr string
	if err := json.Unmarshal(data, &projectTypeStr); err != nil {
		return err
	}

	switch projectTypeStr {
	case "Enterprise":
		*p = Enterprise
	case "Website":
		*p = Website
	default:
		return fmt.Errorf("invalid project type: %s", projectTypeStr)
	}

	return nil
}
