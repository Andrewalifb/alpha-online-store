package dto

import "errors"

// Address Failed Messages
const (
	MESSAGE_FAILED_CREATE_ADDRESS = "failed to create user address"
	MESSAGE_FAILED_UPDATE_ADDRESS = "failed to update user address"
	MESSAGE_FAILED_READ_ADDRESS   = "failed to read user address"
)

// Address Success Messages
const (
	MESSAGE_SUCCESS_CREATE_ADDRESS = "success to create user address"
	MESSAGE_SUCCESS_UPDATE_ADDRESS = "success to update user address"
	MESSAGE_SUCCESS_READ_ADDRESS   = "success to read user address"
)

// Address Custom Errors
var (
	ErrCrateAddressUser  = errors.New(MESSAGE_FAILED_UPDATE_ADDRESS)
	ErrUpdateAddressUser = errors.New(MESSAGE_FAILED_UPDATE_ADDRESS)
	ErrReadAddressUser   = errors.New(MESSAGE_FAILED_READ_ADDRESS)
)
