package routes

import (
	"github.com/Andrewalifb/alpha-online-store/product-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(r *fiber.App, categoryController controller.CategoryController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	routesV1.Post("/categories", categoryController.CreateCategory)
	routesV1.Get("/categories/:id", categoryController.GetCategoryByID)
	routesV1.Get("/categories", categoryController.GetAllCategories)
}
