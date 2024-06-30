package controller

import (
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/order-services/dto"
	"github.com/Andrewalifb/alpha-online-store/order-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/order-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	CreateTransactions(c *fiber.Ctx) error
	GetTransactionByID(c *fiber.Ctx) error
	GetTransactionsByUserID(c *fiber.Ctx) error
	UpdateTransactionStatus(c *fiber.Ctx) error
}

type transactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) TransactionController {
	return &transactionController{
		service: service,
	}
}

func (ctrl *transactionController) CreateTransactions(c *fiber.Ctx) error {
	var request dto.CreateTransactionsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.CreateTransactions(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSACTIONS, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSACTIONS, response))
}

func (ctrl *transactionController) GetTransactionByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetTransactionByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TRANSACTION_BY_ID, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TRANSACTION_BY_ID, response))
}

func (ctrl *transactionController) GetTransactionsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid user ID", err.Error(), nil))
	}

	response, err := ctrl.service.GetTransactionsByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TRANSACTIONS_BY_USER_ID, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TRANSACTIONS_BY_USER_ID, response))
}

func (ctrl *transactionController) UpdateTransactionStatus(c *fiber.Ctx) error {
	var request dto.TransactionPaymentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.UpdateTransactionStatus(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TRANSACTION_STATUS, err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TRANSACTION_STATUS, response))
}
