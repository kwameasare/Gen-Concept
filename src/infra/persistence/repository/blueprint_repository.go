package repository

import (
	"context"

	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/infra/persistence/database"

	"github.com/google/uuid"
)

type BlueprintRepository struct {
	*BaseRepository[model.Blueprint]
}

func NewBlueprintRepository(cfg *config.Config) repository.BlueprintRepository {
	return &BlueprintRepository{
		BaseRepository: NewBaseRepository[model.Blueprint](cfg, []database.PreloadEntity{
			{Entity: "Functionalities"},
			{Entity: "Functionalities.Operations"},
			{Entity: "Libraries"},
			{Entity: "Libraries.ExposedFunctionalities"},
		}),
	}
}

// CreateWithRelationships creates a blueprint including its nested Functionalities and Libraries
func (r *BlueprintRepository) CreateWithRelationships(ctx context.Context, blueprint model.Blueprint) (model.Blueprint, error) {
	tx := r.database.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the blueprint (basic fields only first)
	blueprintToCreate := model.Blueprint{
		StandardName: blueprint.StandardName,
		Type:         blueprint.Type,
		Description:  blueprint.Description,
	}

	if err := tx.Create(&blueprintToCreate).Error; err != nil {
		tx.Rollback()
		return model.Blueprint{}, err
	}

	// Create functionalities with operations
	for i := range blueprint.Functionalities {
		blueprint.Functionalities[i].BlueprintID = blueprintToCreate.ID
		if err := tx.Create(&blueprint.Functionalities[i]).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}
	}

	// Handle Libraries (many-to-many) - Manually create junction table entries
	if len(blueprint.Libraries) > 0 {
		var libraryUUIDs []uuid.UUID
		for _, lib := range blueprint.Libraries {
			libraryUUIDs = append(libraryUUIDs, lib.Uuid)
		}

		// Fetch the actual library records from database (excluding soft-deleted ones)
		var actualLibraries []model.Library
		if err := tx.Where("uuid IN ? AND deleted_by IS NULL", libraryUUIDs).Find(&actualLibraries).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}

		// Manually create junction table entries
		for _, lib := range actualLibraries {
			junction := model.BlueprintLibrary{
				BlueprintID:     blueprintToCreate.ID,
				LibraryID:       lib.ID,
				RequiredVersion: lib.Version,
			}
			if err := tx.Create(&junction).Error; err != nil {
				tx.Rollback()
				return model.Blueprint{}, err
			}
		}
	}

	tx.Commit()

	// Reload the blueprint with all relationships
	var result model.Blueprint
	if err := r.database.WithContext(ctx).
		Preload("Functionalities").
		Preload("Functionalities.Operations").
		Preload("Libraries").
		Preload("Libraries.ExposedFunctionalities").
		Where("uuid = ?", blueprintToCreate.Uuid).
		First(&result).Error; err != nil {
		return model.Blueprint{}, err
	}

	return result, nil
}

// UpdateWithRelationships updates a blueprint including its nested Functionalities and Libraries
func (r *BlueprintRepository) UpdateWithRelationships(ctx context.Context, blueprintUuid uuid.UUID, blueprint model.Blueprint) (model.Blueprint, error) {
	tx := r.database.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get existing blueprint
	var existing model.Blueprint
	if err := tx.Where("uuid = ? and deleted_by is null", blueprintUuid).First(&existing).Error; err != nil {
		tx.Rollback()
		return model.Blueprint{}, err
	}

	// Update basic fields
	existing.StandardName = blueprint.StandardName
	existing.Type = blueprint.Type
	existing.Description = blueprint.Description

	if err := tx.Save(&existing).Error; err != nil {
		tx.Rollback()
		return model.Blueprint{}, err
	}

	// Handle Functionalities (one-to-many with cascading delete)
	// First, get all functionality IDs for this blueprint
	var functionalityIDs []uint
	if err := tx.Model(&model.Functionality{}).Where("blueprint_id = ?", existing.ID).Pluck("id", &functionalityIDs).Error; err != nil {
		tx.Rollback()
		return model.Blueprint{}, err
	}

	// Hard delete all operations for these functionalities
	if len(functionalityIDs) > 0 {
		if err := tx.Unscoped().Where("functionality_id IN ?", functionalityIDs).Delete(&model.FunctionalOperation{}).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}
	}

	// Now hard delete all existing functionalities
	if err := tx.Unscoped().Where("blueprint_id = ?", existing.ID).Delete(&model.Functionality{}).Error; err != nil {
		tx.Rollback()
		return model.Blueprint{}, err
	}

	// Create new functionalities with operations
	for i := range blueprint.Functionalities {
		blueprint.Functionalities[i].BlueprintID = existing.ID
		if err := tx.Create(&blueprint.Functionalities[i]).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}
	}

	// Handle Libraries (many-to-many) - Manually manage junction table
	if len(blueprint.Libraries) > 0 {
		var libraryUUIDs []uuid.UUID
		for _, lib := range blueprint.Libraries {
			libraryUUIDs = append(libraryUUIDs, lib.Uuid)
		}

		// First, delete all existing blueprint-library associations
		if err := tx.Where("blueprint_id = ?", existing.ID).Delete(&model.BlueprintLibrary{}).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}

		// Fetch the actual library records from database (excluding soft-deleted ones)
		var actualLibraries []model.Library
		if err := tx.Where("uuid IN ? AND deleted_by IS NULL", libraryUUIDs).Find(&actualLibraries).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}

		// Manually create new junction table entries
		for _, lib := range actualLibraries {
			junction := model.BlueprintLibrary{
				BlueprintID:     existing.ID,
				LibraryID:       lib.ID,
				RequiredVersion: lib.Version,
			}
			if err := tx.Create(&junction).Error; err != nil {
				tx.Rollback()
				return model.Blueprint{}, err
			}
		}
	} else {
		// Clear all library associations if none provided
		if err := tx.Where("blueprint_id = ?", existing.ID).Delete(&model.BlueprintLibrary{}).Error; err != nil {
			tx.Rollback()
			return model.Blueprint{}, err
		}
	}

	tx.Commit()

	// Reload the blueprint with all relationships
	var result model.Blueprint
	if err := r.database.WithContext(ctx).
		Preload("Functionalities").
		Preload("Functionalities.Operations").
		Preload("Libraries").
		Preload("Libraries.ExposedFunctionalities").
		Where("uuid = ?", blueprintUuid).
		First(&result).Error; err != nil {
		return model.Blueprint{}, err
	}

	return result, nil
}
