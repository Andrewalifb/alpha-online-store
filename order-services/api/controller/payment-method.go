package controller

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/order-services/dto"
	"github.com/Andrewalifb/alpha-online-store/order-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/order-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PaymentMethodController interface {
	CreatePaymentMethod(c *fiber.Ctx) error
	GetAllPaymentMethods(c *fiber.Ctx) error
	GetPaymentMethodByID(c *fiber.Ctx) error
}

type paymentMethodController struct {
	service service.PaymentMethodService
}

func NewPaymentMethodController(service service.PaymentMethodService) PaymentMethodController {
	return &paymentMethodController{
		service: service,
	}
}

func (ctrl *paymentMethodController) CreatePaymentMethod(c *fiber.Ctx) error {
	var request dto.CreatePaymentMethodRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreatePaymentMethod(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PAYMENT_METHOD, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PAYMENT_METHOD, response))
}

func (ctrl *paymentMethodController) GetAllPaymentMethods(c *fiber.Ctx) error {
	response, err := ctrl.service.GetAllPaymentMethods()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_PAYMENT_METHODS, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_PAYMENT_METHODS, response))
}

func (ctrl *paymentMethodController) GetPaymentMethodByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetPaymentMethodByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PAYMENT_METHOD_BY_ID, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PAYMENT_METHOD_BY_ID, response))
}
