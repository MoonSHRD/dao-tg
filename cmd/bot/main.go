// Binary bot-echo implements basic example for bot.
package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/MoonSHRD/dao-tg/internal/bot"
	"github.com/MoonSHRD/dao-tg/internal/config"
	"github.com/MoonSHRD/dao-tg/internal/logger"
	"github.com/MoonSHRD/dao-tg/internal/store"
)

func main() {
	var app *fx.App

	app = fx.New(
		config.Module,
		logger.Module,
		store.Module,
		bot.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)

	app.Run()
}