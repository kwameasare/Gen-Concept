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

type FileUsecase struct {
	base *BaseUsecase[model.File, dto.CreateFile, dto.UpdateFile, dto.File]
}

func NewFileUsecase(cfg *config.Config, repository repository.FileRepository) *FileUsecase {
	return &FileUsecase{
		base: NewBaseUsecase[model.File, dto.CreateFile, dto.UpdateFile, dto.File](cfg, repository),
	}
}

// Create
func (u *FileUsecase) Create(ctx context.Context, req dto.CreateFile) (dto.File, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *FileUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.UpdateFile) (dto.File, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *FileUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *FileUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.File, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *FileUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.File], error) {
	return s.base.GetByFilter(ctx, req)
}
