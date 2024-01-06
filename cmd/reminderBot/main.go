package main

import (
	"context"
	"os/signal"
	"reminderBot/internal/app"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	app.RunApp(ctx)
}
