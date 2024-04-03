package ideastorage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	deleteIdeaQuery = `DELETE FROM ideas.ideas WHERE idea_id = $1`
)

func (s *IdeaStorage) DeleteObject(id int) error {
	_, err := s.store.Query(deleteIdeaQuery, id)
	if err != nil {
		return err
	}

	return nil
}
