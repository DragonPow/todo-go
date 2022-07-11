package usecase

import (
	"context"
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
	if isExists, err := u.userRepo.CheckExists(ctx, user_info.Username, tx); err != nil {
		return domain.User{}, err
	} else if isExists {
		return domain.User{}, domain.NewDomainError(domain.UsernameIsExists)
	}

	new_user, new_user_err := u.userRepo.Create(ctx, append([]interface{}{tx}, args...)...)
	if new_user_err != nil {
		return domain.User{}, new_user_err
	}

	isSuccess = true
	return new_user, new_user_err
}

func (u *userUsecase) ChangePassword(ctx context.Context, new_password string) error {
	return fmt.Errorf("Implement needed")
}

func (u *userUsecase) UpdateName(ctx context.Context, new_name string) error {
	return fmt.Errorf("Implement needed")
}
