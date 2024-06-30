package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Andrewalifb/alpha-online-store/cart-services/api/controller"
	"github.com/Andrewalifb/alpha-online-store/cart-services/config"
	"github.com/Andrewalifb/alpha-online-store/cart-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/cart-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/cart-services/routes"
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
	cartRepository := repository.NewCartRepository(dbConfig.SQLDB)
	cartItemRepository := repository.NewCartItemRepository(dbConfig.SQLDB)

	// Initialize the services
	cartService := service.NewCartService(cartRepository, cartItemRepository)
	cartItemService := service.NewCartItemService(cartItemRepository)

	// Initialize the controllers
	cartController := controller.NewCartController(cartService)
	cartItemController := controller.NewCartItemController(cartItemService)

	// Create a new router
	app := fiber.New()

	// Define routes
	routes.CartRoutes(app, cartController)
	routes.CartItemRoutes(app, cartItemController)

	port := os.Getenv("CART_SERVICE_PORT")

	// Start the server
	log.Println("Server is running on port", port, "...")
	log.Fatal(app.Listen(":" + port))
}
