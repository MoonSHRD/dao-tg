package main

import (
	"context"

	"github.com/gotd/td/tg"
)

// RegisterHandler ...
func RegisterHandler(bot Bot) {
	bot.OnNewMessage(func(ctx context.Context, e tg.Entities, u *tg.UpdateNewMessage) error {
		m, ok := u.Message.(*tg.Message)
		if !ok || m.Out {
			// Outgoing message, not interesting.
			return nil
		}

		// Sending reply.
		_, err := bot.Reply(e, u).Text(ctx, m.Message)
		return err
	})
}
