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

type ProjectHandler struct {
	usecase *usecase.ProjectUsecase
}

func NewProjectHandler(cfg *config.Config) *ProjectHandler {
	return &ProjectHandler{
		usecase: usecase.NewProjectUsecase(cfg, dependency.GetProjectRepository(cfg)),
	}
}

// CreateProject godoc
// @Summary Create a Project
// @Description Create a Project
// @Tags Projects
// @Accept json
// @produces json
// @Param Request body dto.Project true "Create a Project"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.Project} "Project response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/projects/ [post]
// @Security AuthBearer
func (h *ProjectHandler) Create(c *gin.Context) {
	request := new(dto.Project)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// map http request body to usecase input and call use case method
	project, err := h.usecase.Create(c, dto.ToUseCaseProject(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToProjectResponse(project)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// UpdateProject godoc
// @Summary Update a Project
// @Description Update a Project
// @Tags Projects
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.Project true "Update a Project"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Project} "Project response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/projects/{id} [put]
// @Security AuthBearer
func (h *ProjectHandler) Update(c *gin.Context) {
	// bind http request
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}
	request := new(dto.Project)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return

	}
	// map http request body to usecase input and call use case method
	project, err := h.usecase.Update(c, uuid, dto.ToUseCaseProject(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToProjectResponse(project)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteProject godoc
// @Summary Delete a Project
// @Description Delete a Project
// @Tags Projects
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/projects/{id} [delete]
// @Security AuthBearer
func (h *ProjectHandler) Delete(c *gin.Context) {
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

// GetProject godoc
// @Summary Get a Project
// @Description Get a Project
// @Tags Projects
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Project} "Project response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/projects/{id} [get]
// @Security AuthBearer
func (h *ProjectHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	// call use case method
	project, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToProjectResponse(project)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// GetProperties godoc
// @Summary Get Properties
// @Description Get Properties
// @Tags Projects
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.Project]} "Project response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/projects/get-by-filter [post]
// @Security AuthBearer
func (h *ProjectHandler) GetByFilter(c *gin.Context) {
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
	response := filter.PagedList[dto.Project]{
		PageNumber:      properties.PageNumber,
		PageSize:        properties.PageSize,
		TotalRows:       properties.TotalRows,
		TotalPages:      properties.TotalPages,
		HasPreviousPage: properties.HasPreviousPage,
		HasNextPage:     properties.HasNextPage,
	}

	// map usecase response to http response
	items := []dto.Project{}
	for _, item := range *properties.Items {

		items = append(items, dto.ToProjectResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
