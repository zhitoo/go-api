# Simple Go Api

## Routes

you can create your route in api/api.go in Run method

```go
func (s *APIServer) Run() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Accept", "application/json")
		// Go to next middleware:
		return c.Next()
	})

    //without auth routes
	app.Get("/", s.handleHome)
	app.Post("/user", s.handleCreateUser)
	app.Post("/login", s.handleLogin)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Envs.JWTSecretKey)},
	}))

    //auth routes

	app.Get("/auth/user", s.handleGetUser)

	log.Println("JSON API running on port: ", s.listenAddr)
	app.Listen(s.listenAddr)
}
```
