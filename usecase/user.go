package usecase

import (
	"context"
	"errors"
	"fmt"

	"project1/domain"
	"project1/util/db"
)

type userUsecase struct {
	db       db.Database
	userRepo domain.UserRepository
}

func NewUserUsecase(db db.Database, u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		db:       db,
		userRepo: u,
	}
}

func (u *userUsecase) Login(ctx context.Context, username string, password string) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Implement needed")
}

func (u *userUsecase) Create(ctx context.Context, args ...interface{}) (domain.User, error) {
	if len(args) == 0 {
		return domain.User{}, fmt.Errorf("Args is needed")
	}

	isSuccess := false
	tx := u.db.Db.Begin()

	defer func() {
		if isSuccess {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	user_info := args[0].(domain.User)
	if _, err := u.userRepo.GetByUsername(ctx, user_info.Username, tx); err != nil {
		// If username not found, add it
		if errors.Is(err, domain.ErrUserNotExists) {
			new_user, new_user_err := u.userRepo.Create(ctx, append([]interface{}{tx}, args...)...)
			if new_user_err != nil {
				return domain.User{}, new_user_err
			}

			isSuccess = true
			return new_user, new_user_err
		} else {
			// If another error is occured, return error
			return domain.User{}, err
		}
	} else {
		return domain.User{}, domain.ErrUserIsExists
	}
}

func (u *userUsecase) ChangePassword(ctx context.Context, new_password string) error {
	return fmt.Errorf("Implement needed")
}

func (u *userUsecase) UpdateName(ctx context.Context, new_name string) error {
	return fmt.Errorf("Implement needed")
}

func (u *userUsecase) Delete(ctx context.Context, id int32) error {
	return nil
}
