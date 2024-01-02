package telegram

import (
	"log"
	"reminderBot/internal/config"
	"reminderBot/internal/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api       *tgbotapi.BotAPI
	db        *db.Database
	handlers  map[string]func(b *Bot, u *tgbotapi.Update)
	callbacks map[string]func(b *Bot, u *tgbotapi.Update)
}

func New(cfg config.TelegramConfig, db *db.Database) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.BotAPIKey)
	if err != nil {
		return nil, err
	}
	api.Debug = cfg.Debug
	return &Bot{
		api:       api,
		db:        db,
		handlers:  handlers,
		callbacks: callbacks,
	}, nil
}

func (b *Bot) Start() {
	log.Println("Starting bot.")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.api.GetUpdatesChan(u)

	log.Println("Start polling.")
	for update := range updates {
		if update.CallbackQuery != nil {
			go b.handleCallback(&update)
		}

		if update.Message != nil && update.Message.IsCommand() {
			go b.handleCommand(&update)
		}

		if update.Message != nil && update.Message.Location != nil {
			// TODO: fix this shit
			var msgEntity []tgbotapi.MessageEntity
			locCommand := "/location"
			msgEntity = append(msgEntity, tgbotapi.MessageEntity{Type: "bot_command", Offset: 0, Length: len(locCommand)})
			update.Message.Entities = &msgEntity
			update.Message.Text = locCommand
			go b.handleCommand(&update)
		}
	}
}

func (b *Bot) handleCallback(u *tgbotapi.Update) {
	data := u.CallbackQuery.Data
	callback, exists := b.callbacks[data]
	if exists {
		callback(b, u)
	}
}

func (b *Bot) handleCommand(u *tgbotapi.Update) {
	cmd := u.Message.Command()
	handler, exists := b.handlers[cmd]
	if exists {
		handler(b, u)
	}
}
