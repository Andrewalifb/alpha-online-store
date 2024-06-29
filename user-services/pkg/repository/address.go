package repository

import (
	"github.com/Andrewalifb/alpha-online-store/user-services/entity"
	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address entity.Address) (*entity.Address, error)
	UpdateUserAddress(addressID string, address entity.Address) (*entity.Address, error)
	ReadAddress(addressID uint) (*entity.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) CreateAddress(address entity.Address) (*entity.Address, error) {
	result := r.db.Create(&address)
	if result.Error != nil {
		return nil, result.Error
	}

	return &address, nil
}

func (r *addressRepository) UpdateUserAddress(addressID string, address entity.Address) (*entity.Address, error) {
	result := r.db.Model(&entity.Address{}).Where("id = ?", addressID).Updates(address)
	if result.Error != nil {
		return nil, result.Error
	}

	return &address, nil
}

func (r *addressRepository) ReadAddress(addressID uint) (*entity.Address, error) {
	var address entity.Address
	result := r.db.First(&address, addressID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &address, nil
}
