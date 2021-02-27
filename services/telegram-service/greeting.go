package telegram_service

func (GreetingOperation) Execute(update Update) error {
	message := update.Message
	response := "Hi, my name is Tunturin. How are you today? 😻"
	return SendMessage(TelegramMessage{
		ChatId:           message.Chat.Id,
		Text:             response,
		ReplyToMessageId: message.MessageId,
	})
}
