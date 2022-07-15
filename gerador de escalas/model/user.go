package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	PassWord string `json:"password"`
}
