package usecase

import (
	"context"
	"errors"
	"os"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type UserRepository interface {
	FetchByName(ctx context.Context) (res domain.User, err error)
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
	user, err := uu.userRepo.FetchByName(ctx)
	if user.name != name || user.password != password {
		err = errors.New("user not found")
		return 
	}
	return nil
}
