package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type OperatorType int

const (
	Equals OperatorType = iota
	NotEquals
	GreaterThan
	LessThan
	GreaterThanOrEqual
	LessThanOrEqual
	Like
	NotLike
	In
	NotIn
	Between
	NotBetween
	IsNull
	IsNotNull
	Exists
	NotExists
	Contains
	StartsWith
	EndsWith
	RegexMatch
	RegexNotMatch
)

func (o OperatorType) String() string {
	names := [...]string{
		"EQUALS",
		"NOT_EQUALS",
		"GREATER_THAN",
		"LESS_THAN",
		"GREATER_THAN_OR_EQUAL",
		"LESS_THAN_OR_EQUAL",
		"LIKE",
		"NOT_LIKE",
		"IN",
		"NOT_IN",
		"BETWEEN",
		"NOT_BETWEEN",
		"IS_NULL",
		"IS_NOT_NULL",
		"EXISTS",
		"NOT_EXISTS",
		"CONTAINS",
		"STARTS_WITH",
		"ENDS_WITH",
		"REGEX_MATCH",
		"REGEX_NOT_MATCH",
	}
	if o < Equals || int(o) >= len(names) {
		return "UNKNOWN"
	}
	return names[o]
}

func (o OperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

func (o *OperatorType) UnmarshalJSON(data []byte) error {
	var opStr string
	if err := json.Unmarshal(data, &opStr); err != nil {
		return err
	}

	switch opStr {
	case "EQUALS":
		*o = Equals
	case "NOT_EQUALS":
		*o = NotEquals
	case "GREATER_THAN":
		*o = GreaterThan
	case "LESS_THAN":
		*o = LessThan
	case "GREATER_THAN_OR_EQUAL":
		*o = GreaterThanOrEqual
	case "LESS_THAN_OR_EQUAL":
		*o = LessThanOrEqual
	case "LIKE":
		*o = Like
	case "NOT_LIKE":
		*o = NotLike
	case "IN":
		*o = In
	case "NOT_IN":
		*o = NotIn
	case "BETWEEN":
		*o = Between
	case "NOT_BETWEEN":
		*o = NotBetween
	case "IS_NULL":
		*o = IsNull
	case "IS_NOT_NULL":
		*o = IsNotNull
	case "EXISTS":
		*o = Exists
	case "NOT_EXISTS":
		*o = NotExists
	case "CONTAINS":
		*o = Contains
	case "STARTS_WITH":
		*o = StartsWith
	case "ENDS_WITH":
		*o = EndsWith
	case "REGEX_MATCH":
		*o = RegexMatch
	case "REGEX_NOT_MATCH":
		*o = RegexNotMatch
	default:
		return fmt.Errorf("invalid OperatorType: %s", opStr)
	}

	return nil
}

func (o OperatorType) Value() (driver.Value, error) {
	return o.String(), nil
}

func (o *OperatorType) Scan(value interface{}) error {
	if value == nil {
		*o = Equals
		return nil
	}

	var opStr string

	switch v := value.(type) {
	case string:
		opStr = v
	case []byte:
		opStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for OperatorType: %T", value)
	}

	switch opStr {
	case "EQUALS":
		*o = Equals
	case "NOT_EQUALS":
		*o = NotEquals
	case "GREATER_THAN":
		*o = GreaterThan
	case "LESS_THAN":
		*o = LessThan
	case "GREATER_THAN_OR_EQUAL":
		*o = GreaterThanOrEqual
	case "LESS_THAN_OR_EQUAL":
		*o = LessThanOrEqual
	case "LIKE":
		*o = Like
	case "NOT_LIKE":
		*o = NotLike
	case "IN":
		*o = In
	case "NOT_IN":
		*o = NotIn
	case "BETWEEN":
		*o = Between
	case "NOT_BETWEEN":
		*o = NotBetween
	case "IS_NULL":
		*o = IsNull
	case "IS_NOT_NULL":
		*o = IsNotNull
	case "EXISTS":
		*o = Exists
	case "NOT_EXISTS":
		*o = NotExists
	case "CONTAINS":
		*o = Contains
	case "STARTS_WITH":
		*o = StartsWith
	case "ENDS_WITH":
		*o = EndsWith
	case "REGEX_MATCH":
		*o = RegexMatch
	case "REGEX_NOT_MATCH":
		*o = RegexNotMatch
	default:
		return fmt.Errorf("invalid OperatorType: %s", opStr)
	}

	return nil
}
