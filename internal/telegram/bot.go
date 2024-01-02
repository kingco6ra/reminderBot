package telegram

import (
	"log"
	"reminderBot/internal/config"
	"reminderBot/internal/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	db       *db.Database
	handlers map[string]func(b *Bot, u *tgbotapi.Update)
}

func New(cfg config.TelegramConfig, db *db.Database) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.BotAPIKey)
	api.Debug = cfg.Debug
	return &Bot{
		api:      api,
		db:       db,
		handlers: Handlers,
	}, err
}

func (b *Bot) Start() {
	log.Println("Starting bot.")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.api.GetUpdatesChan(u)

	log.Println("Start polling.")
	for update := range updates {
		msg := update.Message
		if msg == nil || !msg.IsCommand() {
			continue
		}
		log.Printf("[%s] %s", msg.From.UserName, msg.Text)
		go b.handleCommand(&update)
	}
}

func (b *Bot) handleCommand(u *tgbotapi.Update) {
	cmd := u.Message.Command()
	handler, exists := b.handlers[cmd]
	if exists {
		handler(b, u)
	}
}
