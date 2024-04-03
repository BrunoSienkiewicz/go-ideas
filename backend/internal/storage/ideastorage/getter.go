package ideastorage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	getIdeaByIDQuery    = `SELECT * FROM ideas.ideas WHERE id = $1`
	getIdeaByFieldQuery = `SELECT * FROM ideas.ideas WHERE $1 = $2`
	getAllIdeasQuery    = `SELECT * FROM ideas.ideas`
)

func (s *IdeaStorage) GetObject(id int) (*types.Idea, error) {
	rows, err := s.store.Query(getIdeaByIDQuery, id)
	if err != nil {
		return nil, err
	}

	idea, err := types.ScanIntoIdea(rows)
	if err != nil {
		return nil, err
	}

	return idea, nil
}

func (s *IdeaStorage) GetAllObjects() ([]*types.Idea, error) {
	rows, err := s.store.Query(getAllIdeasQuery)
	if err != nil {
		return nil, err
	}

	var ideas []*types.Idea
	for rows.Next() {
		idea, err := types.ScanIntoIdea(rows)
		if err != nil {
			return nil, err
		}

		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (s *IdeaStorage) GetObjectByField(field string, value string) (*types.Idea, error) {
	rows, err := s.store.Query(getIdeaByFieldQuery, field, value)
	if err != nil {
		return nil, err
	}

	idea, err := types.ScanIntoIdea(rows)
	if err != nil {
		return nil, err
	}

	return idea, nil
}
