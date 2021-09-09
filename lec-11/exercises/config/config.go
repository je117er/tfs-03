package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
