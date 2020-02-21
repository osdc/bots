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

func ButtonLinks(ID int64, ButtonText string, ButtonUrl string, MessageText string) {
	var button = tbot.NewInlineKeyboardMarkup(
		tbot.NewInlineKeyboardRow(
			tbot.NewInlineKeyboardButtonURL(ButtonText, ButtonUrl),
		),
	)
	msg := tbot.NewMessage(ID, MessageText)
	msg.ReplyMarkup = button
	bot.Send(msg)
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
	var welcomeSlice = make([]string, 0)
	welcomeSlice = append(welcomeSlice,
		fmt.Sprintf("Welcome %s, Welcome to OSDC, enjoy your stay :) Leave: Hopefully you come back to OSDC", User),
		fmt.Sprintf("%s just slid into the server.", User),
		fmt.Sprintf("Welcome, %s. We were expecting you ( ͡° ͜ʖ ͡°)", User),
		fmt.Sprintf("%s just showed up. Hold my beer.", User),
		fmt.Sprintf("%s joined your party.", User),
		fmt.Sprintf("Roses are red, Violets are blue %s joined the chat with you.", User),
		fmt.Sprintf("%s just joined. Can I get a heal.", User),
		fmt.Sprintf("%s is here as the prophecy foretold.", User),
		fmt.Sprintf("%s has joined the battle bus.", User),
		fmt.Sprintf("Welcome %s. Leave your weapons by the door.", User),
		fmt.Sprintf("%s has arrived. Party's over.", User),
		fmt.Sprintf("Welcome %s. We hope you brought pizza.", User),
		fmt.Sprintf("Swoooosh. %s just landed.", User),
		fmt.Sprintf("Is it %s you are looking for?", User),
		fmt.Sprintf("Ready player %s?", User),
		fmt.Sprintf("%s just joined. Everyone, look busy.", User),
		fmt.Sprintf("Brace yourselves. %s just joined the chat.", User),
	)
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(welcomeSlice) - 1
	randomNum := (rand.Intn(max-min+1) + min)
	var welcomeMessage = welcomeSlice[randomNum]
	reply1 := tbot.NewMessage(ID, welcomeMessage)
	reply1.ParseMode = "markdown"
	const introString = "Please introduce yourself."
	reply2 := tbot.NewMessage(ID, introString)
	reply2.ParseMode = "markdown"
	bot.Send(reply1)
	bot.Send(reply2)
}

//extracts the details of the user who sent the message to check whether the user is creator/admin. Returns true in this case else false.
func memberdetails(ID int64, userid int) bool {
	response, _ := bot.GetChatMember(tbot.ChatConfigWithUser{
		ChatID: ID,
		UserID: userid,
	})
	if response.IsCreator() || response.IsAdministrator() {
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
				ButtonLinks(ID, "Click Here", "https://t.me/jiitosdc", "Hello , I'm OSDC Bot, Use /help to know more. To join OSDC group:")
			case "help":
				help(ID)
			case "github":
				ButtonLinks(ID, "Github", "https://github.com/osdc", "Checkout our OSDC projects and feel free to contribute!")
			case "telegram":
				ButtonLinks(ID, "Click Here", "https://t.me/jiitosdc", "Join our telegram group:")
			case "twitter":
				ButtonLinks(ID, "Click Here", "https://twitter.com/osdcjiit", "Check us out on twitter")
			case "website":
				ButtonLinks(ID, "Click Here", "https://osdc.netlify.com", "Website:")
			case "blog":
				ButtonLinks(ID, "Click Here", "https://osdcblog.netlify.com/", "Blogs written by the folks at the Open Source Developers' Club who live in and around JIIT, Noida, India.")
			case "irc":
				ButtonLinks(ID, "Click Here", "https://github.com/osdc/community-committee/wiki/IRC", "Join us on IRC server of Freenode at #jiit-lug. To get started refer our IRC wiki-")
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
