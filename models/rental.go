package models

type Rental struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	UserID    uint    `json:"user_id"`
	Book      Book    `gorm:"foreignKey:BookID" json:"book"`
	BookID    uint    `json:"book_id"`
	RentalFee float64 `json:"rental_fee"`
}
