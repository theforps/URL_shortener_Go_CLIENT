package adapters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAdapter struct {
	bot *tgbotapi.BotAPI
}

// NewTelegramAdapter returns object that works with Telegram
func NewTelegramAdapter(bot *tgbotapi.BotAPI) *TelegramAdapter {
	return &TelegramAdapter{bot: bot}
}

// SendMessage sends messages to telegram's chat
func (ta *TelegramAdapter) SendMessage(chatID int64, message string) (err error) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"

	_, err = ta.bot.Send(msg)

	return err
}
