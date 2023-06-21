package main

import (
	"log"

	"github.com/atifali-pm/go-fiber-api/database"
	"github.com/atifali-pm/go-fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {

	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	// users routes
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	// product routes
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/product/:id", routes.GetProduct)
	app.Put("/api/product/:id", routes.UpdateProduct)
	app.Delete("/api/product/:id", routes.DeleteProduct)

	// order routes
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/order/:id", routes.GetOrder)
	app.Get("/api/orders", routes.GetOrders)
}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
