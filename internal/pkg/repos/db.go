package repos

import (
	cfg "reminderBot/internal/pkg/config"

	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

// NewDB creates a new instance of gorm.DB for connecting to the database.
func NewDB() (*gorm.DB, error) {
	// Open a connection to the database
	db, err := gorm.Open(cfg.Config.PostgresDialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Use Prometheus middleware for collecting database metrics
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

	return db, nil
}
