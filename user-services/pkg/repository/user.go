package repository

import (
	"github.com/Andrewalifb/alpha-online-store/user-services/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user entity.User) (*entity.User, error) {
	result := r.db.Create(&user)
	return &user, result.Error
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}
