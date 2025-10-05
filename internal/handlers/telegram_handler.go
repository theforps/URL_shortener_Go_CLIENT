package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"url_shortener_client/internal/interfaces"
	"url_shortener_client/internal/parsers"
	"url_shortener_client/internal/validators"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgHandler struct {
	botAPI       *tgbotapi.BotAPI
	fetchService interfaces.FetchShortUrl
	sendService  interfaces.SendShortUrl
	context      context.Context
	logger       *slog.Logger
	localization *parsers.Parser
}

func NewTgHandler(
	bot *tgbotapi.BotAPI,
	fs interfaces.FetchShortUrl,
	ss interfaces.SendShortUrl,
	ctx context.Context,
	logger *slog.Logger,
	localization *parsers.Parser) *TgHandler {
	return &TgHandler{
		botAPI:       bot,
		fetchService: fs,
		sendService:  ss,
		context:      ctx,
		logger:       logger,
		localization: localization,
	}
}

func (tr *TgHandler) MessageHandler() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tr.botAPI.GetUpdatesChan(u)

	var msg, errorMsg string

	parsedMessages, err := tr.localization.ParseMessages()
	if err != nil {
		errorMsg = fmt.Sprintf("ошибка парсинга: %v", err)
	}

	for update := range updates {
		if update.Message != nil {
			messageText := update.Message.Text

			if update.Message.Command() != "" && update.Message.Command() == "start" {
				msg = parsedMessages.HelloMes
			} else if strings.HasPrefix(messageText, "/create ") {
				userUrl := strings.Replace(messageText, "/create ", "", 1)

				if validators.IsValidUrl(userUrl) {

					redirectUrl, err := tr.fetchService.FetchUrl(userUrl)

					if err != nil {
						errorMsg = fmt.Sprintf("не удалось вытянуть URL для редиректа: %v", err)
						msg = parsedMessages.WrongMes
					} else {
						msg = fmt.Sprintf(parsedMessages.GoodMes, redirectUrl.URL, strconv.Itoa(redirectUrl.DayLife*24)) //string(rune(redirectUrl.LifeHours)))
					}
				} else {
					errorMsg = "не удалось спарсить пользовательский URL"
				}
			}

			if msg == "" {
				msg = parsedMessages.BadMes
				errorMsg = fmt.Sprintf("не удалось спарсить сообщение и вытянуть URL: %s", messageText)
			}
			err := tr.sendService.SendTextMessage(update.Message.Chat.ID, msg)
			if err != nil {
				errorMsg = fmt.Sprintf("ошибка отправки сообщения: %v", err)
			}

			if errorMsg != "" {
				tr.logger.Error(errorMsg)
			}
		}
	}
}
