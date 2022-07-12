package usecase

import (
	"context"
	"fmt"

	"project1/domain"
	"project1/util/db"
)

type taskUsecase struct {
	db       db.Database
	taskRepo domain.TaskRepository
	userRepo domain.UserRepository
	tagRepo  domain.TagRepository
}

func NewTaskUsecase(db db.Database, t domain.TaskRepository, u domain.UserRepository, tag domain.TagRepository) domain.TaskUsecase {
	return &taskUsecase{
		db:       db,
		taskRepo: t,
		userRepo: u,
		tagRepo:  tag,
	}
}

func (t *taskUsecase) Fetch(ctx context.Context, user_id int32, start_index int32, number int32, conditions ...interface{}) ([]domain.Task, error) {
	return nil, fmt.Errorf("Implemeent needed")
}

func (t *taskUsecase) GetByID(ctx context.Context, id int32) (domain.Task, error) {
	new_task, err := t.taskRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}

	return new_task, nil
}

func (t *taskUsecase) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Task, error) {
	isSuccess := false
	tx := t.db.Db.Begin()

	defer func() {
		if isSuccess {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	_, err := t.userRepo.GetByID(ctx, creator_id, tx)
	if err != nil {
		return domain.Task{}, err
	}

	_, err = t.tagRepo.FetchAll(ctx, tx)
	if err != nil {
		return domain.Task{}, err
	}

	task, err := t.taskRepo.Create(ctx, creator_id, append([]interface{}{tx}, args...)...)
	if err != nil {
		return domain.Task{}, err
	}

	// task.UserCreator = user
	// for i, tag := range task.Tags {
	// 	idx := slices.IndexFunc(tags, func(t domain.Tag) bool { return t.ID == tag.ID })
	// 	task.Tags[i] = tags[idx]
	// }

	isSuccess = true
	return task, nil
}

func (t *taskUsecase) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implemeent needed")
}

func (t *taskUsecase) Delete(ctx context.Context, ids []int32) error {
	isSuccess := false
	tx := t.db.Db.Begin()

	defer func() {
		if isSuccess {
			tx.Commit()
		}
		tx.Rollback()
	}()

	// Check task exists
	if err := t.taskRepo.CheckExists(ctx, ids, tx); err != nil {
		return err
	}

	// Delete
	if err := t.taskRepo.Delete(ctx, ids, tx); err != nil {
		return err
	}

	isSuccess = true
	return nil
}

func (t *taskUsecase) DeleteAll(ctx context.Context, creator_id int32) error {
	isSuccess := false
	tx := t.db.Db.Begin()

	defer func() {
		if isSuccess {
			tx.Commit()
		}
		tx.Rollback()
	}()

	// Find id of user
	tasks_id, err := t.taskRepo.GetByUserId(ctx, creator_id, tx)
	if err != nil {
		return err
	}

	// Delete
	if err := t.taskRepo.Delete(ctx, tasks_id, tx); err != nil {
		return err
	}

	isSuccess = true
	return nil
}
