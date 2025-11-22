package repository

import (
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/infra/persistence/database"
)

type OrganizationRepository struct {
	*BaseRepository[model.Organization]
}

func NewOrganizationRepository(cfg *config.Config) repository.OrganizationRepository {
	return &OrganizationRepository{
		BaseRepository: NewBaseRepository[model.Organization](cfg, []database.PreloadEntity{}),
	}
}
