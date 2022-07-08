package domain

import (
	// "gorm.io/gorm"
	"time"
	"context"
)

type Task struct {
	ID int32 `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"column:name;not null"`
	Description string `gorm:"column:description"`
	IsDone bool `gorm:"column:is_done;default:false"`
	DoneAt time.Time `gorm:"column:done_at"`
	CreatedAt time.Time `gorm:"column:created_at;default:time.Now().UTC()"`
	UserCreate User `gorm:"constrain:OnUpdate:NO ACTION,OnDelete:CASCADE"`
	Tags []Tag `gorm:"constrain:OnUpdate:NO ACTION,OnDelete:SET NULL"`
}

type TaskRepository interface {
	Fetch(ctx context.Context, user_id int32, start_index int32, number int32) ([]Task, error)
	FetchWithConditions(ctx context.Context, user_id int32, conditions ...interface{}) ([]Task, error)
	GetByID(ctx context.Context, id int32) (Task, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Task, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, ids []int32) error
}