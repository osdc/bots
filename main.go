package main

import (
	"log"
	"os"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tbot.BotAPI

func github(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://github.com/osdc"))
}

func telegram(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://t.me/jiitosdc"))
}

func help(ID int64) {
	msg := ` Use one of the following commands
	/github - to get a link to OSDC's Github page
	/telegram - to get an invite link for OSDC's Telegram Group
	To contribute to|modify this bot : https://github.com/vaibhavk/osdc-bots
	`
	bot.Send(tbot.NewMessage(ID, msg))
}

func start() {
	var err error
	bot, err = tbot.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func main() {
	start()
	u := tbot.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		ID := update.Message.Chat.ID
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				help(ID)
			case "github":
				github(ID)
			case "telegram":
				telegram(ID)
			default:
				bot.Send(tbot.NewMessage(ID, "I don't know that command"))
			}
		}
		if update.Message.NewChatMembers != nil {
			for _, user := range *(update.Message.NewChatMembers) {
				message := "Welcome "
				if user.UserName != "" {
					message += "@" + user.UserName
				} else {
					message += user.FirstName
				}
				message += ", please introduce yourself"
				bot.Send(tbot.NewMessage(ID, message))
			}
		}
	}
}
