package usecases

import (
	"context"
	"fmt"
	"url_shortener_client/internal/entities"
	"url_shortener_client/internal/interfaces"
)

type SendShortUrlService struct {
	telegram interfaces.TelegramSender
}

func NewSendQuoteService(telegram interfaces.TelegramSender) *SendShortUrlService {
	return &SendShortUrlService{telegram: telegram}
}

func (ssu *SendShortUrlService) SendUrl(ctx context.Context, chatID int64, urlDto *entities.URLdto) error {
	if urlDto == nil {
		return fmt.Errorf("ссылка не может быть nil")
	}

	err := ssu.telegram.SendMessage(ctx, chatID, urlDto.URL)
	if err != nil {
		return fmt.Errorf("не удалось отправить сообщение: %v", err)
	}

	return nil
}
