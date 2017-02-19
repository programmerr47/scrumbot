package main

import (
	"github.com/Syfaro/telegram-bot-api"
)

type UpdateAnalyzer interface {
	Analyze(update tgbotapi.Update) bool
}

type UpdateAnalyzerCompositor struct {
	analyzers[] UpdateAnalyzer
}

func NewUpdateAnalyzerCompositor(analyzers ...UpdateAnalyzer) *UpdateAnalyzerCompositor {
	compositor := new(UpdateAnalyzerCompositor)
	compositor.analyzers = analyzers
	return compositor
}

func (a UpdateAnalyzerCompositor) Analyze(update tgbotapi.Update) bool {
	for _, elem := range a.analyzers {
		success := elem.Analyze(update)

		if (success) {return true}
	}

	return false
}

type CommandAnalyzer struct {
	bot *tgbotapi.BotAPI
	commandRepliers map[string]Replier
}

func NewCommandAnalyzer(bot *tgbotapi.BotAPI, commandRepliers map[string]Replier) *CommandAnalyzer {
	analyzer := new(CommandAnalyzer)
	analyzer.bot = bot
	analyzer.commandRepliers = commandRepliers
	return analyzer
}

func (a CommandAnalyzer) Analyze(update tgbotapi.Update) bool {
	if update.Message.IsCommand() {
		commandReplier := a.commandRepliers[update.Message.Command()]

		if commandReplier != nil {
			msg := commandReplier(update)
			a.bot.Send(msg)
			return true
		}
	}

	return false;
}

type ReplyUpdateAnalyzer struct {
	bot *tgbotapi.BotAPI
	replier Replier
}

func NewReplyUpdateAnalyzer(bot *tgbotapi.BotAPI, replier Replier) *ReplyUpdateAnalyzer {
	analyzer := new(ReplyUpdateAnalyzer)
	analyzer.bot = bot
	analyzer.replier = replier
	return analyzer
}

func (a ReplyUpdateAnalyzer) Analyze(update tgbotapi.Update) bool {
	msg := a.replier(update)
	a.bot.Send(msg)
	return true
}


