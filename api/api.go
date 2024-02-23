package api

import (
	"github.com/robsumoza/go-fiber/router"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	router.SetupRoutes((app))

	app.Listen(":8080")

}
