package app

import (
	"context"
	"log"
	"reminderBot/internal/pkg/repos"
	"reminderBot/internal/pkg/services"
	"reminderBot/internal/pkg/tg"
)

func RunApp(ctx context.Context) {
	log.Println("Start application.")
	db := repos.NewDB()
	usersRepo := repos.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepo)
	remindersRepo := repos.NewRemindersRepository(db)
	remindersService := services.NewReminderService(remindersRepo)

	bot := tg.NewBot(ctx, usersService, remindersService)
	bot.Start()

	<-ctx.Done()
	log.Println("Shutting down...")
}
