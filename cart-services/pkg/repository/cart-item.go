package repository

import (
	"github.com/Andrewalifb/alpha-online-store/cart-services/entity"
	"gorm.io/gorm"
)

type CartItemRepository interface {
	CreateCartItem(cartItem entity.CartItem) (*entity.CartItem, error)
	GetCartItemsByCartID(cartID uint) ([]*entity.CartItem, error)
	UpdateCartItem(cartItem entity.CartItem) (*entity.CartItem, error)
	DeleteCartItem(cartItemID uint) error
	DeleteCartItemsByCartID(cartID uint) error
}

type cartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{
		db: db,
	}
}

func (r *cartItemRepository) CreateCartItem(cartItem entity.CartItem) (*entity.CartItem, error) {
	result := r.db.Create(&cartItem)
	return &cartItem, result.Error
}

func (r *cartItemRepository) GetCartItemsByCartID(cartID uint) ([]*entity.CartItem, error) {
	var cartItems []*entity.CartItem
	result := r.db.Where("cart_id = ?", cartID).Find(&cartItems)
	return cartItems, result.Error
}

func (r *cartItemRepository) UpdateCartItem(cartItem entity.CartItem) (*entity.CartItem, error) {
	result := r.db.Save(&cartItem)
	return &cartItem, result.Error
}

func (r *cartItemRepository) DeleteCartItem(cartItemID uint) error {
	result := r.db.Delete(&entity.CartItem{}, cartItemID)
	return result.Error
}

func (r *cartItemRepository) DeleteCartItemsByCartID(cartID uint) error {
	result := r.db.Where("cart_id = ?", cartID).Delete(&entity.CartItem{})
	return result.Error
}
