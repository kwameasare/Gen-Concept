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
	C
	Php
	Ruby
	Swift
	ObjectiveC
	Scala
	Haskell
	Lua
	Perl
	R
	Matlab
	VisualBasic
	Dart
	Elixir
	Erlang
	FSharp
	Groovy
	Julia
	Lisp
	Scheme
	Prolog
	Fortran
	Cobol
	Assembly
	Smalltalk
)

// String method for pretty printing
func (p ProgrammingLanguage) String() string {
	names := [...]string{
		"Golang",
		"Python",
		"Java",
		"JavaScript",
		"TypeScript",
		"Csharp",
		"Cpp",
		"Kotlin",
		"Rust",
		"C",
		"Php",
		"Ruby",
		"Swift",
		"ObjectiveC",
		"Scala",
		"Haskell",
		"Lua",
		"Perl",
		"R",
		"Matlab",
		"VisualBasic",
		"Dart",
		"Elixir",
		"Erlang",
		"FSharp",
		"Groovy",
		"Julia",
		"Lisp",
		"Scheme",
		"Prolog",
		"Fortran",
		"Cobol",
		"Assembly",
		"Smalltalk",
	}

	if p < Golang || int(p) >= len(names) {
		return "Unknown"
	}
	return names[p]
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
	case "C":
		*p = C
	case "Php":
		*p = Php
	case "Ruby":
		*p = Ruby
	case "Swift":
		*p = Swift
	case "ObjectiveC":
		*p = ObjectiveC
	case "Scala":
		*p = Scala
	case "Haskell":
		*p = Haskell
	case "Lua":
		*p = Lua
	case "Perl":
		*p = Perl
	case "R":
		*p = R
	case "Matlab":
		*p = Matlab
	case "VisualBasic":
		*p = VisualBasic
	case "Dart":
		*p = Dart
	case "Elixir":
		*p = Elixir
	case "Erlang":
		*p = Erlang
	case "FSharp":
		*p = FSharp
	case "Groovy":
		*p = Groovy
	case "Julia":
		*p = Julia
	case "Lisp":
		*p = Lisp
	case "Scheme":
		*p = Scheme
	case "Prolog":
		*p = Prolog
	case "Fortran":
		*p = Fortran
	case "Cobol":
		*p = Cobol
	case "Assembly":
		*p = Assembly
	case "Smalltalk":
		*p = Smalltalk
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
	case "C":
		*p = C
	case "Php":
		*p = Php
	case "Ruby":
		*p = Ruby
	case "Swift":
		*p = Swift
	case "ObjectiveC":
		*p = ObjectiveC
	case "Scala":
		*p = Scala
	case "Haskell":
		*p = Haskell
	case "Lua":
		*p = Lua
	case "Perl":
		*p = Perl
	case "R":
		*p = R
	case "Matlab":
		*p = Matlab
	case "VisualBasic":
		*p = VisualBasic
	case "Dart":
		*p = Dart
	case "Elixir":
		*p = Elixir
	case "Erlang":
		*p = Erlang
	case "FSharp":
		*p = FSharp
	case "Groovy":
		*p = Groovy
	case "Julia":
		*p = Julia
	case "Lisp":
		*p = Lisp
	case "Scheme":
		*p = Scheme
	case "Prolog":
		*p = Prolog
	case "Fortran":
		*p = Fortran
	case "Cobol":
		*p = Cobol
	case "Assembly":
		*p = Assembly
	case "Smalltalk":
		*p = Smalltalk
	default:
		return fmt.Errorf("invalid programming language: %s", programmingLanguageStr)
	}

	return nil
}
