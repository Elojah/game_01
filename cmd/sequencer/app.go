package main

import (
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

type app struct {
	game.Services
	id game.ID

	subs map[game.ID]game.Subscription

	limit int
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit

	return nil
}

func (a *app) Start() {
	logger := log.With().Str("coord", a.id.String()).Logger()

	sub, err := a.CreateSubscription(a.id.String(), a.limit)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	a.subs = make(map[game.ID]game.Subscription)
	a.subs[a.id] = sub
	for {
		select {
		case msg := <-sub.Ch:
			go a.AddListener(msg)
		}
	}
}

func (a *app) Close() {
	for _, l := range a.subs {
		l.Close()
	}
}

func (a *app) AddListener(msg *nats.Msg) {
	logger := log.With().Str("event", a.id.String()).Logger()

	var listenerS storage.Listener
	if _, err := listenerS.Unmarshal(msg.Data); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal listener")
		return
	}
	id := listenerS.Domain().ID

	sub, err := a.CreateSubscription(id.String(), a.limit)
	if err != nil {
		logger.Error().Err(err).Str("id", id.String()).Msg("failed to sub")
		return
	}
	a.subs[id] = sub
	logger.Info().Str("id", id.String()).Msg("listening")

	a.Listen(sub.Ch, id)
}

func (a *app) Listen(ch game.MsgChan, id game.ID) {
	r := a.newReplayer(id)
	logger := log.With().Str("listener", id.String()).Logger()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				r.Close()
				return
			}
			var eventS storage.Event
			if _, err := eventS.Unmarshal(msg.Data); err != nil {
				logger.Error().Err(err).Msg("error unmarshaling event")
				break
			}
			event := eventS.Domain()
			if err := a.CreateEvent(event, id); err != nil {
				logger.Error().Err(err).Msg("error creating event")
				break
			}
			logger.Info().Str("event", event.ID.String()).Msg("event received")
			r.Trigger <- event.TS
		}
	}
}

func (a *app) Play(event game.Event) {
	logger := log.With().
		Str("event", event.ID.String()).
		Int("ts", int(event.TS.UnixNano())).
		Logger()
	logger.Info().Msg("play event")
}
