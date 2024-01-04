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
	MetricsPort       string
}

var Config *config

func init() {
	if err := godotenv.Load(); err != nil {
		//log.Fatal("No .env file found")
	}
	Config = New()
	log.Println("Load .env file completed.")
}

func New() *config {
	botDebug, err := strconv.ParseBool(getEnv("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}

	// dsn user=gorm password=gorm dbname=gorm port=9920
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		getEnv("PG_USER"), getEnv("PG_PASSWORD"), getEnv("PG_DB_NAME"), getEnv("PG_HOST"), getEnv("PG_PORT"),
	)
	pgDialector := postgres.New(postgres.Config{DSN: dsn})

	return &config{
		BotAPIKey:         getEnv("TELEGRAM_BOT_API_KEY"),
		BotDebug:          botDebug,
		PostgresDialector: pgDialector,
		MetricsHost:       getEnv("METRICS_HOST"),
		MetricsPort:       getEnv("METRICS_PORT"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Variable %s not found.\n", key)
	return ""
}
