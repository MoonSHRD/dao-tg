// Binary bot-echo implements basic example for bot.
package main

import (
	"context"
	"errors"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/MoonSHRD/dao-tg/config"
	"github.com/MoonSHRD/dao-tg/contract"
	"github.com/MoonSHRD/dao-tg/contract/gnosis"
	"github.com/MoonSHRD/dao-tg/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	Gnosis *gnosis.Gnosis
}

// NewBot ...
func NewBot(lifecycle fx.Lifecycle, config config.Telegram, gnosisSafe *gnosis.Gnosis, logger *zap.Logger) (Bot, error) {
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
		Gnosis:           gnosisSafe,
	}, nil
}

func main() {
	var app *fx.App

	app = fx.New(
		fx.Supply(
			config.Ethereum{
				Gateway: "wss://rinkeby.infura.io/ws/v3/4c19e44066c24f4694d596996f373cf2",
			},
			config.Gnosis{
				ContractAddress: common.HexToAddress("0x2De84e4c52B1E7B4DaEa4b64CE6d907cE601Ee0d"),
			},
		),
		// config.Module,
		logger.Module,
		contract.Module,
		// fx.Provide(
		// 	NewBot,
		// ),
		// fx.Invoke(
		// 	RegisterHandler,
		// ),

		fx.Invoke(func(gnosisSafe *gnosis.Gnosis, logger *zap.Logger) error {
			iter, err := gnosisSafe.FilterSafeSetup(&bind.FilterOpts{Start: 10567901}, nil)
			if err != nil {
				return err
			}

			for iter.Next() {
				logger.Info("safe setup event",
					zap.String("Initiator", iter.Event.Initiator.String()),
					zap.Any("Owners", iter.Event.Owners),
					zap.Any("Threshold", iter.Event.Threshold),
					zap.Any("Initializer", iter.Event.Initializer),
				)
			}

			return nil
		}),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)

	app.Run()
}
