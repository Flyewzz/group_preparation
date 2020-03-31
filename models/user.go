package models

type User struct {
	Id       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
