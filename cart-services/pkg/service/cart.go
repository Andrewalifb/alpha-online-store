package service

import (
	"fmt"

	"github.com/Andrewalifb/alpha-online-store/cart-services/dto"
	"github.com/Andrewalifb/alpha-online-store/cart-services/entity"
	"github.com/Andrewalifb/alpha-online-store/cart-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/cart-services/utils"
)

type CartService interface {
	CreateCart(cart dto.AddToCartRequest) (*dto.AddToCartResponse, error)
	GetCartsByUserID(userID uint) ([]*dto.GetAllCartByUserIdResponse, error)
	DeleteCart(cartID uint) (*dto.RemoveCartResponse, error)
	UpdateCart(cartID uint, cartItem dto.CartItemRequest) (*dto.UpdateCartResponse, error)
	CheckoutCarts(cartIDs dto.CheckoutRequest) (*dto.CheckoutResponse, error)
	DeleteCartsByUserID(userID uint) (*dto.RemoveCartResponse, error)
}

type cartService struct {
	cartRepo     repository.CartRepository
	cartItemRepo repository.CartItemRepository
}

func NewCartService(cartRepo repository.CartRepository, cartItemRepo repository.CartItemRepository) CartService {
	return &cartService{
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo,
	}
}

func (s *cartService) CreateCart(cart dto.AddToCartRequest) (*dto.AddToCartResponse, error) {
	cartEntity := entity.Cart{
		UserID: cart.UserID,
	}

	createCartResult, err := s.cartRepo.CreateCart(cartEntity)
	if err != nil {
		return nil, err
	}

	productResult, err := utils.GetProduct(fmt.Sprint(cart.CartItem.ProductID))
	if err != nil {
		return nil, err
	}

	cartItemEntity := entity.CartItem{
		CartID:    createCartResult.ID,
		ProductID: cart.CartItem.ProductID,
		Quantity:  cart.CartItem.Quantity,
		Price:     productResult.Data.Price,
	}

	createCartItemResult, err := s.cartItemRepo.CreateCartItem(cartItemEntity)
	if err != nil {
		return nil, err
	}

	return &dto.AddToCartResponse{
		CartID: createCartResult.ID,
		CartItem: dto.CartItemResponse{
			ID:        createCartItemResult.ID,
			CartID:    createCartItemResult.CartID,
			ProductID: createCartItemResult.ProductID,
			Quantity:  createCartItemResult.Quantity,
			Price:     createCartItemResult.Price,
		},
	}, nil
}

func (s *cartService) GetCartsByUserID(userID uint) ([]*dto.GetAllCartByUserIdResponse, error) {
	carts, err := s.cartRepo.GetCartsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.GetAllCartByUserIdResponse
	for _, cart := range carts {
		cartItems, err := s.cartItemRepo.GetCartItemsByCartID(cart.ID)
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

		responses = append(responses, &dto.GetAllCartByUserIdResponse{
			ID:        cart.ID,
			UserID:    cart.UserID,
			CartItems: cartItemResponses,
		})
	}

	return responses, nil
}

func (s *cartService) DeleteCart(cartID uint) (*dto.RemoveCartResponse, error) {
	err := s.cartRepo.DeleteCart(cartID)
	if err != nil {
		return nil, err
	}

	return &dto.RemoveCartResponse{
		RemovedCartID: cartID,
	}, nil
}

func (s *cartService) UpdateCart(cartID uint, cartItem dto.CartItemRequest) (*dto.UpdateCartResponse, error) {
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

	return &dto.UpdateCartResponse{
		CartID: updateCartItemResult.CartID,
		CartItem: dto.CartItemResponse{
			ID:        updateCartItemResult.ID,
			CartID:    updateCartItemResult.CartID,
			ProductID: updateCartItemResult.ProductID,
			Quantity:  updateCartItemResult.Quantity,
			Price:     updateCartItemResult.Price,
		},
	}, nil
}

func (s *cartService) CheckoutCarts(cartIDs dto.CheckoutRequest) (*dto.CheckoutResponse, error) {
	var checkoutResponse dto.CheckoutResponse
	var transactionItems []dto.TransactionItemRequest
	var totalPrice float64

	for _, cartID := range cartIDs.CartIDs {
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

			transactionItems = append(transactionItems, dto.TransactionItemRequest{
				ProductID: cartItem.ProductID,
				Quantity:  cartItem.Quantity,
				Price:     cartItem.Price,
			})

			totalPrice += (float64(cartItem.Quantity) * cartItem.Price)

			productResult, err := utils.GetProduct(fmt.Sprint(cartItem.ProductID))
			if err != nil {
				return nil, err
			}

			decreaseProductQty := dto.UpdateProductRequest{
				Name:        productResult.Data.Name,
				Description: productResult.Data.Description,
				Price:       productResult.Data.Price,
				Inventory:   productResult.Data.Inventory - cartItem.Quantity,
				ImageURL:    productResult.Data.ImageURL,
				CategoryID:  productResult.Data.Category.ID,
				Status:      productResult.Data.Status,
			}

			_, err = utils.UpdateProductById(fmt.Sprint(cartItem.ProductID), &decreaseProductQty)
			if err != nil {
				return nil, err
			}
		}

		checkoutResponse.Carts = append(checkoutResponse.Carts, struct {
			CartID    uint                   `json:"cart_id"`
			CartItems []dto.CartItemResponse `json:"cart_items"`
		}{
			CartID:    cartID,
			CartItems: cartItemResponses,
		})

		err = s.cartRepo.DeleteCart(cartID)
		if err != nil {
			return nil, err
		}
	}

	transactionData := dto.CreateTransactionsRequest{
		Transactions: dto.TransactionRequest{
			UserID:          cartIDs.UserId,
			Address:         "",
			Total:           totalPrice,
			TransactionItem: transactionItems,
		},
	}
	_, err := utils.CreateTransactions(&transactionData)
	if err != nil {
		return nil, err
	}

	return &checkoutResponse, nil
}

func (s *cartService) DeleteCartsByUserID(userID uint) (*dto.RemoveCartResponse, error) {
	carts, err := s.cartRepo.GetCartsByUserID(userID)
	if err != nil {
		return nil, err
	}

	for _, cart := range carts {
		err = s.cartItemRepo.DeleteCartItemsByCartID(cart.ID)
		if err != nil {
			return nil, err
		}
	}

	err = s.cartRepo.DeleteCartsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.RemoveCartResponse{
		RemovedCartID: userID,
	}, nil
}
