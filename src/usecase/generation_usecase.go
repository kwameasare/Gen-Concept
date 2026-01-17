package usecase

import (
	"context"
	"gen-concept-api/config"
	"gen-concept-api/domain/repository"
	"gen-concept-api/domain/service"

	"github.com/google/uuid"
)

type GenerationUsecase struct {
	blueprintRepo     repository.BlueprintRepository
	generationService *service.GenerationService
}

func NewGenerationUsecase(cfg *config.Config, blueprintRepo repository.BlueprintRepository, genService *service.GenerationService) *GenerationUsecase {
	return &GenerationUsecase{
		blueprintRepo:     blueprintRepo,
		generationService: genService,
	}
}

func (u *GenerationUsecase) Generate(ctx context.Context, blueprintID uuid.UUID, inputs map[string]string) (string, error) {
	// 1. Fetch Blueprint
	blueprint, err := u.blueprintRepo.GetByUuidWithRelationships(ctx, blueprintID)
	if err != nil {
		return "", err
	}

	// 2. Generate
	return u.generationService.GenerateCode(blueprint, inputs)
}
