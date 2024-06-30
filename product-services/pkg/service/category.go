package service

import (
	"github.com/Andrewalifb/alpha-online-store/product-services/dto"
	"github.com/Andrewalifb/alpha-online-store/product-services/entity"
	"github.com/Andrewalifb/alpha-online-store/product-services/pkg/repository"
)

type CategoryService interface {
	CreateCategory(category dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error)
	GetAllCategories() (*dto.ReadAllCategoriesResponse, error)
	GetCategoryByID(id uint) (*dto.CategoryDTO, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) CreateCategory(category dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error) {
	categoryEntity := entity.Category{
		Name:        category.Name,
		Description: category.Description,
	}

	createCategoryResult, err := s.categoryRepo.CreateCategory(categoryEntity)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCategoryResponse{
		ID:          createCategoryResult.ID,
		Name:        createCategoryResult.Name,
		Description: createCategoryResult.Description,
		CreatedAt:   createCategoryResult.CreatedAt,
		UpdatedAt:   createCategoryResult.UpdatedAt,
	}, nil
}

func (s *categoryService) GetAllCategories() (*dto.ReadAllCategoriesResponse, error) {
	categories, err := s.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	var categoryResponses []dto.CreateCategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.CreateCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		})
	}

	return &dto.ReadAllCategoriesResponse{
		Categories: categoryResponses,
	}, nil
}

func (s *categoryService) GetCategoryByID(id uint) (*dto.CategoryDTO, error) {
	category, err := s.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
