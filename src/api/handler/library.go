package handler

import (
	"gen-concept-api/api/dto"
	"gen-concept-api/api/helper"
	"gen-concept-api/config"
	"gen-concept-api/dependency"
	"gen-concept-api/domain/filter"
	"gen-concept-api/infra/git"
	"gen-concept-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LibraryHandler struct {
	usecase *usecase.LibraryUsecase
}

func NewLibraryHandler(cfg *config.Config) *LibraryHandler {
	return &LibraryHandler{
		usecase: usecase.NewLibraryUsecase(cfg, dependency.GetLibraryRepository(cfg), git.NewGitHubProvider()),
	}
}

func (h *LibraryHandler) Create(c *gin.Context) {
	request := new(dto.Library)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	library, err := h.usecase.Create(c, dto.ToUseCaseLibrary(*request))
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToLibraryResponse(library)
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

func (h *LibraryHandler) Update(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	request := new(dto.Library)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	library, err := h.usecase.Update(c, uuid, dto.ToUseCaseLibrary(*request))
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToLibraryResponse(library)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

func (h *LibraryHandler) Delete(c *gin.Context) {
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

func (h *LibraryHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	library, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToLibraryResponse(library)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

func (h *LibraryHandler) GetByFilter(c *gin.Context) {
	req := new(filter.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	libraries, err := h.usecase.GetByFilter(c, *req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := filter.PagedList[dto.Library]{
		PageNumber:      libraries.PageNumber,
		PageSize:        libraries.PageSize,
		TotalRows:       libraries.TotalRows,
		TotalPages:      libraries.TotalPages,
		HasPreviousPage: libraries.HasPreviousPage,
		HasNextPage:     libraries.HasNextPage,
	}

	items := []dto.Library{}
	for _, item := range *libraries.Items {
		items = append(items, dto.ToLibraryResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

func (h *LibraryHandler) Discover(c *gin.Context) {
	request := struct {
		RepositoryURL string `json:"repositoryUrl" binding:"required"`
		Token         string `json:"token"`
	}{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	library, err := h.usecase.DiscoverAndImport(c, request.RepositoryURL, request.Token)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToLibraryResponse(library)
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}
