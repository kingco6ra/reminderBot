package app

import (
	"context"
	"log"
	"reminderBot/internal/pkg/config"
	"reminderBot/internal/pkg/repos"
	"reminderBot/internal/pkg/telegram"
)

func RunApp(ctx context.Context) {
	log.Println("Start application.")

	db, err := repos.NewDB(config.Cfg.Database.GetPostgresDialector(), config.Cfg.Metrics.Port)
	if err != nil {
		log.Fatal("Failed to create connection to DB:", err)
	}

	usersRepo, err := repos.NewUsersRepository(db)
	if err != nil {
		log.Fatalln("Error during migration: ", err)
	}

	remindersRepo, err := repos.NewRemindersRepository(db)
	if err != nil {
		log.Fatalln("Error during migration: ", err)
	}

	bot, err := telegram.NewBot(config.Cfg.Telegram.Token, config.Cfg.Telegram.Debug, usersRepo, remindersRepo)
	if err != nil {
		log.Fatalln(err)
	}

	bot.Start(ctx)
	<-ctx.Done()
	log.Println("Shutting down...")
}
