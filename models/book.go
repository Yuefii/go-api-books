package models

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"subtitle" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}
