package dto

type OrganizationDto struct {
	Name             string `json:"name" binding:"required"`
	Description      string `json:"description"`
	Domain           string `json:"domain" binding:"required"`
	SubscriptionPlan string `json:"subscriptionPlan"`
}

type OnboardOrganizationRequest struct {
	Organization OrganizationDto        `json:"organization" binding:"required"`
	User         RegisterUserByUsername `json:"user" binding:"required"`
}
