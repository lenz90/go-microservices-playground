package kafka

import (
	"context"
	"encoding/json"
	"time"

	kgo "github.com/segmentio/kafka-go"
)

type Producer struct{ w *kgo.Writer }

func New(broker string) *Producer {
	return &Producer{w: &kgo.Writer{Addr: kgo.TCP(broker), Topic: "events.payments"}}
}
func (p *Producer) Close() error { return p.w.Close() }
func (p *Producer) PaymentCreated(ctx context.Context, id, userID string, amount float64) error {
	b, _ := json.Marshal(map[string]any{"type": "payment.created", "payment_id": id, "user_id": userID, "amount": amount, "at": time.Now().UTC()})
	return p.w.WriteMessages(ctx, kgo.Message{Key: []byte(id), Value: b})
}
