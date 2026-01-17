package ai

import (
	"fmt"
	"gen-concept-api/domain/service"
)

type MockAIProvider struct{}

func NewMockAIProvider() service.AIProvider {
	return &MockAIProvider{}
}

func (p *MockAIProvider) GenerateContent(prompt string) (string, error) {
	return fmt.Sprintf("// [AI GENERATED Content for prompt: %s]", prompt), nil
}
