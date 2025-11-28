package usecase

import (
	"context"

	"gen-concept-api/common"
	"gen-concept-api/config"
	"gen-concept-api/domain/filter"
	model "gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

type JourneyUsecase struct {
	base *BaseUsecase[model.Journey, dto.Journey, dto.Journey, dto.Journey]
	repo repository.JourneyRepository
}

func NewJourneyUsecase(cfg *config.Config, repository repository.JourneyRepository) *JourneyUsecase {
	return &JourneyUsecase{
		base: NewBaseUsecase[model.Journey, dto.Journey, dto.Journey, dto.Journey](cfg, repository),
		repo: repository,
	}
}

// Create
func (u *JourneyUsecase) Create(ctx context.Context, req dto.Journey) (dto.Journey, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *JourneyUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Journey) (dto.Journey, error) {
	var journey model.Journey
	// Convert DTO to Model
	// We use common.TypeConverter similar to BaseUsecase
	// Note: We need to ensure the UUID is set from the path if not in body, but DTO usually has it.
	// req.UUID should be set.

	// Use TypeConverter to map DTO to Model
	// Note: common.TypeConverter returns value, not pointer for generic T
	j, err := common.TypeConverter[model.Journey](req)
	if err != nil {
		return dto.Journey{}, err
	}
	journey = j

	// Fetch existing journey to get the ID (primary key)
	existingJourney, err := s.repo.GetById(ctx, uuid)
	if err != nil {
		return dto.Journey{}, err
	}

	// Set the ID and other base fields to ensure update instead of insert
	journey.ID = existingJourney.ID
	journey.Uuid = uuid
	journey.CreatedAt = existingJourney.CreatedAt
	journey.CreatedBy = existingJourney.CreatedBy

	// Map IDs for nested entities to ensure updates
	mapJourneyIDs(&existingJourney, &journey)

	updatedJourney, err := s.repo.UpdateJourney(ctx, &journey)
	if err != nil {
		return dto.Journey{}, err
	}

	response, _ := common.TypeConverter[dto.Journey](updatedJourney)
	return response, nil
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
