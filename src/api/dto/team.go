package dto

import (
	usecaseDto "gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

type Team struct {
	Uuid           uuid.UUID `json:"uuid"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OrganizationID uint      `json:"organizationID"`
}

type CreateTeam struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description"`
	OrganizationID uint   `json:"organizationID" validate:"required"`
}

type UpdateTeam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ToUseCaseCreateTeam
func ToUseCaseCreateTeam(t CreateTeam) usecaseDto.CreateTeam {
	return usecaseDto.CreateTeam{
		Name:           t.Name,
		Description:    t.Description,
		OrganizationID: t.OrganizationID,
	}
}

// ToUseCaseUpdateTeam
func ToUseCaseUpdateTeam(t UpdateTeam) usecaseDto.UpdateTeam {
	return usecaseDto.UpdateTeam{
		Name:        t.Name,
		Description: t.Description,
	}
}

// ToTeamResponse
func ToTeamResponse(t usecaseDto.Team) Team {
	return Team{
		Uuid:           t.Uuid,
		Name:           t.Name,
		Description:    t.Description,
		OrganizationID: t.OrganizationID,
	}
}
