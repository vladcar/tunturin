package telegram_service

func (Inline) Execute(update Update) error {
	message := update.Message
	row := []InlineKeyboardButton{
		{Text: "😻", CallbackData: "G"},
		{Text: "😿", CallbackData: "B"},
	}
	inlineKeyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{row},
	}

	return SendMessage(TelegramMessage{
		ChatId:      message.Chat.Id,
		Text:        "How are you feeling today?",
		ReplyMarkup: &inlineKeyboard,
	})
}
