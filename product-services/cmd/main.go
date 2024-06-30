package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Andrewalifb/alpha-online-store/product-services/api/controller"
	"github.com/Andrewalifb/alpha-online-store/product-services/config"
	"github.com/Andrewalifb/alpha-online-store/product-services/pkg/repository"
	"github.com/Andrewalifb/alpha-online-store/product-services/pkg/service"
	"github.com/Andrewalifb/alpha-online-store/product-services/routes"
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
	rdConfig := config.NewConfigRedis()

	// Initialize the repositories
	categoryRepository := repository.NewCategoryRepository(dbConfig.SQLDB)
	productRepository := repository.NewProductRepository(dbConfig.SQLDB, rdConfig.RedisDB)

	// Initialize the services
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository, categoryRepository)

	// Initialize the controllers
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)

	// Create a new router
	app := fiber.New()

	// Define routes
	routes.CategoryRoutes(app, categoryController)
	routes.ProductRoutes(app, productController)

	port := os.Getenv("PRODUCT_SERVICE_PORT")

	// Start the server
	log.Println("Server is running on port", port, "...")
	log.Fatal(app.Listen(":" + port))
}
