package repository

import (
	"context"
	"fmt"
	"project1/domain"
	"project1/util/db"
)

type tagRepository struct {
	Conn db.Database
}

func NewTagRepository(conn db.Database) domain.TagRepository {
	return &tagRepository{Conn: conn}
}

func (t *tagRepository) FetchAll(ctx context.Context) ([]domain.Tag, error) {
	return nil, fmt.Errorf("Implemeent needed")
}
func (t *tagRepository) GetByID(ctx context.Context, id int32) (domain.Tag, error) {
	return domain.Tag{}, fmt.Errorf("Implemeent needed")
}
func (t *tagRepository) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Tag, error) {
	return domain.Tag{}, fmt.Errorf("Implemeent needed")
}
func (t *tagRepository) Delete(ctx context.Context, ids []int32) error {
	return fmt.Errorf("Implemeent needed")
}
