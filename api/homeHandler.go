package api

import (
	"github.com/gofiber/fiber/v2"
)

func (s *APIServer) handleHome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
