package repository

import (
	"context"
	"fmt"
	"project1/domain"
	"project1/util/db"
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

func (u *userRepository) Create(ctx context.Context, args ...interface{}) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Implemeent needed")
}

func (u *userRepository) Update(ctx context.Context, id int32, args ...interface{}) error {
	return fmt.Errorf("Implemeent needed")
}

func (u *userRepository) Delete(ctx context.Context, ids []int32) error {
	return fmt.Errorf("Implemeent needed")
}
