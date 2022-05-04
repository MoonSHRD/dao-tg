package bot

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/MoonSHRD/dao-tg/internal/config"
	"github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
	"go.uber.org/fx"
	tg "gopkg.in/telebot.v3"
)

// RegisterHandler ...
func RegisterHandler(b *tg.Bot, store store.Store) {
	b.Handle(tg.OnText, func(c tg.Context) error {
		return c.Send("Hello world!")
	})

	b.Handle("/transaction", func(ctx tg.Context) error {
		args := ctx.Args()
		if len(args) < 3 {
			return ctx.Reply("unsuffishent agrs")
		}

		typ := args[0]
		safeAddr := args[1]
		safeTx := args[2]

		t, _, err := gnosis.GetTransaction(http.DefaultClient, typ, safeAddr, safeTx)
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
		safeAddr := args[0]

		t, _, err := gnosis.GetHistory(http.DefaultClient, safeAddr)
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		nice, err := json.MarshalIndent(t, "", "\t")
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
		safeAddr := args[0]

		t, _, err := gnosis.GetQueue(http.DefaultClient, safeAddr)
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		nice, err := json.MarshalIndent(t, "", "\t")
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
