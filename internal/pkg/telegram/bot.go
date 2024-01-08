package telegram

import (
	"context"
	"log"
	cfg "reminderBot/internal/pkg/config"
	"reminderBot/internal/pkg/models"
	"reminderBot/internal/pkg/repos"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot represents a structure for working with Telegram bot.
type Bot struct {
	ctx          context.Context
	api          *tgbotapi.BotAPI
	usersRepo    repos.UserRepoInterface
	reminderRepo repos.ReminderRepoInterface
}

// NewBot creates a new instance of Bot.
func NewBot(usersRepo repos.UserRepoInterface, reminderRepo repos.ReminderRepoInterface) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Config.BotAPIKey)
	if err != nil {
		return nil, err
	}

	api.Debug = cfg.Config.BotDebug

	return &Bot{
		api:          api,
		usersRepo:    usersRepo,
		reminderRepo: reminderRepo,
	}, nil
}

// Start launches bot and begins listening for updates.
func (b *Bot) Start(ctx context.Context) {
	log.Println("Start telegram bot.")

	b.massRemind(ctx)
	b.pollingUpdates(ctx)

	<-b.ctx.Done()
	log.Println("Stop bot.")
}

// pollingUpdates polls updates from users.
func (b *Bot) pollingUpdates(ctx context.Context) {
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
			go b.handleUpdate(&update)
		case <-ctx.Done():
			log.Println("Stop polling telegram updates.")
			return
		}
	}
}

// remind sends reminders for scheduled events.
func (b *Bot) remind(ctx context.Context, r models.Reminder) {
	withDeadline, cancel := context.WithDeadline(ctx, r.ReminderTime)
	defer cancel()

	select {
	case <-withDeadline.Done():
		b.sendReminder(r)
	case <-b.ctx.Done():
		return
	}
}

// massRemind starts mass reminders for all uncompleted events.
func (b *Bot) massRemind(ctx context.Context) {
	log.Println("Start mass remind.")

	reminders := b.reminderRepo.GetAllUncompletedReminders()

	for _, r := range reminders {
		go b.remind(ctx, r)
	}

	log.Printf("End mass remind. Reminders count - %d.", len(reminders))
}

// sendReminder sends a reminder message to the specified user.
func (b *Bot) sendReminder(r models.Reminder) {
	msg := tgbotapi.NewMessage(int64(r.TelegramUserID), r.Description)
	b.api.Send(msg)
}

// handleUpdate handle commands & callbacks.
func (b *Bot) handleUpdate(u *tgbotapi.Update) {
	if u.CallbackQuery != nil {
		b.handleCallback(u)
	} else if u.Message.IsCommand() {
		b.handleCommand(u)
	} else if u.Message.Location != nil {
		// TODO: fix this shit
		var msgEntity []tgbotapi.MessageEntity
		locCommand := "/location"
		msgEntity = append(msgEntity, tgbotapi.MessageEntity{Type: "bot_command", Offset: 0, Length: len(locCommand)})
		u.Message.Entities = &msgEntity
		u.Message.Text = locCommand
		b.handleCommand(u)
	} else {
		log.Println("Unknown update type: ", u)
	}
}

// handleCommand handling user command.
func (b *Bot) handleCommand(u *tgbotapi.Update) {
	cmd := u.Message.Command()
	cmdHandler, exists := commandHandlers[cmd]

	if exists {
		cmdHandler(b, u)
	}
}

// handleCallback handling user callback.
func (b *Bot) handleCallback(u *tgbotapi.Update) {
	clb := callback(u.CallbackQuery.Data)
	clbHandler, exists := callbackHandlers[clb]

	if exists {
		clbHandler(b, u)
	}
}
