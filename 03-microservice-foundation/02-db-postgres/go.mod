module playground/03-microservice-foundation/02-db-postgres

go 1.22

require (
 github.com/go-chi/chi/v5 v5.1.0
 github.com/jackc/pgx/v5 v5.6.0
)
replace github.com/go-chi/chi/v5 => ../../third_party/github.com/go-chi/chi/v5
replace github.com/jackc/pgx/v5 => ../../third_party/github.com/jackc/pgx/v5
