package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func GetLangButtons() *tgbotapi.InlineKeyboardMarkup {
	buttons := []tgbotapi.InlineKeyboardButton{
		{Text: "Русский", CallbackData: &RUSSIAN_LANGUAGE_CALLBACK},
		{Text: "English", CallbackData: &ENGLISH_LANGUAGE_CALLBACK},
	}
	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{buttons},
	}
}
