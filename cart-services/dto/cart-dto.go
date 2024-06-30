package dto

import "time"

// Add to cart
type CartItemRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type AddToCartRequest struct {
	UserID   uint            `json:"user_id"`
	CartItem CartItemRequest `json:"cart_item"`
}

type AddToCartResponse struct {
	CartID   uint             `json:"cart_id"`
	CartItem CartItemResponse `json:"cart_item"`
}

// Get all cart details by user id
type GetAllCartByUserIdRequest struct {
	UserID uint `json:"user_id"`
}

type CartItemResponse struct {
	ID        uint      `json:"id"`
	CartID    uint      `json:"cart_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllCartByUserIdResponse struct {
	ID        uint               `json:"id"`
	UserID    uint               `json:"user_id"`
	CartItems []CartItemResponse `json:"cart_items"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// Check out multiple cart id
type CheckoutRequest struct {
	UserId  uint   `json:"user_id"`
	CartIDs []uint `json:"cart_ids"`
}

type CheckoutResponse struct {
	Carts []struct {
		CartID    uint               `json:"cart_id"`
		CartItems []CartItemResponse `json:"cart_items"`
	} `json:"carts"`
}

// Update cart
type UpdateCartRequest struct {
	CartID   uint            `json:"cart_id"`
	CartItem CartItemRequest `json:"cart_item"`
}

type UpdateCartResponse struct {
	CartID   uint             `json:"cart_id"`
	CartItem CartItemResponse `json:"cart_item"`
}

// Remove single cart
type RemoveCartRequest struct {
	CartID uint `json:"cart_id"`
}

type RemoveCartResponse struct {
	RemovedCartID uint `json:"removed_cart_id"`
}
