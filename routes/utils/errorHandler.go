package utils

import "github.com/gofiber/fiber/v2"

func ErrorHandlder(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(err.Error())
}
