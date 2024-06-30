package repository

import (
	"github.com/Andrewalifb/alpha-online-store/product-services/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category entity.Category) (*entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	GetAllCategories() ([]*entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) CreateCategory(category entity.Category) (*entity.Category, error) {
	result := r.db.Create(&category)
	return &category, result.Error
}

func (r *categoryRepository) GetCategoryByID(id uint) (*entity.Category, error) {
	var category entity.Category
	result := r.db.First(&category, id)
	return &category, result.Error
}

func (r *categoryRepository) GetAllCategories() ([]*entity.Category, error) {
	var categories []*entity.Category
	result := r.db.Find(&categories)
	return categories, result.Error
}
