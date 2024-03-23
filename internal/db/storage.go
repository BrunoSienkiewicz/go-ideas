package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetIdea(id int) (*Idea, error)
	AddIdea(idea *Idea) error
	UpdateIdea(idea *Idea) error
	DeleteIdea(id int) error
	ListIdeas() ([]*Idea, error)
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

func (s *PostgresStorage) GetIdea(id int) (*Idea, error) {
	return nil, nil
}

func (s *PostgresStorage) AddIdea(idea *Idea) error {
	return nil
}

func (s *PostgresStorage) UpdateIdea(idea *Idea) error {
	return nil
}

func (s *PostgresStorage) DeleteIdea(id int) error {
	return nil
}

func (s *PostgresStorage) ListIdeas() ([]*Idea, error) {
	return nil, nil
}
