package repository

import (
	"github.com/Andrewalifb/alpha-online-store/cart-services/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart entity.Cart) (*entity.Cart, error)
	GetCartByUserID(userID uint) (*entity.Cart, error)
	DeleteCart(cartID uint) error
	UpdateCart(cart entity.Cart) (*entity.Cart, error)
	GetCartsByUserID(userID uint) ([]*entity.Cart, error)
	DeleteCartsByUserID(userID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) CreateCart(cart entity.Cart) (*entity.Cart, error) {
	result := r.db.Create(&cart)
	return &cart, result.Error
}

func (r *cartRepository) GetCartByUserID(userID uint) (*entity.Cart, error) {
	var cart entity.Cart
	result := r.db.Where("user_id = ?", userID).First(&cart)
	return &cart, result.Error
}

func (r *cartRepository) DeleteCart(cartID uint) error {
	result := r.db.Delete(&entity.Cart{}, cartID)
	return result.Error
}

func (r *cartRepository) UpdateCart(cart entity.Cart) (*entity.Cart, error) {
	result := r.db.Save(&cart)
	return &cart, result.Error
}

func (r *cartRepository) GetCartsByUserID(userID uint) ([]*entity.Cart, error) {
	var carts []*entity.Cart
	result := r.db.Where("user_id = ?", userID).Find(&carts)
	return carts, result.Error
}

func (r *cartRepository) DeleteCartsByUserID(userID uint) error {
	result := r.db.Where("user_id = ?", userID).Delete(&entity.Cart{})
	return result.Error
}
