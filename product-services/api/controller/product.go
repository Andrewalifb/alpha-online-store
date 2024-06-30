package controller

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/product-services/dto"
	"github.com/Andrewalifb/alpha-online-store/product-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/product-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	GetProductsByCategoryID(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (ctrl *productController) CreateProduct(c *fiber.Ctx) error {
	var createProductRequest dto.CreateProductRequest
	if err := c.BodyParser(&createProductRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(createProductRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreateProduct(createProductRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to create new product", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to create new product", response))
}

func (ctrl *productController) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid product ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get product by ID", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get product by ID", response))
}

func (ctrl *productController) GetAllProducts(c *fiber.Ctx) error {
	response, err := ctrl.service.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get all products", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get all products", response))
}

func (ctrl *productController) GetProductsByCategoryID(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("categoryID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid category ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetProductsByCategoryID(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get products by category ID", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get products by category ID", response))
}

func (ctrl *productController) UpdateProduct(c *fiber.Ctx) error {
	var updateProductRequest dto.UpdateProductRequest
	if err := c.BodyParser(&updateProductRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	productID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid product ID", err.Error(), nil))
	}

	response, err := ctrl.service.UpdateProduct(uint(productID), updateProductRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to update product", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to update product", response))
}
