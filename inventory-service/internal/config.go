package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
	DBSSLMODE  string
}

type Config struct {
	AppPort  string
	DBConfig DBConfig
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	dbcfg := DBConfig{
		DBPORT:     getEnv("DB_PORT", ""),
		DBHOST:     getEnv("DB_HOST", ""),
		DBNAME:     getEnv("DB_NAME", ""),
		DBPASSWORD: getEnv("DB_PASSWORD", ""),
		DBUSER:     getEnv("DB_USER", ""),
		DBSSLMODE:  getEnv("DB_SSLMODE", ""),
	}
	cfg := &Config{
		AppPort:  getEnv("APP_PORT", "3000"),
		DBConfig: dbcfg,
	}

	return cfg
}

func (cfg *Config) CreateDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBConfig.DBUSER,
		cfg.DBConfig.DBUSER,
		cfg.DBConfig.DBHOST,
		cfg.DBConfig.DBPORT,
		cfg.DBConfig.DBNAME,
		cfg.DBConfig.DBSSLMODE,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
