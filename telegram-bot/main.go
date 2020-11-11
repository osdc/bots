package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	tbot "github.com/go-telegram-bot-api/telegram-bot-api"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bot *tbot.BotAPI

//ButtonLinks sends the parsed arguements as an inline buttons to the specified chat ID
func ButtonLinks(ID int64, ButtonText string, ButtonURL string, MessageText string) {
	var button = tbot.NewInlineKeyboardMarkup(
		tbot.NewInlineKeyboardRow(
			tbot.NewInlineKeyboardButtonURL(ButtonText, ButtonURL),
		),
	)
	msg := tbot.NewMessage(ID, MessageText)
	msg.ReplyMarkup = button
	bot.Send(msg)
}

//Streams tweets with the help of Twitter API , bot sends the link of any new tweet or tweet related activity to the specified group on telegram.
func tweet() {
	//Initializing twitter API
	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	params := &twitter.StreamFilterParams{
		Follow: []string{os.Getenv("OSDC_TWITTER_ID")},
	}
	stream, _ := client.Streams.Filter(params)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.IDStr)
		GroupID, _ := strconv.ParseInt(os.Getenv("OSDC_GROUP_ID"), 10, 64)

		bot.Send(tbot.NewMessage(GroupID, "New Tweet Alert: \n https://twitter.com/osdcjiit/status/"+tweet.IDStr))
	}
	demux.HandleChan(stream.Messages)

}

func me(user string, ID int64, msgtext string) {

	s := strings.SplitN(msgtext, " ", 2)
	if len(s) == 2 {
		bot.Send(tbot.NewMessage(ID, "* @"+user+" "+s[1]+" *"))
	}
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
	/addmeetup* - Add the details of next OSDC Meetup.` + "\n" + `Format : /addmeetup <Title> <DD/MM/YYYY Hr:Min> <Venue>
	/nextmeetup - Get the details of next OSDC Meetup
	To contribute to|modify this bot : https://github.com/osdc/bots
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

//extracts the details of the user who sent the message to check whether the user is creator/admin. Returns true in this case else false.
func memberdetails(ID int64, userid int) bool {
	response, _ := bot.GetChatMember(tbot.ChatConfigWithUser{
		ChatID: ID,
		UserID: userid,
	})
	if response.IsCreator() || response.IsAdministrator() {
		return true
	}
	return false
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
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Running twitter stream concurrently
	go tweet()

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
					addmeetup(ID, update.Message.Text, *client)
				} else {
					bot.Send(tbot.NewMessage(ID, "Sorry, only admins can add the details of next meetup."))
				}
			case "nextmeetup":
				nextmeetup(ID, *client)
			case "save":
				check := memberdetails(ID, update.Message.From.ID)
				if check == true {
					savenote(ID, update.Message.Text, *client)
				} else {
					bot.Send(tbot.NewMessage(ID, "Sorry, only admins can add the details of next meetup."))
				}
			case "notes":
				fetchallnotes(ID, *client)
			case "deletenote":
				deletenote(ID, update.Message.Text, *client)
			case "fetchnote":
				fetchnote(ID, update.Message.Text, *client)
			case "paste":
				paste(ID, update.Message)
			case "me":
				me(update.Message.From.UserName, ID, update.Message.Text)
				bot.DeleteMessage(tbot.NewDeleteMessage(ID, update.Message.MessageID))
			default:
				{
					msg, err := bot.Send(tbot.NewMessage(ID, "I don't know this command"))

					log.Print(err)
					timer1 := time.NewTimer(5 * time.Second)
					<-timer1.C

					bot.DeleteMessage(tbot.NewDeleteMessage(ID, msg.MessageID))
				}
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
