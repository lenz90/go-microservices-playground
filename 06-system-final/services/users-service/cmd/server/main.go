package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"playground/06-system-final/services/users-service/internal/config"
	"playground/06-system-final/services/users-service/internal/http"
	"playground/06-system-final/services/users-service/internal/kafka"
)

func main() {
	cfg := config.Load()
	producer := kafka.NewProducer(cfg.Broker)
	defer producer.Close()
	handler := httpx.New(producer)
	srv := &http.Server{Addr: ":" + cfg.Port, Handler: handler.Router()}
	go srv.ListenAndServe()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
