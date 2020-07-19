package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func paste(ID int64, message *tbot.Message) {
	var resp *http.Response
	var err error
	if message.ReplyToMessage.Text != "" {
		data := message.ReplyToMessage.Text
		log.Println("text", data)
		resp, err = http.Post("https://paste.rs/", "application/octet-stream", strings.NewReader(data))
	} else {
		data := message.ReplyToMessage.Document
		filelink, errf := bot.GetFileDirectURL(data.FileID)
		if errf != nil {
			log.Fatalln(errf)
		}
		filebytes, _ := http.Get(filelink)
		file, _ := ioutil.ReadAll(filebytes.Body)
		resp, err = http.Post("https://paste.rs/", "application/octet-stream", bytes.NewBuffer(file))
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
