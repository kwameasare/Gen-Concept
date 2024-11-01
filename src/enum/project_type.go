package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ProjectType int

const (
	Enterprise ProjectType = iota
	Website
	MobileApp
	DesktopApp
	Microservice
	API
	Library
	CommandLineTool
	Game
	DataProcessing
	AIML
	EmbeddedSystem
	IoT
	Blockchain
	SaaS
	PaaS
	ECommerce
	SocialMediaPlatform
	CRM
	ERP
	CMS
	UtilityTool
	Plugin
	Middleware
	Script
	TestSuite
	NetworkService
	Firmware
	Simulator
	Virtualization
	MonitoringTool
	SecurityTool
	Documentation
	Chatbot
	DataVisualization
)

func (p ProjectType) String() string {
	names := [...]string{
		"Enterprise",
		"Website",
		"MobileApp",
		"DesktopApp",
		"Microservice",
		"API",
		"Library",
		"CommandLineTool",
		"Game",
		"DataProcessing",
		"AI/ML",
		"EmbeddedSystem",
		"IoT",
		"Blockchain",
		"SaaS",
		"PaaS",
		"ECommerce",
		"SocialMediaPlatform",
		"CRM",
		"ERP",
		"CMS",
		"UtilityTool",
		"Plugin",
		"Middleware",
		"Script",
		"TestSuite",
		"NetworkService",
		"Firmware",
		"Simulator",
		"Virtualization",
		"MonitoringTool",
		"SecurityTool",
		"Documentation",
		"Chatbot",
		"DataVisualization",
	}

	if p < Enterprise || int(p) >= len(names) {
		return "Unknown"
	}
	return names[p]
}

func (p ProjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

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
	case "MobileApp":
		*p = MobileApp
	case "DesktopApp":
		*p = DesktopApp
	case "Microservice":
		*p = Microservice
	case "API":
		*p = API
	case "Library":
		*p = Library
	case "CommandLineTool":
		*p = CommandLineTool
	case "Game":
		*p = Game
	case "DataProcessing":
		*p = DataProcessing
	case "AI/ML":
		*p = AIML
	case "EmbeddedSystem":
		*p = EmbeddedSystem
	case "IoT":
		*p = IoT
	case "Blockchain":
		*p = Blockchain
	case "SaaS":
		*p = SaaS
	case "PaaS":
		*p = PaaS
	case "ECommerce":
		*p = ECommerce
	case "SocialMediaPlatform":
		*p = SocialMediaPlatform
	case "CRM":
		*p = CRM
	case "ERP":
		*p = ERP
	case "CMS":
		*p = CMS
	case "UtilityTool":
		*p = UtilityTool
	case "Plugin":
		*p = Plugin
	case "Middleware":
		*p = Middleware
	case "Script":
		*p = Script
	case "TestSuite":
		*p = TestSuite
	case "NetworkService":
		*p = NetworkService
	case "Firmware":
		*p = Firmware
	case "Simulator":
		*p = Simulator
	case "Virtualization":
		*p = Virtualization
	case "MonitoringTool":
		*p = MonitoringTool
	case "SecurityTool":
		*p = SecurityTool
	case "Documentation":
		*p = Documentation
	case "Chatbot":
		*p = Chatbot
	case "DataVisualization":
		*p = DataVisualization
	default:
		return fmt.Errorf("invalid project type: %s", projectTypeStr)
	}

	return nil
}

func (p ProjectType) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p *ProjectType) Scan(value interface{}) error {
	if value == nil {
		*p = Enterprise // Default to Enterprise or handle as needed
		return nil
	}

	var projectTypeStr string

	switch v := value.(type) {
	case string:
		projectTypeStr = v
	case []byte:
		projectTypeStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for ProjectType: %T", value)
	}

	switch projectTypeStr {
	case "Enterprise":
		*p = Enterprise
	case "Website":
		*p = Website
	case "MobileApp":
		*p = MobileApp
	case "DesktopApp":
		*p = DesktopApp
	case "Microservice":
		*p = Microservice
	case "API":
		*p = API
	case "Library":
		*p = Library
	case "CommandLineTool":
		*p = CommandLineTool
	case "Game":
		*p = Game
	case "DataProcessing":
		*p = DataProcessing
	case "AI/ML":
		*p = AIML
	case "EmbeddedSystem":
		*p = EmbeddedSystem
	case "IoT":
		*p = IoT
	case "Blockchain":
		*p = Blockchain
	case "SaaS":
		*p = SaaS
	case "PaaS":
		*p = PaaS
	case "ECommerce":
		*p = ECommerce
	case "SocialMediaPlatform":
		*p = SocialMediaPlatform
	case "CRM":
		*p = CRM
	case "ERP":
		*p = ERP
	case "CMS":
		*p = CMS
	case "UtilityTool":
		*p = UtilityTool
	case "Plugin":
		*p = Plugin
	case "Middleware":
		*p = Middleware
	case "Script":
		*p = Script
	case "TestSuite":
		*p = TestSuite
	case "NetworkService":
		*p = NetworkService
	case "Firmware":
		*p = Firmware
	case "Simulator":
		*p = Simulator
	case "Virtualization":
		*p = Virtualization
	case "MonitoringTool":
		*p = MonitoringTool
	case "SecurityTool":
		*p = SecurityTool
	case "Documentation":
		*p = Documentation
	case "Chatbot":
		*p = Chatbot
	case "DataVisualization":
		*p = DataVisualization
	default:
		return fmt.Errorf("invalid project type: %s", projectTypeStr)
	}

	return nil
}
