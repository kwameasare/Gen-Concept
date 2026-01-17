package dto

import (
	usecaseDto "gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

// Library is a DTO for library-related data
type Library struct {
	Uuid                   uuid.UUID              `json:"uuid"`
	Name                   string                 `json:"standardName"`
	Version                string                 `json:"version"`
	Description            string                 `json:"description"`
	RepositoryURL          string                 `json:"repositoryURL"`
	Namespace              string                 `json:"namespace"`
	ExposedFunctionalities []LibraryFunctionality `json:"exposedFunctionalities"`
	OrganizationID         *uint                  `json:"organizationID,omitempty"`
	TeamID                 *uint                  `json:"teamID,omitempty"`
}

type LibraryFunctionality struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
}

// ToUseCaseLibrary converts API DTO to usecase DTO
func ToUseCaseLibrary(library Library) usecaseDto.Library {
	functionalities := make([]usecaseDto.LibraryFunctionality, len(library.ExposedFunctionalities))
	for i, f := range library.ExposedFunctionalities {
		functionalities[i] = usecaseDto.LibraryFunctionality{
			Uuid:        f.Uuid,
			Name:        f.Name,
			Type:        f.Type,
			Description: f.Description,
		}
	}

	return usecaseDto.Library{
		Uuid:                   library.Uuid,
		Name:                   library.Name,
		Version:                library.Version,
		Description:            library.Description,
		RepositoryURL:          library.RepositoryURL,
		Namespace:              library.Namespace,
		ExposedFunctionalities: functionalities,
		OrganizationID:         library.OrganizationID,
		TeamID:                 library.TeamID,
	}
}

// ToLibraryResponse converts usecase DTO to API response DTO
func ToLibraryResponse(library usecaseDto.Library) Library {
	functionalities := make([]LibraryFunctionality, len(library.ExposedFunctionalities))
	for i, f := range library.ExposedFunctionalities {
		functionalities[i] = LibraryFunctionality{
			Uuid:        f.Uuid,
			Name:        f.Name,
			Type:        f.Type,
			Description: f.Description,
		}
	}

	return Library{
		Uuid:                   library.Uuid,
		Name:                   library.Name,
		Version:                library.Version,
		Description:            library.Description,
		RepositoryURL:          library.RepositoryURL,
		Namespace:              library.Namespace,
		ExposedFunctionalities: functionalities,
		OrganizationID:         library.OrganizationID,
		TeamID:                 library.TeamID,
	}
}
