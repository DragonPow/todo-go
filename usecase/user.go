package usecase

import (
	"context"
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

func (u *userUsecase) Login(ctx context.Context, username string, password string) (domain.User, error)

func (u *userUsecase) ChangePassword(ctx context.Context, new_password string) error

func (u *userUsecase) UpdateName(ctx context.Context, new_name string) error
