package controller

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/cart-services/dto"
	"github.com/Andrewalifb/alpha-online-store/cart-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/cart-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CartController interface {
	CreateCart(c *fiber.Ctx) error
	GetCartsByUserID(c *fiber.Ctx) error
	DeleteCart(c *fiber.Ctx) error
	UpdateCart(c *fiber.Ctx) error
	CheckoutCarts(c *fiber.Ctx) error
	DeleteCartsByUserID(c *fiber.Ctx) error
}

type cartController struct {
	service service.CartService
}

func NewCartController(service service.CartService) CartController {
	return &cartController{
		service: service,
	}
}

func (ctrl *cartController) CreateCart(c *fiber.Ctx) error {
	var addToCartRequest dto.AddToCartRequest
	if err := c.BodyParser(&addToCartRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(addToCartRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreateCart(addToCartRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to create new cart", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to create new cart", response))
}

func (ctrl *cartController) GetCartsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid user ID", err.Error(), nil))
	}

	responses, err := ctrl.service.GetCartsByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to get carts by user ID", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to get carts by user ID", responses))
}

func (ctrl *cartController) DeleteCart(c *fiber.Ctx) error {
	cartID, err := strconv.ParseUint(c.Params("cartID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid cart ID", err.Error(), nil))
	}

	response, err := ctrl.service.DeleteCart(uint(cartID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to delete cart", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to delete cart", response))
}

func (ctrl *cartController) UpdateCart(c *fiber.Ctx) error {
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

	response, err := ctrl.service.UpdateCart(uint(cartID), cartItemRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to update cart", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to update cart", response))
}

func (ctrl *cartController) CheckoutCarts(c *fiber.Ctx) error {
	var checkoutRequest dto.CheckoutRequest
	if err := c.BodyParser(&checkoutRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(checkoutRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CheckoutCarts(checkoutRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to checkout carts", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to checkout carts", response))
}

func (ctrl *cartController) DeleteCartsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid user ID", err.Error(), nil))
	}

	response, err := ctrl.service.DeleteCartsByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to delete carts", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to delete carts", response))
}
