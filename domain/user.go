package domain

import (
	// "gorm.io/gorm"
	"context"
	"time"
)

type User struct {
	ID        int32     `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"column:usernames;not null;unique"`
	Password  string    `gorm:"column:password;not null"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;"`
}

type UserRepository interface {
	GetByUsernameAndPassword(ctx context.Context, username string, password string) (User, error)
	GetByID(ctx context.Context, id int32) (User, error)
	Create(ctx context.Context, args ...interface{}) (User, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, ids []int32) error
}

type UserUsecase interface {
	Login(ctx context.Context, username string, password string) (User, error)
	ChangePassword(ctx context.Context, new_password string) error
	UpdateName(ctx context.Context, new_name string) error
}
