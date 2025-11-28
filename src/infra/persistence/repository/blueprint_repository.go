package repository

import (
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/infra/persistence/database"
)

type BlueprintRepository struct {
	*BaseRepository[model.Blueprint]
}

func NewBlueprintRepository(cfg *config.Config) repository.BlueprintRepository {
	return &BlueprintRepository{
		BaseRepository: NewBaseRepository[model.Blueprint](cfg, []database.PreloadEntity{
			{Entity: "Functionalities"},
			{Entity: "Functionalities.Operations"},
		}),
	}
}
