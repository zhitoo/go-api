package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhitoo/gobank/models"
	"github.com/zhitoo/gobank/requests"
	"github.com/zhitoo/gobank/utils"
)

func (s *APIServer) handleGetAccount(c *fiber.Ctx) error {
	return nil
}

func (s *APIServer) handleCreateAccount(c *fiber.Ctx) error {
	payload := new(requests.RegisterUser)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	// Validation
	errs := s.validator.Validate(payload)
	if errs != nil {
		c.Status(422)
		return c.JSON(errs)
	}

	hash, err := utils.Encrypt(payload.Password)

	if err != nil {
		return err
	}

	user, _ := s.storage.GetUserByUserName(payload.UserName)

	if user.ID != 0 {
		return c.Status(400).SendString("User Exists!")
	}

	user = &models.User{
		LastName:  payload.LastName,
		FirstName: payload.FirstName,
		UserName:  payload.UserName,
		Password:  hash,
	}
	user, err = s.storage.CreateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
