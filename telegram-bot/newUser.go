package main

import (
	"fmt"
	"math/rand"
	"time"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func welcome(user tbot.User, ID int64) {
	User := fmt.Sprintf("[%v](tg://user?id=%v)", user.FirstName, user.ID)
	var botSlice = make([]string, 0)
	botSlice = append(botSlice,
		"helping the hackers in here.",
		"a bot made by the geeks for the geeks.",
		"an Autobot on Earth to promote open source.",
		"a distant cousin of the mars rover.",
		"a friendly bot written in Golang.",
	)
	var quesSlice = make([]string, 0)
	quesSlice = append(quesSlice,
		"which language do you work with?",
		"what do you want to learn?",
		"what is your current operating system?",
		"if you are new to open source.",
		"which technology fascinates you the most?",
		"what have you been exploring recently?",
	)

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(botSlice) - 1
	randomNum1 := (rand.Intn(max-min+1) + min)
	randomNum2 := (rand.Intn(max-min+1) + min)
	var welcomeMessage1 = fmt.Sprintf(botSlice[randomNum1])
	var welcomeMessage2 = fmt.Sprintf(quesSlice[randomNum2])
	reply := tbot.NewMessage(ID, "Welcome "+User+", I am the OSDC-bot, "+welcomeMessage1+
		" Please introduce yourself. To start with, you can tell us "+welcomeMessage2)
	reply.ParseMode = "markdown"
	bot.Send(reply)
}
