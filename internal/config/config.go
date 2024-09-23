package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host        string
	Port        string
	PostgresURL string
}

func New() *Config {
	
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file", slog.String("err msg:", err.Error()))
		slog.Info("shutdown application")
		os.Exit(1)
	}

	cfg := &Config{
		Host: os.Getenv("PASTEBIN_HOST"),
		Port: os.Getenv("PASTEBIN_PORT"),
		PostgresURL: os.Getenv("POSTGRES_URL"),
	}

	return cfg

}