package main

import (
	"time"

	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

type app struct {
	game.Services

	subject string
	bufsize int

	listenerSub  *game.Subscription
	listenerChan game.MsgChan
	subs         [](*game.Subscription)
}

func (a *app) Dial(c Config) error {
	a.subject = c.Subject
	a.bufsize = c.Bufsize
	return nil
}

func (a *app) Start() {
	logger := log.With().Str("coord", a.subject).Logger()

	sub, ch, err := a.ReceiveEvent(a.subject, a.bufsize)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	a.listenerSub = (*game.Subscription)(sub)
	a.listenerChan = (game.MsgChan)(ch)
	for {
		select {
		case msg := <-ch:
			go a.AddListener(msg)
		}
	}
}

func (a *app) Close() {
	for _, s := range a.subs {
		s.Unsubscribe()
	}
	a.listenerSub.Unsubscribe()
	close(a.listenerChan)
}

func (a *app) AddListener(msg *nats.Msg) {
	logger := log.With().Str("event", msg.Subject).Logger()

	var listenerS storage.Listener
	if _, err := listenerS.Unmarshal(msg.Data); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal listener")
		return
	}
	listener := listenerS.Domain()
	id := listener.ID

	sub, ch, err := a.ReceiveEvent(id.String(), a.bufsize)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	a.subs = append(a.subs, (*game.Subscription)(sub))
	logger.Info().Str("id", id.String()).Msg("listening")

	go a.Listen(ch, id)
}

func (a *app) Listen(ch game.MsgChan, id game.ID) {
	var last time.Time
	var closer chan struct{}
	logger := log.With().Str("listener", id.String()).Logger()
	for {
		select {
		case msg := <-ch:
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
			if event.TS.Before(last) {
				closer <- struct{}{}
				// close closer somewhere...
			}
			closer = a.ReplayFrom(event, id)
		}
	}
}

func (a *app) ReplayFrom(event game.Event, id game.ID) chan struct{} {
	closer := make(chan struct{}, 0)
	logger := log.With().Str("replay", event.ID.String()).Logger()
	go func() {
		events, err := a.ListEvent(game.EventBuilder{
			Key:   id.String(),
			Start: event.TS.UnixNano(),
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to replay events")
			return
		}
		for _, e := range events {
			select {
			case _, _ = <-closer:
				return
			default:
				logger.Info().Msg("replay event" + e.TS.String())
			}
		}
	}()
	return closer
}
