package config

import (
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/fx"
)

// Config ...
type Config struct {
	fx.Out

	Telegram Telegram
	Ethereum Ethereum
	Gnosis   Gnosis
}

// Telegram ...
type Telegram struct {
	AppID    int
	AppHash  string
	BotToken string
}

// Ethereum ...
type Ethereum struct {
	Gateway string
}

// Gnosis ...
type Gnosis struct {
	ContractAddress common.Address
}

// New Config struct filled from env
func New() Config {
	return Config{
		Telegram: Telegram{
			AppID:    StringToInt(os.Getenv("TELEGRAM_APP_ID")),
			AppHash:  os.Getenv("TELEGRAM_APP_HASH"),
			BotToken: os.Getenv("TELEGRAM_TOKEN"),
		},
		Ethereum: Ethereum{
			Gateway: os.Getenv("ETHEREUM_GATEWAY"),
		},
		Gnosis: Gnosis{
			ContractAddress: common.HexToAddress(os.Getenv("GNOSIS_CONTRACT")),
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
