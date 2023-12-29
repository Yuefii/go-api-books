package models

import "time"

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"subtitle" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
