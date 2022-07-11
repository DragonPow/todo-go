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

func (u *userUsecase) ChangePassword(ctx context.Context, id int32, new_password string) error {
	return fmt.Errorf("Implement needed")
}

func (u *userUsecase) Update(ctx context.Context, id int32, args ...interface{}) error {
	if len(args) != 1 {
		return fmt.Errorf("Args is needed")
	}

	value := args[0].(map[string]interface{})
	isSuccess := false
	tx := u.db.Db.Begin()

	defer func() {
		if isSuccess {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	// Check user is exists
	if _, err := u.userRepo.GetByID(ctx, id, tx); err != nil {
		return err
	}

	isSuccess = true
	return u.userRepo.Update(ctx, id, []interface{}{tx, value}...)
}

func (u *userUsecase) Delete(ctx context.Context, id int32) error {
	if err := u.userRepo.Delete(ctx, id); err != nil {
		return err
	} else {
		return nil
	}
}
