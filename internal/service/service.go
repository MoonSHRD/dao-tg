package service

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/BenJetson/humantime"
	"github.com/MoonSHRD/dao-tg/internal/models"
	storage "github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/MoonSHRD/dao-tg/pkg/escape"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
	"github.com/akrylysov/pogreb"
	"go.uber.org/fx"
	"go.uber.org/zap"
	tg "gopkg.in/telebot.v3"
)

// Module ...
var Module = fx.Module("service", fx.Invoke(Run))

// Run ...
func Run(lifecycle fx.Lifecycle, bot *tg.Bot, client *gnosis.Gnosis, store storage.Store, logger *zap.Logger) error {
	bg, cancel := context.WithCancel(context.Background())
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func(bg context.Context) {
					for {
						if err := Process(bot, client, store, logger); err != nil {
							logger.Error("process error", zap.Error(err))
						}

						logger.Info("service iteration")

						select {
						case <-bg.Done():
							logger.Info("background context done")
							return
						case <-time.After(10 * time.Second):
							continue
						}
					}
				}(bg)

				return nil
			},
			OnStop: func(ctx context.Context) error {
				cancel()
				return nil
			},
		},
	)

	return nil
}

func Process(bot *tg.Bot, client *gnosis.Gnosis, store storage.Store, logger *zap.Logger) error {
	it := store.Items()
	for {
		key, val, err := it.Next()
		if err == pogreb.ErrIterationDone {
			break
		}
		if err != nil {
			logger.Fatal("failed to iterate over store", zap.Error(err))
			return err
		}

		recipient, err := storage.FromGob[models.Recipient](val)
		if err != nil {
			logger.Fatal("failed to get from gob", zap.Error(err))
			continue
		}

		for i, sub := range recipient.Subscriptions {
			network := gnosis.Network(sub.Network)
			c := client.WithNetwork(network)
			safe, err := c.Safe(sub.SafeAddress)
			if err != nil {
				logger.Fatal("failed to get safe", zap.Error(err))
				continue
			}

			ordering := "modified"
			modified := sub.LastUpdated.Format(time.RFC3339)
			result, err := c.MultisigTransactions(sub.SafeAddress, gnosis.MultisigOptions{
				Ordering:   &ordering,
				ModifiedGt: &modified,
			})
			if err != nil {
				logger.Fatal("failed to get result", zap.Error(err))
				continue
			}

			for _, multisig := range result.Results {
				log.Println(multisig.Modified.Format(time.RFC3339))
				recipient.Subscriptions[i].LastUpdated = multisig.Modified.Add(1 * time.Second)

				// Building notification
				var b strings.Builder

				if multisig.Safe == multisig.To {
					b.WriteString("Multisignature in *" + sub.Label + "*\n")
				} else {
					b.WriteString(fmt.Sprintf("From: *%s*\n", sub.Label))

					var verb string
					if multisig.IsExecuted {
						verb = "Sent"
					} else {
						verb = "Send"
					}

					value := new(big.Int)
					value, ok := value.SetString(multisig.Value, 10)
					if !ok {
						logger.Fatal("failed to parse multisig value", zap.Error(err))
						continue
					}
					eth := toDecimal(value, 18)

					link := escape.Markdown.Replace(gnosis.ExplorerLink(network, multisig.To))
					b.WriteString(fmt.Sprintf("%s *%s ETH* to:\n", verb, escape.Markdown.Replace(eth.String())))
					b.WriteString(fmt.Sprintf("[%s](%s)\n", multisig.To, link))
				}

				b.WriteByte('\n')

				remain := safe.Threshold - int64(len(multisig.Confirmations))
				thresholdReached := remain == 0
				if multisig.ConfirmationsRequired == nil {
					if thresholdReached {
						b.WriteString("*Execution needed*\n")
					} else {
						b.WriteString(fmt.Sprintf("*Confirmations Required* — %d more\n", remain))
					}
				} else {
					b.WriteString("*Executed by:*\n")
					link := escape.Markdown.Replace(gnosis.ExplorerLink(network, *multisig.Executor))
					b.WriteString(fmt.Sprintf("[%s](%s)\n", *multisig.Executor, link))
				}

				b.WriteByte('\n')

				b.WriteString("Confirmations:\n")
				for _, conf := range multisig.Confirmations {
					link := escape.Markdown.Replace(gnosis.ExplorerLink(network, conf.Owner))
					b.WriteString(fmt.Sprintf("[%s](%s)\n", conf.Owner, link))
					b.WriteString("\t\t— Signed " + humantime.Since(conf.SubmissionDate) + "\n")
				}

				menu := tg.ReplyMarkup{}
				btn := menu.URL("Gnosis",
					fmt.Sprintf("https://gnosis-safe.io/app/%[1]s:%[2]s/transactions/multisig_%[2]s_%[3]s",
						network.AsPrefix(), safe.Address, multisig.SafeTxHash))
				menu.Inline(
					menu.Row(btn),
				)

				_, err := bot.Send(pocketRecipient{string(key)}, b.String(), &menu, tg.ModeMarkdownV2, tg.NoPreview)
				if err != nil {
					logger.Warn("failed to send message", zap.Error(err))
					return err
				}

				// Saving
				// TODO: put in a separate function
				data, err := storage.ToGob(recipient)
				if err != nil {
					logger.Info("Failed to convert into gob", zap.Error(err))
					return err
				}

				if err = store.Put(key, data); err != nil {
					logger.Info("Failed to put into store", zap.Error(err))
					return err
				}
			}
		}
	}

	return nil
}

type pocketRecipient struct {
	recipient string
}

func (r pocketRecipient) Recipient() string {
	return r.recipient
}
