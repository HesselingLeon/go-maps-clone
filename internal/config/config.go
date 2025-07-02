package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMODE  string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseSSLMODE:  os.Getenv("DATABASE_SSLMODE"),
	}
}
