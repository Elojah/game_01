package main

import (
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

type app struct {
	game.Services

	subject string
	bufsize int
}

func (a *app) Dial(c Config) error {
	a.subject = c.Subject
	a.bufsize = c.Bufsize
	return nil
}

func (a *app) Start() {

	logger := log.With().Str("player", a.subject).Logger()

	_, ch, err := a.ReceiveEvent(a.subject, a.bufsize)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	select {
	case msg := <-ch:
		go a.Listen(msg)
	}
}

func (a *app) Listen(msg *nats.Msg) {
	println(msg.Subject)
}
