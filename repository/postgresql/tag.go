package repository

import (
	"context"
	"project1/domain"
	"project1/util/db"
)

type tagRepository struct {
	Conn db.Database
}

func NewTagRepository(conn db.Database) domain.TagRepository {
	return &tagRepository{Conn: conn}
}

func (t *tagRepository) FetchAll(ctx context.Context) ([]domain.Tag, error)
func (t *tagRepository) GetByID(ctx context.Context, id int32) (domain.Tag, error)
func (t *tagRepository) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Tag, error)
func (t *tagRepository) Update(ctx context.Context, id int32, args ...interface{}) error
func (t *tagRepository) Delete(ctx context.Context, ids []int32) error
