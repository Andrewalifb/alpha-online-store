package routes

import (
	"github.com/Andrewalifb/alpha-online-store/user-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r *fiber.App, posUserController controller.UserController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	// Routes is not protected by the JWT middleware
	routesV1.Post("/register", posUserController.Register)
	routesV1.Post("/login", posUserController.Login)
}
