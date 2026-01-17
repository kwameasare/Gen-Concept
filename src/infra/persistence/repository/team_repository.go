package repository

import (
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/infra/persistence/database"
)

type TeamRepository struct {
	*BaseRepository[model.Team]
}

func NewTeamRepository(cfg *config.Config) repository.TeamRepository {
	return &TeamRepository{
		BaseRepository: NewBaseRepository[model.Team](cfg, []database.PreloadEntity{
			{Entity: "Organization"},
			{Entity: "Users"},
		}),
	}
}
