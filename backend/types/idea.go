package types

import "database/sql"

type Idea struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
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

func ScanIntoIdea(rows *sql.Rows) (*Idea, error) {
	idea := new(Idea)
	if err := rows.Scan(&idea.ID, &idea.Name, &idea.Category); err != nil {
		return nil, err
	}

	return idea, nil
}
