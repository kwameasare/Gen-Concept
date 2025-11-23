package repository

import (
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/infra/persistence/database"
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
