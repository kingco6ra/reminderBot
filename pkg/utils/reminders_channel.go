package utils

import "reminderBot/internal/models"

func MakeRemindersChannel() chan models.Reminder {
	return make(chan models.Reminder)
}
