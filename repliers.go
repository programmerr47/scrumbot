package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

type Replier func(update tgbotapi.Update) tgbotapi.MessageConfig

func ReplySame(update tgbotapi.Update) tgbotapi.MessageConfig {
	return reply(update, update.Message.Text)
}

func reply(update tgbotapi.Update, text string) tgbotapi.MessageConfig {
	log.Printf("Reply with message: %s", text)
	msg := message(update, text)
	msg.ReplyToMessageID = update.Message.MessageID
	return msg
}

func message(update tgbotapi.Update, text string) tgbotapi.MessageConfig {
	log.Printf("Answer with text: %s", text)
	return tgbotapi.NewMessage(update.Message.Chat.ID, text)
}
