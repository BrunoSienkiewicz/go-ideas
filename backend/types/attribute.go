package types

import "database/sql"

type Attribute struct {
	ApiObject
	ID     int    `json:"id"`
	IdeaId int    `json:"idea_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type DbAttribute struct {
	DbObject
	Attribute_id int
	Idea_id      int
	Name         string
	Value        string
}

func NewAttribute(name string, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: value,
	}
}

func ScanIntoAttribute(rows *sql.Rows) (*DbAttribute, error) {
	attribute := new(DbAttribute)
	if err := rows.Scan(&attribute.Attribute_id, &attribute.Idea_id, &attribute.Name, &attribute.Value); err != nil {
		return nil, err
	}

	return attribute, nil
}
