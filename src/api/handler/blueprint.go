package handler

import (
	"gen-concept-api/api/dto"
	"gen-concept-api/api/helper"
	"gen-concept-api/config"
	"gen-concept-api/dependency"
	"gen-concept-api/domain/filter"
	"gen-concept-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlueprintHandler struct {
	usecase *usecase.BlueprintUsecase
}

func NewBlueprintHandler(cfg *config.Config) *BlueprintHandler {
	return &BlueprintHandler{
		usecase: usecase.NewBlueprintUsecase(cfg, dependency.GetBlueprintRepository(cfg)),
	}
}

// CreateBlueprint godoc
// @Summary Create a Blueprint
// @Description Create a Blueprint
// @Tags Blueprints
// @Accept json
// @produces json
// @Param Request body dto.Blueprint true "Create a Blueprint"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.Blueprint} "Blueprint response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/Blueprints/ [post]
// @Security AuthBearer
func (h *BlueprintHandler) Create(c *gin.Context) {
	request := new(dto.Blueprint)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	validationError:= request.Validate()
	if validationError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, validationError))
		return
	}
	// map http request body to usecase input and call use case method
	Blueprint, err := h.usecase.Create(c, dto.ToUseCaseBlueprint(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToBlueprintResponse(Blueprint)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// UpdateBlueprint godoc
// @Summary Update a Blueprint
// @Description Update a Blueprint
// @Tags Blueprints
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.Blueprint true "Update a Blueprint"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Blueprint} "Blueprint response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Blueprints/{id} [put]
// @Security AuthBearer
func (h *BlueprintHandler) Update(c *gin.Context) {
	// bind http request
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}
	request := new(dto.Blueprint)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return

	}
	// map http request body to usecase input and call use case method
	Blueprint, err := h.usecase.Update(c, uuid, dto.ToUseCaseBlueprint(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToBlueprintResponse(Blueprint)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteBlueprint godoc
// @Summary Delete a Blueprint
// @Description Delete a Blueprint
// @Tags Blueprints
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Blueprints/{id} [delete]
// @Security AuthBearer
func (h *BlueprintHandler) Delete(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	err := h.usecase.Delete(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
}

// GetBlueprint godoc
// @Summary Get a Blueprint
// @Description Get a Blueprint
// @Tags Blueprints
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Blueprint} "Blueprint response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Blueprints/{id} [get]
// @Security AuthBearer
func (h *BlueprintHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	// call use case method
	Blueprint, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToBlueprintResponse(Blueprint)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// GetProperties godoc
// @Summary Get Properties
// @Description Get Properties
// @Tags Blueprints
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.Blueprint]} "Blueprint response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/Blueprints/get-by-filter [post]
// @Security AuthBearer
func (h *BlueprintHandler) GetByFilter(c *gin.Context) {
	req := new(filter.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// call use case method
	properties, err := h.usecase.GetByFilter(c, *req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	response := filter.PagedList[dto.Blueprint]{
		PageNumber:      properties.PageNumber,
		PageSize:        properties.PageSize,
		TotalRows:       properties.TotalRows,
		TotalPages:      properties.TotalPages,
		HasPreviousPage: properties.HasPreviousPage,
		HasNextPage:     properties.HasNextPage,
	}

	// map usecase response to http response
	items := []dto.Blueprint{}
	for _, item := range *properties.Items {

		items = append(items, dto.ToBlueprintResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
