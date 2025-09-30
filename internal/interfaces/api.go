package interfaces

import (
	"context"
	"url_shortener_client/internal/entities"
)

type TelegramSender interface {
	SendMessage(ctx context.Context, chatID int64, message string) error
}

type UrlShortenerAPI interface {
	CreateShortUrl(ctx context.Context, userUrl string) (*entities.URLdto, error)
}
