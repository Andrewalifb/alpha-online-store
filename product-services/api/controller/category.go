package controller

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/product-services/dto"
	"github.com/Andrewalifb/alpha-online-store/product-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/product-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	CreateCategory(c *fiber.Ctx) error
	GetAllCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{
		service: service,
	}
}

func (ctrl *categoryController) CreateCategory(c *fiber.Ctx) error {
	var createCategoryRequest dto.CreateCategoryRequest
	if err := c.BodyParser(&createCategoryRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(createCategoryRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreateCategory(createCategoryRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to create new category", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to create new category", response))
}

func (ctrl *categoryController) GetAllCategories(c *fiber.Ctx) error {
	response, err := ctrl.service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get all categories", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get all categories", response))
}

func (ctrl *categoryController) GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid category ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetCategoryByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get category by ID", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get category by ID", response))
}
