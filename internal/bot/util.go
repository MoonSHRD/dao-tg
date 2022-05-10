package bot

import (
	"log"

	storage "github.com/MoonSHRD/dao-tg/internal/store"
	tg "gopkg.in/telebot.v3"
)

const (
	contextStateKey = "state"
	statePostfix    = ":state"
)

func ManageState(store storage.Store) tg.MiddlewareFunc {
	return func(next tg.HandlerFunc) tg.HandlerFunc {
		return func(ctx tg.Context) error {
			key := []byte(ctx.Recipient().Recipient())

			state, err := getState(store, key)
			if err != nil {
				log.Println(err)
				return ctx.Reply("Internal error happened :(")
			}

			ctx.Set(contextStateKey, state)

			defer func() {
				state, ok := ctx.Get(contextStateKey).(string)
				if !ok {
					return
				}

				if err := setState(store, key, state); err != nil {
					log.Println(err)
				}
			}()

			return next(ctx)
		}
	}
}

// Get current recipient dialog state
func getState(store storage.Store, key []byte, defaultValue ...string) (string, error) {
	var val string
	if len(defaultValue) > 0 {
		val = defaultValue[0]
	}

	stateKey := append(key, statePostfix...)
	exists, err := store.Has(stateKey)
	if err != nil {
		return val, err
	}

	if !exists {
		return val, nil
	}

	data, err := store.Get(stateKey)
	if err != nil {
		return val, err
	}

	return string(data), nil
}

// Set current recipient dialog state
func setState(store storage.Store, key []byte, state string) error {
	stateKey := append(key, statePostfix...)
	return store.Put(stateKey, []byte(state))
}
