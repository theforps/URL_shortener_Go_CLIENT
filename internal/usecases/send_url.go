package usecases

import (
	"errors"
	"fmt"

	"url_shortener_client/internal/interfaces"
)

type SendShortUrlService struct {
	telegram interfaces.TelegramSender
}

func NewSendService(telegram interfaces.TelegramSender) *SendShortUrlService {
	return &SendShortUrlService{telegram: telegram}
}

func (ssu *SendShortUrlService) SendTextMessage(chatID int64, message string) (err error) {
	if message == "" {
		return errors.New("сообщение не может быть пустым")
	}

	err = ssu.telegram.SendMessage(chatID, message)
	if err != nil {
		return fmt.Errorf("не удалось отправить сообщение: %v", err)
	}

	return nil
}
