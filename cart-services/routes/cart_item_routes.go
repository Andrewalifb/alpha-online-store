package routes

import (
	"github.com/Andrewalifb/alpha-online-store/cart-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func CartItemRoutes(r *fiber.App, cartItemController controller.CartItemController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	// Routes is not protected by the JWT middleware
	routesV1.Post("/cartItem/:cartID", cartItemController.CreateCartItem)
	routesV1.Get("/cartItem/:cartID", cartItemController.GetCartItemsByCartID)
	routesV1.Put("/cartItem/:cartID", cartItemController.UpdateCartItem)
	routesV1.Delete("/cartItem/:cartItemID", cartItemController.DeleteCartItem)
}
