package service

import (
	"sort"
	"time"

	"github.com/MoonSHRD/dao-tg/internal/models"
	storage "github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/akrylysov/pogreb"
	"go.uber.org/zap"
)

// ProcessEvents ...
func (service *Service) ProcessEvents() error {
	it := service.store.Items()
	for {
		key, val, err := it.Next()
		if err == pogreb.ErrIterationDone {
			break
		}
		if err != nil {
			service.logger.Error("failed to iterate over store", zap.Error(err))
			return err
		}

		recipient, err := storage.FromGob[models.Recipient](val)
		if err != nil {
			service.logger.Error("failed to get from gob", zap.Error(err))
			continue
		}

		target := newRecipient(string(key))
		for i, sub := range recipient.Subscriptions {
			notifications, err := CollectNotifications(sub,
				service.MultisigTransactionsNotifications,
			)

			// Sorting in ascending order
			sort.Slice(notifications, func(i, j int) bool {
				return notifications[i].Timestamp.Before(notifications[j].Timestamp)
			})

			if err != nil {
				service.logger.Error("failed to get notifications", zap.Error(err))
				continue
			}

			latestTimestamp, err := service.SendNotifications(target, notifications...)

			// Update latest sended notification
			if sub.LastUpdated.Before(latestTimestamp) {
				recipient.Subscriptions[i].LastUpdated = latestTimestamp.Add(time.Second)
			}

			if err != nil {
				service.logger.Error("failed to send notifications", zap.Error(err))
			}
		}

		// Saving
		// TODO: put in a separate function
		data, err := storage.ToGob(recipient)
		if err != nil {
			service.logger.Info("Failed to convert into gob", zap.Error(err))
			return err
		}

		if err = service.store.Put(key, data); err != nil {
			service.logger.Info("Failed to put into store", zap.Error(err))
			return err
		}

		service.logger.Info("processed subscriber", zap.String("recipient", target.recipient))
	}

	return nil
}
