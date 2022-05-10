package bot

import (
	"log"
	"strings"
	"time"

	"github.com/MoonSHRD/dao-tg/internal/models"
	storage "github.com/MoonSHRD/dao-tg/internal/store"
	client "github.com/MoonSHRD/dao-tg/pkg/gnosis"
	tg "gopkg.in/telebot.v3"
)

var (
	menu = &tg.ReplyMarkup{ResizeKeyboard: true, OneTimeKeyboard: true}

	networks = map[string]string{
		"mainnet":   "Ethereum Mainnet",
		"xdai":      "Gnosis Chain",
		"arbitrum":  "Arbitrum",
		"avalanche": "Avalanche",
		"aurora":    "Aurora",
		"bsc":       "Binance Smart Chain",
		"ewc":       "Energy Web Chain",
		"optiprism": "Optiprism",
		"polygon":   "Polygon",
		"goerli":    "Goerli",
		"rinkeby":   "Rinkeby",
		"volta":     "Volta",
	}
)

const (
	SubscribeStateChoosingNetwork = "subscribe-network"
	SubscribeStateSpecifyingSafe  = "subscribe-safe"
	SubscribeStateSpecifyingName  = "subscribe-name"
)

func RegisterSubscribePipeline(b *tg.Bot, gnosis *client.Gnosis, store storage.Store) {
	rows := make([]tg.Row, 0, len(networks))
	// Network Keyboard handler
	for network, label := range networks {
		button := menu.Text(label)
		rows = append(rows, menu.Row(button))
		b.Handle(&button, func(ctx tg.Context) error {
			// TODO: FSM
			ctx.Set(contextStateKey, SubscribeStateSpecifyingSafe+":"+network)
			return ctx.Reply("Please specify safe address\n(starts with `rin:` or `0x`)")
		})
	}

	b.Handle("/subscribe", func(ctx tg.Context) error {
		menu.Reply(rows...)
		return ctx.Reply("What transaction service you would like to use?", menu.ReplyKeyboard)
	})

	b.Handle(tg.OnText, func(ctx tg.Context) error {
		args := ctx.Args()
		if len(args) < 1 {
			return ctx.Reply("Send proper address, or type `dismiss`")
		}

		state, ok := ctx.Get(contextStateKey).(string)
		if !ok || !strings.HasPrefix(state, "subscribe-safe") {
			log.Println("Failed to get state")
			return ctx.Reply("Internal error happened")
		}

		splittedState := strings.SplitN(state, ":", 2)
		if splittedState == nil {
			log.Println("Failed to split " + state)
			return ctx.Reply("Internal error happened")
		}

		network := splittedState[1]
		safe, err := gnosis.WithNetwork(network).Safe(args[0])
		if err != nil {
			if err == client.ErrResponseNotOK {
				return ctx.Reply("Safe does not exist, send proper address, or type `dismiss`")
			}
			return ctx.Reply("Internal error happened")
		}

		key := []byte(ctx.Recipient().Recipient())
		recipient := new(models.Recipient)
		exists, err := store.Has(key)
		if err != nil {
			return ctx.Reply("Internal error happened")
		}

		if exists {
			data, err := store.Get(key)
			if err != nil {
				log.Println("Failed to get recipient from store - " + err.Error())
				return ctx.Reply("Internal error happened")
			}

			*recipient, err = storage.FromGob[models.Recipient](data)
			if err != nil {
				log.Println("Failed to get from gob - " + err.Error())
				return ctx.Reply("Internal error happened")
			}
		}

		recipient.Subscriptions = append(recipient.Subscriptions, models.Subscription{
			Network:     network,
			SafeAddress: safe.Address,
			LastUpdated: time.Now().Unix(),
		})

		data, err := storage.ToGob(recipient)
		if err != nil {
			log.Println("Failed to convert into gob - " + err.Error())
			return ctx.Reply("Internal error happened")
		}

		if err := store.Put(key, data); err != nil {
			log.Println("Failed to save recipient - " + err.Error())
			return ctx.Reply("Internal error happened")
		}

		return ctx.Reply("Safe exists, now specify name for that subscription")
	}, func(next tg.HandlerFunc) tg.HandlerFunc {
		return func(ctx tg.Context) error {
			state, ok := ctx.Get(contextStateKey).(string)
			if !ok || !strings.HasPrefix(state, "subscribe-safe") {
				return nil
			}

			return next(ctx)
		}
	})
}

func Subscribe(net string, safeAddress string, gnosis *client.Gnosis, store storage.Store) error {
	return nil
}
