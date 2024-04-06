package storage

import (
	"fmt"

	_ "github.com/lib/pq"

	db "github.com/BrunoSienkiewicz/go_ideas/internal/db"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

const (
	getCategory      = `SELECT id FROM ideas.categories WHERE category_id = $1`
	getAllCategories = `SELECT * FROM ideas.categories`
	addCategory      = `INSERT INTO ideas.categories (name) VALUES ($1) RETURNING category_id`
	updateCategory   = `UPDATE ideas.categories SET name = $1 WHERE id = $2`
	deleteCategory   = `DELETE FROM ideas.categories WHERE id = $1`
)

type CategoryStorage struct {
	store *db.Postgres
}

func NewCategoryStorage(postgres *db.Postgres) *CategoryStorage {
	return &CategoryStorage{
		store: postgres,
	}
}

func (s *CategoryStorage) GetObject(id int) (*types.DbCategory, error) {
	rows, err := s.store.Query(getCategory, id)
	if err != nil {
		return nil, err
	}

	category, err := types.ScanIntoCategory(rows)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryStorage) GetObjectsByField(field string, value string) ([]*types.DbCategory, error) {
	query := fmt.Sprintf("SELECT * FROM ideas.categories WHERE %s = $1", field)
	rows, err := s.store.Query(query, value)
	if err != nil {
		return nil, err
	}

	var categories []*types.DbCategory
	for rows.Next() {
		category, err := types.ScanIntoCategory(rows)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *CategoryStorage) GetAllObjects() ([]*types.DbCategory, error) {
	rows, err := s.store.Query(getAllCategories)
	if err != nil {
		return nil, err
	}

	var categories []*types.DbCategory
	for rows.Next() {
		category, err := types.ScanIntoCategory(rows)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *CategoryStorage) AddObject(obj *types.DbCategory) error {
	rows, err := s.store.Query(addCategory, obj.Name)
	if err != nil {
		return err
	}

	if !rows.Next() {
		return StorageError{Err: "Unable to add category"}
	}

	err = rows.Scan(&obj.Category_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStorage) UpdateObject(obj *types.DbCategory) error {
	rows, err := s.store.Query(updateCategory, obj.Name, obj.Category_id)
	if err != nil {
		return err
	}

	if !rows.Next() {
		return StorageError{Err: "Unable to update category"}
	}

	return nil
}

func (s *CategoryStorage) UpdateObjectField(id int, field string, value string) error {
	query := fmt.Sprintf("UPDATE ideas.categories SET %s = $1 WHERE id = $2", field)
	_, err := s.store.Query(query, value, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStorage) DeleteObject(id int) error {
	_, err := s.store.Query(deleteCategory, id)
	if err != nil {
		return err
	}

	return nil
}
