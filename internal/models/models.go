package models

type Task struct {
	Id      string `json:"id" db:"id"`
	Date    string `json:"date" db:"date"`
	Title   string `json:"title" validate:"required" db:"title"`
	Comment string `json:"comment,omitempty" db:"comment"`
	Repeat  string `json:"repeat,omitempty" db:"repeat"`
}
