package models

type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Gender string `json:"gender"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}