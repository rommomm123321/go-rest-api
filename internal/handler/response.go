package handler

import "github.com/gofiber/fiber/v2"

func newErrorResponse(c *fiber.Ctx, status int, message string, details interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   message,
		"details": details,
	})
}

func newResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}
