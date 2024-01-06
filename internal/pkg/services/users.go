package services

import (
	"reminderBot/internal/pkg/models"
	"reminderBot/internal/pkg/repos"
)

type UsersService struct {
	repo *repos.UsersRepository
}

func NewUsersService(repo *repos.UsersRepository) *UsersService {
	return &UsersService{repo: repo}
}

// CreateUser create new user.
func (service *UsersService) CreateUser(user *models.User) error {
	return service.repo.CreateUser(user)
}

// GetUserByTelegramID get user by ID.
func (service *UsersService) GetUserByTelegramID(telegramID int) (*models.User, error) {
	return service.repo.GetUserByTelegramID(telegramID)
}

// UpdateUser update user info.
func (service *UsersService) UpdateUser(user *models.User) error {
	return service.repo.UpdateUser(user)
}

// DeleteUser delete user by telegramID.
func (service *UsersService) DeleteUser(telegramID int) error {
	return service.repo.DeleteUser(telegramID)
}
