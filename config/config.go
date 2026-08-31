package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	DatabaseURL string
	REDIS_ADDR  string
}

func Load() Config {
	return Config{
		Port:        getEnv("PORT", "5001"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5433/urlshortener?sslmode=disable"),
		REDIS_ADDR:  getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
