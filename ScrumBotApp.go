package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
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

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)


		var msg = ReplySame(update)
		if update.Message.IsCommand() {
			log.Printf("Command: %s", update.Message.Text)
			switch update.Message.Command() {
				case startCommand:
					msg = reply(update, randString(startAnswers))
				case helpCommand:
					msg = message(update, helpAnswer)
			    default:
					continue
			}
		} else if update.Message.NewChatMember != nil && update.Message.NewChatMember.UserName != "" {
			msg = messageWithMention(update, *update.Message.NewChatMember, newMemberAnswers)
		} else if update.Message.LeftChatMember != nil && update.Message.LeftChatMember.UserName != "" {
			msg = messageWithMention(update, *update.Message.LeftChatMember, leftMemberAnswers)
		}

		bot.Send(msg)
	}
}

func messageWithMention(update tgbotapi.Update, user tgbotapi.User, answers []string) tgbotapi.MessageConfig {
	return message(update, formatIfHas(randString(answers), mentionString(user)))
}

func mentionString(user tgbotapi.User) string {
	if user.FirstName != "" {
		return user.FirstName
	} else {
		return user.UserName
	}
}
