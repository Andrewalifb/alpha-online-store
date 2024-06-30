package routes

import (
	"github.com/Andrewalifb/alpha-online-store/order-services/api/controller"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(r *fiber.App, transactionController controller.TransactionController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1")

	routesV1.Post("/transactions", transactionController.CreateTransactions)
	routesV1.Get("/transactions/:id", transactionController.GetTransactionByID)
	routesV1.Get("/users/:userID/transactions", transactionController.GetTransactionsByUserID)
	routesV1.Put("/transactions/status", transactionController.UpdateTransactionStatus)
}
