package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ProgrammingLanguage int

const (
	Golang ProgrammingLanguage = iota
	Python
	Java
	JavaScript
	TypeScript
	Csharp
	Cpp
	Kotlin
	Rust
)

// String method for pretty printing
func (p ProgrammingLanguage) String() string {
	return [...]string{"Golang", "Python", "Java", "JavaScript", "TypeScript", "Csharp", "Cpp", "Kotlin", "Rust"}[p]
}

// MarshalJSON for custom JSON encoding
func (p ProgrammingLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// UnmarshalJSON for custom JSON decoding
func (p *ProgrammingLanguage) UnmarshalJSON(data []byte) error {
	var programmingLanguageStr string
	if err := json.Unmarshal(data, &programmingLanguageStr); err != nil {
		return err
	}

	switch programmingLanguageStr {
	case "Golang":
		*p = Golang
	case "Python":
		*p = Python
	case "Java":
		*p = Java
	case "JavaScript":
		*p = JavaScript
	case "TypeScript":
		*p = TypeScript
	case "Csharp":
		*p = Csharp
	case "Cpp":
		*p = Cpp
	case "Kotlin":
		*p = Kotlin
	case "Rust":
		*p = Rust
	default:
		return fmt.Errorf("invalid programming language: %s", programmingLanguageStr)
	}

	return nil
}

// Implement the driver.Valuer interface
func (p ProgrammingLanguage) Value() (driver.Value, error) {
	return p.String(), nil
}

// Implement the sql.Scanner interface
func (p *ProgrammingLanguage) Scan(value interface{}) error {
	if value == nil {
		*p = Golang // Default to Golang if needed
		return nil
	}

	var programmingLanguageStr string

	switch v := value.(type) {
	case string:
		programmingLanguageStr = v
	case []byte:
		programmingLanguageStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for ProgrammingLanguage: %T", value)
	}

	switch programmingLanguageStr {
	case "Golang":
		*p = Golang
	case "Python":
		*p = Python
	case "Java":
		*p = Java
	case "JavaScript":
		*p = JavaScript
	case "TypeScript":
		*p = TypeScript
	case "Csharp":
		*p = Csharp
	case "Cpp":
		*p = Cpp
	case "Kotlin":
		*p = Kotlin
	case "Rust":
		*p = Rust
	default:
		return fmt.Errorf("invalid programming language: %s", programmingLanguageStr)
	}

	return nil
}
