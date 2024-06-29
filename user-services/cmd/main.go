package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Andrewalifb/alpha-online-store/user-services/api/controller"
	"github.com/Andrewalifb/alpha-online-store/user-services/config"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/user-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/user-services/routes"
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
	userRepository := repository.NewUserRepository(dbConfig.SQLDB)
	addressRepository := repository.NewAddressRepository(dbConfig.SQLDB)

	// Initialize the services
	userService := service.NewUserService(userRepository, addressRepository)
	addressService := service.NewAddressService(addressRepository)

	// Initialize the controllers
	userController := controller.NewUserController(userService)
	addressController := controller.NewAddressController(addressService)

	// Create a new router
	app := fiber.New()

	// Define routes
	routes.UserRoutes(app, userController)
	routes.AddressRoutes(app, addressController)

	port := os.Getenv("USER_SERVICE_PORT")

	// Start the server
	log.Println("Server is running on port", port, "...")
	log.Fatal(app.Listen(":" + port))
}
