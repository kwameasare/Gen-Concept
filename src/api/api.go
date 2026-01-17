package api

import (
	"fmt"

	"gen-concept-api/api/middleware"
	"gen-concept-api/api/router"
	validation "gen-concept-api/api/validation"
	"gen-concept-api/config"
	"gen-concept-api/docs"
	"gen-concept-api/pkg/logging"
	"gen-concept-api/pkg/metrics"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	RegisterValidators()
	RegisterPrometheus()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(middleware.Prometheus())
	r.Use(gin.CustomRecovery(middleware.ErrorHandler) /*middleware.TestMiddleware()*/, middleware.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Test
		health := v1.Group("/health")
		testRouter := v1.Group("/test" /*middleware.Authentication(cfg), middleware.Authorization([]string{"admin"})*/)

		// User
		users := v1.Group("/users")
		organizations := v1.Group("/organizations")

		// Base
		files := v1.Group("/files", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))

		// Property
		properties := v1.Group("/properties", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		propertyCategories := v1.Group("/property-categories", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))

		//Project
		projects := v1.Group("/projects", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		entities := v1.Group("/entities", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		entityFields := v1.Group("/entity-fields", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		// Journey
		journeys := v1.Group("/journeys", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		// Blueprints
		blueprints := v1.Group("/blueprints", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		// Libraries
		libraries := v1.Group("/libraries", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))

		// Test
		router.Health(health)
		router.TestRouter(testRouter)

		// User
		router.User(users, cfg)
		router.Organization(organizations, cfg)

		// Base
		router.File(files, cfg)
		// Property
		router.Property(properties, cfg)
		router.PropertyCategory(propertyCategories, cfg)
		r.Static("/static", "./uploads")

		//Project
		router.Project(projects, cfg)
		// Entity
		router.Entity(entities, cfg)
		// Entity Fields
		router.EntityField(entityFields, cfg)
		// Journey
		router.Journey(journeys, cfg)
		// Blueprints
		router.Blueprint(blueprints, cfg)
		// Libraries
		// Libraries
		router.Library(libraries, cfg)

		// Teams
		teams := v1.Group("/teams", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		router.Team(teams, cfg)

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		router.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
}
