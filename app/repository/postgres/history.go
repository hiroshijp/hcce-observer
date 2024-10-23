package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/hiroshijp/try-clean-arch/domain"
)

type HistoryRepository struct {
	Conn *sql.DB
}

func NewHistoryRepository(conn *sql.DB) *HistoryRepository {
	return &HistoryRepository{conn}
}

func (hr *HistoryRepository) Fetch(ctx context.Context, num int) (res []domain.History, err error) {
	rows, err := hr.Conn.QueryContext(ctx, `SELECT * FROM histories ORDER BY visited_at DESC LIMIT $1`, num)
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

func (hr *HistoryRepository) Store(ctx context.Context, history *domain.History) (err error) {
	query := `INSERT INTO histories (visitor_id, visited_from, visited_at) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := hr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	history.VisitedAt = time.Now()
	err = stmt.QueryRow(history.Visitor.ID, history.VisitedFrom, history.VisitedAt).Scan(&history.ID)
	if err != nil {
		return err
	}
	return
}
