package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

type Event struct {
	Type, UserID string
	At           time.Time
}

func main() {
	c := flag.Int("count", 1, "")
	i := flag.Duration("interval", time.Second, "")
	flag.Parse()
	b := os.Getenv("KAFKA_BROKER")
	if b == "" {
		b = "localhost:9092"
	}
	w := &kafka.Writer{Addr: kafka.TCP(b), Topic: "events.users"}
	defer w.Close()
	for x := 0; x < *c; x++ {
		id := fmt.Sprintf("u-%d", x+1)
		v, _ := json.Marshal(Event{"user.created", id, time.Now()})
		_ = w.WriteMessages(context.Background(), kafka.Message{Key: []byte(id), Value: v})
		time.Sleep(*i)
	}
}
