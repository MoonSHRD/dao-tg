package bot

import (
	"context"
	"encoding/json"
	"time"

	"github.com/MoonSHRD/dao-tg/internal/config"
	"github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis/client"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis/client/multisig_transactions"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis/client/safes"
	"go.uber.org/fx"
	tg "gopkg.in/telebot.v3"
)

// RegisterHandler ...
func RegisterHandler(b *tg.Bot, gnosis *client.Gnosis, store store.Store) {
	b.Handle(tg.OnText, func(c tg.Context) error {
		return c.Send("Hello world!")
	})

	b.Handle("/transaction", func(ctx tg.Context) error {
		args := ctx.Args()
		if len(args) < 1 {
			return ctx.Reply("unsuffishent agrs")
		}

		params := multisig_transactions.NewMultisigTransactionsReadParams().WithSafeTxHash(args[0])
		t, err := gnosis.MultisigTransactions.MultisigTransactionsRead(params, nil)
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		nice, err := json.MarshalIndent(t, "", "\t")
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		return ctx.Reply(string(nice))
	})

	b.Handle("/history", func(ctx tg.Context) error {
		args := ctx.Args()
		if len(args) < 1 {
			return ctx.Reply("unsuffishent agrs")
		}

		executed := "true"
		params := safes.NewSafesMultisigTransactionsListParams().WithDefaults().WithAddress(args[0]).WithExecuted(&executed)
		t, err := gnosis.Safes.SafesMultisigTransactionsList(params, nil)
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		nice, err := json.MarshalIndent(t.GetPayload().Results, "", "\t")
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		return ctx.Reply(string(nice))
	})

	b.Handle("/queue", func(ctx tg.Context) error {
		args := ctx.Args()
		if len(args) < 1 {
			return ctx.Reply("unsuffishent agrs")
		}

		executed := "false"
		params := safes.NewSafesMultisigTransactionsListParams().WithDefaults().WithAddress(args[0]).WithExecuted(&executed)
		t, err := gnosis.Safes.SafesMultisigTransactionsList(params, nil)
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		nice, err := json.MarshalIndent(t.GetPayload().Results, "", "\t")
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		return ctx.Reply(string(nice))
	})
}

// New ...
func New(lifecycle fx.Lifecycle, cfg config.Telegram) (*tg.Bot, error) {
	b, err := tg.NewBot(tg.Settings{
		Token:  cfg.BotToken,
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go b.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			b.Stop()
			return nil
		},
	})

	return b, nil
}

// Module ...
var Module = fx.Module("bot",
	fx.Provide(New),
	fx.Invoke(RegisterHandler),
)
