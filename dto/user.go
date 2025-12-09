package dto

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required, min=5"`
}
