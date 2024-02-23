package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/robsumoza/go-fiber/docs"
)

// Setup swagger
func AddSwaggerRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
