package repository

import (
	"context"
	"errors"
	"fmt"
	"project1/domain"
	"project1/util/db"

	"gorm.io/gorm"
)

type userRepository struct {
	Conn db.Database
}

func NewUserRepository(conn db.Database) domain.UserRepository {
	return &userRepository{Conn: conn}
}

func (u *userRepository) GetByUsernameAndPassword(ctx context.Context, username string, password string) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Implemeent needed")
}

func (u *userRepository) GetByID(ctx context.Context, id int32) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Implemeent needed")
}

func (u *userRepository) GetByUsername(ctx context.Context, username string, args ...interface{}) (domain.User, error) {
	if len(args) != 1 {
		return domain.User{}, fmt.Errorf("Args is required")
	}

	tx := args[0].(*gorm.DB)
	var user domain.User

	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotExists
		} else {
			return domain.User{}, err
		}
	}
	return user, nil
}

func (u *userRepository) Create(ctx context.Context, args ...interface{}) (domain.User, error) {
	if len(args) != 2 {
		return domain.User{}, fmt.Errorf("Args is required")
	}

	tx := args[0].(*gorm.DB)
	new_user := args[1].(domain.User)

	if err := tx.Create(&new_user).Error; err != nil {
		return domain.User{}, err
	}

	return new_user, nil
}

func (u *userRepository) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implemeent needed")
}

func (u *userRepository) Delete(ctx context.Context, id int32) error {
	return fmt.Errorf("Implemeent needed")
}
