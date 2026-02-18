module playground/04-redis/01-cache-aside

go 1.22

require (
 github.com/go-chi/chi/v5 v5.1.0
 github.com/redis/go-redis/v9 v9.6.1
)
replace github.com/go-chi/chi/v5 => ../../third_party/github.com/go-chi/chi/v5
replace github.com/redis/go-redis/v9 => ../../third_party/github.com/redis/go-redis/v9
