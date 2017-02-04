package main

import "github.com/Syfaro/telegram-bot-api"

type Replier func(update tgbotapi.Update) tgbotapi.MessageConfig

func ReplySame(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	return msg
}
