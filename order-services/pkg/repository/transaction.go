package repository

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (*entity.Transaction, error)
	GetTransactionByID(id uint) (*entity.Transaction, error)
	GetTransactionsByUserID(userID uint) ([]entity.Transaction, error)
	UpdateTransactionStatus(transaction entity.Transaction) (*entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreateTransaction(transaction entity.Transaction) (*entity.Transaction, error) {
	result := r.db.Create(&transaction)
	return &transaction, result.Error
}

func (r *transactionRepository) GetTransactionByID(id uint) (*entity.Transaction, error) {
	var transaction entity.Transaction
	result := r.db.First(&transaction, id)
	return &transaction, result.Error
}

func (r *transactionRepository) GetTransactionsByUserID(userID uint) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	result := r.db.Where("user_id = ?", userID).Find(&transactions)
	return transactions, result.Error
}

func (r *transactionRepository) UpdateTransactionStatus(transaction entity.Transaction) (*entity.Transaction, error) {
	result := r.db.Model(&transaction).Update("status", transaction.Status)
	return &transaction, result.Error
}
