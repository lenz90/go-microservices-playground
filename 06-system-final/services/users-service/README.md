# users-service

## What you learn
- REST endpoint + Kafka event publishing.

## How to run
```bash
go run ./cmd/server
```

## Key files
- `internal/http/handler.go`
- `internal/kafka/producer.go`

## Exercises
- Persist users to Postgres.

## Next step
Add idempotency checks.
