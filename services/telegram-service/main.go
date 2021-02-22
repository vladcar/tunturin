package telegram_service

import (
	"encoding/json"
	"fmt"
)

func init() {
	// initialize stuff here
}

type TelegramMessage struct {
	id string
}

//todo implement
func SendMessage(msg string) (TelegramMessage, error) {
	var tgMessage TelegramMessage
	err := json.Unmarshal([]byte(msg), &tgMessage)
	if err != nil {
		fmt.Println("error:", err)
	}
	return tgMessage, nil
}
