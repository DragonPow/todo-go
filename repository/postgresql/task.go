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
	Conn db.Database
}

func NewTaskRepository(conn db.Database) domain.TaskRepository {
	return &taskRepository{
		Conn: conn,
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
	return nil, fmt.Errorf("Implement needed")
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

func (t *taskRepository) GetTransaction(args []interface{}, minNumberLen int) (*gorm.DB, error) {
	var tx *gorm.DB
	if len(args) <= minNumberLen {
		tx = t.Conn.Db
	} else {
		arg0, ok := args[0].(*gorm.DB)
		if ok {
			tx = arg0
		} else {
			return nil, fmt.Errorf("Args tx is needed")
		}
	}
	return tx, nil
}

func (t *taskRepository) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implement needed")
}

func (t *taskRepository) Delete(ctx context.Context, ids int32, args ...interface{}) error {
	return fmt.Errorf("Implement needed")
}
