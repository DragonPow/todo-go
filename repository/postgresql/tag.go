package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"

	"project1/domain"
	"project1/util/db"
)

type tagRepository struct {
	postgre_repository
}

func NewTagRepository(conn db.Database) domain.TagRepository {
	return &tagRepository{
		postgre_repository: *newRepository(conn),
	}
}

func (t *tagRepository) FetchAll(ctx context.Context, args ...interface{}) ([]domain.Tag, error) {
	tx, err := t.GetTransaction(args, 0)
	if err != nil {
		return nil, err
	}

	var tags []domain.Tag
	if err := tx.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *tagRepository) GetByID(ctx context.Context, id int32, args ...interface{}) (domain.Tag, error) {
	tx, err := t.GetTransaction(args, 0)
	if err != nil {
		return domain.Tag{}, err
	}

	var tag domain.Tag
	if err := tx.First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Tag{}, domain.ErrTagNotExists
		}
		return domain.Tag{}, err
	}

	return tag, nil
}

func (t *tagRepository) Create(ctx context.Context, args ...interface{}) (domain.Tag, error) {
	tx, err := t.GetTransaction(args, 1)
	if err != nil {
		return domain.Tag{}, err
	}

	new_tag := args[len(args)-1].(domain.Tag)

	// Update
	if err := tx.Create(&new_tag).Error; err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok && errors.Is(err, pgError) {
			// Duplicate value
			if pgError.Code == "23505" {
				return domain.Tag{}, domain.ErrTagValueDuplicated
			}
		}
		return domain.Tag{}, err
	}

	return new_tag, nil
}

func (t *tagRepository) Update(ctx context.Context, args ...interface{}) (domain.Tag, error) {
	tx, err := t.GetTransaction(args, 1)
	if err != nil {
		return domain.Tag{}, err
	}

	new_tag_info := args[len(args)-1].(map[string]interface{})

	// Update
	var tag domain.Tag
	if err := tx.First(&tag).Updates(new_tag_info).Error; err != nil {
		return domain.Tag{}, err
	}

	return tag, nil
}

func (t *tagRepository) Delete(ctx context.Context, id int32, args ...interface{}) error {
	tx, err := t.GetTransaction(args, 0)
	if err != nil {
		return err
	}

	tag := domain.Tag{ID: id}
	if err := tx.Delete(&tag).Error; err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok && errors.Is(err, pgError) {
			// Tag still another reference
			if pgError.Code == "23503" {
				return domain.ErrTagStillReference
			}
		}
		return err
	}

	return nil
}
