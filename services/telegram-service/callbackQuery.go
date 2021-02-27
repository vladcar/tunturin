package telegram_service

func (c CallbackQueryOperation) Execute(update Update) error {
	query := update.CallbackQuery
	message := query.Message
	var text string
	if query.Data == "G" {
		text = "Thank you 😻"
	} else if query.Data == "B" {
		text = "😿"
	}
	err := AnswerCallbackQuery(CallbackQueryAnswer{
		query.Id,
		text,
		true,
	})
	if err != nil {
		return err
	}
	return EditMessageText(EditMessage{
		message.Chat.Id,
		message.MessageId,
		"Response accepted 😻",
	})
}
