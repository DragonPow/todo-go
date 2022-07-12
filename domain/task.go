package domain

import (
	// "gorm.io/gorm"
	"context"
	"time"
)

type Task struct {
	ID          int32     `form:"-" gorm:"primaryKey;autoIncrement"`
	Name        string    `form:"name" gorm:"column:name;not null"`
	Description string    `form:"description" gorm:"column:description"`
	IsDone      bool      `form:"is_done" gorm:"column:is_done;default:false"`
	DoneAt      time.Time `form:"-" gorm:"column:done_at"`
	CreatedAt   time.Time `form:"-" gorm:"column:created_at"`

	CreatorId   int32 `form:"-"`
	UserCreator User  `form:"-" gorm:"foreignKey:CreatorId;constrain:OnUpdate:NO ACTION,OnDelete:CASCADE"`

	Tags []Tag `form:"tags" gorm:"many2many:task_tags;constrain:OnUpdate:NO ACTION,OnDelete:SET NULL"`
}

type TaskRepository interface {
	Fetch(ctx context.Context, user_id int32, start_index int32, number int32, args ...interface{}) ([]Task, error)
	GetByID(ctx context.Context, id int32, args ...interface{}) (Task, error)
	GetByUserId(ctx context.Context, creator_id int32, args ...interface{}) ([]int32, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Task, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, id int32, args ...interface{}) error
}

type TaskUsecase interface {
	Fetch(ctx context.Context, user_id int32, start_index int32, number int32, conditions ...interface{}) ([]Task, error)
	GetByID(ctx context.Context, id int32) (Task, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Task, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, ids int32) error
	DeleteAll(ctx context.Context, creator_id int32) error
}
