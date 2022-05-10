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
	BotToken string
}

// New Config struct filled from env
func New() Config {
	return Config{
		Telegram: Telegram{
			BotToken: os.Getenv("BOT_TOKEN"),
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

// Module ...
var Module = fx.Module("config", fx.Provide(New))
