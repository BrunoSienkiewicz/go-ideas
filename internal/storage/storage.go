package storage

import (
	"database/sql"
	_ "github.com/lib/pq"

	config "github.com/BrunoSienkiewicz/go_ideas/config"
)

type Storage[T any] interface {
	GetObject(id int) (*T, error)
	AddObject(obj *T) error
	UpdateObject(obj *T) error
	DeleteObject(id int) error
	ListObjects() ([]*T, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	cfg := config.NewConfig()
	db, err := sql.Open("postgres", cfg.GetDbConnectionString())

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
