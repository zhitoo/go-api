package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zhitoo/go-api/config"
	"github.com/zhitoo/go-api/requests"
	"github.com/zhitoo/go-api/storage"

	jwtware "github.com/gofiber/contrib/jwt"
)

type ApiError struct {
	Message string
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
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Accept", "application/json")
		// Go to next middleware:
		return c.Next()
	})

	app.Get("/", s.handleHome)
	app.Post("/user", s.handleCreateUser)
	app.Post("/login", s.handleLogin)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Envs.JWTSecretKey)},
	}))

	app.Get("/auth/user", s.handleGetUser)

	log.Println("JSON API running on port: ", s.listenAddr)
	app.Listen(s.listenAddr)
}
