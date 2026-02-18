package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func router() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	m.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"echo": r.URL.Query().Get("msg")})
	})
	return m
}
func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	s := &http.Server{Addr: ":" + p, Handler: router()}
	go s.ListenAndServe()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	ctx, can := context.WithTimeout(context.Background(), time.Second)
	defer can()
	_ = s.Shutdown(ctx)
}
