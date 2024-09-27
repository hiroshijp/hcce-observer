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

func (ur *UserRepository) FetchByName(ctx context.Context, name string) (res domain.User, err error) {
	query := `SELECT name, password FROM users WHERE name=$1`
	return ur.getOne(ctx, query, name)
}

func (ur *UserRepository) getOne(ctx context.Context, query string, args ...interface{}) (res domain.User, err error) {
	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.User{}
	err = row.Scan(
		&res.Name,
		&res.Password,
	)
	return
}

func (ur *UserRepository) Store(ctx context.Context, user *domain.User) (err error) {
	query := `INSERT INTO users (name, password, is_admin) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	err = stmt.QueryRow(user.Name, user.Password, user.IsAdmin).Scan(&user.ID)
	if err != nil {
		return err
	}
	return
}
