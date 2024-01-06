package services

import (
	"reminderBot/internal/pkg/models"
	"reminderBot/internal/pkg/repos"
)

type RemindersService struct {
	repo *repos.RemindersRepository
}

func NewReminderService(repo *repos.RemindersRepository) *RemindersService {
	service := &RemindersService{repo: repo}
	return service
}

// CreateReminder create new user.
func (service *RemindersService) CreateReminder(reminder *models.Reminder) error {
	return service.repo.CreateReminder(reminder)
}

// GetAllUncompletedReminders returning all uncompleted reminders for remind.
func (service *RemindersService) GetAllUncompletedReminders() []models.Reminder {
	return service.repo.GetAllUncompletedReminders()
}

// GetUserReminders returning user reminders with selected status.
func (service *RemindersService) GetUserReminders(telegramUserID int) []models.Reminder {
	return service.repo.GetUserReminders(telegramUserID)
}
