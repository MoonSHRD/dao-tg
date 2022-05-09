// Binary bot-echo implements basic example for bot.
package main

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/MoonSHRD/dao-tg/internal/bot"
	"github.com/MoonSHRD/dao-tg/internal/config"
	"github.com/MoonSHRD/dao-tg/internal/logger"
	"github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
)

func main() {
	var app *fx.App

	app = fx.New(
		config.Module,
		logger.Module,
		// TODO: Constructor for gnosis config
		fx.Supply(gnosis.Config{
			Base:   "https://safe-transaction.rinkeby.gnosis.io/api/v1",
			Client: http.DefaultClient,
		}),
		gnosis.Module,
		store.Module,
		bot.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)

	app.Run()
}
