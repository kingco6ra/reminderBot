package repos

import "reminderBot/internal/pkg/models"

type UserRepoInterface interface {
	CreateUser(*models.User) error
	GetUserByTelegramID(int) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int) error
}

type ReminderRepoInterface interface {
	CreateReminder(reminder *models.Reminder) error
	GetAllUncompletedReminders() []models.Reminder
	GetUserReminders(telegramUserID int) []models.Reminder
}
