package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *APIServer) Run() {
	app := fiber.New()
	app.Get("/", s.handleHome)
	app.Post("/account", s.handleCreateAccount)
	app.Get("/account/{id}", s.handleGetAccount)

	log.Println("JSON API running on port: ", s.listenAddr)
	app.Listen(s.listenAddr)
}
