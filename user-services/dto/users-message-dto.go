package dto

import "errors"

// User Failed Messages
const (
	MESSAGE_FAILED_REGISTER_NEW_USER = "failed to register new user"
	MESSAGE_FAILED_LOGIN_USER        = "failed to login"
)

// User Success Messages
const (
	MESSAGE_SUCCESS_REGISTER_NEW_USER = "success to register new user"
	MESSAGE_SUCCESS_LOGIN_USER        = "success to login"
)

// User Custom Errors
var (
	ErrRegisterNewUser = errors.New(MESSAGE_FAILED_REGISTER_NEW_USER)
	ErrLoginUser       = errors.New(MESSAGE_FAILED_LOGIN_USER)
)
