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

type LibraryUsecase struct {
	base *BaseUsecase[model.Library, dto.Library, dto.Library, dto.Library]
}

func NewLibraryUsecase(cfg *config.Config, repository repository.LibraryRepository) *LibraryUsecase {
	return &LibraryUsecase{
		base: NewBaseUsecase[model.Library, dto.Library, dto.Library, dto.Library](cfg, repository),
	}
}

// Create
func (u *LibraryUsecase) Create(ctx context.Context, req dto.Library) (dto.Library, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *LibraryUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Library) (dto.Library, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *LibraryUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *LibraryUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Library, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *LibraryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Library], error) {
	return s.base.GetByFilter(ctx, req)
}
