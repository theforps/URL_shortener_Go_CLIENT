package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"url_shortener_client/internal/config"

	"url_shortener_client/internal/entities"
)

type UrlShortenerAPI struct {
	client *http.Client
}

// NewUrlShortener creates API object that provides access to the creation of short links
func NewUrlShortener() *UrlShortenerAPI {
	return &UrlShortenerAPI{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// CreateShortUrl creates short URL from user's link
func (us *UrlShortenerAPI) CreateShortUrl(ctx context.Context, configuration *config.Config, userUrl string) (redirectUrl *entities.Query, err error) {

	urlDtoQuery := entities.Query{
		URL: userUrl,
	}

	jsonBody, err := json.Marshal(&urlDtoQuery)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания JSON тела запроса: %v", err)
	}

	request, err := http.NewRequestWithContext(ctx, "POST", configuration.ShortenerURL, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания POST запроса: %v", err)
	}

	response, err := us.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к API: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API вернул отрицательный статус: %v", err)
	}

	if err = json.NewDecoder(response.Body).Decode(&urlDtoQuery); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа с API: %v", err)
	}

	return &urlDtoQuery, nil
}
