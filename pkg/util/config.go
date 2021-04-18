package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVar user viper to get ENV variables
func GetEnvVar(key string) string {
	// Init logger for a human-friendly, colorized output

	err := godotenv.Load("config.env")

	if err != nil {
		fmt.Printf("Error while reading config file %v", err)
	}

	value := os.Getenv(key)

	return value
}
