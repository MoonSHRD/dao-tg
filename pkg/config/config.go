package config

import (
	"log"
	"os"
	"strconv"

	"go.uber.org/fx"
)

// Config ...
type Config struct {
	fx.Out

	Telegram Telegram
}

// Telegram ...
type Telegram struct {
	AppID    int
	AppHash  string
	BotToken string
}

// New Config struct filled from env
func New() Config {
	return Config{
		Telegram: Telegram{
			AppID:    StringToInt(os.Getenv("TELEGRAM_APP_ID")),
			AppHash:  os.Getenv("TELEGRAM_APP_HASH"),
			BotToken: os.Getenv("TELEGRAM_TOKEN"),
		},
	}
}

// StringToInt ...
func StringToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
