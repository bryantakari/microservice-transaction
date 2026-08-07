package main

import (
	"inventory-service/internal"
	"os"
)

func main() {
	log := internal.NewLogger()
	log.Info(
		"inventory service starting",
	)
	conf := internal.LoadConfig()

	db, err := internal.NewDbConnection(conf)

	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	stats := db.Stats()
	log.Info("Connection Created ", "stats", stats)

	app := internal.AppStart(log)

	if err := app.Listen(":" + conf.AppPort); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}

}
