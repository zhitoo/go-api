package api

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/zhitoo/gobank/requests"
)

func (s *APIServer) handleGetAccount(c *fiber.Ctx) error {
	return nil
}

func (s *APIServer) handleCreateAccount(c *fiber.Ctx) error {
	payload := new(requests.RegisterAccount)
	if err := c.BodyParser(payload); err != nil {
		return err
	}
	account, err := s.storage.CreateAccount(payload.FirstName, payload.LastName, uint64(rand.Intn(100000000)))
	if err != nil {
		return err
	}
	return c.JSON(account)
}
