package config

import "os"

type Config struct{ Port, Broker string }

func Load() Config {
	return Config{Port: getenv("PORT", "8081"), Broker: getenv("KAFKA_BROKER", "localhost:9092")}
}
func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
