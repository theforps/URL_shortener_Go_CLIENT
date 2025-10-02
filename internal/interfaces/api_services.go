package interfaces

import "url_shortener_client/internal/entities"

type FetchShortUrl interface {
	FetchUrl(userUrl string) (*entities.Query, error)
}

type SendShortUrl interface {
	SendTextMessage(chatID int64, message string) error
}
