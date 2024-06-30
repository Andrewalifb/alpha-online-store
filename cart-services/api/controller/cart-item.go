package controller

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/cart-services/dto"
	"github.com/Andrewalifb/alpha-online-store/cart-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/cart-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CartItemController interface {
	CreateCartItem(c *fiber.Ctx) error
	GetCartItemsByCartID(c *fiber.Ctx) error
	UpdateCartItem(c *fiber.Ctx) error
	DeleteCartItem(c *fiber.Ctx) error
}

type cartItemController struct {
	service service.CartItemService
}

func NewCartItemController(service service.CartItemService) CartItemController {
	return &cartItemController{
		service: service,
	}
}

func (ctrl *cartItemController) CreateCartItem(c *fiber.Ctx) error {
	cartID, err := strconv.ParseUint(c.Params("cartID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid cart ID", err.Error(), nil))
	}

	var cartItemRequest dto.CartItemRequest
	if err := c.BodyParser(&cartItemRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(cartItemRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreateCartItem(uint(cartID), cartItemRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to create new cart item", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to create new cart item", response))
}

func (ctrl *cartItemController) GetCartItemsByCartID(c *fiber.Ctx) error {
	cartID, err := strconv.ParseUint(c.Params("cartID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid cart ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetCartItemsByCartID(uint(cartID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get cart items by cart ID", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get cart items by cart ID", response))
}

func (ctrl *cartItemController) UpdateCartItem(c *fiber.Ctx) error {
	cartID, err := strconv.ParseUint(c.Params("cartID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid cart ID", err.Error(), nil))
	}

	var cartItemRequest dto.CartItemRequest
	if err := c.BodyParser(&cartItemRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(cartItemRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.UpdateCartItem(uint(cartID), cartItemRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to update cart item", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to update cart item", response))
}

func (ctrl *cartItemController) DeleteCartItem(c *fiber.Ctx) error {
	cartItemID, err := strconv.ParseUint(c.Params("cartItemID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid cart item ID", err.Error(), nil))
	}

	response, err := ctrl.service.DeleteCartItem(uint(cartItemID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to delete cart item", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to delete cart item", response))
}
