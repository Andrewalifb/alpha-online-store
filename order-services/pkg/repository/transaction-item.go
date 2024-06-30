package repository

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/entity"
	"gorm.io/gorm"
)

type TransactionItemRepository interface {
	CreateTransactionItem(transactionItem entity.TransactionItem) (*entity.TransactionItem, error)
	GetTransactionItemsByTransactionID(transactionID uint) ([]entity.TransactionItem, error)
}

type transactionItemRepository struct {
	db *gorm.DB
}

func NewTransactionItemRepository(db *gorm.DB) TransactionItemRepository {
	return &transactionItemRepository{
		db: db,
	}
}

func (r *transactionItemRepository) CreateTransactionItem(transactionItem entity.TransactionItem) (*entity.TransactionItem, error) {
	result := r.db.Create(&transactionItem)
	return &transactionItem, result.Error
}

func (r *transactionItemRepository) GetTransactionItemsByTransactionID(transactionID uint) ([]entity.TransactionItem, error) {
	var transactionItems []entity.TransactionItem
	result := r.db.Where("transaction_id = ?", transactionID).Find(&transactionItems)
	return transactionItems, result.Error
}
