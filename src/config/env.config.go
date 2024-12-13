package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DB     string
	Secret string
}

func Keys() Config {
	godotenv.Load()
	return Config{
		Port:   getEnv("PORT", ":4001"),
		DB:     getEnv("MONGODB_URI", "null"),
		Secret: getEnv("SECRET_KEY", "null"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
