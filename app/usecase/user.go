package usecase

import (
	"context"
	"errors"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type UserRepository interface {
	FetchByName(ctx context.Context, name string) (res domain.User, err error)
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
	user, err := uu.userRepo.FetchByName(ctx, name)
	if err != nil {
		return err
	}

	if user.Password != password {
		err = errors.New("password is not correct")
		return err
	}
	return nil
}
