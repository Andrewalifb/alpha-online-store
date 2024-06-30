package dto

import "time"

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ReadAllCategoriesResponse struct {
	Categories []CreateCategoryResponse `json:"categories"`
}
