package service

import (
	"errors"

	"github.com/Andrewalifb/alpha-online-store/user-services/dto"
	"github.com/Andrewalifb/alpha-online-store/user-services/entity"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/user-services/utils"
)

type UserService interface {
	Register(user dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(username, password string) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo    repository.UserRepository
	addressRepo repository.AddressRepository
}

func NewUserService(userRepo repository.UserRepository, addressRepo repository.AddressRepository) UserService {
	return &userService{
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}

func (s *userService) Register(user dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword := utils.HashPassword(user.User.Password)

	userEntity := entity.User{
		Username:       user.User.Username,
		HashedPassword: hashedPassword,
		Email:          user.User.Email,
		FirstName:      user.User.FirstName,
		LastName:       user.User.LastName,
		PhoneNumber:    user.User.PhoneNumber,
		Role:           user.User.Role,
	}

	createUserResult, err := s.userRepo.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}

	addressEntity := entity.Address{
		UserID:        createUserResult.ID,
		Street:        user.Address.Street,
		SubDistrict:   user.Address.SubDistrict,
		District:      user.Address.District,
		CityOrRegency: user.Address.CityOrRegency,
		Province:      user.Address.Province,
		Country:       user.Address.Country,
		PostalCode:    user.Address.PostalCode,
		IsDefault:     false,
	}

	createAddressResult, err := s.addressRepo.CreateAddress(addressEntity)
	if err != nil {
		return nil, err
	}

	userResponse := dto.UserResponse{
		ID:          createUserResult.ID,
		Username:    createUserResult.Username,
		Email:       createUserResult.Email,
		FirstName:   createUserResult.FirstName,
		LastName:    createUserResult.LastName,
		PhoneNumber: createUserResult.PhoneNumber,
		Role:        createUserResult.Role,
		CreatedAt:   createUserResult.CreatedAt,
		UpdatedAt:   createUserResult.UpdatedAt,
	}

	addressResponse := dto.AddressResponse{
		ID:            createAddressResult.ID,
		UserID:        createAddressResult.UserID,
		Street:        createAddressResult.Street,
		SubDistrict:   createAddressResult.SubDistrict,
		District:      createAddressResult.District,
		CityOrRegency: createAddressResult.CityOrRegency,
		Province:      createAddressResult.Province,
		Country:       createAddressResult.Country,
		PostalCode:    createAddressResult.PostalCode,
		IsDefault:     createAddressResult.IsDefault,
		CreatedAt:     createAddressResult.CreatedAt,
		UpdatedAt:     createAddressResult.UpdatedAt,
	}

	return &dto.RegisterResponse{
		User:    userResponse,
		Address: addressResponse,
	}, nil
}

func (s *userService) Login(username, password string) (*dto.LoginResponse, error) {

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.HashedPassword) {
		return nil, errors.New("invalid password")
	}

	token := utils.GenerateJWT(user.Username)

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
