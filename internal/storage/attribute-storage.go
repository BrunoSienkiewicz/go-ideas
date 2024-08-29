package storage

import (
	"fmt"

	_ "github.com/lib/pq"

	db "github.com/BrunoSienkiewicz/go_ideas/internal/db"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	getAttributeByIDQuery    = `SELECT * FROM ideas.attributes WHERE attribute_id = $1`
	getAttributeByFieldQuery = `SELECT * FROM ideas.attributes WHERE $1 = $2`
	getAllAttributesQuery    = `SELECT * FROM ideas.attributes`

	addAttributeQuery = `INSERT INTO ideas.attributes (attribute_id, name, value, idea_id) VALUES ($1, $2, $3, $4) RETURNING attribute_id`

	updateAttributeQuery = `UPDATE ideas.attributes SET name = $1, value = $2 WHERE attribute_id = $3 RETURNING attribute_id`

	deleteAttributeQuery = `DELETE FROM ideas.attributes WHERE attribute_id = $1`
)

type AttributeStorage struct {
	store *db.Postgres
}

func NewAttributeStorage(postgres *db.Postgres) *AttributeStorage {
	return &AttributeStorage{
		store: postgres,
	}
}

func (s *AttributeStorage) GetObject(id int) (*types.DbAttribute, error) {
	rows, err := s.store.Query(getAttributeByIDQuery, id)
	if err != nil {
		return nil, err
	}

	attribute, err := types.ScanIntoAttribute(rows)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}

func (s *AttributeStorage) GetObjectsByField(field string, value string) ([]*types.DbAttribute, error) {
	query := fmt.Sprintf("SELECT * FROM ideas.attributes WHERE %s = $1", field)
	rows, err := s.store.Query(query, value)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, NotFoundError{Err: "Attribute with " + field + ": " + value + " Not Found"}
	}

	var attributes []*types.DbAttribute
	for rows.Next() {
		attribute, err := types.ScanIntoAttribute(rows)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attribute)
	}

	return attributes, nil
}

func (s *AttributeStorage) GetAllObjects() ([]*types.DbAttribute, error) {
	rows, err := s.store.Query(getAllAttributesQuery)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, NotFoundError{Err: "No attributes found"}
	}

	var attributes []*types.DbAttribute
	for rows.Next() {
		attribute, err := types.ScanIntoAttribute(rows)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attribute)
	}

	return attributes, nil
}

func (s *AttributeStorage) AddObject(attribute *types.DbAttribute) (*types.DbAttribute, error) {
	rows, err := s.store.Query(addAttributeQuery, attribute.Name, attribute.Value, attribute.Idea_id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, NotFoundError{Err: "Unable to add attribute"}
	}
	rows.Scan(&attribute.Attribute_id)

	return attribute, nil
}

func (s *AttributeStorage) UpdateObject(attribute *types.DbAttribute) (*types.DbAttribute, error) {
	rows, err := s.store.Query(updateAttributeQuery, attribute.Name, attribute.Value, attribute.Attribute_id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, InvalidFieldError{Err: "Unable to update attribute"}
	}

	return attribute, nil
}

func (s *AttributeStorage) UpdateObjectField(id int, field string, value string) (*types.DbAttribute, error) {
	query := `UPDATE ideas.attributes SET $1 = $2 WHERE attribute_id = $3 RETURNING attribute_id`
	rows, err := s.store.Query(query, field, value, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, InvalidFieldError{Err: "Unable to update attribute"}
	}

	return s.GetObject(id)
}

func (s *AttributeStorage) DeleteObject(id int) error {
	_, err := s.store.Query(deleteAttributeQuery, id)
	if err != nil {
		return err
	}

	return nil
}
