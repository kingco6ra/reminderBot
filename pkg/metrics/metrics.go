package metrics

import (
	"fmt"
	"log"
	"net/http"
	cfg "reminderBot/internal/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var TelegramCommandsCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "telegram_commands_total",
		Help: "Total number of Telegram commands received",
	},
	[]string{"command"},
)

var NewUsersCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "new_users",
		Help: "New users who are not in the database",
	},
)

func IncCommand(command string) {
	TelegramCommandsCounter.WithLabelValues(command).Inc()
}

func init() {
	prometheus.MustRegister(TelegramCommandsCounter)
	prometheus.MustRegister(NewUsersCounter)
	go Listen()
}

func Listen() error {
	address := fmt.Sprintf("%s:%s", cfg.Config.MetricsHost, cfg.Config.MetricsPort)
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Printf("Metrics server is starting at %s\n", address)
	return http.ListenAndServe(address, mux)
}
