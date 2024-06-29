package dto

import "time"

type AddressCreateRequest struct {
	UserID        uint   `json:"user_id"`
	Street        string `json:"street"`
	SubDistrict   string `json:"sub_district,omitempty"`
	District      string `json:"district,omitempty"`
	CityOrRegency string `json:"city_or_regency"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	PostalCode    string `json:"postal_code,omitempty"`
}

type AddressUpdateRequest struct {
	Street        string `json:"street"`
	SubDistrict   string `json:"sub_district,omitempty"`
	District      string `json:"district,omitempty"`
	CityOrRegency string `json:"city_or_regency"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	PostalCode    string `json:"postal_code,omitempty"`
	IsDefault     bool   `json:"is_default"`
}

type AddressResponse struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"user_id"`
	Street        string    `json:"street"`
	SubDistrict   string    `json:"sub_district,omitempty"`
	District      string    `json:"district,omitempty"`
	CityOrRegency string    `json:"city_or_regency"`
	Province      string    `json:"province"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code,omitempty"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateAddressRequest struct {
	Address AddressCreateRequest `json:"address"`
}

type CreateAddressResponse struct {
	Address AddressResponse `json:"address"`
}

type ReadAddressResponse struct {
	Address AddressResponse `json:"address"`
}

type UpdateAddressRequest struct {
	Address AddressUpdateRequest `json:"address"`
}

type UpdateAddressResponse struct {
	Address AddressResponse `json:"address"`
}
