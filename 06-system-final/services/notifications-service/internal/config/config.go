package config

import "os"

type Config struct{ Broker, RedisAddr string }

func Load() Config {
	return Config{Broker: getenv("KAFKA_BROKER", "localhost:9092"), RedisAddr: getenv("REDIS_ADDR", "localhost:6379")}
}
func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
