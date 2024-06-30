package routes

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func PaymentMethodRoutes(r *fiber.App, paymentMethodController controller.PaymentMethodController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	routesV1.Post("/payment-methods", paymentMethodController.CreatePaymentMethod)
	routesV1.Get("/payment-methods", paymentMethodController.GetAllPaymentMethods)
	routesV1.Get("/payment-method/:id", paymentMethodController.GetPaymentMethodByID)
}
