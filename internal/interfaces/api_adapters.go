package interfaces

import (
	"context"
	"url_shortener_client/internal/config"
	"url_shortener_client/internal/entities"
)

type TelegramSender interface {
	SendMessage(chatID int64, message string) error
}

type UrlShortenerAPI interface {
	CreateShortUrl(ctx context.Context, configuration *config.Config, userUrl string) (*entities.Query, error)
}
