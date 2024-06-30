package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := os.Getenv(key)
	if env == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return env
}
