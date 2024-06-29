package controller

import (
	"github.com/Andrewalifb/alpha-online-store/user-services/dto"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/user-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AddressController interface {
	CreateAddress(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	ReadAddress(c *fiber.Ctx) error
}

type addressController struct {
	service service.AddressService
}

func NewAddressController(service service.AddressService) AddressController {
	return &addressController{
		service: service,
	}
}

func (ctrl *addressController) CreateAddress(c *fiber.Ctx) error {
	var createAddressRequest dto.CreateAddressRequest
	if err := c.BodyParser(&createAddressRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(createAddressRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreateAddress(createAddressRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to create new address", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to create new address", response))
}

func (ctrl *addressController) UpdateAddress(c *fiber.Ctx) error {
	var updateAddressRequest dto.UpdateAddressRequest
	if err := c.BodyParser(&updateAddressRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(updateAddressRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	addressID := c.Params("id")
	response, err := ctrl.service.UpdateAddress(addressID, updateAddressRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to update address", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to update address", response))
}

func (ctrl *addressController) ReadAddress(c *fiber.Ctx) error {
	addressID := c.Params("id")
	response, err := ctrl.service.ReadAddress(addressID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to read address", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to read address", response))
}
