package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
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
	Date  time.Time
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
	if len(args) == 8 {
		//converting the arguments into integers to pass them to time.Date
		day, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln(err)
		}
		month, err := strconv.Atoi(args[3])
		if err != nil {
			log.Fatalln(err)
		}
		year, err := strconv.Atoi(args[4])
		if err != nil {
			log.Fatalln(err)
		}
		hour, err := strconv.Atoi(args[5])
		if err != nil {
			log.Fatalln(err)
		}
		minutes, err := strconv.Atoi(args[6])
		if err != nil {
			log.Fatalln(err)
		}
		data := meetupdata{
			Name:  args[1],
			Date:  time.Date(year, time.Month(month), day, hour, minutes, 0, 0, time.Local),
			Venue: args[7],
		}
		_, err = collection.DeleteMany(context.TODO(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
		}
		_, err = collection.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
		}
		bot.Send(tbot.NewMessage(ID, "Meetup added successfully."))
		nextmeetup(ID, client)

		//making a formatted string from time to pass it to s.At(). This will
		//specify the time at which the reminder is sent. Currently 2 hours
		//before the meetup
		remindTime := fmt.Sprintf("%d:%d:00", hour-1, minutes)

		//making a new scheduler
		s1 := gocron.NewScheduler(time.Local)

		//assigning the scheduler a task and then starting it
		s1.Every(1).Day().At(remindTime).Do(reminder, ID, client, s1)
		s1.Start()
	} else {
		bot.Send(tbot.NewMessage(ID, "Please provide the details of meetup in this format - /addmeetup <Title> <DD MM YY Hr Min> <Venue> \nTIME IN 24Hr FORMAT"))
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
	//the string inside Format method is a sample string to specify the display
	//format of meetupTimeString
	timeString := result.Date.Local().Format("Mon Jan _2 15:04:05 2006")
	nxtMeetupData := "Details of next OSDC Meetup :" + "\n" + "Title -" + "\t" + result.Name + "\n" + "Date -" + "\t" + timeString + "\n" + "Venue - " + result.Venue
	bot.Send(tbot.NewMessage(ID, nxtMeetupData))
}

func reminder(ID int64, client mongo.Client, s1 *gocron.Scheduler) {
	collection := client.Database("test").Collection("meetups")
	var data meetupdata
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	//the string inside Format method is a sample string to specify the display
	//format of meetupTimeString
	timeString := data.Date.Local().Format("Mon Jan _2 15:04:05 2006")
	nxtMeetupData := "MEETUP REMINDER!" + "\n" + "Title -" + "\t" + data.Name + "\n" + "Date -" + "\t" + timeString + "\n" + "Venue - " + data.Venue
	if !time.Now().Local().Before(data.Date) {
		//stops the scheduler when meetup is done
		s1.Clear()
	} else {
		bot.Send(tbot.NewMessage(ID, nxtMeetupData))
	}
}
