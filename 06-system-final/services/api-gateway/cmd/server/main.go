package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"playground/06-system-final/services/api-gateway/internal/config"
	httpx "playground/06-system-final/services/api-gateway/internal/http"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	h, err := httpx.New(cfg.UsersBaseURL, cfg.PaymentsBaseURL)
	if err != nil {
		panic(err)
	}
	srv := &http.Server{Addr: ":" + cfg.Port, Handler: h.Router()}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server", "err", err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
