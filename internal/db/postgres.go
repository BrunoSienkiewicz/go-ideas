package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	config "github.com/BrunoSienkiewicz/go_ideas/config"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (*Postgres, error) {
	cfg := config.NewConfig()
	db, err := sql.Open("postgres", cfg.GetDbConnectionString())

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (s *Postgres) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return rows, nil
}

func (s *Postgres) Close() {
	s.db.Close()
}
