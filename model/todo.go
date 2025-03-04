package model

type Todo struct {
	Title string `json:"title" validate:"required,min=5,max=20"`
	Text  string `json:"text" validate:"required,min=5,max=20"`
}