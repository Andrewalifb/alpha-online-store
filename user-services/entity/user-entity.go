package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"type:varchar(50);unique;not null" json:"username"`
	HashedPassword string `gorm:"type:varchar(255);not null" json:"hashed_password"`
	Email          string `gorm:"type:varchar(255);unique;not null" json:"email"`
	FirstName      string `gorm:"type:varchar(100)" json:"first_name"`
	LastName       string `gorm:"type:varchar(100)" json:"last_name"`
	PhoneNumber    string `gorm:"type:varchar(20)" json:"phone_number"`
	Role           string `gorm:"type:varchar(50);not null" json:"role"`
}
