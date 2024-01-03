package telegram

import (
	"reminderBot/internal/languages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var menuButtonsOrder = []callback{setRemind, getReminds, help, settings}
var menuButtons = map[languages.Language]map[callback]string{
	languages.RUSSIAN: {
		setRemind:  "Поставить напоминание 🔔",
		getReminds: "Список напоминаний 📝",
		help:       "Помощь ❓",
		settings:   "Настройки 🛠",
	},
	languages.ENGLISH: {
		setRemind:  "Set a reminder 🔔",
		getReminds: "List of reminders 📝",
		help:       "Help ❓",
		settings:   "Settings 🛠",
	},
}

func getMenuButtons(lang languages.Language) *tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, v := range menuButtonsOrder {
		clb := string(v)
		button := tgbotapi.InlineKeyboardButton{
			Text:         menuButtons[lang][v],
			CallbackData: &clb,
		}
		row := []tgbotapi.InlineKeyboardButton{button}
		rows = append(rows, row)
	}
    
	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}
