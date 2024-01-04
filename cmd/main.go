package main

import (
	"log"
	"reminderBot/internal/repos"
	"reminderBot/internal/services"
	"reminderBot/pkg/utils"

	tgbot "reminderBot/internal/telegram"
)

func main() {
	db, err := repos.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	usersRepo := repos.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepo)

	remindersChannel := utils.MakeRemindersChannel()
	remindersRepo := repos.NewRemindersRepository(db)
	remindersService := services.NewReminderService(remindersRepo, &remindersChannel)

	bot, err := tgbot.New(usersService, remindersService, &remindersChannel)
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}
