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

type ProjectUsecase struct {
	base *BaseUsecase[model.Project, dto.Project, dto.Project, dto.Project]
}

func NewProjectUsecase(cfg *config.Config, repository repository.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{
		base: NewBaseUsecase[model.Project, dto.Project, dto.Project, dto.Project](cfg, repository),
	}
}

// Create
func (u *ProjectUsecase) Create(ctx context.Context, req dto.Project) (dto.Project, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *ProjectUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Project) (dto.Project, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *ProjectUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *ProjectUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Project, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *ProjectUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Project], error) {
	return s.base.GetByFilter(ctx, req)
}
