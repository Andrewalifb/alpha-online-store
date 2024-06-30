package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
