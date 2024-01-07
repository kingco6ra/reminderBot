package tg

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

	metrics.IncCommand(u.Mpackage times

		import (
			"errors"
			"strings"
			"time"
		
			"github.com/olebedev/when"
			"github.com/olebedev/when/rules/en"
			"github.com/olebedev/when/rules/ru"
		)
		
		var ErrorTimePattern = errors.New("ErrorTimePattern")
		
		func ParseReminderTime(now time.Time, text string) (*time.Time, error) {
			w := when.New(nil)
			w.Add(ru.All...)
			w.Add(en.All...)
		
			r, e := w.Parse(strings.TrimSpace(text), now)
			if r == nil {
				return nil, ErrorTimePattern
			}
			
			if e != nil {
				return nil, e
			}
		
			return &r.Time, nil
		}
		essage.Command())
}

// remindHandler set reminder for user.
func remindHandler(b *Bot, u *tgbotapi.Update) {
}
