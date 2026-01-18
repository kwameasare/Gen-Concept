package handler

import (
	"fmt"
	"gen-concept-api/api/helper"
	"gen-concept-api/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JourneyGraphHandler struct {
	service *service.JourneyGraphService
}

func NewJourneyGraphHandler() *JourneyGraphHandler {
	return &JourneyGraphHandler{
		service: service.NewJourneyGraphService(),
	}
}

func (h *JourneyGraphHandler) GetGraph(c *gin.Context) {
	journeyIDStr := c.Param("id")
	journeyID, err := uuid.Parse(journeyIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, fmt.Errorf("invalid journey id")))
		return
	}

	level := c.Query("level")
	parentIDStr := c.Query("parent_id")

	var parentID *uuid.UUID
	if parentIDStr != "" {
		parsed, err := uuid.Parse(parentIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, fmt.Errorf("invalid parent_id")))
			return
		}
		parentID = &parsed
	}

	nodes, edges, err := h.service.GetGraph(journeyID, level, parentID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	response := gin.H{
		"nodes": nodes,
		"edges": edges,
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
