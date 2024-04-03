package ideastorage

import (
	_ "github.com/lib/pq"

	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
)

type IdeaStorage struct {
	store            *storage.PostgresStorage
	attributeStorage *storage.AttributeStorage
}

func NewIdeaStorage(postgres *storage.PostgresStorage, attAttributeStorage *storage.AttributeStorage) *IdeaStorage {
	return &IdeaStorage{
		store:            postgres,
		attributeStorage: attAttributeStorage,
	}
}

func (s *IdeaStorage) handleCategory(category string) (*int, error) {
	query := `SELECT * FROM ideas.categories WHERE name = $1`
	rows, err := s.store.Query(query, category)
	if err != nil {
		return nil, err
	}

	category_id := 0
	if rows.Next() {
		rows.Scan(&category_id, &category)
	} else {
		query = `INSERT INTO ideas.categories (name) VALUES ($1) RETURNING category_id`
		rows, err := s.store.Query(query, category)
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			rows.Scan(&category_id, &category)
		}
	}

	return &category_id, nil
}
