module playground/06-system-final/services/payments-service

go 1.22

require (
    github.com/jackc/pgx/v5 v5.6.0
    github.com/segmentio/kafka-go v0.4.47
)

replace github.com/segmentio/kafka-go => ../../../third_party/github.com/segmentio/kafka-go
replace github.com/jackc/pgx/v5 => ../../../third_party/github.com/jackc/pgx/v5
