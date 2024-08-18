package handler

import (
	"gen-concept-api/api/dto"
	_ "gen-concept-api/api/helper"
	"gen-concept-api/config"
	"gen-concept-api/dependency"
	_ "gen-concept-api/domain/filter"
	"gen-concept-api/usecase"

	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	usecase *usecase.GearboxUsecase
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		usecase: usecase.NewGearboxUsecase(cfg, dependency.GetGearboxRepository(cfg)),
	}
}

// CreateGearbox godoc
// @Summary Create a Gearbox
// @Description Create a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body dto.CreateGearboxRequest true "Create a Gearbox"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateGearbox, dto.ToGearboxResponse, h.usecase.Create)
}

// UpdateGearbox godoc
// @Summary Update a Gearbox
// @Description Update a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a Gearbox"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateGearbox, dto.ToGearboxResponse, h.usecase.Update)
}

// DeleteGearbox godoc
// @Summary Delete a Gearbox
// @Description Delete a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetGearbox godoc
// @Summary Get a Gearbox
// @Description Get a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToGearboxResponse, h.usecase.GetById)
}

// GetGearboxes godoc
// @Summary Get Gearboxes
// @Description Get Gearboxes
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.GearboxResponse]} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/get-by-filter [post]
// @Security AuthBearer
func (h *GearboxHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToGearboxResponse, h.usecase.GetByFilter)
}
