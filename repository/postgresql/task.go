package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"project1/domain"
	"project1/util/db"
)

type taskRepository struct {
	postgre_repository
}

func NewTaskRepository(conn db.Database) domain.TaskRepository {
	return &taskRepository{
		postgre_repository: *newRepository(conn),
	}
}

func (t *taskRepository) Fetch(ctx context.Context, user_id int32, start_index int32, number int32, args ...interface{}) ([]domain.Task, error) {
	return nil, fmt.Errorf("Implement needed")
}

func (t *taskRepository) GetByID(ctx context.Context, id int32, args ...interface{}) (domain.Task, error) {
	tx, err := t.GetTransaction(args, 1)
	if err != nil {
		return domain.Task{}, err
	}

	var task domain.Task
	if err := tx.Preload("UserCreator").Preload("Tags").First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Task{}, domain.ErrTaskNotExists
		}

		return domain.Task{}, err
	}

	return task, nil
}

func (t *taskRepository) GetByUserId(ctx context.Context, creator_id int32, args ...interface{}) ([]int32, error) {
	tx, err := t.GetTransaction(args, 0)
	if err != nil {
		return nil, err
	}

	// Get all task by user id
	var tasks []domain.Task
	if err := tx.Where("cretor_id IN ?", creator_id).Find(&tasks).Error; err != nil {
		return nil, err
	}

	// Map to task model to int
	tasks_ids := []int32{}
	for _, tasks := range tasks {
		tasks_ids = append(tasks_ids, tasks.ID)
	}

	return tasks_ids, nil
}

func (t *taskRepository) Create(ctx context.Context, creator_id int32, args ...interface{}) (domain.Task, error) {
	tx, err := t.GetTransaction(args, 1)
	if err != nil {
		return domain.Task{}, err
	}

	new_task := args[len(args)-1].(domain.Task)
	if err := tx.Omit("Tags").Create(&new_task).Error; err != nil {
		return domain.Task{}, err
	}

	if err := tx.Model(&new_task).Association("Tags").Append(&new_task.Tags); err != nil {
		return domain.Task{}, err
	}

	// if err := tx.Model(&new_task).Association("UserCreator").Replace(&new_task.UserCreator); err != nil {
	// 	return domain.Task{}, err
	// }

	return new_task, nil
}

func (t *taskRepository) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implement needed")
}

func (t *taskRepository) Delete(ctx context.Context, ids []int32, args ...interface{}) error {
	tx, err := t.GetTransaction(args, 0)
	if err != nil {
		return err
	}

	if err := tx.Delete(&domain.Task{}, ids).Error; err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) CheckExists(ctx context.Context, ids []int32, args ...interface{}) error {
	tx, err := t.GetTransaction(args, 0)
	if err != nil {
		return err
	}

	var tasks []domain.Task
	if err := tx.Where("Id IN ?", ids).Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) != len(ids) {
		return domain.ErrTagNotExists
	}

	return nil
}
