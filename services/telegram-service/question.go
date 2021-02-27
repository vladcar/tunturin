package telegram_service

func (QuestionOperation) Execute(update Update) error {
	return SendMessage(TelegramMessage{
		ChatId:           update.Message.Chat.Id,
		Text:             "Sorry, I don't know how to handle questions yet 😿 I will get there soon...",
		ReplyToMessageId: update.Message.MessageId,
	})
}
