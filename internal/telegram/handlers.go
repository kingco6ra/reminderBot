package telegram

import (
	"reminderBot/internal/db"
	"reminderBot/internal/languages"
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
	user := db.User{UserID: int64(u.Message.From.ID)}
	b.db.UsersTable.Insert(user)

	msg := tgbotapi.NewMessage(user.UserID, WelcomeMessage[languages.Language(u.Message.From.LanguageCode)])
	b.api.Send(msg)
}

func menuHandler(b *Bot, u *tgbotapi.Update) {
	lang := languages.Language(u.Message.From.LanguageCode)
	msg := tgbotapi.NewMessage(int64(u.Message.From.ID), MenuMessage[lang])
	msg.ReplyMarkup = getMenuButtons(lang)
	b.api.Send(msg)
}

func locationHandler(b *Bot, u *tgbotapi.Update) {
	lat := u.Message.Location.Latitude
	Lon := u.Message.Location.Longitude
	tz := utils.GetTimeZoneByLatLon(lat, Lon)

	user := db.User{
		UserID:   int64(u.Message.From.ID),
		Lat:      &lat,
		Lon:      &Lon,
		Timezone: &tz,
	}
	b.db.UsersTable.Update(user)
	menuHandler(b, u)
}
