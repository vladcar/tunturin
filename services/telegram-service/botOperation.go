package telegram_service

import (
	"log"
	"strings"
)

type BotOperation interface {
	Execute(update Update) error
}

type GreetingOperation struct{}
type QuestionOperation struct{}
type DefaultOperation struct{}
type Inline struct{}
type CallbackQueryOperation struct{}

//todo refactor this
func GetOperation(update Update) BotOperation {
	if update.CallbackQuery != nil {
		return CallbackQueryOperation{}
	}

	text := strings.ReplaceAll(update.Message.Text, "/", "")
	if text == "hi" || text == "Hi" {
		log.Println("responding with greeting...")
		return GreetingOperation{}
	} else if strings.Contains(text, "?") {
		log.Println("responding with question...")
		return QuestionOperation{}
	} else if text == "op" {
		log.Println("responding with inline keyboard...")
		return Inline{}
	} else {
		return DefaultOperation{}
	}
}

func (DefaultOperation) Execute(Update) error {
	return nil
}
