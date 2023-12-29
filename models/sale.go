package models

type Sale struct {
	ID       uint    `gorm:"primaryKey" json:"rental_fee"`
	User     User    `gorm:"foreignKey:UserID" json:"user"`
	UserID   uint    `json:"user_id"`
	Book     Book    `gorm:"foreignKey:BookID" json:"book"`
	BookID   uint    `json:"book_id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
