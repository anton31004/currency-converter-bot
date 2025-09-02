package main

import (
	"currency-converter-bot/api"
	"currency-converter-bot/config"
	"currency-converter-bot/storage"
	"currency-converter-bot/telegram"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := telegram.InitBot(cfg.BotToken)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	err = storage.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	err = storage.CreateTable()
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	api.API = cfg.ExchangeRate
	telegram.StartBot(bot)

}
