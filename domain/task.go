package domain

import (
	// "gorm.io/gorm"
	"context"
	"time"
)

type Task struct {
	ID          int32     `json:"id" form:"-" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" form:"name" gorm:"column:name;not null"`
	Description string    `json:"description" form:"description" gorm:"column:description"`
	IsDone      bool      `json:"is_done" form:"is_done" gorm:"column:is_done;default:false"`
	DoneAt      time.Time `json:"done_at" form:"-" gorm:"column:done_at"`
	CreatedAt   time.Time `json:"created_at" form:"-" gorm:"column:created_at"`

	CreatorId   int32 `form:"-"`
	UserCreator User  `json:"user_creator" form:"-" gorm:"foreignKey:CreatorId;constrain:OnUpdate:NO ACTION,OnDelete:CASCADE"`

	Tags []Tag `json:"tags" form:"tags" gorm:"many2many:task_tags"`
}

type TaskRepository interface {
	Fetch(ctx context.Context, user_id int32, start_index int32, number int32, args ...interface{}) ([]Task, error)
	GetByID(ctx context.Context, id int32, args ...interface{}) (Task, error)
	CheckExists(ctx context.Context, ids []int32, args ...interface{}) error
	GetByUserId(ctx context.Context, creator_id int32, args ...interface{}) ([]int32, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Task, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, ids []int32, args ...interface{}) error
}

type TaskUsecase interface {
	Fetch(ctx context.Context, user_id int32, start_index int32, number int32, conditions ...interface{}) ([]Task, error)
	GetByID(ctx context.Context, id int32) (Task, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Task, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, ids []int32) error
	DeleteAll(ctx context.Context, creator_id int32) error
}
