package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken     string
	ExchangeRate string
	Database     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("config.env")
	if err != nil {
		return nil, err
	}
	config := &Config{
		BotToken:     os.Getenv("BOT_TOKEN"),
		ExchangeRate: os.Getenv("EXCHANGE_RATE_API"),
		Database:     os.Getenv("DATABASE_URL"),
	}

	return config, nil
}
