package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var err error

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return err
	}
	return nil
}

// GetEnv Initialize env variables
func getEnv(key string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return "NotFound"
}

func init() {
	err = LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
}
