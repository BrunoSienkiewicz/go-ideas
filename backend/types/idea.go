package types

import "database/sql"

type Idea struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

type DbIdea struct {
	DbObject
	Idea_id     int
	Name        string
	Category_id int
}

type CreateIdeaRequest struct {
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

type UpdateIdeaRequest struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

type DeleteIdeaRequest struct {
	ID int `json:"id"`
}

func NewIdea(name string, category string, attributes []Attribute) *Idea {
	return &Idea{
		Name:       name,
		Category:   category,
		Attributes: attributes,
	}
}

func ScanIntoIdea(rows *sql.Rows) (*DbIdea, error) {
	idea := new(DbIdea)
	if err := rows.Scan(&idea.Idea_id, &idea.Name, &idea.Category_id); err != nil {
		return nil, err
	}

	return idea, nil
}
