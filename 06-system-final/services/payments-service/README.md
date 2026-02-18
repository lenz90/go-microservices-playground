# payments-service

## What you learn
- Postgres write path + Kafka event publishing.

## How to run
```bash
go run ./cmd/server
```

## Key files
- `internal/storage/postgres.go`
- `internal/http/handler.go`
- `internal/kafka/producer.go`

## Exercises
- Add GET payment endpoint.

## Next step
Add outbox pattern.
