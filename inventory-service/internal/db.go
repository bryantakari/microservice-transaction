package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDbConnection(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.CreateDBUrl())
	if nil != err || nil == db {
		fmt.Println("db connection failed")
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {

		fmt.Println("error ping")
		return nil, err
	}

	return db, nil
}
