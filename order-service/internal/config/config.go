package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBURL      string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		AppPort:    getEnv("APP_PORT", "3000"),
		DBPORT:     getEnv("DB_PORT", ""),
		DBHOST:     getEnv("DB_HOST", ""),
		DBNAME:     getEnv("DB_NAME", ""),
		DBPASSWORD: getEnv("DB_PASSWORD", ""),
		DBUSER:     getEnv("DB_USER", ""),
	}

	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
