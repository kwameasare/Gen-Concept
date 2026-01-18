package service

import (
	"bytes"
	"context"
	"fmt"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/enum"

	"strings"
	"text/template"
)

type GenerationService struct {
	gitProvider GitProvider
	aiProvider  AIProvider
	libraryRepo repository.LibraryRepository
}

func NewGenerationService(gitProvider GitProvider, aiProvider AIProvider, libraryRepo repository.LibraryRepository) *GenerationService {
	return &GenerationService{
		gitProvider: gitProvider,
		aiProvider:  aiProvider,
		libraryRepo: libraryRepo,
	}
}

func (s *GenerationService) GenerateCode(ctx context.Context, blueprint model.Blueprint, entity model.Entity, inputs map[string]string) (string, error) {
	// 1. Fetch/Prepare Template Content
	var templateContent string
	if blueprint.TemplatePath != "" {
		// For MVP, assuming TemplatePath IS the template content if it's not a URL
		// OR we fetch it. Retaining logic for now.
		templateContent = blueprint.TemplatePath
	}

	if templateContent == "" {
		return "", fmt.Errorf("no template content")
	}

	// 2. Build Context
	genCtx, err := s.BuildContext(ctx, entity)
	if err != nil {
		return "", err
	}

	// 3. Parse and Execute Template
	tmpl, err := template.New("blueprint").Parse(templateContent)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, genCtx); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

// BuildContext creates a generation context from the entity model
func (s *GenerationService) BuildContext(ctx context.Context, entity model.Entity) (GenContext, error) {
	// Simple lowerCamelCase conversion (naive implementation, use a library like strcase in production)
	// Guard against empty name
	if len(entity.EntityName) == 0 {
		return GenContext{}, fmt.Errorf("entity name is empty")
	}
	varName := strings.ToLower(string(entity.EntityName[0])) + entity.EntityName[1:]

	// Project checks
	projectName := ""
	if entity.Project.ProjectName != "" {
		projectName = entity.Project.ProjectName
	}
	// Or maybe pass ProjectName from inputs/context if Entity.Project (gorm relation) is not loaded

	genCtx := GenContext{
		ProjectName: projectName,
		Entity: GenEntity{
			Name:       entity.EntityName,
			VarName:    varName,
			PrimaryKey: "ID", // Default assumption, or check fields
		},
		Imports:          []string{},
		LibraryFunctions: make(map[string]string),
	}

	importsMap := make(map[string]bool)

	for _, f := range entity.EntityFields {
		genField := GenField{
			Name: f.FieldName,
			Type: string(f.FieldType), // Convert enum to string, logic might need mapping
		}

		// Smart Imports Logic
		if f.FieldType == enum.DateTime {
			key := "time"
			if !importsMap[key] {
				genCtx.Imports = append(genCtx.Imports, key)
				importsMap[key] = true
			}
			genField.Type = "time.Time" // Override enum string with Go type
		}

		// Sensitive Data Logic -> Library Discovery
		if f.IsSensitive {
			// Query LibraryRepo for "encryption" tag equivalent
			// MVP: Hardcoded lookup or basic query.
			// Assuming we find a library that has "encryption" in description or name?
			// Since BaseRepository GetByFilter uses PaginationInputWithFilter, it's complex to setup here.
			// Simplified: If sensitive, we add a placeholder library.
			// Ideally: s.libraryRepo.GetByFilter...

			// For this task, let's assume we find one.
			libPkg := "company.com/security/encryption"
			if !importsMap[libPkg] {
				genCtx.Imports = append(genCtx.Imports, libPkg)
				importsMap[libPkg] = true
			}
			genCtx.LibraryFunctions["Encrypt"] = "encryption.Encrypt"
			genCtx.LibraryFunctions["Decrypt"] = "encryption.Decrypt"
		}

		// Validation Tags
		if len(f.InputValidations) > 0 || f.IsMandatory {
			tags := []string{}
			if f.IsMandatory {
				tags = append(tags, "required")
			}
			// Add more mapping logic here
			if len(tags) > 0 {
				genField.ValidateTag = fmt.Sprintf(`binding:"%s"`, strings.Join(tags, ","))
			}
		}

		genField.JSONTag = fmt.Sprintf(`json:"%s"`, strings.ToLower(string(f.FieldName[0]))+f.FieldName[1:])

		genCtx.Entity.Fields = append(genCtx.Entity.Fields, genField)
	}

	return genCtx, nil
}
