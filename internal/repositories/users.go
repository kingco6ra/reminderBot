// TODO: add contexts.
package db

import (
	cfg "reminderBot/internal/config"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserSchema struct {
	ID         uint      `gorm:"primaryKey"`
	TelegramID int       `gorm:"primaryKey;unique"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	Timezone   *string
	Latitude   *float64
	Longitude  *float64
}

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	if err := db.AutoMigrate(&UserSchema{}, cfg.Config.MigrationPath); err != nil {
		log.Fatal(err)
	}
	return &UsersRepository{db: db}
}

// CreateUser create new user.
func (repo *UsersRepository) CreateUser(user *UserSchema) error {
	return repo.db.Create(user).Error
}

// GetUserByTelegramID get user by ID.
func (repo *UsersRepository) GetUserByTelegramID(telegramID int) (*UserSchema, error) {
	var user UserSchema
	err := repo.db.Where("telegram_id = ?", telegramID).First(&user).Error
	return &user, err
}

// UpdateUser update user info.
func (repo *UsersRepository) UpdateUser(user *UserSchema) error {
	return repo.db.Model(&UserSchema{}).Where("telegram_id = ?", user.TelegramID).Updates(user).Error
}

// DeleteUser delete user by telegramID.
func (repo *UsersRepository) DeleteUser(telegramID int) error {
	return repo.db.Where("telegram_id = ?", telegramID).Delete(&UserSchema{}).Error
}
