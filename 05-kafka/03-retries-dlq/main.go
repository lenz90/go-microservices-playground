package main

import (
	"context"
	"flag"
	"github.com/segmentio/kafka-go"
	"os"
	"strconv"
)

const (
	mainTopic  = "events.payments"
	retryTopic = "events.payments.retry"
	dlqTopic   = "events.payments.dlq"
)

func retryCount(h []kafka.Header) int {
	for _, x := range h {
		if x.Key == "x-retry-count" {
			n, _ := strconv.Atoi(string(x.Value))
			return n
		}
	}
	return 0
}
func produce(b string) error {
	w := &kafka.Writer{Addr: kafka.TCP(b), Topic: mainTopic}
	defer w.Close()
	return w.WriteMessages(context.Background(), kafka.Message{Key: []byte("p1"), Value: []byte(`{"payment_id":"p1"}`)})
}
func consume(b string, n int) error { _ = n; return nil }
func main() {
	m := flag.String("mode", "producer", "")
	n := flag.Int("max-retry", 3, "")
	flag.Parse()
	b := os.Getenv("KAFKA_BROKER")
	if b == "" {
		b = "localhost:9092"
	}
	if *m == "producer" {
		_ = produce(b)
	} else {
		_ = consume(b, *n)
	}
}
