package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint    `gorm:"not null" json:"user_id"`
	CartID          uint    `gorm:"not null" json:"cart_id"`
	Total           float64 `gorm:"type:decimal(10,2);not null" json:"total"`
	PaymentMethodID *uint   `json:"payment_method_id"`
	Status          string  `gorm:"type:varchar(50);not null;default:'PENDING'" json:"status"`
	Address         string  `gorm:"type:varchar(500);not null" json:"address"`
}
