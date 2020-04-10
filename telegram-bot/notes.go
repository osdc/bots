package main

import (
	"context"
	"log"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type notesData struct {
	Name    string
	Content string
}

func savenote(ID int64, msgtext string, client mongo.Client) {

	collection := client.Database("test").Collection("SavedNote")

	args := strings.Fields(msgtext)

	if len(args) >= 3 {
		data := notesData{
			Name:    args[1],
			Content: strings.Join(args[2:], " "),
		}
		log.Print(data)
		_, err := collection.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
		}
		bot.Send(tbot.NewMessage(ID, "Note added successfully."))
	} else {
		bot.Send(tbot.NewMessage(ID, "Please provide the details of note in this format - /save <Title> <Content>"))
	}
}
