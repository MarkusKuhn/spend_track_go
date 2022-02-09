package utilities

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	return os.Getenv(key)
}
