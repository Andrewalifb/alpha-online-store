package dto

import "errors"

// TransactionItem Failed Messages
const (
	MESSAGE_FAILED_CREATE_TRANSACTION_ITEM                = "failed to create transaction item"
	MESSAGE_FAILED_GET_TRANSACTION_ITEM_BY_ID             = "failed to get transaction item by ID"
	MESSAGE_FAILED_GET_TRANSACTION_ITEM_BY_TRANSACTION_ID = "failed to get transaction items by transaction ID"
)

// TransactionItem Success Messages
const (
	MESSAGE_SUCCESS_CREATE_TRANSACTION_ITEM                = "success to create transaction item"
	MESSAGE_SUCCESS_GET_TRANSACTION_ITEM_BY_ID             = "success to get transaction item by ID"
	MESSAGE_SUCCESS_GET_TRANSACTION_ITEM_BY_TRANSACTION_ID = "success to get transaction items by transaction ID"
)

// TransactionItem Custom Errors
var (
	ErrCreateTransactionItem              = errors.New(MESSAGE_FAILED_CREATE_TRANSACTION_ITEM)
	ErrGetTransactionItemByID             = errors.New(MESSAGE_FAILED_GET_TRANSACTION_ITEM_BY_ID)
	ErrGetTransactionItemsByTransactionID = errors.New(MESSAGE_SUCCESS_GET_TRANSACTION_ITEM_BY_TRANSACTION_ID)
)
