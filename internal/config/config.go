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

type MetricsConfig struct {
	Host string
	Port string
}

type Config struct {
	TelegramConfig TelegramConfig
	PostgresConfig pgx.ConnConfig
	MetricsConfig  MetricsConfig
}

func New() (*Config, error) {
	pgPort, err := strconv.Atoi(getEnv("PG_PORT"))
	if err != nil {
		return nil, err
	}

	pgCfg := pgx.ConnConfig{
		User:     getEnv("PG_USER"),
		Password: getEnv("PG_PASSWORD"),
		Host:     getEnv("PG_HOST"),
		Port:     uint16(pgPort),
		Database: getEnv("PG_DB_NAME"),
	}

	tgCfg := TelegramConfig{
		BotAPIKey: getEnv("TELEGRAM_BOT_API_KEY"),
	}

	debug, err := strconv.ParseBool(getEnv("DEBUG"))
	if err != nil {
		return nil, err
	}

	tgCfg.Debug = debug

	metricsCfg := MetricsConfig{
		Host: getEnv("METRICS_HOST"),
		Port: getEnv("METRICS_PORT"),
	}

	return &Config{
		TelegramConfig: tgCfg,
		PostgresConfig: pgCfg,
		MetricsConfig:  metricsCfg,
	}, nil
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Variable %s not found.\n", key)
	return ""
}
