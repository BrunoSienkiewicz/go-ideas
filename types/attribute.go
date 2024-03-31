package types

import "database/sql"

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewAttribute(name string, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: value,
	}
}

func ScanIntoAttribute(rows *sql.Rows) (*Attribute, error) {
	attribute := new(Attribute)
	if err := rows.Scan(&attribute.Name, &attribute.Value); err != nil {
		return nil, err
	}

	return attribute, nil
}
