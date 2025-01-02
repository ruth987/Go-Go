package models

type Book struct {
	ID int
	Title string
	Author string
	Status string
}

func NewBook(id int, title string, author string) Book {
	return Book{
		ID: id,
		Title: title,
		Author: author,
		Status: "Available",
	}
}