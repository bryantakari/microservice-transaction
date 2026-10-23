package app

import (
	"inventory-service/internal/inventory"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func AppStart(inventoryService inventory.Service, log *slog.Logger) *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:               "inventory-service",
		DisableStartupMessage: true,
	})
	app.Use(cors.New())
	// Global middleware
	app.Use(recover.New())
	inventoryHandler := inventory.NewHandler(inventoryService, log)
	inventoryHandler.RegisterRoutes(app)
	log.Info("inventory-service ready")
	return app
}
