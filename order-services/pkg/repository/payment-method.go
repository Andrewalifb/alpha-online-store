package repository

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/entity"
	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	CreatePaymentMethod(paymentMethod entity.PaymentMethod) (*entity.PaymentMethod, error)
	GetAllPaymentMethods() ([]entity.PaymentMethod, error)
	GetPaymentMethodByID(id uint) (*entity.PaymentMethod, error)
}

type paymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}

func (r *paymentMethodRepository) CreatePaymentMethod(paymentMethod entity.PaymentMethod) (*entity.PaymentMethod, error) {
	result := r.db.Create(&paymentMethod)
	return &paymentMethod, result.Error
}

func (r *paymentMethodRepository) GetAllPaymentMethods() ([]entity.PaymentMethod, error) {
	var paymentMethods []entity.PaymentMethod
	result := r.db.Find(&paymentMethods)
	return paymentMethods, result.Error
}

func (r *paymentMethodRepository) GetPaymentMethodByID(id uint) (*entity.PaymentMethod, error) {
	var paymentMethod entity.PaymentMethod
	result := r.db.First(&paymentMethod, id)
	return &paymentMethod, result.Error
}
