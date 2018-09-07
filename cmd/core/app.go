package main

import (
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

type app struct {
	AbilityStore         ability.Store
	AbilityTemplateStore ability.TemplateStore
	ability.FeedbackStore

	account.TokenStore

	EntityTemplateStore entity.TemplateStore
	EntityStore         entity.Store
	entity.PermissionStore

	infra.QSequencerStore
	infra.CoreStore

	EventQStore event.QStore
	EventStore  event.Store

	sector.EntitiesStore
	SectorStore sector.Store

	id ulid.ID

	subs map[ulid.ID]*infra.Subscription
	seqs map[ulid.ID]*Sequencer

	limit         int
	moveTolerance float64
}

func (a *app) Dial(c Config) error {
	a.id = c.ID
	a.limit = c.Limit
	a.moveTolerance = c.MoveTolerance
	go a.Run()
	return nil
}

func (a *app) Run() {
	logger := log.With().Str("core", a.id.String()).Logger()

	a.subs = make(map[ulid.ID]*infra.Subscription)
	a.subs[a.id] = a.SubscribeSequencer(a.id)
	go func(sub *infra.Subscription) {
		for msg := range sub.Channel() {
			go a.Sequencer(msg)
		}
	}(a.subs[a.id])

	a.seqs = make(map[ulid.ID]*Sequencer)

	if err := a.SetCore(infra.Core{ID: a.id}); err != nil {
		logger.Error().Err(err).Msg("failed to set core")
		return
	}
}

func (a *app) Close() {
	for _, s := range a.subs {
		s.Unsubscribe()
	}
	for _, s := range a.seqs {
		s.Close()
	}
}

func (a *app) Sequencer(msg *infra.Message) {
	logger := log.With().Str("core", a.id.String()).Logger()

	var sequencer infra.Sequencer
	if err := sequencer.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal sequencer")
		return
	}
	logger = logger.With().Str("sequencer", sequencer.ID.String()).Logger()

	switch sequencer.Action {
	case infra.Open:
		a.seqs[sequencer.ID] = NewSequencer(sequencer.ID, a.limit, a.Apply)
		a.seqs[sequencer.ID].EventStore = a.EventStore
		a.seqs[sequencer.ID].EntityStore = a.EntityStore
		a.seqs[sequencer.ID].Run()

		a.subs[sequencer.ID] = a.EventQStore.SubscribeEvent(sequencer.ID)

		go func(seq *Sequencer, sub *infra.Subscription) {
			for msg := range sub.Channel() {
				seq.Handler(msg)
			}
		}(a.seqs[sequencer.ID], a.subs[sequencer.ID])

		logger.Info().Msg("sequencer up")
	case infra.Close:
		seq, ok := a.seqs[sequencer.ID]
		if !ok {
			logger.Error().Msg("sequencer not found")
			return
		}
		seq.Close()
		sub, ok := a.subs[sequencer.ID]
		if !ok {
			logger.Error().Msg("subscription not found")
			return
		}
		if err := sub.Unsubscribe(); err != nil {
			logger.Error().Err(err).Msg("failed to unsubscribe")
			return
		}
		delete(a.seqs, sequencer.ID)
		delete(a.subs, sequencer.ID)
		logger.Info().Msg("sequencer down")
	}
}

func (a *app) Apply(id ulid.ID, e event.E) {
	logger := log.With().
		Str("core", a.id.String()).
		Str("sequencer", id.String()).
		Int64("ts", e.TS.UnixNano()).
		Str("type", e.Action.Type()).
		Logger()

	switch e.Action.GetValue().(type) {
	case *event.MoveSource:
		if err := a.MoveSource(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case *event.MoveTarget:
		if err := a.MoveTarget(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case *event.CastSource:
		if err := a.CastSource(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case *event.PerformSource:
		if err := a.PerformSource(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case *event.PerformTarget:
		if err := a.PerformTarget(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case *event.FeedbackTarget:
		logger.Error().Msg("not implemented")
	default:
		logger.Error().Msg("unrecognized action")
	}
	logger.Info().Msg("event applied")
}
