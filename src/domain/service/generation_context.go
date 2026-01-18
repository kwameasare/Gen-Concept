package service

// GenContext holds all data required for smart code generation
type GenContext struct {
	ProjectName      string
	Entity           GenEntity
	Imports          []string
	LibraryFunctions map[string]string // Map of key (e.g. "Encrypt") to Function Name
}

// GenEntity represents the entity model for generation
type GenEntity struct {
	Name       string
	VarName    string // lowerCamelCase name
	Fields     []GenField
	PrimaryKey string
}

// GenField represents a field within the entity
type GenField struct {
	Name        string
	Type        string
	JSONTag     string
	ValidateTag string
}
