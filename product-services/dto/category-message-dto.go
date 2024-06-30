package dto

import "errors"

// Category Failed Messages
const (
	MESSAGE_FAILED_CREATE_CATEGORY   = "failed to create product category"
	MESSAGE_FAILED_READ_ALL_CATEGORY = "failed to update product category"
)

// Category Success Messages
const (
	MESSAGE_SUCCESS_CREATE_CATEGORY   = "success to create product category"
	MESSAGE_SUCCESS_READ_ALL_CATEGORY = "success to update product category"
)

// Category Custom Errors
var (
	ErrCrateProductCategory   = errors.New(MESSAGE_SUCCESS_CREATE_CATEGORY)
	ErrReadAllProductCategory = errors.New(MESSAGE_SUCCESS_READ_ALL_CATEGORY)
)
