package main

import (
	"fmt"
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
	fmt.Println("Starting Gen-Concept API...")
	cfg := config.GetConfig()
	fmt.Println("Config loaded successfully")

	logger := logging.NewLogger(cfg)
	fmt.Println("Logger initialized")

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	fmt.Println("Redis connected")

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	fmt.Println("Database connected")

	migration.Up1()
	fmt.Println("Migrations completed")

	api.InitServer(cfg)
}
