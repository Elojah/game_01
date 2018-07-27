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
	ability.FeedbackMapper
	AbilityTemplateMapper ability.TemplateMapper
	AbilityMapper         ability.Mapper

	account.TokenMapper

	EntityTemplateMapper entity.TemplateMapper
	EntityMapper         entity.Mapper

	infra.QListenerMapper
	event.QMapper
	EventMapper event.Mapper

	infra.CoreMapper

	entity.PermissionMapper

	sector.EntitiesMapper
	SectorMapper sector.Mapper

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
	a.subs[a.id] = a.SubscribeListener(a.id)
	go func(sub *infra.Subscription) {
		for msg := range sub.Channel() {
			go a.AddListener(msg)
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

func (a *app) AddListener(msg *infra.Message) {
	logger := log.With().Str("core", a.id.String()).Logger()

	var listener infra.Listener
	if _, err := listener.Unmarshal([]byte(msg.Payload)); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal listener")
		return
	}
	logger = logger.With().Str("listener", listener.ID.String()).Uint8("action", uint8(listener.Action)).Logger()

	switch listener.Action {
	case infra.Open:
		a.seqs[listener.ID] = NewSequencer(listener.ID, a.limit, a.EventMapper, a.Apply)
		a.seqs[listener.ID].Run()

		a.subs[listener.ID] = a.SubscribeEvent(listener.ID)

		go func(seq *Sequencer, sub *infra.Subscription) {
			for msg := range sub.Channel() {
				seq.Handler(msg)
			}
		}(a.seqs[listener.ID], a.subs[listener.ID])

		logger.Info().Msg("listening")
	case infra.Close:
		seq, ok := a.seqs[listener.ID]
		if !ok {
			logger.Error().Msg("listener not found")
			return
		}
		seq.Close()
		sub, ok := a.subs[listener.ID]
		if !ok {
			logger.Error().Msg("subscription not found")
			return
		}
		if err := sub.Unsubscribe(); err != nil {
			logger.Error().Err(err).Msg("failed to unsubscribe")
			return
		}
		delete(a.seqs, listener.ID)
		delete(a.subs, listener.ID)
	}
}

func (a *app) Apply(id ulid.ID, e event.E) {
	logger := log.With().
		Str("core", a.id.String()).
		Str("listener", id.String()).
		Int64("ts", e.TS.UnixNano()).
		Str("type", event.String(e.Action)).
		Logger()

	switch e.Action.(type) {
	case event.Move:
		if err := a.Move(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case event.Cast:
		if err := a.Cast(id, e); err != nil {
			logger.Error().Err(err).Msg("event rejected")
		}
	case event.Feedback:
		logger.Error().Msg("not implemented")
	default:
		logger.Error().Msg("unrecognized action")
	}
	logger.Info().Msg("event applied")
}
