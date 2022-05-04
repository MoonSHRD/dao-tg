package service

import (
	"bytes"
	"encoding/gob"

	"github.com/MoonSHRD/dao-tg/internal/models"
	"github.com/MoonSHRD/dao-tg/internal/store"
	"github.com/akrylysov/pogreb"
	"go.uber.org/fx"
	"go.uber.org/zap"
	tg "gopkg.in/telebot.v3"
)

// New ...
func New(lifecycle fx.Lifecycle, bot *tg.Bot, store store.Store, logger *zap.Logger) error {

	it := store.FilterItems([]byte("sub:"))
	for {
		key, val, err := it.Next()
		if err == pogreb.ErrIterationDone {
			break
		}
		if err != nil {
			logger.Fatal("failed to iterate over store", zap.Error(err))
			return err
		}

		buf := bytes.NewBuffer(val)
		dec := gob.NewDecoder(buf)
		sub := new(models.SubscriberFeed)
		if err := dec.Decode(sub); err != nil {
			logger.Fatal("failed to decode record from store", zap.Error(err))
			return err
		}

		logger.Info("sub record", zap.String("key", string(key)), zap.Any("val", sub))
	}

	return nil
}

// Module ...
var Module = fx.Module("service", fx.Provide(New))
