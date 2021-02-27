package telegram_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var telegramApi string

func init() {
	telegramApi = os.ExpandEnv("https://api.telegram.org/bot${BOT_KEY}")
}

func SendMessage(msg TelegramMessage) error {
	resp, err := call("sendMessage", msg)
	defer resp.Body.Close()
	return err
}

func AnswerCallbackQuery(msg CallbackQueryAnswer) error {
	resp, err := call("answerCallbackQuery", msg)
	defer resp.Body.Close()
	return err
}

type EditMessage struct {
	ChatId    int    `json:"chat_id"`
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
}

func EditMessageText(msg EditMessage) error {
	resp, err := call("editMessageText", msg)
	defer resp.Body.Close()
	return err
}

func call(path string, v interface{}) (*http.Response, error) {
	log.Printf("Sending: %v", v)
	body, _ := json.Marshal(v)
	response, err := http.Post(fmt.Sprintf("%v/%v", telegramApi, path),
		"application/json",
		bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Failed to send message %v", err)
	}
	var r Response
	if err2 := json.NewDecoder(response.Body).Decode(&r); err2 != nil {
		log.Fatalf("Failed to decode message body %v", err2)
	}
	log.Printf("status: %v", response.Status)
	log.Println(r)
	return response, err
}

type Response struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}
