package telegram

import (
	"log"
	cfg "reminderBot/internal/config"
	"reminderBot/internal/models"
	"reminderBot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type updateType int

const (
	cmd updateType = iota
	clb
)

type Bot struct {
	api              *tgbotapi.BotAPI
	usersService     *services.UsersService
	remindersService *services.RemindersService
	remindersChannel *chan models.Reminder
	handlers         map[string]func(b *Bot, u *tgbotapi.Update)
	callbacks        map[callback]func(b *Bot, u *tgbotapi.Update)
}

func New(usersService *services.UsersService, remindersService *services.RemindersService, remindersChannel *chan models.Reminder) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Config.BotAPIKey)
	if err != nil {
		return nil, err
	}
	api.Debug = cfg.Config.BotDebug
	return &Bot{
		api:              api,
		usersService:     usersService,
		remindersService: remindersService,
		remindersChannel: remindersChannel,
		handlers:         handlers,
		callbacks:        callbacks,
	}, nil
}

func (b *Bot) Start() {
	log.Println("Starting bot.")
	go b.pollingRemindersChannel()
	b.pollingUpdates()
}

// pollingUpdates polling user messages.
func (b *Bot) pollingUpdates() {
	log.Println("Start polling.")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.api.GetUpdatesChan(u)

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

// pollingRemindersChannel reading reminders channel for notify user in messenger.
func (b *Bot) pollingRemindersChannel() {
	log.Println("Start reminders channel polling.")
	for reminder := range *b.remindersChannel {
		msg := tgbotapi.NewMessage(int64(reminder.TelegramUserID), reminder.Description)
		b.api.Send(msg)
	}
	log.Println("End reminders channel polling.")
}

// handleUpdate handle commands & callbacks.
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
