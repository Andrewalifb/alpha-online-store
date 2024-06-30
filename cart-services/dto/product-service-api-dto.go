package dto

import "time"

type ReadProductApiResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ProductDTO `json:"product"`
	} `json:"data"`
}

type UpdateProductApiResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		UpdateProductResponse `json:"product"`
	} `json:"data"`
}

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
