package repos

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

// NewDB creates a new instance of gorm.DB for connecting to the database.
func NewDB(dsn string, metricsPort uint32) (*gorm.DB, error) {
	// Open a connection to the database
	dialector := postgres.New(postgres.Config{DSN: dsn})
	db, err := gorm.Open(dialector, &gorm.Config{})
	
	if err != nil {
		return nil, err
	}

	// Use Prometheus middleware for collecting database metrics
	db.Use(
		prometheus.New(
			prometheus.Config{
				DBName:         "PostgreSQL",
				HTTPServerPort: metricsPort,
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
