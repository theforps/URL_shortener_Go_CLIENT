package parsers

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"url_shortener_client/internal/config"
	"url_shortener_client/internal/entities"
	"url_shortener_client/pkg"
)

type Parser struct {
	configuration *config.Config
}

func NewParser(configuration *config.Config) *Parser {
	return &Parser{configuration: configuration}
}

func (mp *Parser) ParseMessages() (scenario *entities.Scenario, err error) {

	jsonByte := pkg.GetMessageContent()
	if jsonByte == nil {
		return nil, errors.New("ошибка чтения данных из файла со сценарием сообщений")
	}

	localization := entities.Localization{Scenarios: []*entities.Scenario{}}

	err = json.NewDecoder(bytes.NewReader(jsonByte)).Decode(&localization)
	if err != nil {
		return nil, fmt.Errorf("ошибка конвертации сценария сообщений в объект")
	}

	currentLang := mp.configuration.Language

	for _, item := range localization.Scenarios {
		if item.Language == currentLang {
			scenario = item
		}
	}

	return scenario, nil
}
