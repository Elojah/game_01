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

	seq := NewSequencer(a.id, a.EventService, a.Apply)
	a.seqs[id] = seq

	sub, err := a.CreateSubscription(id.String(), seq.MsgHandler)
	if err != nil {
		logger.Error().Err(err).Str("id", id.String()).Msg("failed to sub")
		return
	}
	a.subs[id] = sub

	logger.Info().Str("id", id.String()).Msg("listening")
}

func (a *app) Apply(event game.Event) {
	logger := log.With().
		Str("event", event.ID.String()).
		Int("ts", int(event.TS.UnixNano())).
		Logger()
	switch event.Action.(type) {
	case game.Damage:
		logger.Info().Str("type", "damage").Msg("apply action")
	case game.DamageInflict:
		logger.Info().Str("type", "damage_inflict").Msg("apply action")
	case game.Heal:
		logger.Info().Str("type", "heal").Msg("apply action")
	case game.HealInflict:
		logger.Info().Str("type", "heal_inflict").Msg("apply action")
	}
}