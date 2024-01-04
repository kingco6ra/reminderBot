package main

import (
	"log"
	"reminderBot/internal/repos"

	tgbot "reminderBot/internal/telegram"
)

func main() {
	db, err := repos.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	usersRepo := repos.NewUsersRepository(db)
	bot, err := tgbot.New(usersRepo)
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}
