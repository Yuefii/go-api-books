package models

type Book struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Rating      int      `json:"rating"`
	Stock       int      `json:"stock"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
	CategoryID  *uint    `json:"category_id"`
}
