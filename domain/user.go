package domain

import (
	// "gorm.io/gorm"
	"context"
	"time"
)

type User struct {
	ID        int32     `form:"-" json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `form:"username" json:"username" binding:"required" gorm:"column:username;not null;unique"`
	Password  string    `form:"password" json:"password" binding:"required" gorm:"column:password;not null"`
	Name      string    `form:"name" json:"name" binding:"required" gorm:"column:name;not null"`
	CreatedAt time.Time `form:"-" json:"created_at" gorm:"column:created_at;"`
}

type UserRepository interface {
	GetByUsernameAndPassword(ctx context.Context, username string, password string, args ...interface{}) (User, error)
	GetByID(ctx context.Context, id int32, args ...interface{}) (User, error)
	Create(ctx context.Context, args ...interface{}) (User, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, id int32, args ...interface{}) error
	GetByUsername(ctx context.Context, username string, args ...interface{}) (User, error)
}

type UserUsecase interface {
	Login(ctx context.Context, username string, password string) (User, error)
	Create(ctx context.Context, args ...interface{}) (User, error)
	Update(ctx context.Context, id int32, args ...interface{}) error
	Delete(ctx context.Context, id int32) error
	GetByID(ctx context.Context, id int32) (User, error)
}
