package storage

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Payment struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Store struct{ DB *sql.DB }

func Open(dsn string) (*Store, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS payments (id TEXT PRIMARY KEY, user_id TEXT NOT NULL, amount DOUBLE PRECISION NOT NULL)`)
	if err != nil {
		return nil, err
	}
	return &Store{DB: db}, nil
}
func (s *Store) Close() error { return s.DB.Close() }
func (s *Store) Create(ctx context.Context, p Payment) error {
	_, err := s.DB.ExecContext(ctx, `INSERT INTO payments(id,user_id,amount) VALUES($1,$2,$3)`, p.ID, p.UserID, p.Amount)
	return err
}
