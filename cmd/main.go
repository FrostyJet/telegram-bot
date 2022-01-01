package main

import (
	"log"
	"os"

	"github.com/frostyjet/telegram-bot/pkg/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		log.Fatal("environment variable TELEGRAM_TOKEN is not set")
		return
	}

	bot := bot.NewApplication(token)
	bot.Run()
}
