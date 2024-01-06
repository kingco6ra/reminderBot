package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	TelegramID int       `gorm:"primaryKey;unique"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	Timezone   *string
	Latitude   *float64
	Longitude  *float64
	Reminders  []Reminder `gorm:"foreignKey:TelegramUserID;references:TelegramID"`
}

type Reminder struct {
	gorm.Model
	TelegramUserID int
	Description    string
	ReminderTime      time.Time
	Completed      bool
}
