package routes

import (
	"github.com/Andrewalifb/alpha-online-store/cart-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func CartRoutes(r *fiber.App, cartController controller.CartController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	// Routes is not protected by the JWT middleware
	routesV1.Post("/cart", cartController.CreateCart)
	routesV1.Get("/cart/:userID", cartController.GetCartsByUserID)
	routesV1.Delete("/cart/:cartID", cartController.DeleteCart)
	routesV1.Put("/cart/:cartID", cartController.UpdateCart)
	routesV1.Post("/cart/checkout", cartController.CheckoutCarts)
	routesV1.Delete("/carts/user/:userID", cartController.DeleteCartsByUserID)
}
