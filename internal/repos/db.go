package repos

import (
	cfg "reminderBot/internal/config"

	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(cfg.Config.PostgresDialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
