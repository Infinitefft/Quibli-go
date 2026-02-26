package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var AppConfig *Config

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	AppConfig = &Config{
		ServerPort: getEnv("SERVER_PORT", "3002"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5433"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),
		DBName:     getEnv("DB_NAME", "postgres"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetDBDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
	)
}
