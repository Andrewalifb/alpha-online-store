package service

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/user-services/dto"
	"github.com/Andrewalifb/alpha-online-store/user-services/entity"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/repository"
)

type AddressService interface {
	CreateAddress(request dto.CreateAddressRequest) (*dto.CreateAddressResponse, error)
	UpdateAddress(addressID string, request dto.UpdateAddressRequest) (*dto.UpdateAddressResponse, error)
	ReadAddress(addressID string) (*dto.ReadAddressResponse, error)
}

type addressService struct {
	addressRepo repository.AddressRepository
}

func NewAddressService(addressRepo repository.AddressRepository) AddressService {
	return &addressService{
		addressRepo: addressRepo,
	}
}

func (s *addressService) CreateAddress(request dto.CreateAddressRequest) (*dto.CreateAddressResponse, error) {

	addressEntity := entity.Address{
		UserID:        request.Address.UserID,
		Street:        request.Address.Street,
		SubDistrict:   request.Address.SubDistrict,
		District:      request.Address.District,
		CityOrRegency: request.Address.CityOrRegency,
		Province:      request.Address.Province,
		Country:       request.Address.Country,
		PostalCode:    request.Address.PostalCode,
		IsDefault:     request.Address.IsDefault,
	}

	addressEntityPtr, err := s.addressRepo.CreateAddress(addressEntity)
	if err != nil {
		return nil, err
	}

	response := dto.AddressResponse{
		ID:            addressEntityPtr.ID,
		UserID:        addressEntityPtr.UserID,
		Street:        addressEntityPtr.Street,
		SubDistrict:   addressEntityPtr.SubDistrict,
		District:      addressEntityPtr.District,
		CityOrRegency: addressEntityPtr.CityOrRegency,
		Province:      addressEntityPtr.Province,
		Country:       addressEntityPtr.Country,
		PostalCode:    addressEntityPtr.PostalCode,
		IsDefault:     addressEntityPtr.IsDefault,
		CreatedAt:     addressEntityPtr.CreatedAt,
		UpdatedAt:     addressEntityPtr.UpdatedAt,
	}

	return &dto.CreateAddressResponse{
		Address: response,
	}, nil
}

func (s *addressService) UpdateAddress(addressID string, request dto.UpdateAddressRequest) (*dto.UpdateAddressResponse, error) {

	resultAddressID, err := strconv.ParseUint(addressID, 10, 0)
	if err != nil {
		return nil, err
	}

	convertedAddressID := uint(resultAddressID)

	address, err := s.addressRepo.ReadAddress(convertedAddressID)
	if err != nil {
		return nil, err
	}

	addressEntity := entity.Address{
		UserID:        address.UserID,
		Street:        request.Address.Street,
		SubDistrict:   request.Address.SubDistrict,
		District:      request.Address.District,
		CityOrRegency: request.Address.CityOrRegency,
		Province:      request.Address.Province,
		Country:       request.Address.Country,
		PostalCode:    request.Address.PostalCode,
		IsDefault:     request.Address.IsDefault,
	}

	addressEntityPtr, err := s.addressRepo.UpdateUserAddress(addressID, addressEntity)
	if err != nil {
		return nil, err
	}

	response := dto.AddressResponse{
		ID:            addressEntityPtr.ID,
		UserID:        addressEntityPtr.UserID,
		Street:        addressEntityPtr.Street,
		SubDistrict:   addressEntityPtr.SubDistrict,
		District:      addressEntityPtr.District,
		CityOrRegency: addressEntityPtr.CityOrRegency,
		Province:      addressEntityPtr.Province,
		Country:       addressEntityPtr.Country,
		PostalCode:    addressEntityPtr.PostalCode,
		IsDefault:     addressEntityPtr.IsDefault,
		CreatedAt:     addressEntityPtr.CreatedAt,
		UpdatedAt:     addressEntityPtr.UpdatedAt,
	}

	return &dto.UpdateAddressResponse{
		Address: response,
	}, nil
}

func (s *addressService) ReadAddress(addressID string) (*dto.ReadAddressResponse, error) {

	resultAddressID, err := strconv.ParseUint(addressID, 10, 0)
	if err != nil {
		return nil, err
	}

	convertedAddressID := uint(resultAddressID)

	address, err := s.addressRepo.ReadAddress(convertedAddressID)
	if err != nil {
		return nil, err
	}

	response := dto.AddressResponse{
		ID:            address.ID,
		UserID:        address.UserID,
		Street:        address.Street,
		SubDistrict:   address.SubDistrict,
		District:      address.District,
		CityOrRegency: address.CityOrRegency,
		Province:      address.Province,
		Country:       address.Country,
		PostalCode:    address.PostalCode,
		IsDefault:     address.IsDefault,
		CreatedAt:     address.CreatedAt,
		UpdatedAt:     address.UpdatedAt,
	}

	return &dto.ReadAddressResponse{
		Address: response,
	}, nil
}
