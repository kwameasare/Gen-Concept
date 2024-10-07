package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type RelationType int

const (
	OneToOne RelationType = iota
	OneToMany
	ManyToOne
	ManyToMany
	NoRelation
)

// String method for pretty printing

func (r RelationType) String() string {
	return [...]string{"OneToOne", "OneToMany", "ManyToOne", "ManyToMany","No Relation"}[r]
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
		*r = NoRelation
	}

	return nil
}

// Implement the driver.Valuer interface
func (r RelationType) Value() (driver.Value, error) {
	return r.String(), nil
}

// Implement the sql.Scanner interface
func (r *RelationType) Scan(value interface{}) error {
	if value == nil {
		*r = NoRelation
		return nil
	}

	var relationTypeStr string

	switch v := value.(type) {
	case string:
		relationTypeStr = v
	case []byte:
		relationTypeStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for RelationType: %T", value)
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
	case "No Relation":
		*r = NoRelation
	default:
		return fmt.Errorf("invalid RelationType: %s", relationTypeStr)
	}

	return nil
}