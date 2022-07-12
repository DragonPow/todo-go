package domain

import (
	"context"
	"time"
)

type Tag struct {
	ID          int32     `form:"-" json:"id" gorm:"primaryKey;autoIncrement"`
	Value       string    `form:"value" json:"value" gorm:"column:value;not null;unique"`
	Description string    `form:"description" json:"description" gorm:"column:description"`
	CreatedAt   time.Time `form:"-" json:"created_at" gorm:"column:created_at"`
}

type TagRepository interface {
	FetchAll(ctx context.Context, args ...interface{}) ([]Tag, error)
	GetByID(ctx context.Context, id int32, args ...interface{}) (Tag, error)
	Create(ctx context.Context, args ...interface{}) (Tag, error)
	Delete(ctx context.Context, ids int32, args ...interface{}) error
}

type TagUsecase interface {
	FetchAll(ctx context.Context) ([]Tag, error)
	GetByID(ctx context.Context, id int32) (Tag, error)
	Create(ctx context.Context, args ...interface{}) (Tag, error)
	Delete(ctx context.Context, ids int32) error
}
