package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID        uint   `gorm:"not null" json:"user_id"`
	Street        string `gorm:"type:varchar(255);not null" json:"street"`
	SubDistrict   string `gorm:"type:varchar(100)" json:"sub_district"`
	District      string `gorm:"type:varchar(100)" json:"district"`
	CityOrRegency string `gorm:"type:varchar(100);not null" json:"city_or_regency"`
	Province      string `gorm:"type:varchar(100);not null" json:"province"`
	Country       string `gorm:"type:varchar(100);not null" json:"country"`
	PostalCode    string `gorm:"type:varchar(10)" json:"postal_code"`
	IsDefault     bool   `gorm:"not null;default:false" json:"is_default"`
}
