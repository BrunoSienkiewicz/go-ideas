package ideastorage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	addIdeaQuery = `INSERT INTO ideas.ideas (name, category_id) VALUES ($1, $2) RETURNING idea_id`
)

func (s *IdeaStorage) AddObject(obj *types.Idea) error {
	category_id, err := s.handleCategory(obj.Category)
	if err != nil {
		return err
	}

	rows, err := s.store.Query(addIdeaQuery, obj.Name, *category_id)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&obj.ID)
	}

	return nil
}
