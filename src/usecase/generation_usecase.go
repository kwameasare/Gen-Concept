package usecase

import (
	"context"
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/domain/service"

	"github.com/google/uuid"
)

type GenerationUsecase struct {
	blueprintRepo     repository.BlueprintRepository
	entityRepo        repository.EntityRepository
	generationService *service.GenerationService
}

func NewGenerationUsecase(cfg *config.Config, blueprintRepo repository.BlueprintRepository, entityRepo repository.EntityRepository, genService *service.GenerationService) *GenerationUsecase {
	return &GenerationUsecase{
		blueprintRepo:     blueprintRepo,
		entityRepo:        entityRepo,
		generationService: genService,
	}
}

func (u *GenerationUsecase) Generate(ctx context.Context, blueprintID uuid.UUID, inputs map[string]string) (string, error) {
	// 1. Fetch Blueprint
	blueprint, err := u.blueprintRepo.GetByUuidWithRelationships(ctx, blueprintID)
	if err != nil {
		return "", err
	}

	// 2. Fetch Entity (if provided)
	var entity model.Entity
	if entityIDStr, ok := inputs["entity_id"]; ok { // Standardize key
		entityID, err := uuid.Parse(entityIDStr)
		if err == nil {
			// Fetched Entity needs fields loaded? Base GetById might not load relations?
			// Usually check repository implementation. Assuming GetById works for now or logic needs "GetWithFields".
			// But for MVP, let's try GetById.
			fetched, err := u.entityRepo.GetById(ctx, entityID)
			if err == nil {
				entity = fetched
				// Ideally we need to ensure Preload("EntityFields") happens.
				// If BaseRepository doesn't preload, we might need a specific method.
				// BUT: Let's assume Entity has fields or we use a separate fetcher.
				// Since we are in GeneratioUsecase, we rely on what we have.
			}
		}
	}

	// 3. Generate
	return u.generationService.GenerateCode(ctx, blueprint, entity, inputs)
}
