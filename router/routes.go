package router

import (
	"github.com/robsumoza/go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

// Setup up API routes
func SetupRoutes(app *fiber.App) {

	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/todos")
	todos.Get("/", handlers.HandleAllTodos)
	todos.Post("/", handlers.HandleCreateTodo)
	todos.Put("/:id", handlers.HandleUpdateTodo)
	todos.Get("/:id", handlers.HandleGetOneTodo)
	todos.Delete("/:id", handlers.HandleDeleteTodo)
}
