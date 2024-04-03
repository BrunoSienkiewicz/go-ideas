package ideastorage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	updateIdeaQuery = `UPDATE ideas.ideas SET name = $1, category_id = $2 WHERE idea_id = $3`
)

func (s *IdeaStorage) UpdateObject(obj *types.Idea) error {
	category_id, err := s.handleCategory(obj.Category)
	if err != nil {
		return err
	}

	_, err = s.store.Query(updateIdeaQuery, obj.Name, *category_id, obj.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *IdeaStorage) UpdateObjectField(id int, field string, value string) error {
	query := `UPDATE ideas.ideas SET ` + field + ` = $1 WHERE idea_id = $2`
	_, err := s.store.Query(query, value, id)
	if err != nil {
		return err
	}

	return nil
}
