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

type EntityUsecase struct {
	base *BaseUsecase[model.Entity, dto.Entity, dto.Entity, dto.Entity]
}

func NewEntityUsecase(cfg *config.Config, repository repository.EntityRepository) *EntityUsecase {
	return &EntityUsecase{
		base: NewBaseUsecase[model.Entity, dto.Entity, dto.Entity, dto.Entity](cfg, repository),
	}
}

// Create
func (u *EntityUsecase) Create(ctx context.Context, req dto.Entity) (dto.Entity, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *EntityUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Entity) (dto.Entity, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *EntityUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *EntityUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Entity, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *EntityUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Entity], error) {
	return s.base.GetByFilter(ctx, req)
}
