package routes

import (
	"github.com/Andrewalifb/alpha-online-store/product-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(r *fiber.App, productController controller.ProductController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	routesV1.Post("/products", productController.CreateProduct)
	routesV1.Get("/products/:id", productController.GetProductByID)
	routesV1.Get("/products", productController.GetAllProducts)
	routesV1.Get("/categories/:categoryID/products", productController.GetProductsByCategoryID)
	routesV1.Put("/products/:id", productController.UpdateProduct)
}
