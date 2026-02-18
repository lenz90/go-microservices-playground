package config

import "os"

type Config struct{ Port, UsersBaseURL, PaymentsBaseURL string }

func Load() Config {
	c := Config{Port: getenv("PORT", "8090"), UsersBaseURL: getenv("USERS_BASE_URL", "http://localhost:8081"), PaymentsBaseURL: getenv("PAYMENTS_BASE_URL", "http://localhost:8082")}
	return c
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
