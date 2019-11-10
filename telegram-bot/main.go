package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/anaskhan96/soup"
	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tbot.BotAPI

func start(ID int64) {
	bot.Send(tbot.NewMessage(ID, "Hello , I'm OSDC Bot, Use /help to know more. To join OSDC group: https://t.me/jiitosdc"))
}

func github(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://github.com/osdc"))
}

func telegram(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://t.me/jiitosdc"))
}

func twitter(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://twitter.com/osdcjiit"))
}

func website(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://osdc.netlify.com"))
}

func blog(ID int64) {
	bot.Send(tbot.NewMessage(ID, "https://osdc.github.io/blog"))
}

func irc(ID int64) {
	bot.Send(tbot.NewMessage(ID, "Join us on IRC server of Freenode at #jiit-lug. To get started refer our IRC wiki- https://github.com/osdc/community-committee/wiki/IRC ."))
}

//scraping xkcd strip URL from its website with the help of a random generated integer and then sending it as a photo using NewPhotoShare Telegram API method.
func xkcd(ID int64) {
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 2000
	randomnum := (rand.Intn(max-min+1) + min)
	xkcdurl := "https://xkcd.com/" + strconv.Itoa(randomnum)
	resp, err := soup.Get(xkcdurl)
	if err == nil {
		doc := soup.HTMLParse(resp)
		links := doc.Find("div", "id", "comic").FindAll("img")
		for _, link := range links {
			imglink := link.Attrs()["src"]
			fullurl := "https:" + imglink
			fmt.Println(fullurl)
			bot.Send(tbot.NewPhotoShare(ID, fullurl))
		}
	} else {
		fullurl := "https://imgs.xkcd.com/comics/operating_systems.png"
		bot.Send(tbot.NewPhotoShare(ID, fullurl))
	}
}

func help(ID int64) {
	msg := ` Use one of the following commands to :
	/github - Get a link to OSDC's Github page.
	/telegram - Get an invite link for OSDC's Telegram Group.
	/twitter - Get the link of OSDC's twitter account.
	/website - Get the link of the official website of OSDC.
	/blog - Get the link of the OSDC blog.
	/irc - Find us on IRC.
	/xkcd - Get a xkcd comic for you.
	/dlmeetups - Get the list of upcoming meetups of Delhi/NCR communities.
	/addmeetup* - Add the details of next OSDC Meetup.` + "\n" + `Format : /addmeetup <Title> <DD/MM/YYYY>
	/nextmeetup - Get the details of next OSDC Meetup
	To contribute to|modify this bot : https://github.com/vaibhavk/osdc-bots
	* - Only for channel admins.
	`
	bot.Send(tbot.NewMessage(ID, msg))
}

//list of meetup groups urlnames,keeps sending one by one to getmeetups() function in meetup.go
func dlmeetups(ID int64) {
	finallist = ""
	urlnames := []string{"ilugdelhi", "pydelhi", "GDGNewDelhi", "gdgcloudnd", "Paytm-Build-for-India", "jslovers", "Gurgaon-Go-Meetup", "Mozilla_Delhi", "PyDataDelhi", "React-Delhi-NCR", "OWASP-Delhi-NCR-Chapter"}
	for _, each := range urlnames {
		url := "http://api.meetup.com/" + each + "/events"
		getMeetups(url)
	}
	bot.Send(tbot.NewMessage(ID, finallist))
}

func welcome(user tbot.User, ID int64) {
	User := fmt.Sprintf("[%v](tg://user?id=%v)", user.FirstName, user.ID)
	reply := tbot.NewMessage(ID, "**Welcome** "+User+", please introduce yourself")
	reply.ParseMode = "markdown"
	bot.Send(reply)
}

//extracts the details of the user who sent the message to check whether the user is creator/admin. Returns true in this case else false.
func memberdetails(ID int64, userid int) bool {
	response, _ := bot.GetChatMember(tbot.ChatConfigWithUser{
		ChatID: ID,
		UserID: userid,
	})
	if response.Status == "creator" || response.Status == "admin" {
		return true
	} else {
		return false
	}
}

func kickUser(user int, ID int64) {
	bot.KickChatMember(tbot.KickChatMemberConfig{
		ChatMemberConfig: tbot.ChatMemberConfig{
			ChatID: ID,
			UserID: user,
		},
		UntilDate: time.Now().Add(time.Hour * 24).Unix(),
	})
}

func main() {
	var err error
	bot, err = tbot.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		fmt.Println("error in auth")
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
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

			case "start":
				start(ID)
			case "help":
				help(ID)
			case "github":
				github(ID)
			case "telegram":
				telegram(ID)
			case "twitter":
				twitter(ID)
			case "website":
				website(ID)
			case "blog":
				blog(ID)
			case "irc":
				irc(ID)
			case "xkcd":
				xkcd(ID)
			case "dlmeetups":
				dlmeetups(ID)
			case "addmeetup":
				check := memberdetails(ID, update.Message.From.ID)
				if check == true {
					addmeetup(ID, update.Message.Text)
				} else {
					bot.Send(tbot.NewMessage(ID, "Sorry, only admins can add the details of next meetup."))
				}
			case "nextmeetup":
				nextmeetup(ID)
			default:
				bot.Send(tbot.NewMessage(ID, "I don't know that command"))
			}
		}
		if update.Message.NewChatMembers != nil {
			for _, user := range *(update.Message.NewChatMembers) {
				if user.IsBot && user.UserName != "osdcbot" {
					go kickUser(user.ID, ID)
				} else {
					go welcome(user, ID)
				}
			}
		}
	}
}
