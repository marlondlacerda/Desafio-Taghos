package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(".env file not exists")
	}
	return os.Getenv(key)
}

// configs.GetEnvOrDefault("DATABASE_MONGODB", "localhost:27017")
func GetEnvOrDefault(key, defaultValue string) string {
	if value := Env(key); value != "" {
		return value
	}
	return defaultValue
}

const (
	ServerName = "Desafio Taghos"
)

var (
	// Host é o host do servidor
	Host = GetEnvOrDefault("HOST", "127.0.0.1:3000")
)
