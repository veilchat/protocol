package protocol

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() error {
	return godotenv.Load()
}

func Load() {
	if env_err := SetupEnv(); env_err != nil {
		log.Fatalf("Error loading .env: %v\n", env_err)
	}
}
