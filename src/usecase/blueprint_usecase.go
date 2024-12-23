package usecase

import (
	"context"

	"gen-concept-api/config"
	"gen-concept-api/domain/filter"
	model "gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

type BlueprintUsecase struct {
	base *BaseUsecase[model.Blueprint, dto.Blueprint, dto.Blueprint, dto.Blueprint]
}

func NewBlueprintUsecase(cfg *config.Config, repository repository.BlueprintRepository) *BlueprintUsecase {
	return &BlueprintUsecase{
		base: NewBaseUsecase[model.Blueprint, dto.Blueprint, dto.Blueprint, dto.Blueprint](cfg, repository),
	}
}

// Create
func (u *BlueprintUsecase) Create(ctx context.Context, req dto.Blueprint) (dto.Blueprint, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *BlueprintUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Blueprint) (dto.Blueprint, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *BlueprintUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *BlueprintUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Blueprint, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *BlueprintUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Blueprint], error) {
	return s.base.GetByFilter(ctx, req)
}
