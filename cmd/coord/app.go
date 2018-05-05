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

	logger := log.With().Str("coord", a.subject).Logger()

	_, ch, err := a.ReceiveEvent(a.subject, a.bufsize)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	for {
		select {
		case msg := <-ch:
			go a.AddListener(msg)
		}
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
		logger.Error().Str("expected", "listener").Msg("invalid action type")
		return
	}

	listener := event.Action.(game.Listener)
	id := listener.ID.String()
	_, ch, err := a.ReceiveEvent(id, a.bufsize)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	logger.Info().Str("id", id).Msg("listening")
	go a.Play(ch)
}

func (a *app) Play(ch chan *nats.Msg) {
	var current uint64
	for {
		select {
		case msg := <-ch:
			logger := log.With().Str("listener", msg.Subject).Logger()
			var event storage.Event
			if _, err := event.Unmarshal(msg.Data); err != nil {
				logger.Error().Err(err).Msg("error unmarshaling event")
				break
			}

			_ = event
			_ = current
		}
	}
}
