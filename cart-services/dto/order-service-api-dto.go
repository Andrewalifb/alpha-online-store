package dto

import "time"

type TransactionItemRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type TransactionRequest struct {
	UserID          uint                     `json:"user_id"`
	Address         string                   `json:"address"`
	Total           float64                  `json:"total"`
	TransactionItem []TransactionItemRequest `json:"transaction_item"`
}

type CreateTransactionsRequest struct {
	Transactions TransactionRequest `json:"transactions"`
}

type TransactionItemResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID              uint                      `json:"id"`
	UserID          uint                      `json:"user_id"`
	PaymentMethodID uint                      `json:"payment_method_id"`
	Address         string                    `json:"address"`
	TransactionItem []TransactionItemResponse `json:"transaction_item"`
	Total           float64                   `json:"total"`
	Status          string                    `json:"status"`
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
}

type CreateTransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}
