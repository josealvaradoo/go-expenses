package router

import (
	"github.com/gofiber/fiber"
	"github.com/josealvaradoo/expenses/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Users
	users := v1.Group("/users")

	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Post("/", handlers.NewUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)

	// Currencies
	currencies := v1.Group("/currencies")

	currencies.Get("/", handlers.GetCurrencies)
	currencies.Get("/:id", handlers.GetCurrency)
	currencies.Post("/", handlers.NewCurrency)
	currencies.Put("/:id", handlers.UpdateCurrency)
	currencies.Delete("/:id", handlers.DeleteCurrency)
}
