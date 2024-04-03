package storage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
	_ "github.com/lib/pq"
)

type AttributeStorage struct {
	store *PostgresStorage
}

func NewAttributeStorage(postgres *PostgresStorage) *AttributeStorage {
	return &AttributeStorage{
		store: postgres,
	}
}

const (
	attributeSelect = `SELECT * FROM ideas.attributes WHERE idea_id = $1`
	attributeInsert = `INSERT INTO ideas.attributes (idea_id, name, value) VALUES ($1, $2, $3)`
	attributeUpdate = `UPDATE ideas.attributes SET name = $1, value = $2 WHERE attribute_id = $3`
	attributeDelete = `DELETE FROM ideas.attributes WHERE attribute_id = $1`
)

func (s *AttributeStorage) GetObject(id int) (*types.Attribute, error) {
	rows, err := s.store.db.Query(attributeSelect, id)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	attribute, err := types.ScanIntoAttribute(rows)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}

func (s *AttributeStorage) AddObject(obj *types.Attribute) error {
	_, err := s.store.db.Exec(attributeInsert, obj.IdeaId, obj.Name, obj.Value)
	return err
}

func (s *AttributeStorage) UpdateObject(obj *types.Attribute) error {
	_, err := s.store.db.Exec(attributeUpdate, obj.Name, obj.Value, obj.Id)
	return err
}

func (s *AttributeStorage) DeleteObject(id int) error {
	_, err := s.store.db.Exec(attributeDelete, id)
	return err
}

func (s *AttributeStorage) ListObjects() ([]*types.Attribute, error) {
	rows, err := s.store.db.Query(`SELECT * FROM ideas.attributes`)
	if err != nil {
		return nil, err
	}

	var attributes []*types.Attribute
	for rows.Next() {
		attribute, err := types.ScanIntoAttribute(rows)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attribute)
	}

	return attributes, nil
}
