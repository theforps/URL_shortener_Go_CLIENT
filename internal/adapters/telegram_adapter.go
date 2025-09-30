package adapters

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAdapter struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramAdapter(botToken string) (*TelegramAdapter, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &TelegramAdapter{bot: bot}, nil
}

func (ta *TelegramAdapter) SendMessage(ctx context.Context, chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)

	_, err := ta.bot.Send(msg)

	return err
}
