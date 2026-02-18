# go-microservices-playground

A complete hands-on playground to learn Go from basics to a mini event-driven microservice system.

## Learning path

1. `00-go-basics`: language fundamentals, interfaces, and concurrency.
2. `01-cli-and-testing`: CLI patterns and testing strategies.
3. `02-http-api`: HTTP APIs from `net/http` to middleware and config/logging.
4. `03-microservice-foundation`: clean architecture, Postgres CRUD, and migrations.
5. `04-redis`: cache-aside and fixed-window rate limiting.
6. `05-kafka`: producer/consumer and retries with DLQ.
7. `06-system-final`: a mini system with gateway + users + payments + notifications.

## Shared infrastructure

Use root compose for learning modules:

```bash
./scripts/up.sh
./scripts/down.sh
```

Services and ports:
- Redpanda Kafka: `localhost:9092`
- Redpanda Console: `localhost:8080`
- Redis: `localhost:6379`
- Postgres: `localhost:5432`

## Environment variables

Copy and edit:

```bash
cp env.example .env
```

Code defaults to localhost values but always reads from env, making it two-workspace friendly.

## How to run modules

Each module folder contains its own `go.mod` and `README.md` with exact run commands.

Example:

```bash
cd 02-http-api/02-router
go test ./...
go run .
```

## Two Codespaces workflows

### Option A: both run independently
- In Codespace A and B, run `./scripts/up.sh`.
- Run module code in each space against local infra.

### Option B: split client/server workspaces
- Codespace A runs `06-system-final` stack and exposes gateway port `8090` publicly.
- Codespace B points `API_GATEWAY_URL` (or curl target) to A's exposed URL.
- Gateway downstream URLs are configurable via env:
  - `USERS_BASE_URL`
  - `PAYMENTS_BASE_URL`

## Quick validation

Run tests module-by-module:

```bash
find . -name go.mod -not -path './.git/*' -execdir go test ./... \;
```
