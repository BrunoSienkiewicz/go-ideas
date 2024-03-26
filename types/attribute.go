package types

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
