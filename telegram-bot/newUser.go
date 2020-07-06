package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type newUser struct {
	UserName     string
	FirstName    string
	UserID       int
	Introduction bool
	JoinDate     time.Time
}

func welcome(user tbot.User, ID int64, client mongo.Client) {
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

	collection := client.Database("test").Collection("user")
	data := newUser{
		UserName:     user.UserName,
		FirstName:    user.FirstName,
		UserID:       user.ID,
		Introduction: false,
		JoinDate:     time.Now(),
	}
	log.Println(data)
	check, err := collection.InsertOne(context.TODO(), data)
	log.Println(check)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(botSlice) - 1
	randomNum1 := (rand.Intn(max-min+1) + min)
	randomNum2 := (rand.Intn(max-min+1) + min)
	var welcomeMessage1 = fmt.Sprintf(botSlice[randomNum1])
	var welcomeMessage2 = fmt.Sprintf(quesSlice[randomNum2])
	reply := tbot.NewMessage(ID, "Welcome "+User+", I am the OSDC-bot, "+welcomeMessage1+
		" Please introduce yourself. To start with, you can tell us "+welcomeMessage2+" Start your message with /introduction so that I can verify you.")
	reply.ParseMode = "markdown"
	bot.Send(reply)
}

func introverify(user *tbot.User, ID int64, client mongo.Client) {
	collection := client.Database("test").Collection("user")
	filter := bson.M{"userid": user.ID, "introduction": false}
	update := bson.M{"$set": bson.M{"introduction": true}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println("no documents")
			bot.Send(tbot.NewMessage(ID, "No need to introduce again, I already know you."))
			return
		}
		log.Fatal(err)
	}
	if result.MatchedCount > 0 {
		bot.Send(tbot.NewMessage(ID, "Thanks for introducing yourself, You are now a verified member of OSDC :) "))
	} else {
		bot.Send(tbot.NewMessage(ID, "No need to introduce again, I already know you."))
	}

}
