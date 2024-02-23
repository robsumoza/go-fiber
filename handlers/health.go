package handlers

import "github.com/gofiber/fiber/v2"

// @Summary Show the health status of the API.
// @Description get the status of server.
// @Tags healthcheck
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/health [get]
func HandleHealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
