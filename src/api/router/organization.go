package router

import (
	"gen-concept-api/api/handler"
	"gen-concept-api/config"
	"gen-concept-api/infra/persistence/repository"
	"gen-concept-api/usecase"

	"github.com/gin-gonic/gin"
)

func Organization(r *gin.RouterGroup, cfg *config.Config) {
	orgRepo := repository.NewOrganizationRepository(cfg)
	userRepo := repository.NewUserRepository(cfg)

	orgUsecase := usecase.NewOrganizationUsecase(cfg, orgRepo, userRepo)
	h := handler.NewOrganizationHandler(cfg, orgUsecase)

	r.POST("/onboard", h.Onboard)
}
