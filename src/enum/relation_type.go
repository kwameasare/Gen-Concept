package enum

import (
	"encoding/json"
	"fmt"
)

type RelationType int

const (
	OneToOne RelationType = iota
	OneToMany
	ManyToOne
	ManyToMany
)

// String method for pretty printing

func (r RelationType) String() string {
	return [...]string{"OneToOne", "OneToMany", "ManyToOne", "ManyToMany"}[r]
}

// MarshalJSON for custom JSON encoding

func (r RelationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON for custom JSON decoding

func (r *RelationType) UnmarshalJSON(data []byte) error {
	var relationTypeStr string
	if err := json.Unmarshal(data, &relationTypeStr); err != nil {
		return err
	}

	switch relationTypeStr {
	case "OneToOne":
		*r = OneToOne
	case "OneToMany":
		*r = OneToMany
	case "ManyToOne":
		*r = ManyToOne
	case "ManyToMany":
		*r = ManyToMany
	default:
		return fmt.Errorf("invalid relation type: %s", relationTypeStr)
	}

	return nil
}
