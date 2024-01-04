package main

import (
	"log"
	"reminderBot/internal/config"
	db "reminderBot/internal/repositories"
	users "reminderBot/internal/repositories"

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
	
	db, err := db.NewDB(cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	usersRepo := users.NewUsersRepository(db)
	bot, err := tgbot.New(cfg.TelegramConfig, usersRepo)
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}
