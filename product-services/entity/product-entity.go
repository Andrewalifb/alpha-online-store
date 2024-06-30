package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Inventory   int     `gorm:"not null" json:"inventory"`
	ImageURL    string  `gorm:"type:varchar(255)" json:"image_url"`
	CategoryID  uint    `json:"category_id"`
	Status      string  `gorm:"type:varchar(50);not null;default:'AVAILABLE'" json:"status"`
}
