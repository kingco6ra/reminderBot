package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TelegramConfig struct {
	BotAPIKey string
	Debug     bool
}

type PostgresConfig struct {
	PostgresDialector gorm.Dialector
}

type MetricsConfig struct {
	Host string
	Port string
}

type Config struct {
	TelegramConfig TelegramConfig
	PostgresConfig PostgresConfig
	MetricsConfig  MetricsConfig
}

func getTelegramConfig() (*TelegramConfig, error) {
	tgCfg := TelegramConfig{
		BotAPIKey: getEnv("TELEGRAM_BOT_API_KEY"),
	}

	debug, err := strconv.ParseBool(getEnv("DEBUG"))
	if err != nil {
		return nil, err
	}

	tgCfg.Debug = debug

	return &tgCfg, nil
}

func getPostgresConfig() PostgresConfig {
	// dsn user=gorm password=gorm dbname=gorm port=9920
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		getEnv("PG_USER"), getEnv("PG_PASSWORD"), getEnv("PG_DB_NAME"), getEnv("PG_HOST"), getEnv("PG_PORT"),
	)
	pgDialector := postgres.New(postgres.Config{DSN: dsn})
	return PostgresConfig{
		PostgresDialector: pgDialector,
	}
}

func getMetricsConfig() MetricsConfig {
	return MetricsConfig{
		Host: getEnv("METRICS_HOST"),
		Port: getEnv("METRICS_PORT"),
	}
}

func New() (*Config, error) {
	tgCfg, err := getTelegramConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		TelegramConfig: *tgCfg,
		PostgresConfig: getPostgresConfig(),
		MetricsConfig:  getMetricsConfig(),
	}, nil
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Variable %s not found.\n", key)
	return ""
}
