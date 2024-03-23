package api

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Idea struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
