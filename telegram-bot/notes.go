package main

import (
	"context"
	"log"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type notesData struct {
	Name    string
	Content string
}

func savenote(ID int64, msgtext string, client mongo.Client) {

	collection := client.Database("test").Collection("SavedNote")

	args := strings.Fields(msgtext)

	var check notesData
	_ = collection.FindOne(context.TODO(), bson.M{"name": args[1]}).Decode(&check)
	if (notesData{}) == check {
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
	} else {
		bot.Send(tbot.NewMessage(ID, "Note with the same name already exists."))
	}
}

func fetchallnotes(ID int64, client mongo.Client) {
	saved := "List of Saved Notes are: "
	collection := client.Database("test").Collection("SavedNote")
	ctx := context.Background()
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		defer cursor.Close(ctx)
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {
			// Declare a result BSON object
			var result notesData
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			} else {
				saved = saved + "\n /" + result.Name
			}
		}
	}
	bot.Send(tbot.NewMessage(ID, saved))
}

func fetchnote(ID int64, msgtext string, client mongo.Client) {
	collection := client.Database("test").Collection("SavedNote")
	var result notesData
	_ = collection.FindOne(context.TODO(), bson.M{"name": msgtext}).Decode(&result)
	log.Print(result.Name)
	if (notesData{}) != result {
		message := "<b>" + result.Name + "</b> : " + result.Content
		messageconfig := tbot.MessageConfig{
			BaseChat: tbot.BaseChat{
				ChatID:           ID,
				ReplyToMessageID: 0,
			},
			Text:                  message,
			ParseMode:             "html",
			DisableWebPagePreview: true,
		}
		bot.Send(messageconfig)
	} else {
		bot.Send(tbot.NewMessage(ID, "I don't know that command"))
	}
}

func deletenote(ID int64, msgtext string, client mongo.Client) {
	collection := client.Database("test").Collection("SavedNote")
	args := strings.Fields(msgtext)
	result, err := collection.DeleteOne(context.TODO(), bson.M{"name": args[1]})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
	if result.DeletedCount == 0 {
		bot.Send(tbot.NewMessage(ID, "Note doesn't exist."))
	} else {
		bot.Send(tbot.NewMessage(ID, "Note with name "+args[1]+" deleted"))
	}
}
