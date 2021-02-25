package telegram_webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"tunturin/services/telegram-service"
)

func init() {
	// initialize stuff here
}

type Chat struct {
	Id int `json:"id"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

func HandleWebhook(update string) error {
	var body Update
	er := json.Unmarshal([]byte(update), &body)
	if er != nil {
		return fmt.Errorf("json unmarshall error %v", er)
	}
	log.Printf("Message body %v", body.Message.Text)
	log.Printf("Sending to chat_id: %d", body.Message.Chat.Id)

	var reply string
	if body.Message.Text == "hi" || body.Message.Text == "/hi" {
		reply = "Hi, my name is tunturin. How are your today? 😻"
	} else {
		return nil
	}

	msg := telegram_service.TelegramMessage{
		ChatId: strconv.Itoa(body.Message.Chat.Id),
		Body:   reply,
	}

	_ = telegram_service.SendMessage(msg)
	return nil
}
