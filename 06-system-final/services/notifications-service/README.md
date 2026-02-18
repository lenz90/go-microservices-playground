# notifications-service

## What you learn
- Kafka consumers over multiple topics.
- Redis write model for latest notification state.

## How to run
```bash
go run ./cmd/server
```

## Key files
- `internal/kafka/consumer.go`
- `internal/storage/redis.go`

## Exercises
- Persist history list instead of only last value.

## Next step
Add retry and DLQ handling.
