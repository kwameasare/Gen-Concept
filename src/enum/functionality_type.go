package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type FunctionalityType int

const (
	Cache FunctionalityType = iota
	Queue
	Database
	LoggingFunctionality
	AuthenticationFunctionality
	EmailFunctionality
	SMSFunctionality
	NotificationFunctionality
	File
	HttpClient
	Scheduler
	Caching
	Search
	Observability
	ValidationFunctionality
)

// String method for pretty printing
func (f FunctionalityType) String() string {
	return [...]string{
		"Cache", "Queue", "Database", "Logging", "Authentication", "Email", "SMS", "Notification", "File",
		"HttpClient", "Scheduler", "Caching", "Search", "Observability", "InputValidation",
	}[f]
}

// MarshalJSON for custom JSON encoding
func (f FunctionalityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

// UnmarshalJSON for custom JSON decoding
func (f *FunctionalityType) UnmarshalJSON(data []byte) error {
	var functionalityTypeStr string
	if err := json.Unmarshal(data, &functionalityTypeStr); err != nil {
		return err
	}

	switch functionalityTypeStr {
	case "Cache":
		*f = Cache
	case "Queue":
		*f = Queue
	case "Database":
		*f = Database
	case "Logging":
		*f = LoggingFunctionality
	case "Authentication":
		*f = AuthenticationFunctionality
	case "Email":
		*f = EmailFunctionality
	case "SMS":
		*f = SMSFunctionality
	case "Notification":
		*f = NotificationFunctionality
	case "File":
		*f = File
	case "HttpClient":
		*f = HttpClient
	case "Scheduler":
		*f = Scheduler
	case "Caching":
		*f = Caching
	case "Search":
		*f = Search
	case "Observability":
		*f = Observability
	case "InputValidation":
		*f = ValidationFunctionality
	default:
		return	fmt.Errorf("invalid FunctionalityType: %s", functionalityTypeStr)

	}

	return nil
}

// Implement the driver.Valuer interface
func (f FunctionalityType) Value() (driver.Value, error) {
	return f.String(), nil
}

// Implement the sql.Scanner interface
func (f *FunctionalityType) Scan(value interface{}) error {
	if value == nil {
		*f = Database
		return nil
	}

	var functionalityTypeStr string
	switch v := value.(type) {
	case string:
		functionalityTypeStr = v
	case []byte:
		functionalityTypeStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for FunctionalityType: %T", value)
	}

	switch functionalityTypeStr {
	case "Cache":
		*f = Cache
	case "Queue":
		*f = Queue
	case "Database":
		*f = Database
	case "Logging":
		*f = LoggingFunctionality
	case "Authentication":
		*f = AuthenticationFunctionality
	case "Email":
		*f = EmailFunctionality
	case "SMS":
		*f = SMSFunctionality
	case "Notification":
		*f = NotificationFunctionality
	case "File":
		*f = File
	case "HttpClient":
		*f = HttpClient
	case "Scheduler":
		*f = Scheduler
	case "Caching":
		*f = Caching
	case "Search":
		*f = Search
	case "Observability":
		*f = Observability
	case "InputValidation":
		*f = ValidationFunctionality
	default:
	return	fmt.Errorf("invalid FunctionalityType: %s", functionalityTypeStr)
	}

	return nil
}