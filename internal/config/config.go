package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser   string
	DBPasswd string
	DBHost   string
	DBPort   string
	DBName   string
}

func LoadConfig() *Config {
	godotenv.Load() // load environment variables from .env file

	return &Config{
		DBUser:   getEnv("DB_USER", "root"),
		DBPasswd: getEnv("DB_PASSWD", "root"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   getEnv("DB_PORT", "3306"),
		DBName:   getEnv("DB_NAME", "go_starter_project"),
	}
}

func getEnv(key string, defaultValue string) string {
	// Use os.LookupEnv to avoid panic if the environment variable is not set
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
