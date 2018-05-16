package main

import (
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

type app struct {
	game.EntityService
	game.EventService
	game.PermissionService
	game.QEventService
	game.SubscriptionService

	id game.ID

	subs map[game.ID]*game.Subscription
	seqs map[game.ID]*Sequencer

	limit int
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit

	return nil
}

func (a *app) Start() {
	logger := log.With().Str("core", a.id.String()).Logger()

	sub, err := a.CreateSubscription(a.id.String(), a.AddListener)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	a.subs = make(map[game.ID]*game.Subscription)
	a.subs[a.id] = sub

	a.seqs = make(map[game.ID]*Sequencer)
}

func (a *app) Close() {
	for _, s := range a.subs {
		s.Unsubscribe()
	}
	for _, s := range a.seqs {
		s.Close()
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

	seq := NewSequencer(a.id, a.limit, a.EventService, a.Apply)
	a.seqs[id] = seq

	sub, err := a.CreateSubscription(id.String(), seq.MsgHandler)
	if err != nil {
		logger.Error().Err(err).Str("id", id.String()).Msg("failed to sub")
		return
	}
	a.subs[id] = sub

	logger.Info().Str("id", id.String()).Msg("listening")
}

func (a *app) Apply(id game.ID, event game.Event) {
	ts := event.TS.UnixNano()
	key := id.String()
	logger := log.With().
		Str("event", key).
		Str("source", event.Source.String()).
		Int64("ts", ts).
		Logger()

	switch event.Action.(type) {
	case game.Move:
		logger.Info().Str("type", "move").Msg("apply action")
		if err := a.Move(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.DamageReceived:
		logger.Info().Str("type", "damage_received").Msg("apply action")
		if err := a.DamageReceived(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.DamageDone:
		logger.Info().Str("type", "damage_done").Msg("apply action")
		if err := a.DamageDone(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.HealReceived:
		logger.Info().Str("type", "heal_received").Msg("apply action")
		if err := a.HealReceived(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.HealDone:
		logger.Info().Str("type", "heal_done").Msg("apply action")
		if err := a.HealDone(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	}
}
