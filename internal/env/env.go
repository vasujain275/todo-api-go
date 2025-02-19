package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) (int, error) {
	if value, ok := os.LookupEnv(key); ok {
		return strconv.Atoi(value)
	}
	return fallback, nil
}
