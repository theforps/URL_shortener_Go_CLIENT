package config

import "os"

type Config struct {
	ShortenerURL string
	TelegramAPI  string
	Language     string
}

func NewConfig() *Config {
	return &Config{
		ShortenerURL: getEnv("URL_SHORTENER_SERVER"),
		TelegramAPI:  getEnv("TELEGRAM_BOT_API"),
		Language:     getEnv("LANGUAGE"),
	}
}

func getEnv(key string) (value string) {
	value, _ = os.LookupEnv(key)
	return
}
