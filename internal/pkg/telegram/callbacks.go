package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type callback string

const (
	setRemind  callback = "setRemind"
	getReminds callback = "getReminds"
	help       callback = "help"
	settings   callback = "settings"
)

var callbacks = map[callback]func(b *Bot, u *tgbotapi.Update){
	setRemind:  func(b *Bot, u *tgbotapi.Update) {},
	getReminds: func(b *Bot, u *tgbotapi.Update) {},
	help:       func(b *Bot, u *tgbotapi.Update) {},
	settings:   func(b *Bot, u *tgbotapi.Update) {},
}
