package store

import (
	"time"

	"github.com/akrylysov/pogreb"
	"go.uber.org/fx"
)

// Store ...
type Store struct {
	*pogreb.DB
}

// FilterItems returns a new FilterItemIterator.
func (s Store) FilterItems(prefix []byte) *FilterItemIterator {
	return &FilterItemIterator{
		prefix: prefix,
		iter:   s.Items(),
	}
}

// New ...
func New() (Store, error) {
	db, err := pogreb.Open("./pogreb", &pogreb.Options{
		BackgroundSyncInterval:       -1,
		BackgroundCompactionInterval: time.Hour,
	})
	if err != nil {
		return Store{}, err
	}
	return Store{db}, nil
}

// Module ...
var Module = fx.Module("store", fx.Provide(New))
