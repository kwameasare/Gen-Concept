package dto

import (
	"gen-concept-api/domain/model"

	"github.com/google/uuid"
)

type Blueprint struct {
	Uuid            uuid.UUID       `json:"uuid"`
	StandardName    string          `json:"standardName"`
	Type            string          `json:"type"`
	Description     string          `json:"description"`
	TemplatePath    string          `json:"templatePath"`
	Placeholders    []Placeholder   `json:"placeholders"`
	Functionalities []Functionality `json:"functionalities"`
	Libraries       []Library       `json:"libraries"`
}

type Placeholder struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	DefaultVal  string    `json:"defaultVal"`
}

type Functionality struct {
	Uuid               uuid.UUID             `json:"uuid"`
	Category           string                `json:"category"`
	Type               string                `json:"type"`
	Provider           string                `json:"provider"`
	ImplementsGenerics bool                  `json:"implementsGenerics"`
	FilePathsCSV       string                `json:"filePathsCSV"`
	Operations         []FunctionalOperation `json:"operations"`
}

type FunctionalOperation struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// FromBlueprintModel converts a domain model Blueprint to a usecase DTO
func FromBlueprintModel(m model.Blueprint) Blueprint {
	return Blueprint{
		Uuid:            m.Uuid,
		StandardName:    m.StandardName,
		Type:            m.Type,
		Description:     m.Description,
		TemplatePath:    m.TemplatePath,
		Placeholders:    FromPlaceholderModels(m.Placeholders),
		Functionalities: FromFunctionalityModels(m.Functionalities),
		Libraries:       FromLibraryModels(m.Libraries),
	}
}

func FromPlaceholderModels(models []model.Placeholder) []Placeholder {
	dtos := make([]Placeholder, len(models))
	for i, m := range models {
		dtos[i] = Placeholder{
			Uuid:        m.Uuid,
			Name:        m.Name,
			Description: m.Description,
			Type:        m.Type,
			DefaultVal:  m.DefaultVal,
		}
	}
	return dtos
}

func FromFunctionalityModels(models []model.Functionality) []Functionality {
	dtos := make([]Functionality, len(models))
	for i, m := range models {
		dtos[i] = Functionality{
			Uuid:               m.Uuid,
			Category:           m.Category,
			Type:               m.Type,
			Provider:           m.Provider,
			ImplementsGenerics: m.ImplementsGenerics,
			FilePathsCSV:       m.FilePathsCSV,
			Operations:         FromOperationModels(m.Operations),
		}
	}
	return dtos
}

func FromOperationModels(models []model.FunctionalOperation) []FunctionalOperation {
	dtos := make([]FunctionalOperation, len(models))
	for i, m := range models {
		dtos[i] = FunctionalOperation{
			Uuid:        m.Uuid,
			Name:        m.Name,
			Description: m.Description,
		}
	}
	return dtos
}

func FromLibraryModels(models []model.Library) []Library {
	dtos := make([]Library, len(models))
	for i, m := range models {
		dtos[i] = FromLibraryModel(m)
	}
	return dtos
}
