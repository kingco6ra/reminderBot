package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Handlers tg command: handler
var Handlers = map[string]func(b *Bot, u *tgbotapi.Update){
	"start": startHandler,
	"help":  helpHandler,
}

// startHandler returning welcome message & insert new user in DB.
func startHandler(b *Bot, u *tgbotapi.Update) {}

// helpHandler returning help message with bot commands.
func helpHandler(b *Bot, u *tgbotapi.Update) {}
