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

type PropertyCategoryUsecase struct {
	base *BaseUsecase[model.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory]
}

func NewPropertyCategoryUsecase(cfg *config.Config, repository repository.PropertyCategoryRepository) *PropertyCategoryUsecase {
	return &PropertyCategoryUsecase{
		base: NewBaseUsecase[model.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory](cfg, repository),
	}
}

// Create
func (u *PropertyCategoryUsecase) Create(ctx context.Context, req dto.CreatePropertyCategory) (dto.PropertyCategory, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *PropertyCategoryUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.UpdatePropertyCategory) (dto.PropertyCategory, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *PropertyCategoryUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *PropertyCategoryUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.PropertyCategory, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *PropertyCategoryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.PropertyCategory], error) {
	return s.base.GetByFilter(ctx, req)
}
