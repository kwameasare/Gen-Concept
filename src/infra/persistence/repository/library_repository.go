package repository

import (
	"gen-concept-api/config"
	"gen-concept-api/domain/contract/repository"
	"gen-concept-api/domain/model"
	"gen-concept-api/infra/persistence/database"
)

type LibraryRepository struct {
	*BaseRepository[model.Library]
}

func NewLibraryRepository(cfg *config.Config) repository.LibraryRepository {
	return &LibraryRepository{
		BaseRepository: NewBaseRepository[model.Library](cfg, []database.PreloadEntity{
			{Entity: "ExposedFunctionalities"},
			{Entity: "Blueprints"},
		}),
	}
}
