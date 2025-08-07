package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("error loading .env file")
		return
	}
}

func GetString(key string, fallBack string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallBack
	}

	return value
}

func GetInt(key string, fallBack int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallBack
	}

	intVal, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("failed to convert to int")
		return fallBack
	}

	return intVal
}
