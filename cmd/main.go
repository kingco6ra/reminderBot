package main

import (
	"log"
	"reminderBot/internal/config"
	"reminderBot/internal/db"
	tgbot "reminderBot/internal/telegram"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error parsing config value: %v\n", err)
	}
	
	db, err := db.New(cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbot.New(cfg.TelegramConfig, db)
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}