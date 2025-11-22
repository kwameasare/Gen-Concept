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

type EntityFieldUsecase struct {
	base       *BaseUsecase[model.EntityField, dto.EntityField, dto.EntityField, dto.EntityField]
	entityRepo repository.EntityRepository
}

func NewEntityFieldUsecase(cfg *config.Config, repository repository.EntityFieldRepository, entityRepo repository.EntityRepository) *EntityFieldUsecase {
	return &EntityFieldUsecase{
		base:       NewBaseUsecase[model.EntityField, dto.EntityField, dto.EntityField, dto.EntityField](cfg, repository),
		entityRepo: entityRepo,
	}
}

// Create
func (u *EntityFieldUsecase) Create(ctx context.Context, req dto.EntityField) (dto.EntityField, error) {
	var response dto.EntityField
	// 1. Find Entity by UUID
	entity, err := u.entityRepo.GetById(ctx, req.EntityUuid)
	if err != nil {
		return response, err
	}

	// 2. Convert DTO to Model
	entityFieldModel, _ := common.TypeConverter[model.EntityField](req)

	// 3. Set EntityID
	entityFieldModel.EntityID = entity.ID

	// 4. Save
	createdField, err := u.base.repository.Create(ctx, entityFieldModel)
	if err != nil {
		return response, err
	}

	// 5. Convert back to DTO
	response, _ = common.TypeConverter[dto.EntityField](createdField)
	response.EntityUuid = req.EntityUuid
	return response, nil
}

// Update
func (s *EntityFieldUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.EntityField) (dto.EntityField, error) {
	return s.base.Update(ctx, uuid, req)
}

// Delete
func (s *EntityFieldUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *EntityFieldUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.EntityField, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *EntityFieldUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.EntityField], error) {
	return s.base.GetByFilter(ctx, req)
}
