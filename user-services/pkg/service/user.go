package service

import (
	"errors"

	"github.com/Andrewalifb/alpha-online-store/user-services/dto"
	"github.com/Andrewalifb/alpha-online-store/user-services/entity"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/user-services/utils"
)

type UserService interface {
	Register(user dto.UserRequest) (*dto.RegisterResponse, error)
	Login(username, password string) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(user dto.UserRequest) (*dto.RegisterResponse, error) {
	hashedPassword := utils.HashPassword(user.Password)

	userEntity := entity.User{
		Username:       user.Username,
		HashedPassword: hashedPassword,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		Role:           user.Role,
	}

	createUserResult, err := s.userRepo.CreateUser(userEntity)
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

	return &dto.RegisterResponse{
		User: userResponse,
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
