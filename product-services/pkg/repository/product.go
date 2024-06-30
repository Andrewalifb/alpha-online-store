package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/Andrewalifb/alpha-online-store/product-services/entity"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product entity.Product) (*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	GetProductsByCategoryID(categoryID uint) ([]*entity.Product, error)
	UpdateProduct(product entity.Product) (*entity.Product, error)
}

type productRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProductRepository(db *gorm.DB, redis *redis.Client) ProductRepository {
	return &productRepository{
		db:    db,
		redis: redis,
	}
}

func (r *productRepository) CreateProduct(product entity.Product) (*entity.Product, error) {
	result := r.db.Create(&product)
	return &product, result.Error
}

func (r *productRepository) GetProductByID(id uint) (*entity.Product, error) {
	productData, err := r.redis.Get(context.Background(), strconv.FormatUint(uint64(id), 10)).Result()

	if err == redis.Nil {

		var product entity.Product
		if err := r.db.First(&product, id).Error; err != nil {
			return nil, err
		}
		productData, err := json.Marshal(product)
		if err != nil {
			return nil, err
		}
		err = r.redis.Set(context.Background(), strconv.FormatUint(uint64(id), 10), productData, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}
		return &product, nil
	} else if err != nil {
		return nil, err
	}
	var product entity.Product
	err = json.Unmarshal([]byte(productData), &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAllProducts() ([]*entity.Product, error) {
	var products []*entity.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *productRepository) GetProductsByCategoryID(categoryID uint) ([]*entity.Product, error) {
	var products []*entity.Product
	result := r.db.Where("category_id = ?", categoryID).Find(&products)
	return products, result.Error
}

func (r *productRepository) UpdateProduct(product entity.Product) (*entity.Product, error) {
	result := r.db.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	err := r.redis.Del(context.Background(), strconv.FormatUint(uint64(product.ID), 10)).Err()
	if err != nil {
		return nil, err
	}
	return &product, nil
}
