package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"project1/domain"
	"project1/util/db"
)

type userRepository struct {
	postgre_repository
}

func NewUserRepository(conn db.Database) domain.UserRepository {
	return &userRepository{
		postgre_repository: *newRepository(conn),
	}
}

func (u *userRepository) GetByUsernameAndPassword(ctx context.Context, username string, password string, args ...interface{}) (domain.User, error) {
	tx, err := u.GetTransaction(args, 0)
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	if err := tx.Where("username=? AND password=?", username, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotExists
		} else {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int32, args ...interface{}) (domain.User, error) {
	tx, err := u.GetTransaction(args, 0)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{ID: id}
	if err := tx.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotExists
		} else {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (u *userRepository) GetByUsername(ctx context.Context, username string, args ...interface{}) (domain.User, error) {
	tx, err := u.GetTransaction(args, 0)
	if err != nil {
		return domain.User{}, err
	}

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
	tx, err := u.GetTransaction(args, 1)
	if err != nil {
		return domain.User{}, err
	}

	new_user := args[1].(domain.User)
	if err := tx.Create(&new_user).Error; err != nil {
		return domain.User{}, err
	}

	return new_user, nil
}

func (u *userRepository) Update(ctx context.Context, id int32, args ...interface{}) error {
	tx, err := u.GetTransaction(args, 1)
	if err != nil {
		return err
	}

	new_user := args[1].(map[string]interface{})

	if err := tx.First(&domain.User{ID: id}).Updates(new_user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(ctx context.Context, id int32, args ...interface{}) error {
	tx, err := u.GetTransaction(args, 0)
	if err != nil {
		return err
	}

	if err := tx.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
