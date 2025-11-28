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

type JourneyHandler struct {
	usecase *usecase.JourneyUsecase
}

func NewJourneyHandler(cfg *config.Config) *JourneyHandler {
	return &JourneyHandler{
		usecase: usecase.NewJourneyUsecase(cfg, dependency.GetJourneyRepository(cfg)),
	}
}

// CreateJourney godoc
// @Summary Create a Journey
// @Description Create a Journey
// @Tags Journeys
// @Accept json
// @produces json
// @Param Request body dto.Journey true "Create a Journey"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.Journey} "Journey response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/journeys/ [post]
// @Security AuthBearer
func (h *JourneyHandler) Create(c *gin.Context) {
	request := new(dto.Journey)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// map http request body to usecase input and call use case method
	journey, err := h.usecase.Create(c, *request.ToUsecaseJourneyDTO())

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := new(dto.Journey)
	response.FromUsecaseJourneyDTO(&journey)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// UpdateJourney godoc
// @Summary Update a Journey
// @Description Update a Journey
// @Tags Journeys
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.Journey true "Update a Journey"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Journey} "Journey response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/journeys/{id} [put]
// @Security AuthBearer
func (h *JourneyHandler) Update(c *gin.Context) {
	// bind http request
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}
	request := new(dto.Journey)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return

	}
	// map http request body to usecase input and call use case method
	journey, err := h.usecase.Update(c, uuid, *request.ToUsecaseJourneyDTO())

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := new(dto.Journey)
	response.FromUsecaseJourneyDTO(&journey)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteJourney godoc
// @Summary Delete a Journey
// @Description Delete a Journey
// @Tags Journeys
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/journeys/{id} [delete]
// @Security AuthBearer
func (h *JourneyHandler) Delete(c *gin.Context) {
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

// GetJourney godoc
// @Summary Get a Journey
// @Description Get a Journey
// @Tags Journeys
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Journey} "Journey response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/journeys/{id} [get]
// @Security AuthBearer
func (h *JourneyHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	// call use case method
	journey, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := new(dto.Journey)
	response.FromUsecaseJourneyDTO(&journey)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// GetJourneysByFilter godoc
// @Summary Get Journeys by Filter
// @Description Get Journeys by Filter
// @Tags Journeys
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.Journey]} "Journey response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/journeys/get-by-filter [post]
// @Security AuthBearer
func (h *JourneyHandler) GetByFilter(c *gin.Context) {
	req := new(filter.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// call use case method
	journeys, err := h.usecase.GetByFilter(c, *req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	response := filter.PagedList[dto.Journey]{
		PageNumber:      journeys.PageNumber,
		PageSize:        journeys.PageSize,
		TotalRows:       journeys.TotalRows,
		TotalPages:      journeys.TotalPages,
		HasPreviousPage: journeys.HasPreviousPage,
		HasNextPage:     journeys.HasNextPage,
	}

	// map usecase response to http response
	items := []dto.Journey{}
	for _, item := range *journeys.Items {
		res := new(dto.Journey)
		res.FromUsecaseJourneyDTO(&item)
		items = append(items, *res)
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
