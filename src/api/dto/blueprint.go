package dto

import (
	"gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

// Blueprint is a DTO for blueprint-related data
type Blueprint struct {
	Uuid            uuid.UUID       `json:"uuid"`
	StandardName    string          `json:"standardName"`
	Type            string          `json:"type"`
	Description     string          `json:"description"`
	Functionalities []Functionality `json:"functionalities"`
	Libraries       []Library       `json:"libraries"`
}

func (b Blueprint) Validate() error {
	if b.StandardName == "" {
		// handle error
	}
	for _, f := range b.Functionalities {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Functionality represents a specific functionality of the blueprint
type Functionality struct {
	Uuid               uuid.UUID             `json:"uuid"`
	Category           string                `json:"category"`
	Type               string                `json:"type"`
	Provider           string                `json:"provider"`
	ImplementsGenerics bool                  `json:"implementsGenerics"`
	FilePathsCSV       string                `json:"filePathsCSV"`
	Operations         []FunctionalOperation `json:"operations"`
}

func (f Functionality) Validate() error {
	if f.Category == "" {
		// handle error
	}
	for _, op := range f.Operations {
		if err := op.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// FunctionalOperation is an individual FunctionalOperation within a functionality
type FunctionalOperation struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (op FunctionalOperation) Validate() error {
	if op.Name == "" {
		// handle error
	}
	return nil
}

// ToUseCaseBlueprint converts this DTO to the usecase Blueprint
func ToUseCaseBlueprint(from Blueprint) dto.Blueprint {
	return dto.Blueprint{
		Uuid:            from.Uuid,
		StandardName:    from.StandardName,
		Type:            from.Type,
		Description:     from.Description,
		Functionalities: ToUseCaseFunctionalities(from.Functionalities),
		Libraries:       ToUseCaseLibraries(from.Libraries),
	}
}

// ToUseCaseFunctionalities converts a slice of Functionality
func ToUseCaseFunctionalities(from []Functionality) []dto.Functionality {
	funcs := make([]dto.Functionality, len(from))
	for i, f := range from {
		funcs[i] = ToUseCaseFunctionality(f)
	}
	return funcs
}

// ToUseCaseFunctionality converts a single Functionality
func ToUseCaseFunctionality(from Functionality) dto.Functionality {
	return dto.Functionality{
		Uuid:               from.Uuid,
		Category:           from.Category,
		Type:               from.Type,
		Provider:           from.Provider,
		ImplementsGenerics: from.ImplementsGenerics,
		FilePathsCSV:       from.FilePathsCSV,
		Operations:         ToUseCaseOperations(from.Operations),
	}
}

// ToUseCaseOperations converts a slice of FunctionalOperation
func ToUseCaseOperations(from []FunctionalOperation) []dto.FunctionalOperation {
	ops := make([]dto.FunctionalOperation, len(from))
	for i, op := range from {
		ops[i] = ToUseCaseOperation(op)
	}
	return ops
}

// ToUseCaseOperation converts a single FunctionalOperation
func ToUseCaseOperation(from FunctionalOperation) dto.FunctionalOperation {
	return dto.FunctionalOperation{
		Uuid:        from.Uuid,
		Name:        from.Name,
		Description: from.Description,
	}
}

// ToUseCaseLibraries converts a slice of Library
func ToUseCaseLibraries(from []Library) []dto.Library {
	libs := make([]dto.Library, len(from))
	for i, lib := range from {
		libs[i] = ToUseCaseLibrary(lib)
	}
	return libs
}

// ToBlueprintResponse converts a usecase Blueprint back to the DTO
func ToBlueprintResponse(from dto.Blueprint) Blueprint {
	return Blueprint{
		Uuid:            from.Uuid,
		StandardName:    from.StandardName,
		Type:            from.Type,
		Description:     from.Description,
		Functionalities: ToFunctionalitiesResponse(from.Functionalities),
		Libraries:       ToLibrariesResponse(from.Libraries),
	}
}

// ToFunctionalitiesResponse converts a slice of usecase Functionalities
func ToFunctionalitiesResponse(from []dto.Functionality) []Functionality {
	funcs := make([]Functionality, len(from))
	for i, f := range from {
		funcs[i] = ToFunctionalityResponse(f)
	}
	return funcs
}

// ToFunctionalityResponse converts a single usecase Functionality
func ToFunctionalityResponse(from dto.Functionality) Functionality {
	return Functionality{
		Uuid:               from.Uuid,
		Category:           from.Category,
		Type:               from.Type,
		Provider:           from.Provider,
		ImplementsGenerics: from.ImplementsGenerics,
		FilePathsCSV:       from.FilePathsCSV,
		Operations:         ToOperationsResponse(from.Operations),
	}
}

// ToOperationsResponse converts a slice of usecase Operations
func ToOperationsResponse(from []dto.FunctionalOperation) []FunctionalOperation {
	ops := make([]FunctionalOperation, len(from))
	for i, op := range from {
		ops[i] = ToOperationResponse(op)
	}
	return ops
}

// ToOperationResponse converts a single usecase FunctionalOperation
func ToOperationResponse(from dto.FunctionalOperation) FunctionalOperation {
	return FunctionalOperation{
		Uuid:        from.Uuid,
		Name:        from.Name,
		Description: from.Description,
	}
}

// ToLibrariesResponse converts a slice of usecase Libraries
func ToLibrariesResponse(from []dto.Library) []Library {
	libs := make([]Library, len(from))
	for i, lib := range from {
		libs[i] = ToLibraryResponse(lib)
	}
	return libs
}
