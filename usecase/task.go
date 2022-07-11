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
}

func NewTaskUsecase(db db.Database, t domain.TaskRepository, u domain.UserRepository) domain.TaskUsecase {
	return &taskUsecase{
		db:       db,
		taskRepo: t,
		userRepo: u,
	}
}

func (t *taskUsecase) Fetch(ctx context.Context, user_id int32, start_index int32, number int32, conditions ...interface{}) ([]domain.Task, error) {
	return nil, fmt.Errorf("Implemeent needed")
}

func (t *taskUsecase) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Task, error) {
	return domain.Task{}, fmt.Errorf("Implemeent needed")
}

func (t *taskUsecase) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implemeent needed")
}

func (t *taskUsecase) Delete(ctx context.Context, ids []int32) error {
	return fmt.Errorf("Implemeent needed")
}

func (t *taskUsecase) DeleteAll(ctx context.Context, creator_id int32) error {
	return fmt.Errorf("Implemeent needed")
}
