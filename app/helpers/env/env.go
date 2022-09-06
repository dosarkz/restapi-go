package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const (
	local      = "local"
	staging    = "staging"
	production = "production"
	env        = "ENV"
)

// MustGet will return the env or panic if not present.
func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" && key != "PORT" {
		log.Fatal("Env key missing " + key)
	}
	return val
}

func MustGetInt(key string) int {
	val := os.Getenv(key)
	if val == "" && key != "PORT" {
		log.Fatal("Env key missing " + key)
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("read env variable: %v", err)
	}
	return intVal
}

func HasEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}
