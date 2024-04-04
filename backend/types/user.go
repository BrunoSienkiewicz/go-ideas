package types

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DbUser struct {
	DbObject
	user_id  int
	name     string
	email    string
	password string
}
