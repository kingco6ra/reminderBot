package db

import (
	"reminderBot/internal/config"

	"gorm.io/gorm"
)

func NewDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(cfg.PostgresDialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
