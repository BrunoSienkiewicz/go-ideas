package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type IdeaStorage struct {
	store *PostgresStorage
}

func NewIdeaStorage() (*IdeaStorage, error) {
	store, err := NewPostgresStorage()
	if err != nil {
		return nil, err
	}

	return &IdeaStorage{store: store}, nil
}

func (s *IdeaStorage) GetObject(id int) (*Idea, error) {
	return nil, nil
}

func (s *IdeaStorage) AddObject(obj *Idea) error {
	return nil
}

func (s *IdeaStorage) UpdateObject(obj *Idea) error {
	return nil
}

func (s *IdeaStorage) DeleteObject(id int) error {
	return nil
}

func (s *IdeaStorage) ListObjects() ([]*Idea, error) {
	return nil, nil
}
