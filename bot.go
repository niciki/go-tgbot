package main

import (
	"bot/commands"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

TELEGRAM_TOKEN := ....

func main() {
	tgbot, err := tgbotapi.NewBotAPI(TELEGRAM_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	tgbot.Debug = true
	log.Printf("Authorized on account %s", tgbot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := tgbot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("Error during open updates chan")
	}
	ch := commands.NewCommandHandler(tgbot)
	//BotHandler := commands.NewCommandHandler(tgbot)
	for update := range updates {
		//BotHandler.Handle(update)
		ch.Handle(update)
	}
}
