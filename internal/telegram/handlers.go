package telegram

import (
	"reminderBot/internal/languages"
	"reminderBot/internal/models"
	"reminderBot/pkg/metrics"
	"reminderBot/pkg/utils"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// handlers tg command: handler
var handlers = map[string]func(b *Bot, u *tgbotapi.Update){
	"start":    startHandler,
	"menu":     menuHandler,
	"location": locationHandler,
	"remind":   remindHandler,
}

// startHandler returning welcome message & insert new user in DB.
func startHandler(b *Bot, u *tgbotapi.Update) {
	user := models.User{TelegramID: u.Message.From.ID}
	b.usersService.CreateUser(&user)
	msg := tgbotapi.NewMessage(int64(user.TelegramID), WelcomeMessage[languages.Language(u.Message.From.LanguageCode)])
	b.api.Send(msg)

	metrics.IncCommand(u.Message.Command())
}

// menuHandler returning menu message with buttons.
func menuHandler(b *Bot, u *tgbotapi.Update) {
	lang := languages.Language(u.Message.From.LanguageCode)
	msg := tgbotapi.NewMessage(int64(u.Message.From.ID), MenuMessage[lang])
	msg.ReplyMarkup = getMenuButtons(lang)
	b.api.Send(msg)

	metrics.IncCommand(u.Message.Command())
}

// locationHandler write user lat/lon and TZ in DB & returning menu.
func locationHandler(b *Bot, u *tgbotapi.Update) {
	lat := u.Message.Location.Latitude
	Lon := u.Message.Location.Longitude
	tz := utils.GetTimeZoneByLatLon(lat, Lon)

	user := models.User{
		TelegramID: u.Message.From.ID,
		Latitude:   &lat,
		Longitude:  &Lon,
		Timezone:   &tz,
	}
	b.usersService.UpdateUser(&user)
	menuHandler(b, u)

	metrics.IncCommand(u.Message.Command())
}

// remindHandler set reminder for user.
func remindHandler(b *Bot, u *tgbotapi.Update) {
	reminder := models.Reminder{
		TelegramUserID: u.Message.From.ID,
		Description:    u.Message.Text,
		ReminderTime:   time.Now().Add(5 * time.Second), // FIXME
		Completed:      false,
	}
	go b.remind(reminder)
	b.remindersService.CreateReminder(&reminder)
}
