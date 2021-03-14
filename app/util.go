package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	mode := os.Getenv("MODE")

	if mode != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Print("WARNING: No .env file found")
		}
	}
}
