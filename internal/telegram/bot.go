package telegram

import (
	"context"
	"log"
	cfg "reminderBot/internal/config"
	"reminderBot/internal/models"
	"reminderBot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// updateType defines the type of update (command or callback).
type updateType int

const (
	cmd updateType = iota // Command
	clb                   // Callback
)

// Bot represents a structure for working with Telegram bot.
type Bot struct {
	ctx              context.Context
	api              *tgbotapi.BotAPI
	usersService     *services.UsersService
	remindersService *services.RemindersService
}

// NewBot creates a new instance of Bot.
func NewBot(ctx context.Context, usersService *services.UsersService, remindersService *services.RemindersService) *Bot {
	api, err := tgbotapi.NewBotAPI(cfg.Config.BotAPIKey)
	if err != nil {
		log.Fatal("Failed to create Bot API:", err)
	}
	api.Debug = cfg.Config.BotDebug
	return &Bot{
		ctx:              ctx,
		api:              api,
		usersService:     usersService,
		remindersService: remindersService,
	}
}

// Start launches bot and begins listening for updates.
func (b *Bot) Start() {
	log.Println("Start telegram bot.")

	b.massRemind()
	b.pollingUpdates()

	<-b.ctx.Done()
	log.Println("Stop bot.")
}

// pollingUpdates polls updates from users.
func (b *Bot) pollingUpdates() {
	log.Println("Start polling telegram updates.")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		log.Fatalln("Failed to get updates:", err)
	}

	for {
		select {
		case update := <-updates:
			if update.CallbackQuery != nil {
				go b.handleUpdate(&update, clb)

			} else if update.Message.IsCommand() {
				go b.handleUpdate(&update, cmd)
				
			} else if update.Message.Location != nil {
				// TODO: fix this shit
				var msgEntity []tgbotapi.MessageEntity
				locCommand := "/location"
				msgEntity = append(msgEntity, tgbotapi.MessageEntity{Type: "bot_command", Offset: 0, Length: len(locCommand)})
				update.Message.Entities = &msgEntity
				update.Message.Text = locCommand
				go b.handleUpdate(&update, cmd)
			}

		case <-b.ctx.Done():
			log.Println("Stop polling telegram updates.")
			return
		}
	}
}

// remind sends reminders for scheduled events.
func (b *Bot) remind(r models.Reminder) {
	withDeadline, cancel := context.WithDeadline(b.ctx, r.ReminderTime)
	defer cancel()

	select {
	case <-withDeadline.Done():
		b.sendReminder(r)
	case <-b.ctx.Done():
		return
	}
}

// massRemind starts mass reminders for all uncompleted events.
func (b *Bot) massRemind() {
	log.Println("Start mass remind.")
	reminders := b.remindersService.GetAllUncompletedReminders()
	for _, r := range reminders {
		go b.remind(r)
	}
	log.Printf("End mass remind. Reminders count - %d.", len(reminders))
}

// sendReminder sends a reminder message to the specified user.
func (b *Bot) sendReminder(r models.Reminder) {
	msg := tgbotapi.NewMessage(int64(r.TelegramUserID), r.Description)
	b.api.Send(msg)
}

// handleUpdate handle commands & callbacks.
func (b *Bot) handleUpdate(u *tgbotapi.Update, ut updateType) {
	var request func(*Bot, *tgbotapi.Update)
	var exists bool

	switch ut {
	case cmd:
		param := u.Message.Command()
		request, exists = handlers[param]
	case clb:
		param := callback(u.CallbackQuery.Data)
		request, exists = callbacks[param]
	default:
		return
	}

	if !exists {
		log.Println("Unknown params for handleUpdate:", u)
	}
	request(b, u)
}
