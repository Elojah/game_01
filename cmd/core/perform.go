package main

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (a *app) PerformSource(id gulid.ID, e event.E) error {

	ps := e.Action.PerformSource
	ts := e.ID.Time()

	// #Retrieve source
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "perform source")
	}

	// #Check if entity is alive
	if source.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "perform source")
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(source.ID, ps.AbilityID)
	if err != nil {
		return errors.Wrap(err, "perform source")
	}

	// #Check cast was not interrupted.
	if source.Cast == nil ||
		source.Cast.AbilityID != ps.AbilityID ||
		source.Cast.TS+ab.CastTime != ts {
		// normal behavior, don't return errors
		return nil
	}

	// #Set ability LastUsed
	// LastUsed is set in PerformSource and not in feedback because there are potentially X multiple targets.
	// If one or all target(s) fails (e.g: too far), we still apply to other targets.
	ab.LastUsed = ts
	if err := a.AbilityStore.SetAbility(ab, source.ID); err != nil {
		return errors.Wrap(err, "perform source")
	}

	// #For all ability components.
	var i uint64
	for cid := range ab.Components {

		// #Retrieve targets for this component.
		target, ok := ps.Targets[cid]
		if !ok {
			return errors.Wrap(
				gerrors.ErrMissingTarget{
					AbilityID:   ab.ID.String(),
					ComponentID: cid,
				},
				"perform source",
			)
		}

		// #Send event to all targets
		e := event.E{
			ID: gulid.NewTimeID(ts + i + 1),
			Action: event.Action{
				PerformTarget: &event.PerformTarget{
					AbilityID:   ab.ID,
					ComponentID: gulid.MustParse(cid),
					Source:      source,
				},
			},
			Trigger: e.ID,
		}

		if len(target.Positions) != 0 {
			return errors.Wrap(gerrors.ErrNotImplementedYet{Version: "0.2.0"}, "perform source")
		}

		var g errgroup.Group
		for _, id := range target.Entities {
			id := id
			g.Go(func() error {
				if err := a.EventQStore.PublishEvent(e, id); err != nil {
					return errors.Wrap(err, "perform source")
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		i++
	}
	return nil
}

func (a *app) PerformTarget(id gulid.ID, e event.E) error {

	pt := e.Action.PerformTarget
	ts := e.ID.Time()

	// #Retrieve previous target state.
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(pt.Source.ID, pt.AbilityID)
	if err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Initialize feedback.
	fb := ability.Feedback{
		ID:          gulid.NewID(),
		AbilityID:   ab.ID,
		ComponentID: pt.ComponentID,
	}

	// #Check component validity.
	cid := pt.ComponentID.String()
	component, ok := ab.Components[cid]
	if !ok {
		return errors.Wrap(
			gerrors.ErrMissingTarget{
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			},
			"perform target",
		)
	}

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(pt.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "perform target")
	}
	if dist > component.Range {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: component.Range,
			},
			"perform target",
		)
	}

	wasDead := target.Dead

	// #Apply all ability components.
	if fb.Effects, err = target.ApplyEffects(&pt.Source, component.Effects); err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Add a spawn event to a freshly dead entity.
	// wasDead is a guard against multiple spawn events.
	if !wasDead && target.Dead {
		sp, err := a.EntitySpawnStore.GetSpawn(target.SpawnID)
		if err != nil {
			return errors.Wrap(err, "perform target")
		}
		if err := a.EventQStore.PublishEvent(event.E{
			ID: gulid.NewTimeID(ts + sp.Duration + 1),
			Action: event.Action{
				Spawn: &event.Spawn{
					ID: target.SpawnID,
				},
			},
			Trigger: e.ID,
		}, id); err != nil {
			return errors.Wrapf(err, "perform target")
		}
	}

	// #Set entity new state.
	if err := a.EntityStore.SetEntity(target, ts); err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Set feedback.
	if err := a.FeedbackStore.SetFeedback(fb); err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Publish feedback to source.
	return errors.Wrap(a.EventQStore.PublishEvent(event.E{
		ID: gulid.NewTimeID(ts + 1),
		Action: event.Action{
			PerformFeedback: &event.PerformFeedback{
				ID:     fb.ID,
				Target: target,
			},
		},
		Trigger: e.ID,
	}, pt.Source.ID),
		"perform target",
	)
}

func (a *app) PerformFeedback(id gulid.ID, e event.E) error {

	ft := e.Action.PerformFeedback
	ts := e.ID.Time()

	// #Retrieve previous source state.
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "perform feedback")
	}

	// #Retrieve feedback.
	fb, err := a.FeedbackStore.GetFeedback(ft.ID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrap(err, "perform feedback")
	}

	// #Apply all ability components.
	if err := source.ApplyEffectFeedbacks(&ft.Target, fb.Effects); err != nil {
		return errors.Wrap(err, "perform feedback")
	}

	// #Set entity new state.
	return errors.Wrap(a.EntityStore.SetEntity(source, ts), "perform feedback")
}
