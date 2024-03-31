package models

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	CurrentPage string `json:"current_page"`
	PagesCount  string `json:"pages_count"`
}

type CreateBookInput struct {
	Name        string `json:"name" binding:"required"`
	Author      string `json:"author" binding:"required"`
	CurrentPage string `json:"current_page" binding:"required"`
	PagesCount  string `json:"pages_count" binding:"required"`
}

type UpdateBookInput struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	CurrentPage string `json:"current_page"`
	PagesCount  string `json:"pages_count"`
}
