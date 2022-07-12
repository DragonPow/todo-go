package usecase

import (
	"context"
	"project1/domain"
	"project1/util/db"
)

type tagUsecase struct {
	db      db.Database
	tagRepo domain.TagRepository
}

func NewTagUsecase(db db.Database, t domain.TagRepository) domain.TagUsecase {
	return &tagUsecase{
		db:      db,
		tagRepo: t,
	}
}

func (t *tagUsecase) FetchAll(ctx context.Context) ([]domain.Tag, error) {
	tags, err := t.tagRepo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *tagUsecase) GetByID(ctx context.Context, id int32) (domain.Tag, error) {
	tag, err := t.tagRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}

func (t *tagUsecase) Create(ctx context.Context, args ...interface{}) (domain.Tag, error) {
	tag, err := t.tagRepo.Create(ctx, args...)
	if err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}

func (t *tagUsecase) Delete(ctx context.Context, id int32) error {
	isSuccess := false
	tx := t.db.Db.Begin()

	defer func() {
		if isSuccess {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	// Check is exists
	if _, err := t.tagRepo.GetByID(ctx, id, tx); err != nil {
		return err
	}

	if err := t.tagRepo.Delete(ctx, id, tx); err != nil {
		return err
	}

	isSuccess = true
	return nil
}
