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

func (vr *VisitorRepository) GetByID(ctx context.Context, id int) (res domain.Visitor, err error) {
	query := `SELECT id, mail FROM visitors WHERE id=$1`
	return vr.getOne(ctx, query, id)
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

func (vr *VisitorRepository) Store(ctx context.Context, visitor *domain.Visitor) (err error) {
	query := `INSERT INTO visitors (mail) VALUES ($1)`
	stmt, err := vr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, visitor.Mail)
	if err != nil {
		return
	}
	lastID, _ := res.LastInsertId()
	visitor.ID = int(lastID)
	return
}

func (vr *VisitorRepository) GetByMail(ctx context.Context, mail string) (res domain.Visitor, err error) {
	query := `SELECT id, mail FROM visitors WHERE mail=$1`
	return vr.getOne(ctx, query, mail)
}
