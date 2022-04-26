// Binary bot-echo implements basic example for bot.
package main

import (
	"context"
	"errors"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/MoonSHRD/dao-tg/config"
	"github.com/MoonSHRD/dao-tg/logger"
	"github.com/gotd/contrib/bg"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/tg"
)

// Bot - pretty fun struct
type Bot struct {
	*telegram.Client
	*tg.UpdateDispatcher
	*message.Sender
}

// NewBot ...
func NewBot(lifecycle fx.Lifecycle, config config.Telegram, logger *zap.Logger) (Bot, error) {
	dispatcher := tg.NewUpdateDispatcher()

	client := telegram.NewClient(config.AppID, config.AppHash, telegram.Options{
		Logger:        logger,
		UpdateHandler: &dispatcher,
	})

	var stopFunc *bg.StopFunc
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Run bot in bg
			stop, err := bg.Connect(client)
			if err != nil {
				return err
			}
			stopFunc = &stop

			// Auth as Bot
			_, err = client.Auth().Bot(ctx, config.BotToken)
			if err != nil {
				return err
			}

			// Checking if all ok
			status, err := client.Auth().Status(ctx)
			if err != nil {
				return err
			}

			if !status.Authorized {
				return errors.New("failed to auth telegram")
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return (*stopFunc)()
		},
	})

	return Bot{
		Client:           client,
		UpdateDispatcher: &dispatcher,
		Sender:           message.NewSender(client.API()),
	}, nil
}

func main() {
	var app *fx.App

	app = fx.New(
		fx.Provide(
			logger.New,
			config.New,
			NewBot,
		),
		fx.Invoke(
			RegisterHandler,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)

	app.Run()
}
