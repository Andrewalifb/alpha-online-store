package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Andrewalifb/alpha-online-store/order-services/api/controller"
	"github.com/Andrewalifb/alpha-online-store/order-services/config"
	"github.com/Andrewalifb/alpha-online-store/order-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/order-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/order-services/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}

	// Initialize the database
	dbConfig := config.NewConfigPostgresql()

	// Initialize the repositories
	paymentMethodRepository := repository.NewPaymentMethodRepository(dbConfig.SQLDB)
	transactionRepository := repository.NewTransactionRepository(dbConfig.SQLDB)
	transactionItemRepository := repository.NewTransactionItemRepository(dbConfig.SQLDB)

	// Initialize the services
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepository)
	transactionService := service.NewTransactionService(transactionRepository, transactionItemRepository)

	// Initialize the controllers
	paymentMethodController := controller.NewPaymentMethodController(paymentMethodService)
	transactionController := controller.NewTransactionController(transactionService)

	// Create a new router
	app := fiber.New()

	// Define routes
	routes.PaymentMethodRoutes(app, paymentMethodController)
	routes.TransactionRoutes(app, transactionController)

	port := os.Getenv("ORDER_SERVICE_PORT")

	// Start the server
	log.Println("Server is running on port", port, "...")
	log.Fatal(app.Listen(":" + port))
}
