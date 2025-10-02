package adapters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAdapter struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramAdapter(bot *tgbotapi.BotAPI) *TelegramAdapter {
	return &TelegramAdapter{bot: bot}
}

func (ta *TelegramAdapter) SendMessage(chatID int64, message string) (err error) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"

	_, err = ta.bot.Send(msg)

	return err
}
