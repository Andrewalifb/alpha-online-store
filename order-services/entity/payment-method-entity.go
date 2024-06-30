package entity

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Method string `gorm:"type:varchar(50);not null" json:"method"`
}
