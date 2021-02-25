package telegram_service

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func init() {
	botKey = os.Getenv("BOT_KEY")
	telegramApi = fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", botKey)
}

var botKey string
var telegramApi string

type TelegramMessage struct {
	ChatId string
	Body   string
}

func SendMessage(msg TelegramMessage) error {
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {msg.ChatId},
			"text":    {msg.Body},
		})

	defer response.Body.Close()
	return err
}
