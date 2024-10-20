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

type JourneyUsecase struct {
	base *BaseUsecase[model.Journey, dto.Journey, dto.Journey, dto.Journey]
}

func NewJourneyUsecase(cfg *config.Config, repository repository.JourneyRepository) *JourneyUsecase {
	return &JourneyUsecase{
		base: NewBaseUsecase[model.Journey, dto.Journey, dto.Journey, dto.Journey](cfg, repository),
	}
}

// Create
func (u *JourneyUsecase) Create(ctx context.Context, req dto.Journey) (dto.Journey, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *JourneyUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Journey) (dto.Journey, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *JourneyUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *JourneyUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Journey, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *JourneyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Journey], error) {
	return s.base.GetByFilter(ctx, req)
}