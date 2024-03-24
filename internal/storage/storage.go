package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=ideas_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
