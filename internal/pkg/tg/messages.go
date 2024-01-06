// Templates for messages.
package tg

import "reminderBot/tools/languages"

type Message map[languages.Language]string

var WelcomeMessage = Message{
	languages.RUSSIAN: "Перед началом работы - отправьте свою локацию чтобы я мог определить Ваш часовой пояс.",
	languages.ENGLISH: "Before starting work, send your location so that I can determine your time zone.",
}

var MenuMessage = Message{
	languages.RUSSIAN: "Выберите пункт из меню:",
	languages.ENGLISH: "Select an item from the menu:",
}
