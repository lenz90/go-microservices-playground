package kafka

import (
	"context"
	"encoding/json"
	"time"

	kgo "github.com/segmentio/kafka-go"
)

type Producer struct{ w *kgo.Writer }

func NewProducer(broker string) *Producer {
	return &Producer{w: &kgo.Writer{Addr: kgo.TCP(broker), Topic: "events.users"}}
}
func (p *Producer) Close() error { return p.w.Close() }
func (p *Producer) UserCreated(ctx context.Context, id, email string) error {
	b, _ := json.Marshal(map[string]any{"type": "user.created", "user_id": id, "email": email, "at": time.Now().UTC()})
	return p.w.WriteMessages(ctx, kgo.Message{Key: []byte(id), Value: b})
}
