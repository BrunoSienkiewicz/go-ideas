package types

import "database/sql"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DbCategory struct {
	DbObject
	Category_id int
	Name        string
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}

func ScanIntoCategory(rows *sql.Rows) (*DbCategory, error) {
	category := new(DbCategory)
	if err := rows.Scan(&category.Category_id, &category.Name); err != nil {
		return nil, err
	}

	return category, nil
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteCategoryRequest struct {
	ID int `json:"id"`
}
