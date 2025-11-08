# Go URL shortener Bot
This project is a Telegram bot written in Go that shortens URLs using an external API.
A user sends a long link, and the bot responds with a short one.

## Application Overview
- Accepts any valid URL from users
- Shortens links via an external API 
- Validates input URLs
- Simple `.env` configuration
- Lightweight and fast

## How to run

### Prerequisites
- [Go 1.21+](https://go.dev/)
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api)
- REST API for link shortening
- [godotenv](https://github.com/joho/godotenv) for environment variables

### Quick start
1. **Clone repository**
```bash
git clone https://github.com/theforps/URL_shortener_Go_CLIENT.git
cd URL_shortener_Go_CLIENT 
```
2. **Set up environment variables:**
```bash
cp .env.example .env
```
3. **Install dependencies:**
```bash
go mod download
go mod verify 
```
4. **Run the application:**
```bash
# Development mode with hot reload
go run cmd/main.go

# Or build and run
go build -o bot cmd/main.go
./bot
```

### Configuration Options
Create a `.env` file with the following variables:
```bash
# bot API key
TELEGRAM_BOT_API="1234568:QWERTY"

# link shortening service
URL_SHORTENER_SERVER="https://example.com/create"

# language of the bot's messages. set "ru" or "en"
LANGUAGE="en"
```