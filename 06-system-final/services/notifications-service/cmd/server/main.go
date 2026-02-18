package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"playground/06-system-final/services/notifications-service/internal/config"
	"playground/06-system-final/services/notifications-service/internal/kafka"
	"playground/06-system-final/services/notifications-service/internal/storage"
)

func main() {
	cfg := config.Load()
	st := storage.New(cfg.RedisAddr)
	c := kafka.New(cfg.Broker, st)
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sig; cancel() }()

	for ctx.Err() == nil {
		pollCtx, cancelPoll := context.WithTimeout(ctx, 1*time.Second)
		c.PollOnce(pollCtx)
		cancelPoll()
		time.Sleep(200 * time.Millisecond)
	}
}
