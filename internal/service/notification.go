package service

import (
	"time"

	"github.com/MoonSHRD/dao-tg/internal/models"
	"go.uber.org/zap"
	tg "gopkg.in/telebot.v3"
)

// Notification - formatted event ready for sending
type Notification struct {
	Timestamp time.Time
	Content   string
	Opts      []interface{}
}

// SendNotifications and return latest sended notification timestamp
func (service *Service) SendNotifications(recipient tg.Recipient, notifications ...Notification) (time.Time, error) {
	latestTimestamp := motherBirth

	for _, message := range notifications {
		_, err := service.bot.Send(recipient, message.Content, message.Opts...)
		if err != nil {
			service.logger.Warn("failed to send message", zap.Error(err))
			return latestTimestamp, err
		}

		if message.Timestamp.After(latestTimestamp) {
			latestTimestamp = message.Timestamp
		}
	}

	return latestTimestamp, nil
}

// CollectNotifications ...
func CollectNotifications(subscription models.Subscription, sources ...func(models.Subscription) ([]Notification, error)) ([]Notification, error) {
	notifications := make([]Notification, 0)

	for _, source := range sources {
		some, err := source(subscription)
		if err != nil {
			return notifications, err
		}
		notifications = append(notifications, some...)
	}

	return notifications, nil
}
