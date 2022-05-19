package service

import (
	"context"
	"time"

	storage "github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
	"go.uber.org/fx"
	"go.uber.org/zap"
	tg "gopkg.in/telebot.v3"
)

// Module ...
var Module = fx.Module("service", fx.Invoke(New))

// Service ...
type Service struct {
	ctx    context.Context
	bot    *tg.Bot
	client *gnosis.Gnosis
	store  storage.Store
	logger *zap.Logger
}

// New ...
func New(lifecycle fx.Lifecycle, bot *tg.Bot, client *gnosis.Gnosis, store storage.Store, logger *zap.Logger) (*Service, error) {
	bg, cancel := context.WithCancel(context.Background())

	service := &Service{
		ctx:    bg,
		bot:    bot,
		client: client,
		store:  store,
		logger: logger,
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go service.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				cancel()
				return nil
			},
		},
	)

	return service, nil
}

// Run ...
func (service *Service) Run() {
	for {
		if err := service.ProcessEvents(); err != nil {
			service.logger.Error("process error", zap.Error(err))
		}

		select {
		case <-service.ctx.Done():
			service.logger.Info("background context done")
			return
		case <-time.After(10 * time.Second):
			continue
		}
	}
}
