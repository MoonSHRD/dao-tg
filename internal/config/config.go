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
	Gnosis   Gnosis
}

// Telegram ...
type Telegram struct {
	BotToken string
}

// Gnosis ...
type Gnosis struct {
	Base string
}

// New Config struct filled from env
func New() Config {
	return Config{
		Telegram: Telegram{
			BotToken: os.Getenv("BOT_TOKEN"),
		},
		Gnosis: Gnosis{
			Base: os.Getenv("GNOSIS_BASE"),
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
