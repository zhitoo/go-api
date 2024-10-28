package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zhitoo/gobank/requests"
	"github.com/zhitoo/gobank/storage"
)

type ApiError struct {
	Error string
}

type APIServer struct {
	listenAddr string
	storage    storage.Storage
	validator  *requests.Validator
}

func NewAPIServer(listenAddr string, storage storage.Storage, validator *requests.Validator) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
		validator:  validator,
	}
}

func (s *APIServer) Run() {
	app := fiber.New()
	app.Get("/", s.handleHome)
	app.Post("/account", s.handleCreateAccount)
	app.Get("/account/{id}", s.handleGetAccount)

	log.Println("JSON API running on port: ", s.listenAddr)
	app.Listen(s.listenAddr)
}
