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

		if multisig.Safe == multisig.To {
			b.WriteString("Multisignature in *" + subscription.Label + "*\n")
		} else {
			b.WriteString(fmt.Sprintf("From: *%s*\n", subscription.Label))

			var verb string
			if multisig.IsExecuted {
				verb = "Sent"
			} else {
				verb = "Send"
			}

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

		notifications = append(notifications, Notification{
			Timestamp: multisig.Modified,
			Content:   b.String(),
			Opts:      []interface{}{&menu, tg.ModeMarkdownV2, tg.NoPreview},
		})
	}

	return notifications, nil
}
