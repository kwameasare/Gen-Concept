package dto

import (
	"gen-concept-api/domain/model"

	"github.com/google/uuid"
)

type Team struct {
	ID             uint      `json:"id"`
	Uuid           uuid.UUID `json:"uuid"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OrganizationID uint      `json:"organizationID"`
	// Users are handled separately usually, but we can include basic info or just IDs
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

// ToModel
func (t CreateTeam) ToModel() model.Team {
	return model.Team{
		Name:           t.Name,
		Description:    t.Description,
		OrganizationID: t.OrganizationID,
	}
}

func FromTeamModel(team model.Team) Team {
	return Team{
		ID:             team.ID,
		Uuid:           team.Uuid,
		Name:           team.Name,
		Description:    team.Description,
		OrganizationID: team.OrganizationID,
	}
}
