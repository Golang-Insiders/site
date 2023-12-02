package main

import (
	"fmt"
	"os"

	"github.com/golang-insiders/site/internal/data"
	"github.com/joho/godotenv"
)

var (
	DEFAULT_PORT = ":3000"
	DEFAULT_ENV  = "dev"
)

type config struct {
	port string
	env  string
	db   data.PostgresConfig
}

func loadConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	cfg.port = os.Getenv("PORT")
	if cfg.port == "" {
		cfg.port = DEFAULT_PORT
	}
	cfg.env = os.Getenv("ENV")
	if cfg.env == "" {
		cfg.env = DEFAULT_ENV
	}

	cfg.db, err = data.LoadDBConfig()
	if err != nil {
		return cfg, fmt.Errorf("Failed to load DB config: %v", err)
	}

	return cfg, nil
}
