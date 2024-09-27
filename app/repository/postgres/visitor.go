package postgres

import (
	"context"
	"database/sql"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type VisitorRepository struct {
	Conn *sql.DB
}

func NewVisitorRepository(conn *sql.DB) *VisitorRepository {
	return &VisitorRepository{
		Conn: conn,
	}
}

func (vr *VisitorRepository) GetByID(ctx context.Context, id int) (res domain.Visitor, err error) {
	query := `SELECT id, mail FROM visitors WHERE id=$1`
	return vr.getOne(ctx, query, id)
}

func (vr *VisitorRepository) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Visitor, err error) {
	stmt, err := vr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return domain.Visitor{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Visitor{}
	err = row.Scan(
		&res.ID,
		&res.Mail,
	)
	return
}

func (vr *VisitorRepository) Store(ctx context.Context, visitor *domain.Visitor) (err error) {
	query := `INSERT INTO visitors (mail) VALUES ($1) RETURNING id`
	stmt, err := vr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	err = stmt.QueryRow(visitor.Mail).Scan(&visitor.ID)
	if err != nil {
		return
	}
	return
}

func (vr *VisitorRepository) GetByMail(ctx context.Context, mail string) (res domain.Visitor, err error) {
	query := `SELECT id, mail FROM visitors WHERE mail=$1`
	return vr.getOne(ctx, query, mail)
}
