package config

import (
	"os"
)

type Config struct {
	DB_URI     string
	JWT_SECRET string
}

func LoadConfig() *Config {
	return &Config{
		DB_URI:     os.Getenv("DB_URI"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}
