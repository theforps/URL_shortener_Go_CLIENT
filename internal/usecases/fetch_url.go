package usecases

import (
	"context"
	"fmt"
	"url_shortener_client/internal/config"
	"url_shortener_client/internal/entities"

	"url_shortener_client/internal/interfaces"
)

type FetchShortUrlService struct {
	urlShort      interfaces.UrlShortenerAPI
	ctx           context.Context
	configuration *config.Config
}

func NewFetchService(ctx context.Context, configuration *config.Config, urlShort interfaces.UrlShortenerAPI) *FetchShortUrlService {
	return &FetchShortUrlService{
		urlShort:      urlShort,
		ctx:           ctx,
		configuration: configuration,
	}
}

func (fsu *FetchShortUrlService) FetchUrl(userUrl string) (redirectUrl *entities.Query, err error) {
	if userUrl == "" {
		return nil, fmt.Errorf("ссылка не может быть пустой")
	}

	redirectUrl, err = fsu.urlShort.CreateShortUrl(fsu.ctx, fsu.configuration, userUrl)
	if err != nil || redirectUrl == nil {
		return nil, fmt.Errorf("не удалось получить сокращенную ссылку : %v", err)
	}

	return redirectUrl, nil
}
