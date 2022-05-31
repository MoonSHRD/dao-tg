package service

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/BenJetson/humantime"
	"github.com/MoonSHRD/dao-tg/internal/models"
	"github.com/MoonSHRD/dao-tg/pkg/escape"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
	"go.uber.org/zap"
	tg "gopkg.in/telebot.v3"
)

// MultisigTransactionsNotifications ...
func (service *Service) MultisigTransactionsNotifications(subscription models.Subscription) ([]Notification, error) {
	transactions, err := service.LatestMultisigTransactions(subscription)
	if err != nil {
		return nil, err
	}

	notifications, err := service.FormatMultisigTransactions(subscription, transactions)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

// LatestMultisigTransactions ...
func (service *Service) LatestMultisigTransactions(subscription models.Subscription) ([]gnosis.MultisigTransaction, error) {
	network := gnosis.Network(subscription.Network)
	client := service.client.WithNetwork(network)

	ordering := "modified"
	modified := subscription.LastUpdated.Format(time.RFC3339)
	result, err := client.MultisigTransactions(subscription.SafeAddress, gnosis.MultisigOptions{
		Ordering:    &ordering,
		ModifiedGte: &modified,
	})
	if err != nil {
		service.logger.Fatal("failed to get result", zap.Error(err))
		return nil, err
	}

	return result.Results, nil
}

// FormatMultisigTransactions ...
func (service *Service) FormatMultisigTransactions(subscription models.Subscription, transactions []gnosis.MultisigTransaction) ([]Notification, error) {
	// Safe context
	network := gnosis.Network(subscription.Network)
	client := service.client.WithNetwork(network)
	safe, err := client.Safe(subscription.SafeAddress)
	if err != nil {
		service.logger.Fatal("failed to get safe", zap.Error(err))
		return nil, err
	}

	notifications := make([]Notification, 0, len(transactions))
	for _, multisig := range transactions {
		// Building notification
		var b strings.Builder

		isOutcoming := multisig.Safe != multisig.To
		isDelegateCall := multisig.Operation == 1

		// Checking for overlapping transactions with same nonce
		// If there is, first executed with deny others
		ordering := "modified"
		modified := multisig.Modified.Format(time.RFC3339)
		confictMultisig, err := client.MultisigTransactions(subscription.SafeAddress, gnosis.MultisigOptions{
			Ordering:   &ordering,
			ModifiedLt: &modified,
			Nonce:      &multisig.Nonce,
		})

		// Is it new multisignature transaction?
		// When add "proposed by" to header
		isNew := len(multisig.Confirmations) == 1
		isRejection := !isOutcoming && len(confictMultisig.Results) > 0
		if isNew {
			proposedBy := multisig.Confirmations[0].Owner
			link := escape.Markdown.Replace(gnosis.ExplorerLink(network, proposedBy))
			if isRejection {
				b.WriteString(escape.Markdown.Replace("On-chain rejection proposed by:\n"))
			} else {
				b.WriteString("Proposed by:\n")
			}
			b.WriteString(fmt.Sprintf("[%s](%s)\n", proposedBy, link))
			b.WriteByte('\n')
		}

		service.logger.Info("rej", zap.Int("results", len(confictMultisig.Results)))
		service.logger.Info("tm", zap.Time("s", multisig.SubmissionDate), zap.Time("m", multisig.Modified))
		service.logger.Info("is", zap.Bool("new", isNew), zap.Bool("out", isOutcoming))

		// Notification headder, with info about from and to
		if isDelegateCall {
			b.WriteString("Delegate call in *" + subscription.Label + "*\n")
		} else if isOutcoming {
			b.WriteString(fmt.Sprintf("From: *%s*\n", subscription.Label))

			var verb string
			if multisig.IsExecuted {
				verb = "Sent"
			} else {
				verb = "Send"
			}

			// Formatting ethereum value
			value := new(big.Int)
			value, ok := value.SetString(multisig.Value, 10)
			if !ok {
				service.logger.Error("failed to parse multisig value")
				continue
			}
			eth := toDecimal(value, defaultEtherDecimals)

			link := escape.Markdown.Replace(gnosis.ExplorerLink(network, multisig.To))
			b.WriteString(fmt.Sprintf("%s *%s ETH* to:\n", verb, escape.Markdown.Replace(eth.String())))
			b.WriteString(fmt.Sprintf("[%s](%s)\n", multisig.To, link))
		} else {
			b.WriteString("Multisignature in *" + subscription.Label + "*\n")
		}

		// Checking status, is it executed, is all confirmations collected?
		remain := safe.Threshold - int64(len(multisig.Confirmations))
		thresholdReached := remain == 0
		if multisig.ConfirmationsRequired == nil {
			b.WriteByte('\n')
			if thresholdReached {
				b.WriteString("*Execution needed*\n")
			} else {
				b.WriteString(fmt.Sprintf("*Confirmations Required* — %d more\n", remain))
			}
		} else {
			b.WriteByte('\n')
			b.WriteString("*Executed by:*\n")
			link := escape.Markdown.Replace(gnosis.ExplorerLink(network, *multisig.Executor))
			b.WriteString(fmt.Sprintf("[%s](%s)\n", *multisig.Executor, link))
		}

		if !isNew {
			b.WriteByte('\n')
			b.WriteString("Confirmations:\n")
			for _, conf := range multisig.Confirmations {
				link := escape.Markdown.Replace(gnosis.ExplorerLink(network, conf.Owner))
				b.WriteString(fmt.Sprintf("[%s](%s)\n", conf.Owner, link))
				b.WriteString("\t\t— Signed " + humantime.Since(conf.SubmissionDate) + "\n")
			}
		}

		if err == nil && len(confictMultisig.Results) > 0 {
			b.WriteByte('\n')

			b.WriteString("Conficts:\n")

			for _, conflict := range confictMultisig.Results {
				isOutcoming := conflict.Safe != conflict.To
				isDelegateCall := conflict.Operation == 1
				remain := safe.Threshold - int64(len(conflict.Confirmations))
				isThresholdReached := remain == 0

				b.WriteString("\t")
				if isDelegateCall {
					b.WriteString("Delegate call:\n")
				} else if isOutcoming {
					// Formatting ethereum value?
					value := new(big.Int)
					value, ok := value.SetString(multisig.Value, 10)
					if !ok {
						service.logger.Error("failed to parse multisig value")
						continue
					}
					eth := toDecimal(value, defaultEtherDecimals)

					b.WriteString(fmt.Sprintf("Send *%s ETH* to:\n", escape.Markdown.Replace(eth.String())))
				} else {
					b.WriteString("Multisignature:\n")
				}

				// Link to overlapping transaction
				link := escape.Markdown.Replace(fmt.Sprintf("https://gnosis-safe.io/app/%[1]s:%[2]s/transactions/multisig_%[2]s_%[3]s",
					network.AsPrefix(), safe.Address, conflict.SafeTxHash))
				b.WriteString(fmt.Sprintf("[%s](%s)\n", conflict.To, link))

				confirmations := fmt.Sprintf("%d/%d", len(conflict.Confirmations), safe.Threshold)
				if isThresholdReached {
					b.WriteString("\t\t— " + confirmations + " Ready for execution " + humantime.Since(conflict.Modified) + "\n")
				} else {
					b.WriteString("\t\t— " + confirmations + " Last confirmation " + humantime.Since(conflict.Modified) + "\n")
				}
			}
		} else {
			service.logger.Error("failed to get overlapping multisignature transactions", zap.Int64("nonce", multisig.Nonce), zap.Error(err))
		}

		if multisig.Data != nil {
			b.WriteByte('\n')
			b.WriteString("There is ***Actions*** performed in this multisig, details:")
		}

		// Adding url button to transaction for fast access to confirmation
		menu := tg.ReplyMarkup{}
		btn := menu.URL("Gnosis",
			fmt.Sprintf("https://gnosis-safe.io/app/%[1]s:%[2]s/transactions/multisig_%[2]s_%[3]s",
				network.AsPrefix(), safe.Address, multisig.SafeTxHash))
		menu.Inline(
			menu.Row(btn),
		)

		notifications = append(notifications, Notification{
			Timestamp: multisig.Modified,
			Content:   b.String(),
			Opts:      []interface{}{&menu, tg.ModeMarkdownV2, tg.NoPreview},
		})
	}

	return notifications, nil
}
