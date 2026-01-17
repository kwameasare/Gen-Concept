package service

import (
	"fmt"
	"gen-concept-api/domain/model"
	"strings"
)

type GenerationService struct {
	gitProvider GitProvider
	aiProvider  AIProvider
}

func NewGenerationService(gitProvider GitProvider, aiProvider AIProvider) *GenerationService {
	return &GenerationService{
		gitProvider: gitProvider,
		aiProvider:  aiProvider,
	}
}

func (s *GenerationService) GenerateCode(blueprint model.Blueprint, inputs map[string]string) (string, error) {
	// 1. Fetch Template Content
	var templateContent []byte

	if blueprint.TemplatePath != "" {
		if strings.HasPrefix(blueprint.TemplatePath, "http") {
			// Assume it's a URL (like raw github content or use GitProvider)
			// For simplicity, we can reuse GitProvider if it fits, or just standard HTTP get if we had a generic one.
			// But here we rely on the specific method.
			// Let's assume the TemplatePath IS a repo URL + path, but splitting that is tricky without structure.
			// For MVP, let's assume the TemplatePath is actually the RAW content for now OR we use a simple placeholder string.

			// Actually, let's just treat TemplatePath as the template content itself for the very first step of verification
			// if it doesn't look like a URL/File path, OR implement a simple fetcher.

			// BETTER MVP: Assume TemplatePath IS the template string content if it's short, or we fetch it.
			// Even better: Let's assume we fetch from the library repo if linked?

			// Let's stick to the plan: Read template content.
			// If we don't have a sophisticated file fetcher yet, let's assume inputs["_template_content"] is passed
			// OR we just perform replacement on the string provided in TemplatePath (assuming it might be a small snippet).

			// Real implementation: We need to know WHERE the template is.
			// Let's assume for now the TemplatePath is the actual content for testing purposes or a direct http url.
			templateContent = []byte(blueprint.TemplatePath)
		} else {
			// Treat as direct content for MVP
			templateContent = []byte(blueprint.TemplatePath)
		}
	}

	// If we have no content, we can't generate
	if len(templateContent) == 0 {
		return "", fmt.Errorf("no template content found in blueprint")
	}

	code := string(templateContent)

	// 2. Replace Placeholders
	for _, p := range blueprint.Placeholders {
		key := fmt.Sprintf("{{%s}}", p.Name)
		val, ok := inputs[p.Name]
		if !ok {
			// Use default if available, otherwise fallback to AI
			if p.DefaultVal != "" {
				val = p.DefaultVal
			} else {
				// Fallback to AI
				if p.Description != "" {
					generated, err := s.aiProvider.GenerateContent(p.Description)
					if err == nil {
						val = generated
					} else {
						// Log error? For now fallback empty or error
						return "", fmt.Errorf("missing input for placeholder '%s' and AI generation failed: %v", p.Name, err)
					}
				} else {
					return "", fmt.Errorf("missing input for placeholder: %s (and no description for AI)", p.Name)
				}
			}
		}
		code = strings.ReplaceAll(code, key, val)
	}

	return code, nil
}
