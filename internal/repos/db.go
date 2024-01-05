package repos

import (
	"log"
	cfg "reminderBot/internal/config"

	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(cfg.Config.PostgresDialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to create connection to DB:", err)
	}

	db.Use(
		prometheus.New(
			prometheus.Config{
				DBName:         "PostgreSQL",
				HTTPServerPort: cfg.Config.MetricsPort,
				MetricsCollector: []prometheus.MetricsCollector{
					&prometheus.Postgres{
						VariableNames: []string{"Threads_running"},
					},
				},
			},
		),
	)

	return db
}
