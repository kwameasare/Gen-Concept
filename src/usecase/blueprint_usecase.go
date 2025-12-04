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

type BlueprintUsecase struct {
	base       *BaseUsecase[model.Blueprint, dto.Blueprint, dto.Blueprint, dto.Blueprint]
	repository repository.BlueprintRepository
	cfg        *config.Config
}

func NewBlueprintUsecase(cfg *config.Config, repository repository.BlueprintRepository) *BlueprintUsecase {
	return &BlueprintUsecase{
		base:       NewBaseUsecase[model.Blueprint, dto.Blueprint, dto.Blueprint, dto.Blueprint](cfg, repository),
		repository: repository,
		cfg:        cfg,
	}
}

// Create - Full implementation with nested relationships
func (u *BlueprintUsecase) Create(ctx context.Context, req dto.Blueprint) (dto.Blueprint, error) {
	var response dto.Blueprint

	// Convert DTO to model
	blueprintModel, _ := common.TypeConverter[model.Blueprint](req)

	// Manually ensure libraries are converted (for debugging)
	if len(req.Libraries) > 0 {
		libraries := make([]model.Library, len(req.Libraries))
		for i, lib := range req.Libraries {
			libraries[i] = model.Library{
				BaseModel: model.BaseModel{
					Uuid: lib.Uuid,
				},
				Name:          lib.Name,
				Version:       lib.Version,
				Description:   lib.Description,
				RepositoryURL: lib.RepositoryURL,
				Namespace:     lib.Namespace,
			}
		}
		blueprintModel.Libraries = libraries
	}

	// Use the custom repository method to create with relationships
	created, err := u.repository.CreateWithRelationships(ctx, blueprintModel)
	if err != nil {
		return response, err
	}

	// Convert result to DTO
	response, _ = common.TypeConverter[dto.Blueprint](created)
	return response, nil
}

// Update - Full implementation with nested relationships
func (s *BlueprintUsecase) Update(ctx context.Context, uuid uuid.UUID, req dto.Blueprint) (dto.Blueprint, error) {
	var response dto.Blueprint

	// Convert DTO to model
	blueprintModel, _ := common.TypeConverter[model.Blueprint](req)

	// Manually ensure libraries are converted (for debugging)
	if len(req.Libraries) > 0 {
		libraries := make([]model.Library, len(req.Libraries))
		for i, lib := range req.Libraries {
			libraries[i] = model.Library{
				BaseModel: model.BaseModel{
					Uuid: lib.Uuid,
				},
				Name:          lib.Name,
				Version:       lib.Version,
				Description:   lib.Description,
				RepositoryURL: lib.RepositoryURL,
				Namespace:     lib.Namespace,
			}
		}
		blueprintModel.Libraries = libraries
	}

	// Use the custom repository method to update with relationships
	updated, err := s.repository.UpdateWithRelationships(ctx, uuid, blueprintModel)
	if err != nil {
		return response, err
	}

	// Convert result to DTO
	response, _ = common.TypeConverter[dto.Blueprint](updated)
	return response, nil
}

// Delete
func (s *BlueprintUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {
	return s.base.Delete(ctx, uuid)
}

// Get By Id
func (s *BlueprintUsecase) GetById(ctx context.Context, uuid uuid.UUID) (dto.Blueprint, error) {
	return s.base.GetById(ctx, uuid)
}

// Get By Filter
func (s *BlueprintUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Blueprint], error) {
	return s.base.GetByFilter(ctx, req)
}
