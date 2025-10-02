package main

import (
	"context"
	_ "embed"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"url_shortener_client/internal/config"
	"url_shortener_client/internal/parsers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"url_shortener_client/internal/adapters"
	"url_shortener_client/internal/handlers"
	"url_shortener_client/internal/usecases"
)

func main() {

	// logger init
	logger := setupLogger()
	logger.Info("запуск приложения")

	// load .env
	if err := godotenv.Load("../.env"); err != nil {
		logger.Error("ошибка загрузки .env файла", slog.Any("error", err))
		return
	}
	logger.Info("переменные окружения загружены")

	// config init
	configuration := config.NewConfig()

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramAPI)
	if err != nil {
		logger.Error("ошибка подключения к боту", slog.Any("error", err))
		return
	}
	logger.Info("произведенно успешное подключение к API бота", slog.String("bot_username", bot.Self.UserName))

	// graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// initializing dependencies
	parser := parsers.NewParser(configuration)
	shortener := adapters.NewUrlShortener()
	tgAdapter := adapters.NewTelegramAdapter(bot)

	sendService := usecases.NewSendService(tgAdapter)
	fetchService := usecases.NewFetchService(ctx, configuration, shortener)

	logger.Info("бот запущен")

	// telegram handler
	botHandler := handlers.NewTgHandler(bot, fetchService, sendService, ctx, logger, parser)
	botHandler.MessageHandler()
}

func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return logger
}
