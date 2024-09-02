package postgres

import (
	"context"
	"database/sql"
	"log"

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

func (hr *VisitorRepository) GetByID(ctx context.Context, id int) (res domain.Visitor, err error) {
	query := `SELECT id, mail FROM visitors WHERE id=$1`
	return hr.getOne(ctx, query, id)
}

func (vr *VisitorRepository) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Visitor, err error) {
	stmt, err := vr.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
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
