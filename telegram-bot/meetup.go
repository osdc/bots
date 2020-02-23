package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var finallist string //variable to store final list of meetups

//-------structs to parse json data from Meetup API
type venue struct {
	Venue string `json:"name"`
}
type group struct {
	GroupName string `json:"name"`
}
type meetuplist struct {
	Name  string `json:"name"`
	Venue venue  `json: "venue"`
	Date  string `json:"local_date"`
	Time  string `json:"local_time"`
	Group group  `json: "group"`
	Link  string `json:"link"`
}

//struct to keep track of OSDC meetups using a JSON file
type meetupdata struct {
	Name  string
	Date  string
	Venue string
}

//fetching details of meetup of the group's url from Meetup API
func getMeetups(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var meetups []meetuplist
	json.Unmarshal([]byte(body), &meetups)
	if (len(meetups)) > 0 {
		finallist = finallist + "Title -" + "\t" + meetups[0].Name + "\n" + "Date -" + "\t" + meetups[0].Date + "\n" + "Link -" + "\t" + meetups[0].Link + "\n\n"
	}
}

//Slicing the arguments (title & date of next meetup) from message text and writing them to a json file.
func addmeetup(ID int64, msgtext string, client mongo.Client) {
	// Get a handle for your collection
	collection := client.Database("test").Collection("meetups")

	args := strings.Fields(msgtext)
	if len(args) == 4 {
		data := meetupdata{
			Name:  args[1],
			Date:  args[2],
			Venue: args[3],
		}
		_, err := collection.DeleteMany(context.TODO(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
		}
		_, err = collection.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
		}
		bot.Send(tbot.NewMessage(ID, "Meetup added successfully."))
	} else {
		bot.Send(tbot.NewMessage(ID, "Please provide the details of meetup in this format - /addmeetup <Title> <Date> <Venue>"))
	}
}

//fetching the details (Next Meetup Title & Date) from the JSON file
func nextmeetup(ID int64, client mongo.Client) {
	collection := client.Database("test").Collection("meetups")
	var result meetupdata
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	nxtmeetupdata := "Details of next OSDC Meetup :" + "\n" + "Title -" + "\t" + result.Name + "\n" + "Date -" + "\t" + result.Date + "\n" + "Venue -" + result.Venue
	bot.Send(tbot.NewMessage(ID, nxtmeetupdata))
}
