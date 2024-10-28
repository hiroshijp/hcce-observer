package postgres

import (
	"context"
	"database/sql"

	"github.com/hiroshijp/hcce-observer/domain"
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
	query := `INSERT INTO users (name, password, admin) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Name, user.Password, user.Admin).Scan(&user.ID)
	if err != nil {
		return err
	}
	return
}

func (ur *UserRepository) Delete(ctx context.Context, name string) (err error) {
	query := `DELETE FROM users WHERE name=$1`
	stmt, err := ur.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, name)
	if err != nil {
		return err
	}
	return
}
