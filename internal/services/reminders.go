package services

import (
	"log"
	"reminderBot/internal/models"
	"reminderBot/internal/repos"
	"time"
)

type RemindersService struct {
	repo    *repos.RemindersRepository
	channel *chan models.Reminder
}

func NewReminderService(repo *repos.RemindersRepository, remindersChannel *chan models.Reminder) *RemindersService {
	service := &RemindersService{
		repo:    repo,
		channel: remindersChannel,
	}
	defer service.sendRemindersToChann()
	return service
}

func (service *RemindersService) sendReminderToChan(reminder models.Reminder) {
	now := time.Now()
	sub := reminder.RemindVia.Sub(now).Seconds()
	sleepDuration := time.Duration(sub * float64(time.Second))
	time.Sleep(sleepDuration)
	*service.channel <- reminder
}

func (service *RemindersService) sendRemindersToChann() {
	log.Println("Start reminders channel sending.")
	uncompletedReminders := service.GetAllUncompletedReminders()
	for _, r := range uncompletedReminders {
		go service.sendReminderToChan(r)
	}
	log.Printf("End reminders channel sending. Count reminders: %d\n", len(uncompletedReminders))
}

// CreateReminder create new user.
func (service *RemindersService) CreateReminder(reminder *models.Reminder) error {
	go service.sendReminderToChan(*reminder)
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
