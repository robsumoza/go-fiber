package router

import (
	"github.com/robsumoza/go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

// Setup up API routes
func SetupRoutes(app *fiber.App) {

	app.Get("/health", handlers.HandleHealthCheck)
}
