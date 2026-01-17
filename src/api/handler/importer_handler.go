package handler

import (
	api_dto "gen-concept-api/api/dto"
	"gen-concept-api/api/helper"
	"gen-concept-api/config"
	"gen-concept-api/domain/service"
	"gen-concept-api/infra/parser"
	"gen-concept-api/usecase/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImporterHandler struct {
	service *service.ImporterService
}

func NewImporterHandler(cfg *config.Config) *ImporterHandler {
	// Initialize dependencies
	goParser := parser.NewGoParser()
	importerService := service.NewImporterService(goParser)

	return &ImporterHandler{
		service: importerService,
	}
}

func (h *ImporterHandler) Parse(c *gin.Context) {
	request := struct {
		Content  string `json:"content" binding:"required"`
		Language string `json:"language" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	blueprint, err := h.service.ImportFromSource(request.Content, request.Language)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// Map domain model to usecase dto, then to api dto
	usecaseDto := dto.FromBlueprintModel(blueprint)
	response := api_dto.ToBlueprintResponse(usecaseDto)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
