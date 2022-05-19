package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/MoonSHRD/dao-tg/internal/models"
	"github.com/MoonSHRD/dao-tg/pkg/escape"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
	"go.uber.org/zap"
	tg "gopkg.in/telebot.v3"
)

// IncomingTransfersNotifications ...
func (service *Service) IncomingTransfersNotifications(subscription models.Subscription) ([]Notification, error) {
	transfers, err := service.LatestIncomingTransfers(subscription)
	if err != nil {
		return nil, err
	}

	// Reversing, so we output them in ascending order
	for i, j := 0, len(transfers)-1; i < j; i, j = i+1, j-1 {
		transfers[i], transfers[j] = transfers[j], transfers[i]
	}

	notifications, err := service.FormatIncomingTransfers(subscription, transfers)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

// LatestIncomingTransfers ...
func (service *Service) LatestIncomingTransfers(subscription models.Subscription) ([]gnosis.Transfer, error) {
	network := gnosis.Network(subscription.Network)
	client := service.client.WithNetwork(network)

	// Getting latest transfers
	executionDate := subscription.LastUpdated.Format(time.RFC3339)
	result, err := client.IncomingTransfers(subscription.SafeAddress, gnosis.TransfersOptions{
		ExecutionDateGte: &executionDate,
	})
	if err != nil {
		service.logger.Fatal("failed to get result", zap.Error(err))
		return nil, err
	}

	return result.Results, nil
}

// FormatIncomingTransfers ...
func (service *Service) FormatIncomingTransfers(subscription models.Subscription, transfers []gnosis.Transfer) ([]Notification, error) {
	// Safe context
	network := gnosis.Network(subscription.Network)
	client := service.client.WithNetwork(network)
	safe, err := client.Safe(subscription.SafeAddress)
	if err != nil {
		service.logger.Fatal("failed to get safe", zap.Error(err))
		return nil, err
	}

	// Formatting
	notifications := make([]Notification, 0, len(transfers))
	for _, transfer := range transfers {
		var b strings.Builder

		amount := escape.Markdown.Replace(strToDecimal(transfer.Value, defaultEtherDecimals).String())
		b.WriteString(fmt.Sprintf("Received *%s %s* from:\n", amount, "ETH"))
		link := escape.Markdown.Replace(gnosis.ExplorerLink(network, transfer.From))
		b.WriteString(fmt.Sprintf("[%s](%s)", transfer.From, link))

		menu := tg.ReplyMarkup{}
		btn := menu.URL("Gnosis",
			fmt.Sprintf("https://gnosis-safe.io/app/%[1]s:%[2]s/transactions/ethereum_%[2]s_%[3]s",
				network.AsPrefix(), safe.Address, transfer.TransactionHash))
		menu.Inline(
			menu.Row(btn),
		)

		notifications = append(notifications, Notification{
			Timestamp: transfer.ExecutionDate,
			Content:   b.String(),
			Opts:      []interface{}{&menu, tg.ModeMarkdownV2, tg.NoPreview},
		})
	}

	return notifications, nil
}
