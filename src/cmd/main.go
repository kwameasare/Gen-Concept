package main

import (
	"gen-concept-api/api"
	"gen-concept-api/config"
	"gen-concept-api/infra/cache"
	database "gen-concept-api/infra/persistence/database"
	"gen-concept-api/infra/persistence/migration"
	"gen-concept-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	api.InitServer(cfg)
}
