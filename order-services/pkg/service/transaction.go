package service

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/dto"
	"github.com/Andrewalifb/alpha-online-store/order-services/entity"
	"github.com/Andrewalifb/alpha-online-store/order-services/pkg/repository"
)

type TransactionService interface {
	CreateTransactions(request dto.CreateTransactionsRequest) (*dto.CreateTransactionsResponse, error)
	GetTransactionByID(id uint) (*dto.TransactionResponse, error)
	GetTransactionsByUserID(userID uint) (*dto.ReadTransactionsByUserIDResponse, error)
	UpdateTransactionStatus(request dto.TransactionPaymentRequest) (*dto.TransactionPaymentResponse, error)
}

type transactionService struct {
	repo     repository.TransactionRepository
	itemRepo repository.TransactionItemRepository
}

func NewTransactionService(repo repository.TransactionRepository, itemRepo repository.TransactionItemRepository) TransactionService {
	return &transactionService{
		repo:     repo,
		itemRepo: itemRepo,
	}
}

func (s *transactionService) CreateTransactions(request dto.CreateTransactionsRequest) (*dto.CreateTransactionsResponse, error) {
	var response dto.CreateTransactionsResponse

	transaction := entity.Transaction{
		UserID:  request.Transactions.UserID,
		Address: request.Transactions.Address,
		Total:   request.Transactions.Total,
	}

	createdTransaction, err := s.repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	var transactionItemResponses []dto.TransactionItemResponse
	for _, transactionItemRequest := range request.Transactions.TransactionItem {
		transactionItem := entity.TransactionItem{
			TransactionID: createdTransaction.ID,
			ProductID:     transactionItemRequest.ProductID,
			Quantity:      transactionItemRequest.Quantity,
			Price:         transactionItemRequest.Price,
		}

		createdTransactionItem, err := s.itemRepo.CreateTransactionItem(transactionItem)
		if err != nil {
			return nil, err
		}

		transactionItemResponse := dto.TransactionItemResponse{
			ID:        createdTransactionItem.ID,
			ProductID: createdTransactionItem.ProductID,
			Quantity:  createdTransactionItem.Quantity,
			Price:     createdTransactionItem.Price,
			CreatedAt: createdTransactionItem.CreatedAt,
			UpdatedAt: createdTransactionItem.UpdatedAt,
		}

		transactionItemResponses = append(transactionItemResponses, transactionItemResponse)
	}

	transactionResponse := dto.TransactionResponse{
		ID:              createdTransaction.ID,
		UserID:          createdTransaction.UserID,
		Address:         createdTransaction.Address,
		Total:           createdTransaction.Total,
		Status:          createdTransaction.Status,
		CreatedAt:       createdTransaction.CreatedAt,
		UpdatedAt:       createdTransaction.UpdatedAt,
		TransactionItem: transactionItemResponses,
	}

	response.Transactions = append(response.Transactions, transactionResponse)

	return &response, nil
}

func (s *transactionService) GetTransactionByID(id uint) (*dto.TransactionResponse, error) {
	transaction, err := s.repo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	transactionItems, err := s.itemRepo.GetTransactionItemsByTransactionID(transaction.ID)
	if err != nil {
		return nil, err
	}

	var transactionItemResponses []dto.TransactionItemResponse
	for _, transactionItem := range transactionItems {
		transactionItemResponses = append(transactionItemResponses, dto.TransactionItemResponse{
			ID:        transactionItem.ID,
			ProductID: transactionItem.ProductID,
			Quantity:  transactionItem.Quantity,
			Price:     transactionItem.Price,
			CreatedAt: transactionItem.CreatedAt,
			UpdatedAt: transactionItem.UpdatedAt,
		})
	}

	response := dto.TransactionResponse{
		ID:              transaction.ID,
		UserID:          transaction.UserID,
		Address:         transaction.Address,
		Total:           transaction.Total,
		Status:          transaction.Status,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		TransactionItem: transactionItemResponses,
	}

	if transaction.PaymentMethodID != nil {
		response.PaymentMethodID = *transaction.PaymentMethodID
	}

	return &response, nil
}

func (s *transactionService) GetTransactionsByUserID(userID uint) (*dto.ReadTransactionsByUserIDResponse, error) {
	transactions, err := s.repo.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var response dto.ReadTransactionsByUserIDResponse
	for _, transaction := range transactions {
		transactionItems, err := s.itemRepo.GetTransactionItemsByTransactionID(transaction.ID)
		if err != nil {
			return nil, err
		}

		var transactionItemResponses []dto.TransactionItemResponse
		for _, transactionItem := range transactionItems {
			transactionItemResponses = append(transactionItemResponses, dto.TransactionItemResponse{
				ID:        transactionItem.ID,
				ProductID: transactionItem.ProductID,
				Quantity:  transactionItem.Quantity,
				Price:     transactionItem.Price,
				CreatedAt: transactionItem.CreatedAt,
				UpdatedAt: transactionItem.UpdatedAt,
			})
		}

		transactionResponse := dto.TransactionResponse{
			ID:              transaction.ID,
			UserID:          transaction.UserID,
			Address:         transaction.Address,
			Total:           transaction.Total,
			Status:          transaction.Status,
			CreatedAt:       transaction.CreatedAt,
			UpdatedAt:       transaction.UpdatedAt,
			TransactionItem: transactionItemResponses,
		}

		if transaction.PaymentMethodID != nil {
			transactionResponse.PaymentMethodID = *transaction.PaymentMethodID
		}

		response.Transactions = append(response.Transactions, transactionResponse)
	}

	return &response, nil
}

func (s *transactionService) UpdateTransactionStatus(request dto.TransactionPaymentRequest) (*dto.TransactionPaymentResponse, error) {

	transaction, err := s.repo.GetTransactionByID(request.TransactionID)
	if err != nil {
		return nil, err
	}

	transaction.Status = "PAID"
	if request.PaymentMethod != nil {
		transaction.PaymentMethodID = request.PaymentMethod
	}

	updatedTransaction, err := s.repo.UpdateTransactionStatus(*transaction)
	if err != nil {
		return nil, err
	}

	transactionItems, err := s.itemRepo.GetTransactionItemsByTransactionID(request.TransactionID)
	if err != nil {
		return nil, err
	}

	var transactionItemResponses []dto.TransactionItemResponse
	for _, transactionItem := range transactionItems {
		transactionItemResponses = append(transactionItemResponses, dto.TransactionItemResponse{
			ID:        transactionItem.ID,
			ProductID: transactionItem.ProductID,
			Quantity:  transactionItem.Quantity,
			Price:     transactionItem.Price,
			CreatedAt: transactionItem.CreatedAt,
			UpdatedAt: transactionItem.UpdatedAt,
		})
	}

	transactionResponse := dto.TransactionResponse{
		ID:              updatedTransaction.ID,
		UserID:          updatedTransaction.UserID,
		Address:         updatedTransaction.Address,
		Total:           updatedTransaction.Total,
		Status:          updatedTransaction.Status,
		CreatedAt:       updatedTransaction.CreatedAt,
		UpdatedAt:       updatedTransaction.UpdatedAt,
		TransactionItem: transactionItemResponses,
	}

	if updatedTransaction.PaymentMethodID != nil {
		transactionResponse.PaymentMethodID = *updatedTransaction.PaymentMethodID
	}

	response := dto.TransactionPaymentResponse{
		Transactions: []dto.TransactionResponse{transactionResponse},
	}

	return &response, nil
}
