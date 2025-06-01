package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

var DB *DBConfig

func LoadDBConfig() {
	DB = &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USERNAME", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
		Database: getEnv("DB_NAME", "mydb"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func ConnectDB() (*pgxpool.Pool, error) {
	if DB == nil {
		LoadDBConfig()
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		DB.Username, DB.Password, DB.Host, DB.Port, DB.Database, DB.SSLMode,
	)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse db config: %w", err)
	}

	// Connect with context
	ctx := context.Background()

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test connection
	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return dbpool, nil

}
