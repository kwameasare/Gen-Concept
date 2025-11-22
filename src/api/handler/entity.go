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

type EntityHandler struct {
	usecase *usecase.EntityUsecase
}

func NewEntityHandler(cfg *config.Config) *EntityHandler {
	return &EntityHandler{
		usecase: usecase.NewEntityUsecase(cfg, dependency.GetEntityRepository(cfg)),
	}
}

// CreateEntity godoc
// @Summary Create a Entity
// @Description Create a Entity
// @Tags Entitys
// @Accept json
// @produces json
// @Param Request body dto.Entity true "Create a Entity"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.Entity} "Entity response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/Entitys/ [post]
// @Security AuthBearer
func (h *EntityHandler) Create(c *gin.Context) {
	request := new(dto.Entity)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	validationError := request.Validate()
	if validationError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, validationError))
		return
	}
	// map http request body to usecase input and call use case method
	Entity, err := h.usecase.Create(c, dto.ToUseCaseEntity(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToEntityResponse(Entity)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// UpdateEntity godoc
// @Summary Update a Entity
// @Description Update a Entity
// @Tags Entitys
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.Entity true "Update a Entity"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Entity} "Entity response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Entitys/{id} [put]
// @Security AuthBearer
func (h *EntityHandler) Update(c *gin.Context) {
	// bind http request
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}
	request := new(dto.Entity)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return

	}
	// map http request body to usecase input and call use case method
	Entity, err := h.usecase.Update(c, uuid, dto.ToUseCaseEntity(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToEntityResponse(Entity)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteEntity godoc
// @Summary Delete a Entity
// @Description Delete a Entity
// @Tags Entitys
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Entitys/{id} [delete]
// @Security AuthBearer
func (h *EntityHandler) Delete(c *gin.Context) {
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

// GetEntity godoc
// @Summary Get a Entity
// @Description Get a Entity
// @Tags Entitys
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Entity} "Entity response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/Entitys/{id} [get]
// @Security AuthBearer
func (h *EntityHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	// call use case method
	Entity, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToEntityResponse(Entity)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// GetProperties godoc
// @Summary Get Properties
// @Description Get Properties
// @Tags Entitys
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.Entity]} "Entity response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/Entitys/get-by-filter [post]
// @Security AuthBearer
func (h *EntityHandler) GetByFilter(c *gin.Context) {
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
	response := filter.PagedList[dto.Entity]{
		PageNumber:      properties.PageNumber,
		PageSize:        properties.PageSize,
		TotalRows:       properties.TotalRows,
		TotalPages:      properties.TotalPages,
		HasPreviousPage: properties.HasPreviousPage,
		HasNextPage:     properties.HasNextPage,
	}

	// map usecase response to http response
	items := []dto.Entity{}
	for _, item := range *properties.Items {

		items = append(items, dto.ToEntityResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
