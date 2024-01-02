package telegram

import (
	"reminderBot/internal/db"
	"reminderBot/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// handlers tg command: handler
var handlers = map[string]func(b *Bot, u *tgbotapi.Update){
	"start":    startHandler,
	"help":     helpHandler,
	"location": locationHandler,
}

// startHandler returning welcome message & insert new user in DB.
func startHandler(b *Bot, u *tgbotapi.Update) {
	user := db.User{UserID: int64(u.Message.From.ID)}
	b.db.UsersTable.Insert(user)

	msg := tgbotapi.NewMessage(user.UserID, SelectLangMessage)
	msg.ReplyMarkup = GetLangButtons()
	b.api.Send(msg)
}

func locationHandler(b *Bot, u *tgbotapi.Update) {
	Lat := u.Message.Location.Latitude
	Lon := u.Message.Location.Longitude
	tz := utils.GetTimeZoneByLatLon(Lat, Lon)

	user := db.User{
		UserID:   int64(u.Message.From.ID),
		Lat:      &Lat,
		Lon:      &Lon,
		Timezone: &tz,
	}
	b.db.UsersTable.Update(user)
}

// helpHandler returning help message with bot commands.
func helpHandler(b *Bot, u *tgbotapi.Update) {}
