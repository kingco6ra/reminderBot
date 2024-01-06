// TODO: add contexts.
package repos

import (
	"log"
	"reminderBot/internal/pkg/models"

	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}
	return &UsersRepository{db: db}
}

// CreateUser create new user.
func (repo *UsersRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

// GetUserByTelegramID get user by ID.
func (repo *UsersRepository) GetUserByTelegramID(telegramID int) (*models.User, error) {
	var user models.User
	err := repo.db.Where("telegram_id = ?", telegramID).First(&user).Error
	return &user, err
}

// UpdateUser update user info.
func (repo *UsersRepository) UpdateUser(user *models.User) error {
	return repo.db.Model(&models.User{}).Where("telegram_id = ?", user.TelegramID).Updates(user).Error
}

// DeleteUser delete user by telegramID.
func (repo *UsersRepository) DeleteUser(telegramID int) error {
	return repo.db.Where("telegram_id = ?", telegramID).Delete(&models.User{}).Error
}
