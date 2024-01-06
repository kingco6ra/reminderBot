package telegram

import (
	"reminderBot/internal/pkg/models"
	lang "reminderBot/tools/languages"
	"reminderBot/tools/metrics"
	tz "reminderBot/tools/timezones"

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

	language := lang.GetLang(u.Message.From.LanguageCode)
	msg := tgbotapi.NewMessage(int64(user.TelegramID), WelcomeMessage[language])
	b.api.Send(msg)

	metrics.IncCommand(u.Message.Command())
}

// menuHandler returning menu message with buttons.
func menuHandler(b *Bot, u *tgbotapi.Update) {
	language := lang.GetLang(u.Message.From.LanguageCode)
	msg := tgbotapi.NewMessage(int64(u.Message.From.ID), MenuMessage[language])
	msg.ReplyMarkup = getMenuButtons(language)
	b.api.Send(msg)

	metrics.IncCommand(u.Message.Command())
}

// locationHandler write user lat/lon and TZ in DB & returning menu.
func locationHandler(b *Bot, u *tgbotapi.Update) {
	lat := u.Message.Location.Latitude
	Lon := u.Message.Location.Longitude
	tz := tz.GetTimeZoneByLatLon(lat, Lon)

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
	// FIXME: mb use nlp model for get time from text?
	// reminder := models.Reminder{
	// 	TelegramUserID: u.Message.From.ID,
	// 	Description:    u.Message.Text,
	// 	ReminderTime:   time.Now().Add(5 * time.Second),
	// 	Completed:      false,
	// }
	// go b.remind(reminder)
	// b.remindersService.CreateReminder(&reminder)
}