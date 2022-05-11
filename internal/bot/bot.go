package bot

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BenJetson/humantime"
	"github.com/MoonSHRD/dao-tg/internal/config"
	"github.com/MoonSHRD/dao-tg/internal/models"
	storage "github.com/MoonSHRD/dao-tg/internal/store"
	client "github.com/MoonSHRD/dao-tg/pkg/gnosis"
	"go.uber.org/fx"
	tg "gopkg.in/telebot.v3"
)

// RegisterHandler ...
func RegisterHandler(b *tg.Bot, gnosis *client.Gnosis, store storage.Store) {
	b.Handle("/start", func(ctx tg.Context) error {
		var b strings.Builder
		b.WriteString("Hello, here is commands:\n")
		b.WriteString("\t/subscribe — Subscribe to notifications from Safe\n")
		b.WriteString("\t/subscriptions — List current subscriptions\n")
		b.WriteString("\t/unsubscribe — Unsubscribe from Safe\n")
		return ctx.Reply(b.String())
	})

	b.Handle("/subscribe", func(ctx tg.Context) error {
		args := ctx.Args()

		if len(args) < 2 {
			var b strings.Builder
			b.WriteString("Usage:\n")
			b.WriteString("/subscribe\n")
			b.WriteString("\tmy-dao\n")
			b.WriteString("\trin:0x2De84e4c52...\n")
			return ctx.Reply(b.String())
		}

		// Input stuff
		safe := strings.SplitN(args[1], ":", 2)
		if safe == nil {
			return ctx.Reply("Provide correct gnosis safe address with prefix,\nexample: rin:0x2De84e4c52...")
		}
		label := args[0]
		address := safe[1]
		network, err := client.NetworkFromPrefix(safe[0])
		if err != nil {
			log.Println(err)
			return ctx.Reply("Network prefix is incorrect, provide correct gnosis safe address with prefix,\nexample: rin:0x2De84e4c52...")
		}

		// Checking if safe exists
		safeInfo, err := gnosis.WithNetwork(network).Safe(address)
		if err != nil {
			if err == client.ErrResponseNotOK {
				return ctx.Reply("Safe does not exists, please check the safe address")
			}
			log.Println("Failed to safe - " + err.Error())
			return ctx.Reply("Failed to subscribe to safe, sorry :(")
		}

		// Getting recipient
		// TODO: put in a separate function
		key := []byte(ctx.Recipient().Recipient())
		data, err := store.Get(key)
		if err != nil {
			log.Println("Failed to get from store - " + err.Error())
			return ctx.Reply("Failed to subscribe to safe, sorry :(")
		}

		var recipient models.Recipient
		if data != nil {
			recipient, err = storage.FromGob[models.Recipient](data)
			if err != nil {
				log.Println("Failed to get from gob - " + err.Error())
				return ctx.Reply("Failed to subscribe to safe, sorry :(")
			}
		}

		for _, sub := range recipient.Subscriptions {
			if label == sub.Label {
				return ctx.Reply("You already have subscription with name `"+mdEscaper.Replace(sub.Label)+"`", tg.ModeMarkdownV2)
			}

			if network.String() == sub.Network && address == sub.SafeAddress {
				return ctx.Reply("You already have subscription with same address")
			}
		}

		// Appending subscription
		recipient.Subscriptions = append(recipient.Subscriptions, models.Subscription{
			Label:       label,
			Network:     network.String(),
			SafeAddress: safeInfo.Address,
			LastUpdated: time.Now().Unix(),
		})

		// Saving
		// TODO: put in a separate function
		data, err = storage.ToGob(recipient)
		if err != nil {
			log.Println("Failed to convert into gob - " + err.Error())
			return ctx.Reply("Failed to subscribe to safe, sorry :(")
		}

		if err = store.Put(key, data); err != nil {
			log.Println("Failed to put into store - " + err.Error())
			return ctx.Reply("Failed to subscribe to safe, sorry :(")
		}

		return ctx.Reply(fmt.Sprintf("Successfully subscribed to *%s*\\!", mdEscaper.Replace(label)), tg.ModeMarkdownV2)
	})

	b.Handle("/subscriptions", func(ctx tg.Context) error {
		// Getting recipient
		// TODO: put in a separate function
		key := []byte(ctx.Recipient().Recipient())
		data, err := store.Get(key)
		if err != nil {
			log.Println("Failed to get from store - " + err.Error())
			return ctx.Reply("Failed to get subscriptions, sorry :(")
		}

		var recipient models.Recipient
		if data != nil {
			recipient, err = storage.FromGob[models.Recipient](data)
			if err != nil {
				log.Println("Failed to get from gob - " + err.Error())
				return ctx.Reply("Failed to get subscriptions, sorry :(")
			}
		}

		// Output
		subscriptionsLength := len(recipient.Subscriptions)
		if subscriptionsLength == 0 {
			return ctx.Reply("You are not currently subscribed to any safe")
		}

		var b strings.Builder
		b.WriteString("List of subscriptions:\n")
		for _, sub := range recipient.Subscriptions {
			label := mdEscaper.Replace(sub.Label)
			lastUpdated := humantime.Since(time.Unix(sub.LastUpdated, 0))
			b.WriteString(fmt.Sprintf("*%s* — Last update was %s\n", label, lastUpdated))

			network := client.Network(sub.Network)
			gnosisSafe := network.AsPrefix() + ":" + sub.SafeAddress
			gnosisLink := fmt.Sprintf("https://gnosis-safe.io/app/%s/home", gnosisSafe)
			b.WriteString(fmt.Sprintf("[%s](%s)\n", gnosisSafe, gnosisLink))
		}
		return ctx.Reply(b.String(), tg.ModeMarkdownV2, tg.NoPreview)
	})

	b.Handle("/unsubscribe", func(ctx tg.Context) error {
		args := ctx.Args()

		if len(args) < 1 {
			var b strings.Builder
			b.WriteString("Usage:\n")
			b.WriteString("/unsubscribe\n")
			b.WriteString("\tmy-dao\n")
			return ctx.Reply(b.String())
		}

		// Input stuff
		label := args[0]

		// Getting recipient
		// TODO: put in a separate function
		key := []byte(ctx.Recipient().Recipient())
		data, err := store.Get(key)
		if err != nil {
			log.Println("Failed to get from store - " + err.Error())
			return ctx.Reply("Failed to unsubscribe from safe, sorry :(")
		}

		var recipient models.Recipient
		if data != nil {
			recipient, err = storage.FromGob[models.Recipient](data)
			if err != nil {
				log.Println("Failed to get from gob - " + err.Error())
				return ctx.Reply("Failed to unsubscribe from safe, sorry :(")
			}
		}

		// Deleting from subs
		for i, sub := range recipient.Subscriptions {
			if label == sub.Label {
				recipient.Subscriptions = append(recipient.Subscriptions[:i], recipient.Subscriptions[i+1:]...)
			}
		}

		// Saving
		// TODO: put in a separate function
		data, err = storage.ToGob(recipient)
		if err != nil {
			log.Println("Failed to convert into gob - " + err.Error())
			return ctx.Reply("Failed to unsubscribe from safe, sorry :(")
		}

		if err = store.Put(key, data); err != nil {
			log.Println("Failed to put into store - " + err.Error())
			return ctx.Reply("Failed to unsubscribe from safe, sorry :(")
		}

		return ctx.Reply(fmt.Sprintf("Successfully unsubscribed from *%s*\\!", mdEscaper.Replace(label)), tg.ModeMarkdownV2)
	})
}

// New ...
func New(lifecycle fx.Lifecycle, cfg config.Telegram) (*tg.Bot, error) {
	b, err := tg.NewBot(tg.Settings{
		Token:  cfg.BotToken,
		Poller: &tg.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go b.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			b.Stop()
			return nil
		},
	})

	return b, nil
}

// Module ...
var Module = fx.Module("bot",
	fx.Provide(New),
	fx.Invoke(RegisterHandler),
)
