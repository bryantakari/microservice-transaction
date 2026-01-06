package app

import (
	"log/slog"
	"os"

	"github.com/bryantakari/microservice-payment/order-service/internal/config"
	"github.com/bryantakari/microservice-payment/order-service/internal/db"
	"github.com/bryantakari/microservice-payment/order-service/internal/order"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(conf *config.Config, log *slog.Logger) *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:               "order-service",
		DisableStartupMessage: true,
	})
	// Global middleware
	app.Use(recover.New())

	db, err := db.NewDbConnection(conf)

	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	stats := db.Stats()
	log.Info("Connection Created ", "stats", stats)
	orderRepository := order.NewRepository(db, log)
	log.Info("order repository created")
	orderService := order.NewService(orderRepository, log)
	log.Info("order service created")
	orderHandler := order.NewHandler(orderService, log)
	log.Info("order handler created")
	orderHandler.RegisterRoutes(app)
	log.Info("order-service ready")
	return app
}
