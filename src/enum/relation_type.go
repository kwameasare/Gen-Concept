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
	SelfReferencing
	Aggregation
	Composition
	Association
	Dependency
	Inheritance
	NoRelation
)

func (r RelationType) String() string {
	names := [...]string{
		"OneToOne",
		"OneToMany",
		"ManyToOne",
		"ManyToMany",
		"SelfReferencing",
		"Aggregation",
		"Composition",
		"Association",
		"Dependency",
		"Inheritance",
		"NoRelation",
	}
	if r < OneToOne || int(r) >= len(names) {
		return "Unknown"
	}
	return names[r]
}

func (r RelationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

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
	case "SelfReferencing":
		*r = SelfReferencing
	case "Aggregation":
		*r = Aggregation
	case "Composition":
		*r = Composition
	case "Association":
		*r = Association
	case "Dependency":
		*r = Dependency
	case "Inheritance":
		*r = Inheritance
	case "NoRelation":
		*r = NoRelation
	default:
		return fmt.Errorf("invalid RelationType: %s", relationTypeStr)
	}

	return nil
}

func (r RelationType) Value() (driver.Value, error) {
	return r.String(), nil
}

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
	case "SelfReferencing":
		*r = SelfReferencing
	case "Aggregation":
		*r = Aggregation
	case "Composition":
		*r = Composition
	case "Association":
		*r = Association
	case "Dependency":
		*r = Dependency
	case "Inheritance":
		*r = Inheritance
	case "NoRelation":
		*r = NoRelation
	default:
		return fmt.Errorf("invalid RelationType: %s", relationTypeStr)
	}

	return nil
}
