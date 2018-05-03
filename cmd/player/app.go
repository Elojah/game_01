package main

import (
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
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
		go a.AddListener(msg)
	}
}

func (a *app) AddListener(msg *nats.Msg) {
	logger := log.With().Str("event", msg.Subject).Logger()

	var eventS storage.Event
	if _, err := eventS.Unmarshal(msg.Data); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal event")
		return
	}

	event := eventS.Domain()
	switch event.Action.(type) {
	case game.Listener:
	default:
		logger.Error().Msg("wrong action type")
		return
	}

	listener := event.Action.(game.Listener)
	_, ch, err := a.ReceiveEvent(listener.ID.String(), a.bufsize)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	select {
	case msg := <-ch:
		go a.Play(msg)
	}
}

func (a *app) Play(msg *nats.Msg) {
	println(msg.Subject)
}
