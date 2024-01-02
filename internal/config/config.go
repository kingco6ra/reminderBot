package config

import (
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)


type TelegramConfig struct {
	BotAPIKey string
	Debug     bool
}

type Config struct {
	TelegramConfig TelegramConfig
	PostgresConfig pgx.ConnConfig
}

func New() (*Config, error) {
	pgPort, err := strconv.Atoi(getEnv("PG_PORT"))
	if err != nil {
		return nil, err
	}

	pgConfig := pgx.ConnConfig{
		User:     getEnv("PG_USER"),
		Password: getEnv("PG_PASSWORD"),
		Host:     getEnv("PG_HOST"),
		Port:     uint16(pgPort),
		Database: getEnv("PG_DB_NAME"),
	}

	tgConfig := TelegramConfig{
		BotAPIKey: getEnv("TELEGRAM_BOT_API_KEY"),
	}

	debug, err := strconv.ParseBool(getEnv("DEBUG"))
	if err != nil {
		return nil, err
	}

	tgConfig.Debug = debug

	return &Config{
		TelegramConfig: tgConfig,
		PostgresConfig: pgConfig,
	}, nil
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Variable %s not found.\n", key)
	return ""
}
