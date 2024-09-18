package usecase

import (
	"context"
	"errors"
	"os"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type UserRepository interface {
	Fetch(ctx context.Context) (res []domain.User, err error)
	Store(ctx context.Context, user *domain.User) (err error)
}

type UserUsecase struct {
	userRepo UserRepository
}

func NewUserUsecase(ur UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uu *UserUsecase) Fetch(ctx context.Context) (res []domain.User, err error) {
	return
}

func (uu *UserUsecase) Store(ctx context.Context, user *domain.User) (err error) {
	return
}

func (uu *UserUsecase) Signin(ctx context.Context, name string, password string) (err error) {
	if os.Getenv("ADMIN_NAME") != name || os.Getenv("ADMIN_PASS") != password {
		err = errors.New("user not found")
		return err
	}

	return nil
}
