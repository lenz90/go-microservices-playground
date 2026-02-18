module playground/06-system-final/services/notifications-service

go 1.22

require (
    github.com/redis/go-redis/v9 v9.6.1
    github.com/segmentio/kafka-go v0.4.47
)

replace github.com/segmentio/kafka-go => ../../../third_party/github.com/segmentio/kafka-go
replace github.com/redis/go-redis/v9 => ../../../third_party/github.com/redis/go-redis/v9
