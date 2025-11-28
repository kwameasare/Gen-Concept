package repository

import (
	"context"
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/infra/persistence/database"

	"gorm.io/gorm"
)

type JourneyRepository struct {
	*BaseRepository[model.Journey]
}

func NewJourneyRepository(cfg *config.Config) repository.JourneyRepository {
	return &JourneyRepository{
		BaseRepository: NewBaseRepository[model.Journey](cfg, []database.PreloadEntity{
			{Entity: "EntityJourneys"},
			{Entity: "EntityJourneys.Operations"},
			{Entity: "EntityJourneys.Operations.BackendJourney"},
			{Entity: "EntityJourneys.Operations.BackendJourney.FieldsInvolved"},
			{Entity: "EntityJourneys.Operations.BackendJourney.RetryConditions"},
			{Entity: "EntityJourneys.Operations.BackendJourney.ResponseActions"},
			{Entity: "EntityJourneys.Operations.Filters"},
			{Entity: "EntityJourneys.Operations.Sort"},
		}),
	}
}

func (r *JourneyRepository) UpdateJourney(ctx context.Context, journey *model.Journey) (*model.Journey, error) {
	tx := r.database.WithContext(ctx).Begin()

	// Use FullSaveAssociations to update nested structures
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(journey).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return journey, nil
}
