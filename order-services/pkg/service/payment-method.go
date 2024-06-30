package service

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/dto"
	"github.com/Andrewalifb/alpha-online-store/order-services/entity"
	"github.com/Andrewalifb/alpha-online-store/order-services/pkg/repository"
)

type PaymentMethodService interface {
	CreatePaymentMethod(request dto.CreatePaymentMethodRequest) (*dto.CreatePaymentMethodResponse, error)
	GetAllPaymentMethods() (*dto.ReadAllPaymentMethodsResponse, error)
	GetPaymentMethodByID(id uint) (*dto.ReadPaymentMethodResponse, error)
}

type paymentMethodService struct {
	repo repository.PaymentMethodRepository
}

func NewPaymentMethodService(repo repository.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{
		repo: repo,
	}
}

func (s *paymentMethodService) CreatePaymentMethod(request dto.CreatePaymentMethodRequest) (*dto.CreatePaymentMethodResponse, error) {
	paymentMethod := entity.PaymentMethod{
		Method: request.Method,
	}
	resultPaymentMethod, err := s.repo.CreatePaymentMethod(paymentMethod)
	if err != nil {
		return nil, err
	}
	response := dto.CreatePaymentMethodResponse{
		ID:        resultPaymentMethod.ID,
		Method:    resultPaymentMethod.Method,
		CreatedAt: resultPaymentMethod.CreatedAt,
		UpdatedAt: resultPaymentMethod.UpdatedAt,
	}
	return &response, nil
}

func (s *paymentMethodService) GetAllPaymentMethods() (*dto.ReadAllPaymentMethodsResponse, error) {
	paymentMethods, err := s.repo.GetAllPaymentMethods()
	if err != nil {
		return nil, err
	}
	var response dto.ReadAllPaymentMethodsResponse
	for _, paymentMethod := range paymentMethods {
		response.PaymentMethods = append(response.PaymentMethods, dto.CreatePaymentMethodResponse{
			ID:        paymentMethod.ID,
			Method:    paymentMethod.Method,
			CreatedAt: paymentMethod.CreatedAt,
			UpdatedAt: paymentMethod.UpdatedAt,
		})
	}
	return &response, nil
}

func (s *paymentMethodService) GetPaymentMethodByID(id uint) (*dto.ReadPaymentMethodResponse, error) {
	paymentMethod, err := s.repo.GetPaymentMethodByID(id)
	if err != nil {
		return nil, err
	}
	response := dto.ReadPaymentMethodResponse{
		ID:        paymentMethod.ID,
		Method:    paymentMethod.Method,
		CreatedAt: paymentMethod.CreatedAt,
		UpdatedAt: paymentMethod.UpdatedAt,
	}
	return &response, nil
}
