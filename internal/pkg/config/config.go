package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	BotAPIKey         string
	BotDebug          bool
	PostgresDialector gorm.Dialector
	MetricsHost       string
	MetricsPort       uint32
}

var Config *config

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	cfg, err := newConfig()
	if err != nil {
		panic(err)
	}

	Config = cfg

	log.Println("Load .env file completed.")
}

func newConfig() (*config, error) {
	botDebug, err := strconv.ParseBool(getEnv("DEBUG"))
	if err != nil {
		return nil, err
	}

	// dsn user=gorm password=gorm dbname=gorm port=9920
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		getEnv("PG_USER"), getEnv("PG_PASSWORD"), getEnv("PG_DB_NAME"), getEnv("PG_HOST"), getEnv("PG_PORT"),
	)
	pgDialector := postgres.New(postgres.Config{DSN: dsn})

	metricsPort, err := strconv.Atoi(getEnv("METRICS_PORT"))
	if err != nil {
		return nil, err
	}

	return &config{
		BotAPIKey:         getEnv("TELEGRAM_BOT_API_KEY"),
		BotDebug:          botDebug,
		PostgresDialector: pgDialector,
		MetricsHost:       getEnv("METRICS_HOST"),
		MetricsPort:       uint32(metricsPort),
	}, nil
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	panic("Variable not found: " + key)
}
