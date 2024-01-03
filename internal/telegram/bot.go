package telegram

import (
	"log"
	"reminderBot/internal/config"
	"reminderBot/internal/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type updateType int

const (
	cmd updateType = iota
	clb
)

type Bot struct {
	api       *tgbotapi.BotAPI
	db        *db.Database
	handlers  map[string]func(b *Bot, u *tgbotapi.Update)
	callbacks map[callback]func(b *Bot, u *tgbotapi.Update)
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
			go b.handleUpdate(&update, clb)
		}

		if update.Message != nil && update.Message.IsCommand() {
			go b.handleUpdate(&update, cmd)
		}

		if update.Message != nil && update.Message.Location != nil {
			// TODO: fix this shit
			var msgEntity []tgbotapi.MessageEntity
			locCommand := "/location"
			msgEntity = append(msgEntity, tgbotapi.MessageEntity{Type: "bot_command", Offset: 0, Length: len(locCommand)})
			update.Message.Entities = &msgEntity
			update.Message.Text = locCommand
			go b.handleUpdate(&update, cmd)
		}
	}
}

func (b *Bot) handleUpdate(u *tgbotapi.Update, t updateType) {
	var request func(*Bot, *tgbotapi.Update)
	var exists bool

	switch t {
	case cmd:
		param := u.Message.Command()
		request, exists = b.handlers[param]
	case clb:
		param := callback(u.CallbackQuery.Data)
		request, exists = b.callbacks[param]
	default:
		return
	}

	if exists {
		request(b, u)
	}
}
