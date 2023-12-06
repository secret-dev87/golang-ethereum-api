package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/secret-dev87/golang-ethereum-api/pkg/env"
	"github.com/secret-dev87/golang-ethereum-api/pkg/router"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	app := fiber.New(fiber.Config{
		AppName:       "Go Ethereum API",
		CaseSensitive: true,
	})

	//Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	//Setup routes
	router.InstallRouter(app)

	return app
}
