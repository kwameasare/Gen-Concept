package handler

import (
	"gen-concept-api/api/helper"
	"gen-concept-api/config"
	"gen-concept-api/dependency"
	"gen-concept-api/domain/service"
	gen_ai "gen-concept-api/infra/ai"
	"gen-concept-api/infra/git"
	"gen-concept-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenerationHandler struct {
	usecase *usecase.GenerationUsecase
}

func NewGenerationHandler(cfg *config.Config) *GenerationHandler {
	// Initialize dependencies manually for now, or via dependency package
	blueprintRepo := dependency.GetBlueprintRepository(cfg)
	gitProvider := git.NewGitHubProvider() // Should probably be singleton or passed in
	aiProvider := gen_ai.NewMockAIProvider()
	genService := service.NewGenerationService(gitProvider, aiProvider)

	return &GenerationHandler{
		usecase: usecase.NewGenerationUsecase(cfg, blueprintRepo, genService),
	}
}

func (h *GenerationHandler) Preview(c *gin.Context) {
	request := struct {
		BlueprintID string            `json:"blueprintId" binding:"required"`
		Inputs      map[string]string `json:"inputs"`
	}{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	blueprintUUID, err := uuid.Parse(request.BlueprintID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	code, err := h.usecase.Generate(c, blueprintUUID, request.Inputs)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(map[string]string{"code": code}, true, 0))
}
