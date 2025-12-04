package repository

import (
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
)

type LibraryRepository interface {
	repository.BaseRepository[model.Library]
}
