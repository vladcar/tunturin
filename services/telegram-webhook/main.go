package telegram_webhook

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	tg "tunturin/services/telegram-service"
)

var allowedChats map[int]struct{}

func init() {
	chats := os.Getenv("ALLOWED_CHATS")
	fmt.Printf("allowed chats ids: %v", chats)
	allowedChats = make(map[int]struct{})
	for _, s := range strings.Split(chats, ",") {
		id, _ := strconv.Atoi(s)
		allowedChats[id] = struct{}{}
	}
}

func HandleWebhook(update string) error {
	var body tg.Update
	er := json.Unmarshal([]byte(update), &body)
	if er != nil {
		return fmt.Errorf("json unmarshall error %v", er)
	}

	chatId := getChatId(body)
	if notAllowed(chatId) {
		return nil
	}
	var op = tg.GetOperation(body)
	return op.Execute(body)
}

func getChatId(body tg.Update) int {
	var chatId int
	if body.Message != nil {
		chatId = body.Message.Chat.Id
	} else {
		chatId = body.CallbackQuery.Message.Chat.Id
	}
	return chatId
}

func notAllowed(chatId int) bool {
	_, chatAllowed := allowedChats[chatId]
	if !chatAllowed {
		return true
	}
	return false
}
