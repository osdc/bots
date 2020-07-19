package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func paste(ID int64, message *tbot.Message) {
	data := message.ReplyToMessage.Text
	log.Println(data)
	resp, err := http.Post("https://paste.rs/", "application/octet-stream", strings.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("hello", string(body))
}
