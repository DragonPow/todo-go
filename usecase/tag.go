package usecase

import (
	"context"
	"project1/domain"
)

type tagUsecase struct {
	tagRepo domain.TagRepository
}

func NewTagUsecase(t domain.TagRepository) domain.TagRepository {
	return &tagUsecase{
		tagRepo: t,
	}
}

func (t *tagUsecase) FetchAll(ctx context.Context) ([]domain.Tag, error)
func (t *tagUsecase) GetByID(ctx context.Context, id int32) (domain.Tag, error)
func (t *tagUsecase) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Tag, error)
func (t *tagUsecase) Update(ctx context.Context, id int32, args ...interface{}) error
func (t *tagUsecase) Delete(ctx context.Context, ids []int32) error
