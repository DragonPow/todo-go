package usecase

import (
	"context"
	"fmt"
	"project1/domain"
	"project1/util/db"
)

type tagUsecase struct {
	db      db.Database
	tagRepo domain.TagRepository
}

func NewTagUsecase(db db.Database, t domain.TagRepository) domain.TagRepository {
	return &tagUsecase{
		db:      db,
		tagRepo: t,
	}
}

func (t *tagUsecase) FetchAll(ctx context.Context) ([]domain.Tag, error) {
	return nil, fmt.Errorf("Implemeent needed")
}
func (t *tagUsecase) GetByID(ctx context.Context, id int32) (domain.Tag, error) {
	return domain.Tag{}, fmt.Errorf("Implemeent needed")
}
func (t *tagUsecase) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Tag, error) {
	return domain.Tag{}, fmt.Errorf("Implemeent needed")
}
func (t *tagUsecase) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implemeent needed")
}
func (t *tagUsecase) Delete(ctx context.Context, ids []int32) error {
	return fmt.Errorf("Implemeent needed")
}
