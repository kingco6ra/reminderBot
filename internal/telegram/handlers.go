package telegram

import (
	"reminderBot/internal/languages"
	users "reminderBot/internal/repositories"
	"reminderBot/pkg/metrics"
	"reminderBot/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// handlers tg command: handler
var handlers = map[string]func(b *Bot, u *tgbotapi.Update){
	"start":    startHandler,
	"menu":     menuHandler,
	"location": locationHandler,
}

// startHandler returning welcome message & insert new user in DB.
func startHandler(b *Bot, u *tgbotapi.Update) {
	user := users.UserSchema{TelegramID: u.Message.From.ID}
	b.repo.CreateUser(&user)
	msg := tgbotapi.NewMessage(int64(user.TelegramID), WelcomeMessage[languages.Language(u.Message.From.LanguageCode)])
	b.api.Send(msg)

	metrics.IncCommand(u.Message.Command())
}

func menuHandler(b *Bot, u *tgbotapi.Update) {
	lang := languages.Language(u.Message.From.LanguageCode)
	msg := tgbotapi.NewMessage(int64(u.Message.From.ID), MenuMessage[lang])
	msg.ReplyMarkup = getMenuButtons(lang)
	b.api.Send(msg)

	metrics.IncCommand(u.Message.Command())
}

func locationHandler(b *Bot, u *tgbotapi.Update) {
	lat := u.Message.Location.Latitude
	Lon := u.Message.Location.Longitude
	tz := utils.GetTimeZoneByLatLon(lat, Lon)

	user := users.UserSchema{
		TelegramID: u.Message.From.ID,
		Latitude:   &lat,
		Longitude:  &Lon,
		Timezone:   &tz,
	}
	b.repo.UpdateUser(&user)
	menuHandler(b, u)

	metrics.IncCommand(u.Message.Command())
}
