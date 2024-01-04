package main

import (
	"log"
	repo "reminderBot/internal/repositories"

	tgbot "reminderBot/internal/telegram"
)

func main() {
	db, err := repo.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	usersRepo := repo.NewUsersRepository(db)
	bot, err := tgbot.New(usersRepo)
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}
