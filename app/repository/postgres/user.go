package postgres

import (
	"context"
	"database/sql"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn}
}

func (ur *UserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	return nil, nil
}

func (ur *UserRepository) Store(ctx context.Context, user *domain.User) (err error) {
	return nil
}
