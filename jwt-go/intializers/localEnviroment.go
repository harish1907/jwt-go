package intializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LocalEnvironmentVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
