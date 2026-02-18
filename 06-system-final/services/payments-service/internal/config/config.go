package config

import "os"

type Config struct{ Port, Broker, DSN string }

func Load() Config {
	return Config{Port: getenv("PORT", "8082"), Broker: getenv("KAFKA_BROKER", "localhost:9092"), DSN: getenv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/playground?sslmode=disable")}
}
func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
