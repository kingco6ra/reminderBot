package metrics

import (
	"fmt"
	"log"
	"net/http"
	"reminderBot/internal/pkg/config"

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

func init() {
	prometheus.MustRegister(TelegramCommandsCounter)
	prometheus.MustRegister(NewUsersCounter)

	go Listen()
}

func Listen() error {
	address := fmt.Sprintf("%s:%d", config.Cfg.Metrics.Host, config.Cfg.Metrics.Port)
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Printf("Metrics server is starting at %s\n", address)
	
	return http.ListenAndServe(address, mux)
}
