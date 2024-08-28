package models

type GetUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
