package dto

import "errors"

// Transaction Failed Messages
const (
	MESSAGE_FAILED_CREATE_TRANSACTIONS         = "failed to create transactions"
	MESSAGE_FAILED_GET_TRANSACTION_BY_ID       = "failed to get transaction by ID"
	MESSAGE_FAILED_GET_TRANSACTIONS_BY_USER_ID = "failed to get transactions by user ID"
	MESSAGE_FAILED_UPDATE_TRANSACTION_STATUS   = "failed to update transaction status"
)

// Transaction Success Messages
const (
	MESSAGE_SUCCESS_CREATE_TRANSACTIONS         = "success to create transactions"
	MESSAGE_SUCCESS_GET_TRANSACTION_BY_ID       = "success to get transaction by ID"
	MESSAGE_SUCCESS_GET_TRANSACTIONS_BY_USER_ID = "success to get transactions by user ID"
	MESSAGE_SUCCESS_UPDATE_TRANSACTION_STATUS   = "success to update transaction status"
)

// Transaction Custom Errors
var (
	ErrCreateTransactions      = errors.New(MESSAGE_FAILED_CREATE_TRANSACTIONS)
	ErrGetTransactionByID      = errors.New(MESSAGE_FAILED_GET_TRANSACTION_BY_ID)
	ErrGetTransactionsByUserID = errors.New(MESSAGE_FAILED_GET_TRANSACTIONS_BY_USER_ID)
	ErrUpdateTransactionStatus = errors.New(MESSAGE_FAILED_UPDATE_TRANSACTION_STATUS)
)
