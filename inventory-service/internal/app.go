package internal

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func AppStart(log *slog.Logger) *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:               "inventory-service",
		DisableStartupMessage: true,
	})
	// Global middleware
	app.Use(recover.New())
	log.Info("order-service ready")
	return app
}
