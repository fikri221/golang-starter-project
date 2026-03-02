package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost    string
	Port          string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	JWTSecret     string
	JWTExpiration time.Duration
}

func LoadConfig() *Config {
	godotenv.Load() // load environment variables from .env file

	return &Config{
		PublicHost:    getEnv("PUBLIC_HOST", "http://localhost"),
		Port:          getEnv("PORT", "8080"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "root"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "3306"),
		DBName:        getEnv("DB_NAME", "go_starter_project"),
		JWTSecret:     getEnv("JWT_SECRET", "secret"),
		JWTExpiration: time.Duration(getEnvInt("JWT_EXPIRATION", 3600*24*7)) * time.Second,
	}
}

func getEnv(key string, defaultValue string) string {
	// Use os.LookupEnv to avoid panic if the environment variable is not set
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}
