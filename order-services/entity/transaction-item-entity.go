package entity

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint    `gorm:"not null" json:"transaction_id"`
	ProductID     uint    `gorm:"not null" json:"product_id"`
	Quantity      int     `gorm:"not null" json:"quantity"`
	Price         float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}
