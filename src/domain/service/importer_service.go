package service

import (
	"fmt"
	"gen-concept-api/domain/model"
)

type ImporterService struct {
	parsers map[string]CodeParser
}

func NewImporterService(goParser CodeParser) *ImporterService {
	return &ImporterService{
		parsers: map[string]CodeParser{
			"go": goParser,
		},
	}
}

func (s *ImporterService) ImportFromSource(content string, lang string) (model.Blueprint, error) {
	parser, ok := s.parsers[lang]
	if !ok {
		return model.Blueprint{}, fmt.Errorf("unsupported language: %s", lang)
	}

	metadata, err := parser.Parse(content)
	if err != nil {
		return model.Blueprint{}, err
	}

	// Map metadata to Blueprint
	blueprint := model.Blueprint{
		StandardName: metadata.Name,
		Type:         "STRUCT", // Default type for code import
		Description:  fmt.Sprintf("Imported from %s source", lang),
		TemplatePath: content, // Save the original source as the template (for now)
	}

	for _, field := range metadata.Fields {
		blueprint.Placeholders = append(blueprint.Placeholders, model.Placeholder{
			Name:        field.Name,
			Type:        field.Type,
			Description: fmt.Sprintf("Field extracted from %s", metadata.Name),
		})
	}

	return blueprint, nil
}
