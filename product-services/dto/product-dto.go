package dto

import "time"

type CategoryDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductDTO struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Inventory   int         `json:"inventory"`
	ImageURL    string      `json:"image_url"`
	Category    CategoryDTO `json:"category"`
	Status      string      `json:"status"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Inventory   int     `json:"inventory"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id"`
	Status      string  `json:"status"`
}

type CreateProductResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Inventory   int         `json:"inventory"`
	ImageURL    string      `json:"image_url"`
	Category    CategoryDTO `json:"category"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type ReadProductByIDResponse struct {
	Product ProductDTO `json:"product"`
}

type ReadAllProductsResponse struct {
	Products []ProductDTO `json:"products"`
}

type ReadProductsByCategoryIDResponse struct {
	Products []ProductDTO `json:"products"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Inventory   int     `json:"inventory"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id"`
	Status      string  `json:"status"`
}

type UpdateProductResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Inventory   int         `json:"inventory"`
	ImageURL    string      `json:"image_url"`
	Category    CategoryDTO `json:"category"`
	Status      string      `json:"status"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
