package telegram

import (
	"reminderBot/internal/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	RUSSIAN_LANGUAGE_CALLBACK = "RU"
	ENGLISH_LANGUAGE_CALLBACK = "EN"
)

var (
	TIMEZONE_CALLBACK string = "TIMEZONE_CALLBACK"
)

var callbacks = map[string]func(b *Bot, u *tgbotapi.Update){
	RUSSIAN_LANGUAGE_CALLBACK: onLanguageCallback,
	ENGLISH_LANGUAGE_CALLBACK: onLanguageCallback,
	TIMEZONE_CALLBACK:         onTimeZoneCallback,
}

func onLanguageCallback(b *Bot, u *tgbotapi.Update) {
	userID := int64(u.CallbackQuery.From.ID)
    msgID := u.CallbackQuery.Message.MessageID
	data := u.CallbackQuery.Data
	b.db.UsersTable.Update(db.User{
		UserID:   userID,
		Language: &data,
	})

	msg := tgbotapi.NewEditMessageText(userID, msgID, SelectTZMessage[data])
	b.api.Send(msg)
}

func onTimeZoneCallback(b *Bot, u *tgbotapi.Update) {}
