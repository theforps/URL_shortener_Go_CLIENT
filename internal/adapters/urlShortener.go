package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"url_shortener_client/internal/config"
	"url_shortener_client/internal/entities"
)

type UrlShortenerAPI struct {
	client *http.Client
}

func NewUrlShortener() *UrlShortenerAPI {
	return &UrlShortenerAPI{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (us *UrlShortenerAPI) CreateShortUrl(ctx context.Context, userUrl string, configuration config.Config) (redirectUrl *entities.URLdto, err error) {

	urlDtoRequest := entities.URLdto{
		URL: userUrl,
	}

	jsonBody, err := json.Marshal(&urlDtoRequest)
	if err != nil {
		return nil, errors.New("ошибка создания JSON тела запроса")
	}

	request, err := http.NewRequestWithContext(ctx, configuration.ShortenerURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, errors.New("ошибка создания POST запроса")
	}

	response, err := us.client.Do(request)
	if err != nil {
		return nil, errors.New("ошибка запроса к API")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("API вернул отрицательный статус")
	}

	redirectUrl = &entities.URLdto{}
	if err = json.NewDecoder(response.Body).Decode(redirectUrl); err != nil {
		return nil, errors.New("ошибка декодирования ответа с API")
	}

	return redirectUrl, nil
}
