package main

import (
	nats "github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/storage"
)

type app struct {
	ability.FeedbackMapper
	AbilityTemplateMapper ability.TemplateMapper
	AbilityMapper         ability.Mapper

	account.TokenMapper

	EntityTemplateMapper entity.TemplateMapper
	EntityMapper         entity.Mapper

	event.QListenerMapper
	event.QMapper
	event.SubscriptionMapper
	EventMapper event.Mapper

	game.PermissionMapper

	sector.EntitiesMapper
	SectorMapper sector.Mapper

	id game.ID

	subs map[game.ID]*event.Subscription
	seqs map[game.ID]*Sequencer

	limit         int
	moveTolerance float64

	cores []game.ID
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit
	a.moveTolerance = c.MoveTolerance
	a.cores = make([]game.ID, len(c.Cores))
	copy(a.cores, c.Cores)

	return nil
}

func (a *app) Start() {
	logger := log.With().Str("core", a.id.String()).Logger()

	sub, err := a.SetSubscription(a.id.String(), a.AddListener)
	if err != nil {
		logger.Error().Err(err).Msg("failed to sub")
		return
	}
	a.subs = make(map[game.ID]*event.Subscription)
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

	seq := NewSequencer(a.id, a.limit, a.EventMapper, a.Apply)
	a.seqs[id] = seq
	seq.Start()

	sub, err := a.SetSubscription(id.String(), seq.MsgHandler)
	if err != nil {
		logger.Error().Err(err).Str("listener", id.String()).Msg("failed to sub")
		return
	}
	a.subs[id] = sub

	logger.Info().Str("listener", id.String()).Msg("listening")
}

func (a *app) Apply(id game.ID, e event.E) {
	ts := e.TS.UnixNano()
	key := id.String()
	logger := log.With().
		Str("event", key).
		Str("source", e.Source.String()).
		Int64("ts", ts).
		Logger()

	logger.Info().Str("type", event.ActionString(e.Action)).Msg("apply action")
	switch e.Action.(type) {
	case event.Move:
		if err := a.Move(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case event.Cast:
		if err := a.Cast(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	default:
		logger.Error().Msg("unrecognized action")
	}
}
