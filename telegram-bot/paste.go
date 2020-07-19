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
	if message.ReplyToMessage != nil {
		if message.ReplyToMessage.Text != "" {
			data := message.ReplyToMessage.Text
			resp, err := http.Post("https://paste.rs/", "application/octet-stream", strings.NewReader(data))
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			ButtonLinks(ID, "Click Here", string(body), "Text pasted. Find it below")
		} else if message.ReplyToMessage.Document != nil {
			data := message.ReplyToMessage.Document
			filelink, _ := bot.GetFileDirectURL(data.FileID)
			filebytes, _ := http.Get(filelink)
			file, _ := ioutil.ReadAll(filebytes.Body)
			resp, err := http.Post("https://paste.rs/", "application/octet-stream", bytes.NewBuffer(file))
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			ButtonLinks(ID, "Click Here", string(body), "Document pasted. Find it below")
		} else {
			bot.Send(tbot.NewMessage(ID, "I don't know what to paste. Reply to a text message/document so I can paste it to paste.rs"))
		}
	} else {
		bot.Send(tbot.NewMessage(ID, "I don't know what to paste. Reply to a text message/document and use `/paste` so I can paste it to paste.rs"))
	}
}
