package main

import "github.com/Syfaro/telegram-bot-api"

type Replier func(update tgbotapi.Update) tgbotapi.MessageConfig

func ReplySame(update tgbotapi.Update) tgbotapi.MessageConfig {
	return reply(update, update.Message.Text)
}

func reply(update tgbotapi.Update, text string) tgbotapi.MessageConfig {
	msg := message(update, text)
	msg.ReplyToMessageID = update.Message.MessageID
	return msg
}

func message(update tgbotapi.Update, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, text)
}
