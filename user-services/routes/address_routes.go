package routes

import (
	"github.com/Andrewalifb/alpha-online-store/user-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(r *fiber.App, posUserController controller.AddressController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	// Routes is not protected by the JWT middleware
	routesV1.Post("/address", posUserController.CreateAddress)
	routesV1.Put("/address/:id", posUserController.UpdateAddress)
	routesV1.Get("/address/:id", posUserController.ReadAddress)
}
