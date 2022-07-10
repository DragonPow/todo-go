package repository

import (
	"context"
	"fmt"

	"project1/domain"
	"project1/util/db"
)

type taskRepository struct {
	Conn db.Database
}

func NewTaskRepository(conn db.Database) domain.TaskRepository {
	return &taskRepository{
		Conn: conn,
	}
}

func (t *taskRepository) Fetch(ctx context.Context, user_id int32, start_index int32, number int32, conditions ...interface{}) ([]domain.Task, error) {
	return nil, fmt.Errorf("Implement needed")
}

func (t *taskRepository) GetByID(ctx context.Context, id int32) (domain.Task, error) {
	return domain.Task{}, fmt.Errorf("Implement needed")
}

func (t *taskRepository) GetByUserId(ctx context.Context, creator_id int32) ([]int32, error) {
	return nil, fmt.Errorf("Implement needed")
}

func (t *taskRepository) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Task, error) {
	return domain.Task{}, fmt.Errorf("Implement needed")
}

func (t *taskRepository) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implement needed")
}

func (t *taskRepository) Delete(ctx context.Context, ids []int32) error {
	return fmt.Errorf("Implement needed")
}
