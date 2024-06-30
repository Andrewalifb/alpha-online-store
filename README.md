# Alpha Online Store

Alpha Online Store is a simple online store built as a reference for implementing clean architecture using Golang, PostgreSQL, and Redis. It demonstrates how to structure an application using separate services, each with its own responsibilities.

# Some basic features :
• Customer can view product list by product category
• Customer can add product to shopping cart
• Customers can see a list of products that have been added to the shopping cart
• Customer can delete product list in shopping cart
• Customers can checkout and make payment transactions
• Login and register customers

## Services

The application is divided into several services:

- **User Service**: Manages user-related data, including user information and addresses.
- **Product Service**: Handles product-related data, including product information and categories. To enhance performance, frequently accessed product data is cached using Redis, reducing database load and improving response times.
- **Cart Service**: Manages shopping cart data, including carts and cart items.
- **Order Service**: Handles order-related data, including payment methods, transactions, and transaction items.

## Service Endpoints

Here are the endpoints for each service:

| Service              | HTTP Method | Endpoint                                  | Description                              |
|----------------------|-------------|--------------------------------------------|------------------------------------------|
| User Service         | POST        | /api/v1/register                           | Register a new user                      |
| User Service         | POST        | /api/v1/login                              | Login an existing user                   |
| Address Service      | POST        | /api/v1/address                            | Create a new address for a user          |
| Address Service      | PUT         | /api/v1/address/:id                        | Update an existing address for a user    |
| Address Service      | GET         | /api/v1/address/:id                        | Retrieve an existing address for a user  |
| Category Service     | POST        | /api/v1/categories                         | Create a new category                    |
| Category Service     | GET         | /api/v1/categories/:id                     | Retrieve a category by ID                |
| Category Service     | GET         | /api/v1/categories                         | Retrieve all categories                  |
| Product Service      | POST        | /api/v1/products                           | Create a new product                     |
| Product Service      | GET         | /api/v1/products/:id                       | Retrieve a product by ID                 |
| Product Service      | GET         | /api/v1/products                           | Retrieve all products                    |
| Product Service      | GET         | /api/v1/categories/:categoryID/products    | Retrieve products by category ID         |
| Product Service      | PUT         | /api/v1/products/:id                       | Update a product                         |
| Cart Item Service    | POST        | /api/v1/cartItem/:cartID                   | Create a new cart item                   |
| Cart Item Service    | GET         | /api/v1/cartItem/:cartID                   | Retrieve cart items by cart ID           |
| Cart Item Service    | PUT         | /api/v1/cartItem/:cartID                   | Update a cart item                       |
| Cart Item Service    | DELETE      | /api/v1/cartItem/:cartItemID               | Delete a cart item                       |
| Cart Service         | POST        | /api/v1/cart                               | Create a new cart                        |
| Cart Service         | GET         | /api/v1/cart/:userID                       | Retrieve carts by user ID                |
| Cart Service         | DELETE      | /api/v1/cart/:cartID                       | Delete a cart                            |
| Cart Service         | PUT         | /api/v1/cart/:cartID                       | Update a cart                            |
| Cart Service         | POST        | /api/v1/cart/checkout                      | Checkout carts                           |
| Cart Service         | DELETE      | /api/v1/carts/user/:userID                 | Delete carts by user ID                  |
| Payment Method Service | POST      | /api/v1/payment-methods                    | Create a new payment method              |
| Payment Method Service | GET       | /api/v1/payment-methods                    | Retrieve all payment methods             |
| Payment Method Service | GET       | /api/v1/payment-method/:id                 | Retrieve a payment method by ID          |
| Transaction Service  | POST        | /api/v1/transactions                       | Create a new transaction                 |
| Transaction Service  | GET         | /api/v1/transactions/:id                   | Retrieve a transaction by ID             |
| Transaction Service  | GET         | /api/v1/users/:userID/transactions         | Retrieve transactions by user ID         |
| Transaction Service  | PUT         | /api/v1/transactions/status                | Update transaction status               |

for more detailed each rest api end point request and response, i have added on each services open api documentations.

## Database Schema

The PostgreSQL database schema is well-structured and normalized to ensure data integrity and performance. For more details, please refer to the SQL files in the repository.

## Caching with Redis

To optimize the performance of the Product Service, we use Redis as an in-memory data store for caching. This is particularly beneficial for operations that read product data frequently, such as getting product by ID. By storing this data in Redis, we can significantly reduce the latency and the load on the database.

## Docker Images

The Docker images for each service are available on Docker Hub. You can pull these images and run them without having to build them yourself. Here is the Docker Hub profile where you can find the images: [andrewalifbrata](https://hub.docker.com/u/andrewalifbrata)

To pull an image, you can use the following command:

```bash
docker pull andrewalifbrata/service_name:tag
Replace service_name with the name of the service (e.g., user-service, product-service) and tag with the image tag (usually latest).

After pulling the image, you can run it with Docker or use it in a Docker Compose file.

Getting Started
To get started with the Alpha Online Store, follow these steps:

Clone the repository.

Fill up the .env file with your settings (see the Environment Variables section below).

Install the dependencies.

Set up the database and Redis.

Run the application using Docker with the following commands:

bash
Copy code
docker-compose build --no-cache
docker-compose up
Environment Variables
Here are the environment variables you need to set in your .env file:

dotenv
Copy code
# PostgreSQL Connection
SQL_HOST=your_host
SQL_PORT=your_port
SQL_USER=your_user
SQL_DB_NAME=your_database_name
SQL_PASSWORD=your_password

# Running Port
USER_SERVICE_PORT=your_user_service_port
PRODUCT_SERVICE_PORT=your_product_service_port
CART_SERVICE_PORT=your_cart_service_port
ORDER_SERVICE_PORT=your_order_service_port

# Redis Connection
REDIS_HOST=your_redis_host
REDIS_PORT=your_redis_port

# JWT Auth
SECRET_KEY=your_secret_key
Please replace your_host, your_port, your_user, your_database_name, your_password, your_user_service_port, your_product_service_port, your_cart_service_port, your_order_service_port, your_redis_host, your_redis_port, and your_secret_key with your actual settings.