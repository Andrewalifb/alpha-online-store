package controller

import (
	"github.com/Andrewalifb/alpha-online-store/user-services/dto"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/user-services/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (ctrl *userController) Register(c *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.Register(registerRequest.User)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to register new user", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to register new user", response))
}

func (ctrl *userController) Login(c *fiber.Ctx) error {
	var loginRequest dto.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BuildResponseFailed("Invalid request body", err.Error(), nil))
	}

	response, err := ctrl.service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BuildResponseFailed("Failed to login", err.Error(), nil))
	}

	return c.JSON(utils.BuildResponseSuccess("Success to login", response))
}
