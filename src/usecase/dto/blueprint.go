package dto

import (
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
