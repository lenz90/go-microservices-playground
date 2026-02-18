package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"playground/06-system-final/services/payments-service/internal/config"
	httpx "playground/06-system-final/services/payments-service/internal/http"
	"playground/06-system-final/services/payments-service/internal/kafka"
	"playground/06-system-final/services/payments-service/internal/storage"
)

func main() {
	cfg := config.Load()
	st, err := storage.Open(cfg.DSN)
	if err != nil {
		panic(err)
	}
	defer st.Close()
	pub := kafka.New(cfg.Broker)
	defer pub.Close()
	h := httpx.New(st, pub)
	srv := &http.Server{Addr: ":" + cfg.Port, Handler: h.Router()}
	go srv.ListenAndServe()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
