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

type TeamHandler struct {
	usecase *usecase.TeamUsecase
}

func NewTeamHandler(cfg *config.Config) *TeamHandler {
	return &TeamHandler{
		usecase: usecase.NewTeamUsecase(cfg, dependency.GetTeamRepository(cfg)),
	}
}

func (h *TeamHandler) Create(c *gin.Context) {
	request := new(dto.CreateTeam)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	team, err := h.usecase.Create(c, dto.ToUseCaseCreateTeam(*request))
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToTeamResponse(team)
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

func (h *TeamHandler) Update(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	request := new(dto.UpdateTeam)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	team, err := h.usecase.Update(c, uuid, dto.ToUseCaseUpdateTeam(*request))
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToTeamResponse(team)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

func (h *TeamHandler) Delete(c *gin.Context) {
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

func (h *TeamHandler) GetById(c *gin.Context) {
	uuidStr := c.Params.ByName("id")
	uuid, uuidErr := uuid.Parse(uuidStr)
	if uuidErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, uuidErr))
		return
	}

	team, err := h.usecase.GetById(c, uuid)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := dto.ToTeamResponse(team)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

func (h *TeamHandler) GetByFilter(c *gin.Context) {
	req := new(filter.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	teams, err := h.usecase.GetByFilter(c, *req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := filter.PagedList[dto.Team]{
		PageNumber:      teams.PageNumber,
		PageSize:        teams.PageSize,
		TotalRows:       teams.TotalRows,
		TotalPages:      teams.TotalPages,
		HasPreviousPage: teams.HasPreviousPage,
		HasNextPage:     teams.HasNextPage,
	}

	items := []dto.Team{}
	for _, item := range *teams.Items {
		items = append(items, dto.ToTeamResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
