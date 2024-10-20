package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type BackendJourneyStepType int

const (
	InputValidation BackendJourneyStepType = iota
	Authentication
	Authorization
	BusinessValidation
	DataTransformation
	DataEnrichment
	APICall
	DatabaseOperation
	CacheOperation
	Notification
	Logging
	ExceptionHandling
	TransactionManagement
	ResponseTransformation
	Serialization
	Deserialization
	ReturnStep
)

// String method for pretty printing
func (b BackendJourneyStepType) String() string {
	names := [...]string{
		"INPUT_VALIDATION",
		"AUTHENTICATION",
		"AUTHORIZATION",
		"BUSINESS_VALIDATION",
		"DATA_TRANSFORMATION",
		"DATA_ENRICHMENT",
		"API_CALL",
		"DATABASE_OPERATION",
		"CACHE_OPERATION",
		"NOTIFICATION",
		"LOGGING",
		"EXCEPTION_HANDLING",
		"TRANSACTION_MANAGEMENT",
		"RESPONSE_TRANSFORMATION",
		"SERIALIZATION",
		"DESERIALIZATION",
		"RETURN",
	}

	if b < InputValidation || int(b) >= len(names) {
		return "UNKNOWN"
	}
	return names[b]
}

// MarshalJSON for custom JSON encoding
func (b BackendJourneyStepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

// UnmarshalJSON for custom JSON decoding
func (b *BackendJourneyStepType) UnmarshalJSON(data []byte) error {
	var stepTypeStr string
	if err := json.Unmarshal(data, &stepTypeStr); err != nil {
		return err
	}

	switch stepTypeStr {
	case "INPUT_VALIDATION":
		*b = InputValidation
	case "AUTHENTICATION":
		*b = Authentication
	case "AUTHORIZATION":
		*b = Authorization
	case "BUSINESS_VALIDATION":
		*b = BusinessValidation
	case "DATA_TRANSFORMATION":
		*b = DataTransformation
	case "DATA_ENRICHMENT":
		*b = DataEnrichment
	case "API_CALL":
		*b = APICall
	case "DATABASE_OPERATION":
		*b = DatabaseOperation
	case "CACHE_OPERATION":
		*b = CacheOperation
	case "NOTIFICATION":
		*b = Notification
	case "LOGGING":
		*b = Logging
	case "EXCEPTION_HANDLING":
		*b = ExceptionHandling
	case "TRANSACTION_MANAGEMENT":
		*b = TransactionManagement
	case "RESPONSE_TRANSFORMATION":
		*b = ResponseTransformation
	case "SERIALIZATION":
		*b = Serialization
	case "DESERIALIZATION":
		*b = Deserialization
	case "RETURN":
		*b = ReturnStep
	default:
		return fmt.Errorf("invalid BackendJourneyStepType: %s", stepTypeStr)
	}

	return nil
}

// Implement the driver.Valuer interface
func (b BackendJourneyStepType) Value() (driver.Value, error) {
	return b.String(), nil
}

// Implement the sql.Scanner interface
func (b *BackendJourneyStepType) Scan(value interface{}) error {
	if value == nil {
		*b = ReturnStep
		return nil
	}

	var stepTypeStr string

	switch v := value.(type) {
	case string:
		stepTypeStr = v
	case []byte:
		stepTypeStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for BackendJourneyStepType: %T", value)
	}

	switch stepTypeStr {
	case "INPUT_VALIDATION":
		*b = InputValidation
	case "AUTHENTICATION":
		*b = Authentication
	case "AUTHORIZATION":
		*b = Authorization
	case "BUSINESS_VALIDATION":
		*b = BusinessValidation
	case "DATA_TRANSFORMATION":
		*b = DataTransformation
	case "DATA_ENRICHMENT":
		*b = DataEnrichment
	case "API_CALL":
		*b = APICall
	case "DATABASE_OPERATION":
		*b = DatabaseOperation
	case "CACHE_OPERATION":
		*b = CacheOperation
	case "NOTIFICATION":
		*b = Notification
	case "LOGGING":
		*b = Logging
	case "EXCEPTION_HANDLING":
		*b = ExceptionHandling
	case "TRANSACTION_MANAGEMENT":
		*b = TransactionManagement
	case "RESPONSE_TRANSFORMATION":
		*b = ResponseTransformation
	case "SERIALIZATION":
		*b = Serialization
	case "DESERIALIZATION":
		*b = Deserialization
	case "RETURN":
		*b = ReturnStep
	default:
		return fmt.Errorf("invalid BackendJourneyStepType: %s", stepTypeStr)
	}

	return nil
}
