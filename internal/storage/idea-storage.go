package storage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
	_ "github.com/lib/pq"
)

type IdeaStorage struct {
	store *PostgresStorage
}

func NewIdeaStorage(postgres *PostgresStorage) *IdeaStorage {
	return &IdeaStorage{
		store: postgres,
	}
}

func (s *IdeaStorage) GetObject(id int) (*types.Idea, error) {
	query_idea := `SELECT * FROM ideas WHERE id = $1`
	rows, err := s.store.db.Query(query_idea, id)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	idea, err := types.ScanIntoIdea(rows)
	if err != nil {
		return nil, err
	}

	query_attributes := `SELECT * FROM attributes WHERE idea_id = $1`
	rows, err = s.store.db.Query(query_attributes, id)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	for rows.Next() {
		attribute, err := types.ScanIntoAttribute(rows)
		if err != nil {
			return nil, err
		}
		idea.Attributes = append(idea.Attributes, *attribute)
	}

	return idea, nil
}

func (s *IdeaStorage) AddObject(obj *types.Idea) error {
	query_idea := `INSERT INTO ideas.ideas (name, category_id) VALUES ($1, $2) RETURNING idea_id`

	category_id, err := s.handleCategory(obj.Category)
	if err != nil {
		return err
	}

	rows, err := s.store.db.Query(query_idea, obj.Name, *category_id)
	if err != nil {
		return err
	}

	if rows.Err() != nil {
		return rows.Err()
	}
	if rows.Next() {
		rows.Scan(&obj.ID)
	}

	query_attribute := `INSERT INTO ideas.attributes (idea_id, name, value) VALUES ($1, $2, $3)`
	for _, attribute := range obj.Attributes {
		_, err := s.store.db.Query(query_attribute, obj.ID, attribute.Name, attribute.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *IdeaStorage) UpdateObject(obj *types.Idea) error {
	query_idea := `UPDATE ideas SET name = $1, category = $2 WHERE id = $3`
	_, err := s.store.db.Query(query_idea, obj.Name, obj.Category, obj.ID)
	if err != nil {
		return err
	}

	query_attribute := `UPDATE attributes SET name = $1, value = $2 WHERE idea_id = $3`
	for _, attribute := range obj.Attributes {
		_, err := s.store.db.Query(query_attribute, attribute.Name, attribute.Value, obj.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *IdeaStorage) DeleteObject(id int) error {
	query := `DELETE FROM ideas WHERE id = $1`
	_, err := s.store.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *IdeaStorage) ListObjects() ([]*types.Idea, error) {
	query := `SELECT * FROM ideas`
	rows, err := s.store.db.Query(query)
	if err != nil {
		return nil, err
	}

	var ideas []*types.Idea
	for rows.Next() {
		idea, err := types.ScanIntoIdea(rows)
		if err != nil {
			return nil, err
		}

		query_attributes := `SELECT * FROM attributes WHERE idea_id = $1`
		rows, err = s.store.db.Query(query_attributes, idea.ID)
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		for rows.Next() {
			attribute, err := types.ScanIntoAttribute(rows)
			if err != nil {
				return nil, err
			}
			idea.Attributes = append(idea.Attributes, *attribute)
		}

		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (s *IdeaStorage) handleCategory(category string) (*int, error) {
	query := `SELECT * FROM ideas.categories WHERE name = $1`
	rows, err := s.store.db.Query(query, category)
	if err != nil {
		return nil, err
	}

	category_id := 0
	if rows.Next() {
		rows.Scan(&category_id, &category)
	} else {
		query = `INSERT INTO ideas.categories (name) VALUES ($1) RETURNING category_id`
		rows, err := s.store.db.Query(query, category)
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			rows.Scan(&category_id, &category)
		}
	}

	return &category_id, nil
}
