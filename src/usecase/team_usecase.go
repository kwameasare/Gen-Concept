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

type TeamUsecase struct {
	base *BaseUsecase[model.Team, dto.CreateTeam, dto.UpdateTeam, dto.Team]
}

func NewTeamUsecase(cfg *config.Config, repository repository.TeamRepository) *TeamUsecase {
	return &TeamUsecase{
		base: NewBaseUsecase[model.Team, dto.CreateTeam, dto.UpdateTeam, dto.Team](cfg, repository),
	}
}

// Create
func (u *TeamUsecase) Create(ctx context.Context, req dto.CreateTeam) (dto.Team, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *TeamUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.UpdateTeam) (dto.Team, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *TeamUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *TeamUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Team, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *TeamUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Team], error) {
	return s.base.GetByFilter(ctx, req)
}
