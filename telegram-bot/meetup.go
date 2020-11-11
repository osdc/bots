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
	Venue venue  `json:"venue"`
	Date  string `json:"local_date"`
	Time  string `json:"local_time"`
	Group group  `json:"group"`
	Link  string `json:"link"`
}

//struct to keep track of OSDC meetups using a JSON file
type meetupdata struct {
	Name  string
	Date  time.Time
	Venue string
}

//fetching details of meetup of the group's url from Meetup API
func getMeetups(url string, ID int64) {
	resp, err := http.Get(url)
	if err != nil {
		sendErrorToChan(err, ID, "Error when retrieving url:"+url)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sendErrorToChan(err, ID, "Error when parsing body of url:"+url)
		return
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
	if len(args) == 5 {
		//separating the date and hour arguments so that we can use the
		// dd/mm/yyyy and hh:min format
		separatedDateStr := strings.Split(args[2], "/")
		separatedHourStr := strings.Split(args[3], ":")

		//converting the arguments into integers to pass them to time.Date
		day, err := strconv.Atoi(separatedDateStr[0])
		if err != nil {
			sendErrorToChan(err, ID, "Couldn't convert the day, please try again")
			return
		}
		month, err := strconv.Atoi(separatedDateStr[1])
		if err != nil {
			sendErrorToChan(err, ID, "Couldn't convert the month, please try again")
			return
		}
		year, err := strconv.Atoi(separatedDateStr[2])
		if err != nil {
			sendErrorToChan(err, ID, "Couldn't convert the year, please try again")
			return
		}
		hour, err := strconv.Atoi(separatedHourStr[0])
		if err != nil {
			sendErrorToChan(err, ID, "Couldn't convert the hour, please try again")
			return
		}
		minutes, err := strconv.Atoi(separatedHourStr[1])
		if err != nil {
			sendErrorToChan(err, ID, "Couldn't convert the minute, please try again")
			return
		}
		location, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			sendErrorToChan(err, ID, "Couldn't detect the timezone, please try again")
			return
		}
		data := meetupdata{
			Name:  args[1],
			Date:  time.Date(year, time.Month(month), day, hour, minutes, 0, 0, location),
			Venue: args[4],
		}
		_, err = collection.DeleteMany(context.TODO(), bson.D{{}})
		if err != nil {
			sendErrorToChan(err, ID, "Cannot DeleteMany")
			return
		}
		_, err = collection.InsertOne(context.TODO(), data)
		if err != nil {
			sendErrorToChan(err, ID, "Cannot inser one")
			return
		}
		bot.Send(tbot.NewMessage(ID, "Meetup added successfully."))

		//making a formatted string from time to pass it to s.At(). This will
		//specify the time at which the reminder is sent. Currently 2 hours
		//before the meetup
		remindTime := fmt.Sprintf("%d:%d", hour-2, minutes)

		//making a new scheduler
		s1 := gocron.NewScheduler(location)

		//assigning the scheduler a task and then starting it
		s1.Every(1).Day().At(remindTime).Do(reminder, ID, client, s1)
		s1.Start()
	} else {
		bot.Send(tbot.NewMessage(ID, "Please provide the details of meetup"+
			" in this format - /addmeetup <Title> <DD/MM/YY Hr:Min> <Venue> "+
			"\nTIME IN 24Hr FORMAT"))
	}
}

//fetching the details (Next Meetup Title & Date) from the JSON file
func nextmeetup(ID int64, client mongo.Client) {
	collection := client.Database("test").Collection("meetups")
	var data meetupdata
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&data)
	if err != nil {
		sendErrorToChan(err, ID, "No meetup scheduled.")
		return
	} else {
		//the string inside Format method is a sample string to specify the display
		//format of meetupTimeString
		timeString := data.Date.Local().Format("Mon _2 Jan 2006")
		nxtMeetupData := "Details of next OSDC Meetup :" + "\n" + "Title -" + "\t" +
			data.Name + "\n" + "Date -" + "\t" + timeString + "\n" + "Time -" + "\t" +
			data.Date.Local().Format("15:04") + "\n" + "Venue -" + "\t" + data.Venue
		bot.Send(tbot.NewMessage(ID, nxtMeetupData))
	}
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		sendErrorToChan(err, ID, "Couldn't detect the timezone while saving the meetup, please try again")
		return
	}
	//the string inside Format method is a sample string to specify the display
	//format of meetupTimeString
	timeString := data.Date.In(location).Format("Mon _2 Jan 2006")
	nxtMeetupData := "Details of next OSDC Meetup :" + "\n" + "Title -" + "\t" +
		data.Name + "\n" + "Date -" + "\t" + timeString + "\n" + "Time -" + "\t" +
		data.Date.In(location).Format("15:04") + "\n" + "Venue -" + "\t" + data.Venue
	bot.Send(tbot.NewMessage(ID, nxtMeetupData))
}

func reminder(ID int64, client mongo.Client, s1 *gocron.Scheduler) {
	collection := client.Database("test").Collection("meetups")
	var data meetupdata
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		sendErrorToChan(err, ID, "Couldn't detect the timezone while processing the reminder, please try again")
		return
	}
	//reminder will be sent only if the the meetup is today
	if time.Now().In(location).Day() == data.Date.Day() {
		//the string inside Format method is a sample string to specify the display
		//format of meetupTimeString
		timeString := data.Date.In(location).Format("Mon _2 Jan 2006")
		nxtMeetupData := "MEETUP REMINDER!" + "\n" + "Title -" + "\t" +
			data.Name + "\n" + "Date -" + "\t" + timeString + "\n" + "Time -" +
			"\t" + data.Date.In(location).Format("15:04") + "\n" +
			"Venue -" + "\t" + data.Venue
		bot.Send(tbot.NewMessage(ID, nxtMeetupData))
		//stops the scheduler when reminder is sent
		s1.Clear()
	}
}

func sendErrorToChan(err error, ID int64, errorTxt string) {
	log.Print(err)
	bot.Send(tbot.NewMessage(ID, errorTxt))
}
