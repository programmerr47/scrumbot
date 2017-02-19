package main

import (
	"log"
	"fmt"
	"database/sql"
	"github.com/BurntSushi/toml"
	"github.com/Syfaro/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	var config Config
	_, err := toml.DecodeFile(`config.tolm`, &config)
	fmt.Println(config)
	checkErr(err)

	//temp
	db := applyTestConnectionToDatabase(config)
	defer db.Close()

	initDatabase(db)

	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	checkErr(err)

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
		analyzeUpdate(update, *bot)
	}
}

func applyTestConnectionToDatabase(config Config) *sql.DB {
	db, err := sql.Open("sqlite3", config.DB + "/test.db")
	checkErr(err)

	err = db.Ping()
	checkErr(err)
	return db
}

func initDatabase(db *sql.DB) {
	_, err := db.Exec(`
	PRAGMA FOREIGN_KEYS = ON;

	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		chat_id INTEGER,
		duration_s INTEGER,
		period_s INTEGER,
		date_s INTEGER,
		zone_offset INTEGER
	);

	CREATE TABLE IF NOT EXISTS eventAnnouncements(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		announcement TEXT,
		relative_time_s INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE
	)`)
	checkErr(err)
}

func analyzeUpdate(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	msg := ReplySame(update)
	date := time.Unix(int64(update.Message.Date), 0)
	name, offset := date.Zone()
	log.Printf("date h=%d, m=%d, s=%d, t.n=%s, t.o=%d", date.Hour(), date.Minute(), date.Second(), name, offset)
	if update.Message.IsCommand() {
		log.Printf("Command: %s", update.Message.Text)
		switch update.Message.Command() {
		case startCommand:
			msg = message(update, randString(startAnswers))
		case helpCommand:
			msg = message(update, helpAnswer)
		default:
			return
		}
	} else if update.Message.NewChatMember != nil && update.Message.NewChatMember.UserName != "" {
		msg = messageWithMention(update, *update.Message.NewChatMember, newMemberAnswers)
	} else if update.Message.LeftChatMember != nil && update.Message.LeftChatMember.UserName != "" {
		msg = messageWithMention(update, *update.Message.LeftChatMember, leftMemberAnswers)
	}

	bot.Send(msg)
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

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
