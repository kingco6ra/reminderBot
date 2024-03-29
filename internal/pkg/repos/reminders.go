package repos

import (
	"reminderBot/internal/pkg/models"
	"time"

	"gorm.io/gorm"
)

type RemindersRepository struct {
	db *gorm.DB
}

func NewRemindersRepository(db *gorm.DB) (*RemindersRepository, error) {
	if err := db.AutoMigrate(&models.Reminder{}); err != nil {
		return nil, err
	}

	return &RemindersRepository{db: db}, nil
}

// CreateReminder create new user.
func (repo *RemindersRepository) CreateReminder(reminder *models.Reminder) error {
	return repo.db.Create(reminder).Error
}

// GetAllUncompletedReminders returning all uncompleted reminders for remind.
func (repo *RemindersRepository) GetAllUncompletedReminders() ([]models.Reminder, error) {
	var reminders []models.Reminder

	err := repo.db.Where("completed = ? AND reminder_time >= ?", false, time.Now().UTC()).Find(&reminders).Error

	return reminders, err
}

// GetUserReminders returning user reminders with selected status.
func (repo *RemindersRepository) GetUserReminders(telegramUserID int) ([]models.Reminder, error) {
	var reminders []models.Reminder
	
	err := repo.db.Find(&reminders, models.Reminder{TelegramUserID: telegramUserID}).Error

	return reminders, err
}
