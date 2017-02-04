package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("318390328:AAHHKf7vTCC0hWnrG5N2IGQaQ3-ySxd44zQ")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/") {
			log.Printf("Command: %s", update.Message.Text)
			switch update.Message.Text {
			case startCommand:
				msg := reply(update, randString(startAnswers))
				bot.Send(msg)
				continue
			}
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := ReplySame(update)
		bot.Send(msg)
	}
}
