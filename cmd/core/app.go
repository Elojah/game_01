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
	game.PCService
	game.PCLeftService
	game.PermissionService
	game.QEventService
	game.SubscriptionService
	game.TemplateService
	game.TokenService

	id game.ID

	subs map[game.ID]*game.Subscription
	seqs map[game.ID]*Sequencer

	limit         int
	moveTolerance float64

	listeners []string
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit
	a.moveTolerance = c.MoveTolerance
	a.listeners = make([]string, len(c.Listeners))
	copy(a.listeners, c.Listeners)

	return nil
}

func (a *app) Start() {
	logger := log.With().Str("core", a.id.String()).Logger()

	sub, err := a.SetSubscription(a.id.String(), a.AddListener)
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
	logger := log.With().Str("app", a.id.String()).Logger()

	var listenerS storage.Listener
	if _, err := listenerS.Unmarshal(msg.Data); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal listener")
		return
	}
	id := listenerS.Domain().ID

	seq := NewSequencer(a.id, a.limit, a.EventService, a.Apply)
	a.seqs[id] = seq

	sub, err := a.SetSubscription(id.String(), seq.MsgHandler)
	if err != nil {
		logger.Error().Err(err).Str("listener", id.String()).Msg("failed to sub")
		return
	}
	a.subs[id] = sub

	logger.Info().Str("listener", id.String()).Msg("listening")
}

func (a *app) Apply(id game.ID, event game.Event) {
	ts := event.TS.UnixNano()
	key := id.String()
	logger := log.With().
		Str("event", key).
		Str("source", event.Source.String()).
		Int64("ts", ts).
		Logger()

	logger.Info().Str("type", game.ActionString(event.Action)).Msg("apply action")
	switch event.Action.(type) {
	case game.MoveDone:
		if err := a.MoveDone(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.MoveReceived:
		if err := a.MoveReceived(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.AttackReceived:
		if err := a.AttackReceived(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.AttackDone:
		if err := a.AttackDone(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.HealReceived:
		if err := a.HealReceived(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.HealDone:
		if err := a.HealDone(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.SetEntity:
		if err := a.CreateEntity(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.ConnectPC:
		if err := a.ConnectPC(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case game.SetPC:
		if err := a.CreatePC(event); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	default:
		logger.Error().Msg("unrecognized action")
	}
}
