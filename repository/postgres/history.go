package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type HistoryRepository struct {
	Conn *sql.DB
}

func NewHistoryRepository(conn *sql.DB) *HistoryRepository {
	return &HistoryRepository{conn}
}

func (hr *HistoryRepository) Fetch(ctx context.Context, num int) (res []domain.History, err error) {
	rows, err := hr.Conn.QueryContext(ctx, `SELECT * FROM histories LIMIT $1`, num)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	res = make([]domain.History, 0)
	for rows.Next() {
		h := domain.History{}
		visitorID := 0
		err = rows.Scan(
			&h.ID,
			&visitorID,
			&h.VisitedFrom,
			&h.VisitedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		h.Visitor = domain.Visitor{ID: visitorID}
		res = append(res, h)
	}
	return res, nil
}
