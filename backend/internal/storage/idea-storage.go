package storage

import (
	_ "github.com/lib/pq"

	db "github.com/BrunoSienkiewicz/go_ideas/internal/db"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	getIdeaByIDQuery    = `SELECT * FROM ideas.ideas WHERE id = $1`
	getIdeaByFieldQuery = `SELECT * FROM ideas.ideas WHERE $1 = $2`
	getAllIdeasQuery    = `SELECT * FROM ideas.ideas`

	addIdeaQuery = `INSERT INTO ideas.ideas (name, category_id) VALUES ($1, $2) RETURNING idea_id`

	updateIdeaQuery = `UPDATE ideas.ideas SET name = $1, category_id = $2 WHERE idea_id = $3`

	deleteIdeaQuery = `DELETE FROM ideas.ideas WHERE idea_id = $1`
)

type IdeaStorage struct {
	store            *db.Postgres
	attributeStorage AttributeStorage
}

func NewIdeaStorage(postgres *db.Postgres, attAttributeStorage AttributeStorage) *IdeaStorage {
	return &IdeaStorage{
		store:            postgres,
		attributeStorage: attAttributeStorage,
	}
}

func (s *IdeaStorage) GetObject(id int) (*types.DbIdea, error) {
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

func (s *IdeaStorage) GetAllObjects() ([]*types.DbIdea, error) {
	rows, err := s.store.Query(getAllIdeasQuery)
	if err != nil {
		return nil, err
	}

	var ideas []*types.DbIdea
	for rows.Next() {
		idea, err := types.ScanIntoIdea(rows)
		if err != nil {
			return nil, err
		}

		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (s *IdeaStorage) AddObject(obj *types.DbIdea) error {
	rows, err := s.store.Query(addIdeaQuery, obj.Name, obj.Category_id)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&obj.Idea_id)
	}

	return nil
}

func (s *IdeaStorage) UpdateObject(obj *types.DbIdea) (*types.DbIdea, error) {
	rows, err := s.store.Query(updateIdeaQuery, obj.Name, obj.Category_id, obj.Idea_id)
	if err != nil {
		return nil, err
	}

	idea, err := types.ScanIntoIdea(rows)
	if err != nil {
		return nil, err
	}

	return idea, nil
}

func (s *IdeaStorage) UpdateObjectField(id int, field string, value string) error {
	query := `UPDATE ideas.ideas SET ` + field + ` = $1 WHERE idea_id = $2`
	_, err := s.store.Query(query, value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *IdeaStorage) DeleteObject(id int) error {
	_, err := s.store.Query(deleteIdeaQuery, id)
	if err != nil {
		return err
	}

	return nil
}
