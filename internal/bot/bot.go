package bot

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/MoonSHRD/dao-tg/internal/config"
	"github.com/MoonSHRD/dao-tg/internal/store"
	client "github.com/MoonSHRD/dao-tg/pkg/gnosis"
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

		result, err := gnosis.MultisigTransaction(args[0])
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		nice, err := json.MarshalIndent(result, "", "\t")
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

		safe := args[0]

		executed := true
		queued := false
		result, err := gnosis.Transactions(safe, client.TransactionOptions{
			Executed: &executed,
			Queued:   &queued,
		})
		if err != nil {
			return ctx.Reply("failed" + err.Error())
		}

		var out string
		for _, r := range result.Results {
			switch r.TxType {
			case client.TxTypeEtheriumTransaction:
				out += "tx: " + r.TxHash[:10] + "\n\n"
			case client.TxTypeMultisigTransaction:
				out += "safe_tx: " + r.SafeTxHash[:10] + "\n"
				out += "confirmations needed: " + strconv.FormatInt(*r.ConfirmationsRequired, 10) + "\n"
				out += "confirmations: " + strconv.Itoa(len(r.Confirmations)) + "\n\n"
			}
		}

		return ctx.Reply(out)
	})

	b.Handle("/queue", func(ctx tg.Context) error {
		args := ctx.Args()
		if len(args) < 1 {
			return ctx.Reply("unsuffishent agrs")
		}

		safe := args[0]

		safeInfo, err := gnosis.Safe(safe)
		if err != nil {
			return ctx.Reply("failed " + err.Error())
		}

		executed := false
		result, err := gnosis.MultisigTransactions(safe, client.MultisigOptions{
			Executed: &executed,
		})
		if err != nil {
			return ctx.Reply("failed " + err.Error())
		}

		var out string
		for _, r := range result.Results {
			out += "safe_tx: " + r.SafeTxHash[:10] + "\n"
			if safeInfo.Threshold > int64(len(r.Confirmations)) {
				out += "confirmations needed: " + strconv.FormatInt(safeInfo.Threshold, 10) + "\n"
				out += "confirmations: " + strconv.Itoa(len(r.Confirmations)) + "\n\n"
			} else {
				out += "needs execution" + "\n"
			}
		}

		return ctx.Reply(out)
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
