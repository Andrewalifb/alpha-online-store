package service

import (
	"github.com/Andrewalifb/alpha-online-store/cart-services/dto"
	"github.com/Andrewalifb/alpha-online-store/cart-services/entity"
	"github.com/Andrewalifb/alpha-online-store/cart-services/pkg/repository"
)

type CartItemService interface {
	CreateCartItem(cartID uint, cartItem dto.CartItemRequest) (*dto.CartItemResponse, error)
	GetCartItemsByCartID(cartID uint) ([]dto.CartItemResponse, error)
	UpdateCartItem(cartID uint, cartItem dto.CartItemRequest) (*dto.CartItemResponse, error)
	DeleteCartItem(cartItemID uint) (*dto.RemoveCartResponse, error)
}

type cartItemService struct {
	cartItemRepo repository.CartItemRepository
}

func NewCartItemService(cartItemRepo repository.CartItemRepository) CartItemService {
	return &cartItemService{
		cartItemRepo: cartItemRepo,
	}
}

func (s *cartItemService) CreateCartItem(cartID uint, cartItem dto.CartItemRequest) (*dto.CartItemResponse, error) {
	cartItemEntity := entity.CartItem{
		CartID:    cartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
		Price:     cartItem.Price,
	}

	createCartItemResult, err := s.cartItemRepo.CreateCartItem(cartItemEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CartItemResponse{
		ID:        createCartItemResult.ID,
		CartID:    createCartItemResult.CartID,
		ProductID: createCartItemResult.ProductID,
		Quantity:  createCartItemResult.Quantity,
		Price:     createCartItemResult.Price,
	}, nil
}

func (s *cartItemService) GetCartItemsByCartID(cartID uint) ([]dto.CartItemResponse, error) {
	cartItems, err := s.cartItemRepo.GetCartItemsByCartID(cartID)
	if err != nil {
		return nil, err
	}

	var cartItemResponses []dto.CartItemResponse
	for _, cartItem := range cartItems {
		cartItemResponses = append(cartItemResponses, dto.CartItemResponse{
			ID:        cartItem.ID,
			CartID:    cartItem.CartID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		})
	}

	return cartItemResponses, nil
}

func (s *cartItemService) UpdateCartItem(cartID uint, cartItem dto.CartItemRequest) (*dto.CartItemResponse, error) {
	cartItemEntity := entity.CartItem{
		CartID:    cartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
		Price:     cartItem.Price,
	}

	updateCartItemResult, err := s.cartItemRepo.UpdateCartItem(cartItemEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CartItemResponse{
		ID:        updateCartItemResult.ID,
		CartID:    updateCartItemResult.CartID,
		ProductID: updateCartItemResult.ProductID,
		Quantity:  updateCartItemResult.Quantity,
		Price:     updateCartItemResult.Price,
	}, nil
}

func (s *cartItemService) DeleteCartItem(cartItemID uint) (*dto.RemoveCartResponse, error) {
	err := s.cartItemRepo.DeleteCartItem(cartItemID)
	if err != nil {
		return nil, err
	}

	return &dto.RemoveCartResponse{
		RemovedCartID: cartItemID,
	}, nil
}
