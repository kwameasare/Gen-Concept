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

type PropertyUsecase struct {
	base *BaseUsecase[model.Property, dto.CreateProperty, dto.UpdateProperty, dto.Property]
}

func NewPropertyUsecase(cfg *config.Config, repository repository.PropertyRepository) *PropertyUsecase {
	return &PropertyUsecase{
		base: NewBaseUsecase[model.Property, dto.CreateProperty, dto.UpdateProperty, dto.Property](cfg, repository),
	}
}

// Create
func (u *PropertyUsecase) Create(ctx context.Context, req dto.CreateProperty) (dto.Property, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *PropertyUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.UpdateProperty) (dto.Property, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *PropertyUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *PropertyUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Property, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *PropertyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Property], error) {
	return s.base.GetByFilter(ctx, req)
}
