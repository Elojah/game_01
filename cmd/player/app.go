package main

import (
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
)

type app struct {
	Config
	game.Services
}

func (a *app) Start() {

	logger := log.With().Str("player", a.Subject).Logger()

	_, ch, err := a.ReceiveEvent(a.Subject, a.Bufsize)
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
