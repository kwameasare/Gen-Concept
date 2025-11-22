package handler

import (
	"net/http"

	"gen-concept-api/api/helper"
	"gen-concept-api/config"
	"gen-concept-api/usecase"
	"gen-concept-api/usecase/dto"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	organizationUsecase *usecase.OrganizationUsecase
}

func NewOrganizationHandler(cfg *config.Config, organizationUsecase *usecase.OrganizationUsecase) *OrganizationHandler {
	return &OrganizationHandler{
		organizationUsecase: organizationUsecase,
	}
}

// Onboard godoc
// @Summary Onboard a new organization
// @Description Create a new organization and its first admin user
// @Tags Organization
// @Accept json
// @Produce json
// @Param request body dto.OnboardOrganizationRequest true "Onboard Request"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/organizations/onboard [post]
func (h *OrganizationHandler) Onboard(c *gin.Context) {
	req := dto.OnboardOrganizationRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	org, user, err := h.organizationUsecase.OnboardOrganization(c, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(map[string]interface{}{
		"organization": org,
		"user":         user,
	}, true, helper.Success))
}
