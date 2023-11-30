package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Database     string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func OpenDB(cfg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)

	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func LoadDBConfig() (PostgresConfig, error) {
	var cfg PostgresConfig

	cfg.Host = os.Getenv("PSQL_HOST")
	cfg.Port = os.Getenv("PSQL_PORT")
	cfg.User = os.Getenv("PSQL_USER")
	cfg.Password = os.Getenv("PSQL_PASSWORD")
	cfg.Database = os.Getenv("PSQL_DATABASE")
	cfg.SSLMode = os.Getenv("PSQL_SSLMODE")

	maxOpenConns, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		return cfg, fmt.Errorf("failed to convert MAX_OPEN_CONNS to integer: %v", err)
	}
	cfg.MaxOpenConns = maxOpenConns

	maxIdleConns, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		return cfg, fmt.Errorf("failed to convert MAX_IDLE_CONNS to integer: %v", err)
	}
	cfg.MaxIdleConns = maxIdleConns

	cfg.MaxIdleTime = os.Getenv("MAX_IDLE_TIME")

	return cfg, nil
}
