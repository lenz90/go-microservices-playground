# 06-system-final

Mini system with API gateway, users service, payments service, and notifications consumer.

## Services
- `api-gateway` forwards `/users` and `/payments`.
- `users-service` handles user creation and publishes `user.created` events.
- `payments-service` handles payment creation, stores records in Postgres, and publishes `payment.created` events.
- `notifications-service` consumes both event types and stores latest notification per user in Redis.

## Run
```bash
cd 06-system-final
docker compose up -d --build
```

## Ports
- Gateway: `8090`
- Users: `8081`
- Payments: `8082`
- Kafka: `9092`
- Redis: `6379`
- Postgres: `5432`

## Quick demo
```bash
curl -X POST localhost:8090/users -H 'Content-Type: application/json' -d '{"id":"u1","name":"Ada","email":"ada@example.com"}'
curl -X POST localhost:8090/payments -H 'Content-Type: application/json' -d '{"id":"p1","user_id":"u1","amount":99.5}'
```

## Two Codespaces
- Option A: each codespace runs this compose separately.
- Option B: Codespace A runs compose and exposes gateway URL; Codespace B calls that URL.
