package domain

import (
	"context"
)

type Tag struct {
	ID          int32  `gorm:"primaryKey;autoIncrement"`
	Value       string `gorm:"column:value;not null;unique"`
	Description string `gorm:"column:description"`
}

type TagRepository interface {
	FetchAll(ctx context.Context) ([]Tag, error)
	GetByID(ctx context.Context, id int32) (Tag, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Tag, error)
	Delete(ctx context.Context, ids []int32) error
}

type TagUsecase interface {
	FetchAll(ctx context.Context) ([]Tag, error)
	GetByID(ctx context.Context, id int32) (Tag, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (Tag, error)
	Delete(ctx context.Context, ids []int32) error
}
