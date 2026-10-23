package main

import (
	"os"

	"inventory-service/internal/app"
	"inventory-service/internal/config"
	"inventory-service/internal/db"
	"inventory-service/internal/inventory"
	inventoryRepo "inventory-service/internal/inventory/repository"
	"inventory-service/internal/logger"
)

func main() {
	log := logger.NewLogger()
	log.Info(
		"inventory service starting",
	)
	conf := config.LoadConfig()

	db, err := db.NewDbConnection(conf)

	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	stats := db.Stats()
	log.Info("Connection Created ", "stats", stats)

	inventoryRepository := inventoryRepo.NewInventoryRepository(db, log)
	transactionRepository := inventoryRepo.NewTransactionRepository(db, log)
	inventoryService := inventory.NewService(db, inventoryRepository, transactionRepository, log)

	app := app.AppStart(inventoryService, log)

	if err := app.Listen(":" + conf.AppPort); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}

}
