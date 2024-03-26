package storage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
	_ "github.com/lib/pq"
)

type IdeaStorage struct {
	store *PostgresStorage
}

func (s *IdeaStorage) GetObject(id int) (*types.Idea, error) {
	return nil, nil
}

func (s *IdeaStorage) AddObject(obj *types.Idea) error {
	return nil
}

func (s *IdeaStorage) UpdateObject(obj *types.Idea) error {
	return nil
}

func (s *IdeaStorage) DeleteObject(id int) error {
	return nil
}

func (s *IdeaStorage) ListObjects() ([]*types.Idea, error) {
	return nil, nil
}
