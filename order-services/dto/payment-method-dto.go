package dto

import "time"

// Create New Payment
type CreatePaymentMethodRequest struct {
	Method string `json:"method"`
}

type CreatePaymentMethodResponse struct {
	ID        uint      `json:"id"`
	Method    string    `json:"method"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Read All Payment
type ReadAllPaymentMethodsResponse struct {
	PaymentMethods []CreatePaymentMethodResponse `json:"payment_methods"`
}

// Read Payment By Id
type ReadPaymentMethodResponse CreatePaymentMethodResponse
