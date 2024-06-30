package dto

import "errors"

// PaymentMethod Failed Messages
const (
	MESSAGE_FAILED_CREATE_PAYMENT_METHOD    = "failed to create payment method"
	MESSAGE_FAILED_GET_ALL_PAYMENT_METHODS  = "failed to get all payment methods"
	MESSAGE_FAILED_GET_PAYMENT_METHOD_BY_ID = "failed to get payment method by ID"
)

// PaymentMethod Success Messages
const (
	MESSAGE_SUCCESS_CREATE_PAYMENT_METHOD    = "success to create payment method"
	MESSAGE_SUCCESS_GET_ALL_PAYMENT_METHODS  = "success to get all payment methods"
	MESSAGE_SUCCESS_GET_PAYMENT_METHOD_BY_ID = "success to get payment method by ID"
)

// PaymentMethod Custom Errors
var (
	ErrCreatePaymentMethod  = errors.New(MESSAGE_FAILED_CREATE_PAYMENT_METHOD)
	ErrGetAllPaymentMethods = errors.New(MESSAGE_FAILED_GET_ALL_PAYMENT_METHODS)
	ErrGetPaymentMethodByID = errors.New(MESSAGE_FAILED_GET_PAYMENT_METHOD_BY_ID)
)
