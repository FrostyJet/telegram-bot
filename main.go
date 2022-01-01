package main

import (
	"github.com/frostyjet/telegram-bot/pkg/bot"
)

func main() {
	bot := bot.NewApplication("5068574329:AAFhtd6wGT36v3JxNlB3C0Kjiy15qyCh2Lc")
	bot.Run()
}
