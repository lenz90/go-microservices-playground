package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	b := os.Getenv("KAFKA_BROKER")
	if b == "" {
		b = "localhost:9092"
	}
	r := kafka.NewReader(kafka.ReaderConfig{Brokers: []string{b}, Topic: "events.users", GroupID: "g"})
	defer r.Close()
	ctx, c := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sig; c() }()
	for {
		m, e := r.ReadMessage(ctx)
		if e != nil {
			return
		}
		fmt.Println(m.Offset)
	}
}
