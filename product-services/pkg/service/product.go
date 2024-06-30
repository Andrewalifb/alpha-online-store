package service

import (
	"github.com/Andrewalifb/alpha-online-store/product-services/dto"
	"github.com/Andrewalifb/alpha-online-store/product-services/entity"
	"github.com/Andrewalifb/alpha-online-store/product-services/pkg/repository"
)

type ProductService interface {
	CreateProduct(product dto.CreateProductRequest) (*dto.CreateProductResponse, error)
	GetProductByID(id uint) (*dto.ReadProductByIDResponse, error)
	GetAllProducts() (*dto.ReadAllProductsResponse, error)
	GetProductsByCategoryID(categoryID uint) (*dto.ReadProductsByCategoryIDResponse, error)
	UpdateProduct(id uint, product dto.UpdateProductRequest) (*dto.UpdateProductResponse, error)
}

type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productService) CreateProduct(product dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	productEntity := entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Inventory:   product.Inventory,
		ImageURL:    product.ImageURL,
		CategoryID:  product.CategoryID,
		Status:      product.Status,
	}

	createProductResult, err := s.productRepo.CreateProduct(productEntity)
	if err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.GetCategoryByID(createProductResult.CategoryID)
	if err != nil {
		return nil, err
	}

	return &dto.CreateProductResponse{
		ID:          createProductResult.ID,
		Name:        createProductResult.Name,
		Description: createProductResult.Description,
		Price:       createProductResult.Price,
		Inventory:   createProductResult.Inventory,
		ImageURL:    createProductResult.ImageURL,
		Category: dto.CategoryDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
		Status:    createProductResult.Status,
		CreatedAt: createProductResult.CreatedAt,
		UpdatedAt: createProductResult.UpdatedAt,
	}, nil
}

func (s *productService) GetProductByID(id uint) (*dto.ReadProductByIDResponse, error) {
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.GetCategoryByID(product.CategoryID)
	if err != nil {
		return nil, err
	}

	return &dto.ReadProductByIDResponse{
		Product: dto.ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Inventory:   product.Inventory,
			ImageURL:    product.ImageURL,
			Category: dto.CategoryDTO{
				ID:          category.ID,
				Name:        category.Name,
				Description: category.Description,
			},
			Status: product.Status,
		},
	}, nil
}

func (s *productService) GetAllProducts() (*dto.ReadAllProductsResponse, error) {
	products, err := s.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var productResponses []dto.ProductDTO
	for _, product := range products {
		category, err := s.categoryRepo.GetCategoryByID(product.CategoryID)
		if err != nil {
			return nil, err
		}

		productResponses = append(productResponses, dto.ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Inventory:   product.Inventory,
			ImageURL:    product.ImageURL,
			Category: dto.CategoryDTO{
				ID:          category.ID,
				Name:        category.Name,
				Description: category.Description,
			},
			Status: product.Status,
		})
	}

	return &dto.ReadAllProductsResponse{
		Products: productResponses,
	}, nil
}

func (s *productService) GetProductsByCategoryID(categoryID uint) (*dto.ReadProductsByCategoryIDResponse, error) {
	products, err := s.productRepo.GetProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	var productResponses []dto.ProductDTO
	for _, product := range products {
		category, err := s.categoryRepo.GetCategoryByID(product.CategoryID)
		if err != nil {
			return nil, err
		}

		productResponses = append(productResponses, dto.ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Inventory:   product.Inventory,
			ImageURL:    product.ImageURL,
			Category: dto.CategoryDTO{
				ID:          category.ID,
				Name:        category.Name,
				Description: category.Description,
			},
			Status: product.Status,
		})
	}

	return &dto.ReadProductsByCategoryIDResponse{
		Products: productResponses,
	}, nil
}

func (s *productService) UpdateProduct(id uint, product dto.UpdateProductRequest) (*dto.UpdateProductResponse, error) {
	existingProduct, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.Inventory = product.Inventory
	existingProduct.ImageURL = product.ImageURL
	existingProduct.CategoryID = product.CategoryID
	existingProduct.Status = product.Status

	updateProductResult, err := s.productRepo.UpdateProduct(*existingProduct)
	if err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.GetCategoryByID(updateProductResult.CategoryID)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateProductResponse{
		ID:          updateProductResult.ID,
		Name:        updateProductResult.Name,
		Description: updateProductResult.Description,
		Price:       updateProductResult.Price,
		Inventory:   updateProductResult.Inventory,
		ImageURL:    updateProductResult.ImageURL,
		Category: dto.CategoryDTO{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
		Status:    updateProductResult.Status,
		UpdatedAt: updateProductResult.UpdatedAt,
	}, nil
}
