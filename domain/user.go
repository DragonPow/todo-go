package domain

import (
	// "gorm.io/gorm"
	"time"
	"context"
)

type User struct {
	ID int32 `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"column:usernames;not null;unique"`
	Password string `gorm:"column:password;not null"`
	Name string `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;default:time.Now().UTC()"`
}

type UserRepository interface {
	GetByUsernameAndPassword(ctx context.Context, username string, password string) (User, error)
	GetByID(ctx context.Context, id int32) (User, error)
	Create(ctx context.Context, creator_id int32, args ...interface{}) (User, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, ids []int32) error
}