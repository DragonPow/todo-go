package usecase

import (
	"context"
	"fmt"
	"project1/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Login(ctx context.Context, username string, password string) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Implemeent needed")
}

func (u *userUsecase) ChangePassword(ctx context.Context, new_password string) error {
	return fmt.Errorf("Implemeent needed")
}

func (u *userUsecase) UpdateName(ctx context.Context, new_name string) error {
	return fmt.Errorf("Implemeent needed")
}
