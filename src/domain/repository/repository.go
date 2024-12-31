package repository

import (
	"context"

	"gen-concept-api/domain/filter"
	"gen-concept-api/domain/model"

	"github.com/google/uuid"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, uuid uuid.UUID, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
	GetById(ctx context.Context, uuid uuid.UUID) (TEntity, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
}

type FileRepository interface {
	BaseRepository[model.File]
}

type PropertyCategoryRepository interface {
	BaseRepository[model.PropertyCategory]
}

type PropertyRepository interface {
	BaseRepository[model.Property]
}
type ProjectRepository interface {
	BaseRepository[model.Project]
}

type JourneyRepository interface {
	BaseRepository[model.Journey]
}

type BlueprintRepository interface {
	BaseRepository[model.Blueprint]
}

type EntityRepository interface {
	BaseRepository[model.Entity]
}
type UserRepository interface {
	ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	FetchUserInfo(ctx context.Context, username string, password string) (model.User, error)
	GetDefaultRole(ctx context.Context) (roleId uint, err error)
	CreateUser(ctx context.Context, u model.User) (model.User, error)
}

type RoleRepository interface {
	BaseRepository[model.Role]
}
