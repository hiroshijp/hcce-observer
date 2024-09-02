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
	return &VisitorRepository{conn}
}

func (hr *VisitorRepository) GetByID(ctx context.Context, id int) (res domain.Visitor, err error) {
	row, err := hr.Conn.QueryContext(ctx, `SELECT * FROM visitors WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
		return domain.Visitor{}, err
	}

	defer func() {
		err = row.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	res = domain.Visitor{}
	err = row.Scan(
		&res.ID,
		&res.Mail,
	)

	return res, nil
}
