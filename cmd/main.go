package main

import (
	"context"
	"log"
	"os/signal"
	"reminderBot/internal/repos"
	"reminderBot/internal/services"
	"syscall"

	"reminderBot/internal/telegram"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	runApp(ctx)
}

func runApp(ctx context.Context) {
	db := repos.NewDB()
	usersRepo := repos.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepo)
	remindersRepo := repos.NewRemindersRepository(db)
	remindersService := services.NewReminderService(remindersRepo)

	bot := telegram.NewBot(ctx, usersService, remindersService)
	go bot.Start()
	
	<- ctx.Done()
	log.Println("Shutting down...")
}
