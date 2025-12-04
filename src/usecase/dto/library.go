package dto

import (
	"gen-concept-api/domain/model"

	"github.com/google/uuid"
)

type Library struct {
	Uuid                   uuid.UUID              `json:"uuid"`
	Name                   string                 `json:"standardName"`
	Version                string                 `json:"version"`
	Description            string                 `json:"description"`
	RepositoryURL          string                 `json:"repositoryURL"`
	Namespace              string                 `json:"namespace"`
	ExposedFunctionalities []LibraryFunctionality `json:"exposedFunctionalities"`
}

type LibraryFunctionality struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
}

// ToModel converts usecase DTO to domain model
func (l Library) ToModel() model.Library {
	functionalities := make([]model.LibraryFunctionality, len(l.ExposedFunctionalities))
	for i, f := range l.ExposedFunctionalities {
		functionalities[i] = model.LibraryFunctionality{
			BaseModel: model.BaseModel{
				Uuid: f.Uuid,
			},
			Name:        f.Name,
			Type:        f.Type,
			Description: f.Description,
		}
	}

	return model.Library{
		BaseModel: model.BaseModel{
			Uuid: l.Uuid,
		},
		Name:                   l.Name,
		Version:                l.Version,
		Description:            l.Description,
		RepositoryURL:          l.RepositoryURL,
		Namespace:              l.Namespace,
		ExposedFunctionalities: functionalities,
	}
}

// FromModel converts domain model to usecase DTO
func FromLibraryModel(library model.Library) Library {
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
	}
}
