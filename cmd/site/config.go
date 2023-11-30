package main

import (
	"os"

	"github.com/golang-insiders/site/internal/data"
	"github.com/joho/godotenv"
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
	cfg.env = os.Getenv("ENV")

	cfg.db, err = data.LoadDBConfig()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
