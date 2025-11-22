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

type EntityFieldHandler struct {
	usecase *usecase.EntityFieldUsecase
}

func NewEntityFieldHandler(cfg *config.Config) *EntityFieldHandler {
	return &EntityFieldHandler{
		usecase: usecase.NewEntityFieldUsecase(cfg, dependency.GetEntityFieldRepository(cfg), dependency.GetEntityRepository(cfg)),
	}
}

// CreateEntityField godoc
// @Summary Create an EntityField
// @Description Create an EntityField
// @Tags EntityFields
// @Accept json
// @produces json
// @Param Request body dto.EntityField true "Create an EntityField"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.EntityField} "EntityField response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/entity-fields/ [post]
// @Security AuthBearer
func (h *EntityFieldHandler) Create(c *gin.Context) {
	request := new(dto.EntityField)
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
	entityField, err := h.usecase.Create(c, dto.ToUseCaseEntityField(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToEntityFieldResponse(entityField)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// UpdateEntityField godoc
// @Summary Update an EntityField
// @Description Update an EntityField
// @Tags EntityFields
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.EntityField true "Update an EntityField"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.EntityField} "EntityField response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/entity-fields/{id} [put]
// @Security AuthBearer
func (h *EntityFieldHandler) Update(c *gin.Context) {
	// bind http request
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}
	request := new(dto.EntityField)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return

	}
	// map http request body to usecase input and call use case method
	entityField, err := h.usecase.Update(c, uuid, dto.ToUseCaseEntityField(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToEntityFieldResponse(entityField)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteEntityField godoc
// @Summary Delete an EntityField
// @Description Delete an EntityField
// @Tags EntityFields
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/entity-fields/{id} [delete]
// @Security AuthBearer
func (h *EntityFieldHandler) Delete(c *gin.Context) {
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

// GetEntityField godoc
// @Summary Get an EntityField
// @Description Get an EntityField
// @Tags EntityFields
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.EntityField} "EntityField response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/entity-fields/{id} [get]
// @Security AuthBearer
func (h *EntityFieldHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	// call use case method
	entityField, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToEntityFieldResponse(entityField)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// GetEntityFields godoc
// @Summary Get EntityFields
// @Description Get EntityFields
// @Tags EntityFields
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.EntityField]} "EntityField response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/entity-fields/get-by-filter [post]
// @Security AuthBearer
func (h *EntityFieldHandler) GetByFilter(c *gin.Context) {
	req := new(filter.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// call use case method
	fields, err := h.usecase.GetByFilter(c, *req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	response := filter.PagedList[dto.EntityField]{
		PageNumber:      fields.PageNumber,
		PageSize:        fields.PageSize,
		TotalRows:       fields.TotalRows,
		TotalPages:      fields.TotalPages,
		HasPreviousPage: fields.HasPreviousPage,
		HasNextPage:     fields.HasNextPage,
	}

	// map usecase response to http response
	items := []dto.EntityField{}
	for _, item := range *fields.Items {

		items = append(items, dto.ToEntityFieldResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
