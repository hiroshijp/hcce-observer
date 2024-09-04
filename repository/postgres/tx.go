package postgres

import (
	"context"
	"database/sql"
)

type TxRepository struct {
	Conn *sql.DB
}

func NewTxRepository(conn *sql.DB) *TxRepository {
	return &TxRepository{conn}
}

func (tr *TxRepository) BeginTx(ctx context.Context, f func(ctx context.Context) error) (err error) {
	tx, err := tr.Conn.Begin()
	if err != nil {
		return
	}

	err = f(ctx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	return nil
}
