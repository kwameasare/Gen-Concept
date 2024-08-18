package usecase

import (
	"context"

	"gen-concept-api/config"
	"gen-concept-api/domain/filter"
	model "gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/usecase/dto"
)

type CarModelCommentUsecase struct {
	base *BaseUsecase[model.CarModelComment, dto.CreateCarModelComment, dto.UpdateCarModelComment, dto.CarModelComment]
}

func NewCarModelCommentUsecase(cfg *config.Config, repository repository.CarModelCommentRepository) *CarModelCommentUsecase {
	return &CarModelCommentUsecase{
		base: NewBaseUsecase[model.CarModelComment, dto.CreateCarModelComment, dto.UpdateCarModelComment, dto.CarModelComment](cfg, repository),
	}
}

// Create
func (u *CarModelCommentUsecase) Create(ctx context.Context, req dto.CreateCarModelComment) (dto.CarModelComment, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelCommentUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelComment) (dto.CarModelComment, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelCommentUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelCommentUsecase) GetById(ctx context.Context, id int) (dto.CarModelComment, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelCommentUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelComment], error) {
	return s.base.GetByFilter(ctx, req)
}
