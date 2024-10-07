package dependency

import (
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	contractRepository "gen-concept-api/domain/repository"
	database "gen-concept-api/infra/persistence/database"
	infraRepository "gen-concept-api/infra/persistence/repository"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository(cfg)
}

// func GetProjectRepository(cfg *config.Config) contractRepository.ProjectRepository {
// 	var preloads []database.PreloadEntity = []database.PreloadEntity{
// 		// {Entity: "Company.Country"},
// 		// {Entity: "CarType"},
// 		// {Entity: "Gearbox"},
// 		// {Entity: "CarModelColors.Color"},
// 		// {Entity: "CarModelYears.PersianYear"},
// 		// {Entity: "CarModelYears.CarModelPriceHistories"},
// 		// {Entity: "CarModelProperties.Property.Category"},
// 		// {Entity: "CarModelImages.Image"},
// 		// {Entity: "CarModelComments.User"},
// 	}
// 	return infraRepository.NewBaseRepository[model.Project](cfg, preloads)
// }


func GetFileRepository(cfg *config.Config) contractRepository.FileRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.File](cfg, preloads)
}


func GetPropertyCategoryRepository(cfg *config.Config) contractRepository.PropertyCategoryRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Properties"}}
	return infraRepository.NewBaseRepository[model.PropertyCategory](cfg, preloads)
}

func GetPropertyRepository(cfg *config.Config) contractRepository.PropertyRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Category"}}
	return infraRepository.NewBaseRepository[model.Property](cfg, preloads)
}
func GetProjectRepository(cfg *config.Config) contractRepository.ProjectRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Entities"}, {Entity: "Entities.DependsOnEntities"}, {Entity: "Entities.EntityFields"}, {Entity: "Entities.EntityFields.InputValidations"}}
	return infraRepository.NewBaseRepository[model.Project](cfg, preloads)
}

func GetRoleRepository(cfg *config.Config) contractRepository.RoleRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Role](cfg, preloads)
}
